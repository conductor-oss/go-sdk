import json
import sys
import re

def fix_openapi_spec(input_file, output_file):
    """Fixes known issues in OpenAPI specification"""
    
    print(f"  Loading {input_file}...")
    
    # Load JSON specification
    with open(input_file, 'r') as f:
        spec = json.load(f)
    
    fixes_applied = 0
    
    # Fix 0: Remove /api prefix from all paths (optional, may not exist in all domains)
    if 'paths' in spec:
        print("  Removing /api prefix from paths...")
        new_paths = {}
        for path, methods in spec['paths'].items():
            # Remove /api prefix if it exists
            new_path = path[4:] if path.startswith('/api') else path
            if new_path != path:
                print(f"    Changed: {path} -> {new_path}")
                fixes_applied += 1
            new_paths[new_path] = methods
        spec['paths'] = new_paths
    
    # Fix 1: Path parameters
    if 'paths' in spec:
        print("  Fixing path parameters...")
        for path, methods in spec['paths'].items():
            # Extract parameters from path
            path_params = re.findall(r'\{([^}]+)\}', path)
            
            for method in ['get', 'post', 'put', 'delete', 'patch', 'head', 'options']:
                if method not in methods:
                    continue
                
                operation = methods[method]
                if 'parameters' not in operation:
                    operation['parameters'] = []
                
                # Fix existing parameters
                for param in operation['parameters']:
                    if param.get('name') in path_params and param.get('in') != 'path':
                        print(f"    Fixed parameter {param['name']} in {method.upper()} {path}")
                        param['in'] = 'path'
                        param['required'] = True
                        fixes_applied += 1
                
                # Add missing parameters
                existing_params = {p['name'] for p in operation['parameters'] if 'name' in p}
                for param_name in path_params:
                    if param_name not in existing_params:
                        print(f"    Added parameter {param_name} in {method.upper()} {path}")
                        operation['parameters'].append({
                            "name": param_name,
                            "in": "path",
                            "required": True,
                            "schema": {"type": "string"},
                            "description": f"Path parameter {param_name}"
                        })
                        fixes_applied += 1

    # Fix 1b: Normalize JumpToTask endpoint (with or without /api prefix)
    if 'paths' in spec:
        jump_candidates = [
            p for p in list(spec['paths'].keys())
            if re.match(r'^/(api/)?workflow/\{workflowId\}/jump', p)
        ]
        for jp in jump_candidates:
            methods = spec['paths'].get(jp, {})
            post = methods.get('post') or methods.get('Post')
            if not isinstance(post, dict):
                continue

            # Ensure path ends with literal {taskReferenceName}
            if '{taskReferenceName}' not in jp:
                new_path = jp.rstrip('/') + '/{taskReferenceName}'
                if new_path not in spec['paths']:
                    spec['paths'][new_path] = methods
                if jp in spec['paths']:
                    del spec['paths'][jp]
                jp = new_path
                methods = spec['paths'][jp]
                post = methods.get('post') or methods.get('Post') or post
                fixes_applied += 1

            # Ensure taskReferenceName is a required QUERY parameter
            params = post.get('parameters', []) or []
            # Remove/convert any existing path param definition for taskReferenceName
            for p in params:
                if p.get('name') == 'taskReferenceName':
                    p['in'] = 'query'
                    p['required'] = True
                    p.setdefault('schema', {'type': 'string'})
            # If not present, append as query param
            if not any(p.get('name') == 'taskReferenceName' for p in params):
                params.append({
                    'name': 'taskReferenceName',
                    'in': 'query',
                    'required': True,
                    'schema': {'type': 'string'},
                    'description': 'Target task reference name'
                })
            post['parameters'] = params

            # Ensure requestBody is required application/json with free-form object
            rb = post.get('requestBody')
            if not isinstance(rb, dict):
                post['requestBody'] = {
                    'required': True,
                    'content': {
                        'application/json': {
                            'schema': {
                                'type': 'object',
                                'additionalProperties': True
                            }
                        }
                    }
                }
                fixes_applied += 1
            else:
                rb['required'] = True
                rb.setdefault('content', {}).setdefault('application/json', {}).setdefault(
                    'schema', {'type': 'object', 'additionalProperties': True}
                )
    
    # Fix 2: Remove problematic schemas with circular references
    if 'components' in spec and 'schemas' in spec['components']:
        print("  Processing circular references...")
        problematic_schemas = [
            'Descriptor', 'DescriptorProto', 'EnumDescriptor', 'EnumDescriptorProto',
            'FieldDescriptor', 'FieldDescriptorProto', 'FileDescriptor', 'FileDescriptorProto',
            'MethodDescriptor', 'MethodDescriptorProto', 'OneofDescriptor', 'OneofDescriptorProto',
            'ServiceDescriptor', 'ServiceDescriptorProto', 'Message', 'MessageLite',
            'Parser', 'ParserAny', 'ParserDeclaration', 'ParserDescriptorProto',
            'UninterpretedOption', 'UninterpretedOptionOrBuilder', 'NamePartOrBuilder'
        ]
        
        for schema_name in problematic_schemas:
            if schema_name in spec['components']['schemas']:
                # Replace with simple object instead of deleting
                print(f"    Simplified schema {schema_name}")
                spec['components']['schemas'][schema_name] = {
                    "type": "object",
                    "description": f"Simplified schema for {schema_name} (original had circular references)",
                    "additionalProperties": True
                }
                fixes_applied += 1
    
    # Fix 3: Incorrect properties in array types
    def fix_array_properties(obj, path=""):
        """Recursively fixes properties in array types"""
        nonlocal fixes_applied
        
        if isinstance(obj, dict):
            if obj.get('type') == 'array' and 'properties' in obj:
                print(f"    Fixed array with properties at {path}")
                # Remove properties from array
                del obj['properties']
                fixes_applied += 1
            
            # Recursively process all nested objects
            for key, value in obj.items():
                fix_array_properties(value, f"{path}/{key}")
        elif isinstance(obj, list):
            for i, item in enumerate(obj):
                fix_array_properties(item, f"{path}[{i}]")
    
    print("  Fixing incorrect array schemas...")
    fix_array_properties(spec)
    
    # Fix 4: Add missing required fields
    if 'components' in spec and 'schemas' in spec['components']:
        print("  Checking required fields...")
        for schema_name, schema in spec['components']['schemas'].items():
            if 'properties' in schema and 'required' not in schema:
                # Add empty required array if it doesn't exist
                schema['required'] = []
                fixes_applied += 1
    
    # Fix 5: Handle empty responses
    if 'paths' in spec:
        print("  Fixing empty responses...")
        for path, methods in spec['paths'].items():
            for method in methods.values():
                if isinstance(method, dict) and 'responses' in method:
                    if not method['responses']:
                        method['responses'] = {
                            "200": {
                                "description": "Success"
                            }
                        }
                        fixes_applied += 1
    
    # Save the fixed specification
    print(f"  Saving to {output_file}...")
    with open(output_file, 'w') as f:
        json.dump(spec, f, indent=2)
    
    print(f"  ✅ Applied fixes: {fixes_applied}")
    return fixes_applied

if __name__ == "__main__":
    try:
        input_file = sys.argv[1] if len(sys.argv) > 1 else "../codegen/api-docs.json"
        output_file = sys.argv[2] if len(sys.argv) > 2 else "openapi-fixed.json"
        
        fixes = fix_openapi_spec(input_file, output_file)
        sys.exit(0 if fixes >= 0 else 1)
    except Exception as e:
        print(f"  ❌ Error: {e}")
        sys.exit(1)
