#!/usr/bin/env python3
import json
import sys
import re

def filter_integration_api(input_file, output_file):
    """
    Filter OpenAPI specification to include only Integration API endpoints
    """
    
    print(f"  Loading {input_file}...")
    
    # Load JSON specification
    with open(input_file, 'r') as f:
        spec = json.load(f)
    
    # Count original endpoints
    original_path_count = len(spec['paths']) if 'paths' in spec else 0
    print(f"  Original API has {original_path_count} endpoints")
    
    # Filter paths to include only Integration API endpoints
    if 'paths' in spec:
        filtered_paths = {}
        for path, methods in spec['paths'].items():
            # Keep only paths related to Integration API
            if any(pattern in path.lower() for pattern in ['/integration/', '/integrations/']):
                filtered_paths[path] = methods
        
        # Replace paths with filtered paths
        spec['paths'] = filtered_paths
    
    # Count filtered endpoints
    filtered_path_count = len(spec['paths']) if 'paths' in spec else 0
    print(f"  Filtered API has {filtered_path_count} endpoints (removed {original_path_count - filtered_path_count})")
    
    # Find all referenced schemas
    referenced_schemas = set()
    
    def collect_references(obj):
        """Recursively collect all schema references"""
        if isinstance(obj, dict):
            if '$ref' in obj and obj['$ref'].startswith('#/components/schemas/'):
                schema_name = obj['$ref'].split('/')[-1]
                referenced_schemas.add(schema_name)
            
            # Recursively process all nested objects
            for key, value in obj.items():
                collect_references(value)
        elif isinstance(obj, list):
            for item in obj:
                collect_references(item)
    
    # Collect references from paths
    collect_references(spec['paths'])
    
    # Recursively collect references from referenced schemas
    if 'components' in spec and 'schemas' in spec['components']:
        processed_schemas = set()
        schemas_to_process = referenced_schemas.copy()
        
        while schemas_to_process:
            schema_name = schemas_to_process.pop()
            processed_schemas.add(schema_name)
            
            if schema_name in spec['components']['schemas']:
                schema = spec['components']['schemas'][schema_name]
                collect_references(schema)
                
                # Add newly discovered schemas to processing queue
                new_schemas = referenced_schemas - processed_schemas
                schemas_to_process.update(new_schemas)
    
    # Filter schemas to include only referenced ones
    if 'components' in spec and 'schemas' in spec['components']:
        original_schema_count = len(spec['components']['schemas'])
        
        filtered_schemas = {}
        for schema_name in referenced_schemas:
            if schema_name in spec['components']['schemas']:
                filtered_schemas[schema_name] = spec['components']['schemas'][schema_name]
        
        # Replace schemas with filtered schemas
        spec['components']['schemas'] = filtered_schemas
        
        print(f"  Filtered schemas from {original_schema_count} to {len(filtered_schemas)}")
    
    # Save the filtered specification
    print(f"  Saving to {output_file}...")
    with open(output_file, 'w') as f:
        json.dump(spec, f, indent=2)
    
    print(f"  ✅ Successfully filtered OpenAPI spec to include only Integration API")
    return True

if __name__ == "__main__":
    try:
        input_file = sys.argv[1] if len(sys.argv) > 1 else "openapi-fixed.json"
        output_file = sys.argv[2] if len(sys.argv) > 2 else "openapi-filtered.json"
        
        success = filter_integration_api(input_file, output_file)
        sys.exit(0 if success else 1)
    except Exception as e:
        print(f"  ❌ Error: {e}")
        sys.exit(1)
