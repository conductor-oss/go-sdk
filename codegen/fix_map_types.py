#!/usr/bin/env python3
"""
Script to fix map types in generated Go code.
Replaces map[string]map[string]interface{} with map[string]interface{}.
"""

import os
import re
import sys
from pathlib import Path

def fix_map_types_in_file(file_path):
    """Fix map types in a single file."""
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            content = f.read()
        
        # Keep the original content for comparison
        original_content = content
        
        # Pattern to find map[string]map[string]interface{}
        pattern = re.compile(r'map\[string\]map\[string\]interface\{\}')
        replacement = 'map[string]interface{}'
        
        # Perform the replacement
        new_content = pattern.sub(replacement, content)
        
        # Write back if content changed
        if new_content != original_content:
            with open(file_path, 'w', encoding='utf-8') as f:
                f.write(new_content)
            return True
        return False
        
    except Exception as e:
        print(f"Error processing file {file_path}: {e}")
        return False

def fix_map_types_in_directory(directory_path):
    """Fix map types across all .go files in a directory."""
    directory = Path(directory_path)
    
    if not directory.exists():
        print(f"Directory {directory_path} does not exist!")
        return
    
    print(f"Processing directory: {directory_path}")
    
    fixed_files = []
    total_files = 0
    
    # Recursively iterate over all .go files
    for go_file in directory.rglob("*.go"):
        total_files += 1
        print(f"Processing: {go_file.relative_to(directory)}")
        
        if fix_map_types_in_file(go_file):
            fixed_files.append(str(go_file.relative_to(directory)))
            print(f"  ✓ Fixed")
        else:
            print(f"  - No changes")
    
    print(f"\nResult:")
    print(f"  Total files: {total_files}")
    print(f"  Files fixed: {len(fixed_files)}")
    
    if fixed_files:
        print(f"\nFixed files:")
        for file in fixed_files:
            print(f"  - {file}")

def main():
    if len(sys.argv) != 2:
        print("Usage: python fix_map_types.py <directory_path>")
        print("Example: python fix_map_types.py ../sdk/generated/http/orkes")
        sys.exit(1)
    
    directory_path = sys.argv[1]
    fix_map_types_in_directory(directory_path)

if __name__ == "__main__":
    main()

