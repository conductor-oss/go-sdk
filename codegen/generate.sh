#!/bin/bash

# ========================================================================
# OpenAPI Go SDK Generation via Docker
# Generates Go client and models from OpenAPI specification
#
# REQUIREMENTS:
#   - Docker installed and running
#   - Python 3 for spec fixes
#
# USAGE:
#   cd codegen && ./generate.sh <client_version>
#   
# EXAMPLES:
#   ./generate.sh orkes        # Generates in sdk/generated/http/orkes
#   ./generate.sh conductor       # Generates in sdk/generated/http/conductor
# ========================================================================

set -e

# Check for required argument
if [ $# -eq 0 ]; then
    echo "❌ Error: Client version/name is required!"
    echo ""
    echo "Usage: $0 <client_version>"
    echo ""
    echo "Examples:"
    echo "  $0 orkes        # For Orkes Cloud version"
    echo "  $0 community    # For Community version"
    echo "  $0 v2           # For API v2"
    echo ""
    exit 1
fi

CLIENT_VERSION="$1"

echo "============================================"
echo "🚀 GO SDK GENERATION"
echo "============================================"
echo "Client version: ${CLIENT_VERSION}"
echo ""

# Configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
INPUT_SPEC="${PROJECT_ROOT}/codegen/api-docs.json"
FIXED_SPEC="${SCRIPT_DIR}/openapi-fixed.json"
OUTPUT_DIR="${PROJECT_ROOT}/sdk/generated/http/${CLIENT_VERSION}"
PACKAGE_NAME="${CLIENT_VERSION}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# ========================================================================
# STEP 1: SPECIFICATION FIXES
# ========================================================================
echo -e "${YELLOW}📝 Step 1: Fixing OpenAPI specification...${NC}"

# Run spec fixes via standalone Python script
cd "$SCRIPT_DIR"
python3 "${SCRIPT_DIR}/fix_spec.py" "$INPUT_SPEC" "$FIXED_SPEC"

# If client version is orkes_v4, filter the spec to include only Integration API
if [ "$CLIENT_VERSION" == "orkes_v4" ]; then
    echo -e "${YELLOW}📋 Filtering specification for orkes_v4 (Integration API only)...${NC}"
    python3 "${SCRIPT_DIR}/filter_integration_api.py" "$FIXED_SPEC" "$FIXED_SPEC"
    if [ $? -ne 0 ]; then
        echo -e "${RED}❌ Failed to filter specification${NC}"
        exit 1
    fi
    echo -e "${GREEN}✓ Specification filtered successfully${NC}"
fi

# ========================================================================
# STEP 2: CHECK DOCKER
# ========================================================================
echo ""
echo -e "${YELLOW}🐳 Step 2: Checking Docker...${NC}"

# Check for Docker
if ! command -v docker &> /dev/null; then
    echo -e "  ${RED}❌ Docker is not installed!${NC}"
    echo ""
    echo "  Please install Docker from https://www.docker.com"
    exit 1
fi

# Check if Docker daemon is running
if ! docker info &> /dev/null; then
    echo -e "  ${RED}❌ Docker daemon is not running!${NC}"
    echo ""
    echo "  Please start Docker Desktop or Docker daemon"
    exit 1
fi

echo -e "  ${GREEN}✓ Docker is ready${NC}"

# Pull the image if needed
DOCKER_IMAGE="openapitools/openapi-generator-cli:v7.1.0"
echo "  Checking Docker image..."
if ! docker image inspect "$DOCKER_IMAGE" &> /dev/null; then
    echo "  Pulling $DOCKER_IMAGE..."
    docker pull "$DOCKER_IMAGE"
fi

# ========================================================================
# STEP 3: CLEAN OLD GENERATION
# ========================================================================
echo ""
echo -e "${YELLOW}🧹 Step 3: Cleaning old generation...${NC}"

if [ -d "$OUTPUT_DIR" ]; then
    echo "  Removing old generated files..."
    rm -rf "$OUTPUT_DIR"
fi
mkdir -p "$OUTPUT_DIR"

# ========================================================================
# STEP 4: SDK GENERATION
# ========================================================================
echo ""
echo -e "${YELLOW}🚀 Step 4: Generating SDK...${NC}"
echo "  This may take a few minutes..."

# Docker volume paths
DOCKER_WORKDIR="/local"
DOCKER_SPEC="${DOCKER_WORKDIR}/codegen/openapi-fixed.json"
DOCKER_OUTPUT="${DOCKER_WORKDIR}/sdk/generated/http/${CLIENT_VERSION}"

# Run generation
docker run --rm \
    -v "${PROJECT_ROOT}:${DOCKER_WORKDIR}" \
    "$DOCKER_IMAGE" generate \
    -i "$DOCKER_SPEC" \
    -g go \
    -o "$DOCKER_OUTPUT" \
    --package-name="$PACKAGE_NAME" \
    --skip-validate-spec \
    --git-user-id="conductor-sdk" \
    --git-repo-id="conductor-go/sdk/generated/http/${CLIENT_VERSION}" \
    --global-property models,apis,supportingFiles \
    --additional-properties=packageVersion=1.0.0 \
    --additional-properties=isGoSubmodule=false \
    --additional-properties=withGoMod=false \
    --additional-properties=generateInterfaces=true \
    --additional-properties=structPrefix=true \
    --additional-properties=enumClassPrefix=true \
    2>&1 | while IFS= read -r line; do
        # Filter output for readability
        if [[ $line == *"writing"* ]] || [[ $line == *"Generating"* ]] || [[ $line == *"ERROR"* ]] || [[ $line == *"WARNING"* ]]; then
            echo "  $line"
        fi
    done

# ========================================================================
# STEP 5: CHECK RESULTS
# ========================================================================
echo ""
echo -e "${YELLOW}📋 Step 5: Checking generated files...${NC}"

if [ -d "$OUTPUT_DIR" ]; then
    GO_FILES=$(find "$OUTPUT_DIR" -name "*.go" | wc -l)
    if [ "$GO_FILES" -gt 0 ]; then
        echo -e "  ${GREEN}✅ Generated $GO_FILES Go files${NC}"
        
        # Show file structure
        echo ""
        echo "  File structure:"
        ls -la "$OUTPUT_DIR" | head -20
        
        # Count different types of files
        MODEL_FILES=$(find "$OUTPUT_DIR" -name "model_*.go" | wc -l)
        API_FILES=$(find "$OUTPUT_DIR" -name "api_*.go" | wc -l)
        
        echo ""
        echo "  Generated:"
        echo "    - Models: $MODEL_FILES files"
        echo "    - APIs: $API_FILES files"
    else
        echo -e "  ${RED}❌ No Go files were generated${NC}"
        exit 1
    fi
else
    echo -e "  ${RED}❌ Output directory was not created${NC}"
    exit 1
fi

# ========================================================================
# CLEANUP
# ========================================================================
echo ""
echo -e "${YELLOW}🧹 Step 6: Cleanup...${NC}"

# Remove temporary files (keep fix_spec.py as a maintained script)
rm -f "$FIXED_SPEC"
echo "  Temporary files removed (kept fix_spec.py)"

# ========================================================================
# COMPLETION
# ========================================================================
echo ""
echo "============================================"
echo -e "${GREEN}✅ GENERATION COMPLETED!${NC}"
# ========================================================================
# STEP 7: POST-PROCESSING - FIX MAP TYPES
# ========================================================================
echo ""
echo "============================================"
echo "🔧 POST-PROCESSING: Fixing map types"
echo "============================================"

# Check if Python is available
if command -v python3 &> /dev/null; then
    echo "  Applying generated code fixes..."
    
    # Apply map type fixes
    if [ -f "${SCRIPT_DIR}/fix_map_types.py" ]; then
        if python3 "${SCRIPT_DIR}/fix_map_types.py" "${OUTPUT_DIR}"; then
            echo -e "  ${GREEN}✓ Map types fixed successfully${NC}"
        else
            echo -e "  ${YELLOW}⚠ Warning: Failed to apply map type fixes${NC}"
        fi
    fi
    
    # Apply generated code fixes
    if [ -f "${SCRIPT_DIR}/fix_generated_code.py" ]; then
        if python3 "${SCRIPT_DIR}/fix_generated_code.py" "${OUTPUT_DIR}" --verbose; then
            echo -e "  ${GREEN}✓ Generated code fixes applied successfully${NC}"
        else
            echo -e "  ${YELLOW}⚠ Warning: Failed to apply generated code fixes${NC}"
        fi
    else
        echo -e "  ${YELLOW}⚠ Warning: fix_generated_code.py not found, skipping generated code fixes${NC}"
    fi

    if [ -f "${SCRIPT_DIR}/add_omitempty_to_all_fields.py" ]; then
        if python3 "${SCRIPT_DIR}/add_omitempty_to_all_fields.py" "${OUTPUT_DIR}" --verbose; then
            echo -e "  ${GREEN}✓ Added omitempty tags to all fields successfully${NC}"
        else
            echo -e "  ${YELLOW}⚠ Warning: Failed to add omitempty tags${NC}"
        fi
    fi

else
    echo -e "  ${YELLOW}⚠ Warning: Python3 not found, skipping post-processing${NC}"
    echo "  You can manually run:"
    echo "    python3 ${SCRIPT_DIR}/fix_map_types.py ${OUTPUT_DIR}"
    echo "    python3 ${SCRIPT_DIR}/fix_generated_code.py ${OUTPUT_DIR}"
    echo "    python3 ${SCRIPT_DIR}/add_omitempty_to_all_fields.py ${OUTPUT_DIR}"
fi

# ========================================================================
# STEP 8: POST-PROCESS VERIFY (build + fmt)
# ========================================================================
echo ""
echo -e "${YELLOW}🧪 Verifying build after post-processing...${NC}"

# Keep post-process verification minimal; all content fixes live in fix_generated_code.py

# ========================================================================
# STEP 9: REMOVE UNUSED DIRECTORIES
# ========================================================================
echo ""
echo "============================================"
echo "🗑️ Removing unused files and directories"
echo "============================================"

# Remove api, docs, and test directories
if [ -d "${OUTPUT_DIR}/api" ]; then
    echo "  Removing ${OUTPUT_DIR}/api directory..."
    rm -rf "${OUTPUT_DIR}/api"
    echo -e "  ${GREEN}✓ API directory removed${NC}"
fi

if [ -d "${OUTPUT_DIR}/docs" ]; then
    echo "  Removing ${OUTPUT_DIR}/docs directory..."
    rm -rf "${OUTPUT_DIR}/docs"
    echo -e "  ${GREEN}✓ Docs directory removed${NC}"
fi

if [ -d "${OUTPUT_DIR}/test" ]; then
    echo "  Removing ${OUTPUT_DIR}/test directory..."
    rm -rf "${OUTPUT_DIR}/test"
    echo -e "  ${GREEN}✓ Test directory removed${NC}"
fi

# Remove git_push.sh script
if [ -f "${OUTPUT_DIR}/git_push.sh" ]; then
    echo "  Removing ${OUTPUT_DIR}/git_push.sh file..."
    rm -f "${OUTPUT_DIR}/git_push.sh"
    echo -e "  ${GREEN}✓ git_push.sh removed${NC}"
fi

# Remove README.md file
if [ -f "${OUTPUT_DIR}/README.md" ]; then
    echo "  Removing ${OUTPUT_DIR}/README.md file..."
    rm -f "${OUTPUT_DIR}/README.md"
    echo -e "  ${GREEN}✓ README.md removed${NC}"
fi

# Remove .bak files
BAK_FILES=$(find "${OUTPUT_DIR}" -name "*.bak" -type f | wc -l)
if [ "$BAK_FILES" -gt 0 ]; then
    echo "  Removing ${BAK_FILES} .bak files..."
    find "${OUTPUT_DIR}" -name "*.bak" -type f -delete
    echo -e "  ${GREEN}✓ .bak files removed${NC}"
fi

echo ""
echo -e "${GREEN}✅ Unused files and directories removed successfully!${NC}"

# Now try to build SDK
if (cd "${PROJECT_ROOT}" && go build ./sdk/...); then
    echo -e "  ${GREEN}✓ Build OK${NC}"
else
    echo -e "  ${RED}❌ Build failed after post-processing${NC}"
    # Show a brief hint where to look
    echo "  Hint: Check ${SIG_FILE} and ${GRP_FILE} for expected fixes."
    exit 1
fi

# Optional: go fmt
if command -v go &> /dev/null; then
    (cd "${PROJECT_ROOT}" && go fmt ./... > /dev/null || true)
fi

echo "============================================"
echo ""
echo "Generated SDK location:"
echo -e "  ${GREEN}${OUTPUT_DIR}${NC}"
