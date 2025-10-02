#!/usr/bin/env python3
"""
Script to add omitempty tag to all JSON fields in generated Go models.
This ensures all fields are omitted when empty, regardless of whether they are required or not.

Usage:
    python3 add_omitempty_to_all_fields.py <path_to_generated_code> [--verbose] [--dry-run] [--backup]

Examples:
    python3 add_omitempty_to_all_fields.py ../sdk/generated/http/orkes --verbose
"""

import os
import re
import sys
import argparse
from pathlib import Path

def _write_file(path: Path, content: str, *, dry_run: bool, backup: bool, verbose: bool):
    if dry_run:
        if verbose:
            print(f"  [DRY-RUN] Would write: {path}")
        return
    if backup:
        bak = path.with_suffix(path.suffix + ".bak")
        try:
            bak.write_text(path.read_text(encoding='utf-8'), encoding='utf-8')
        except Exception:
            pass
    path.write_text(content, encoding='utf-8')

def add_omitempty_to_json_tags(file_path, *, verbose=False, dry_run=False, backup=False):
    """Add omitempty tag to all JSON fields in a Go file."""
    try:
        p = Path(file_path)
        if p.suffix != '.go' or not p.name.startswith('model_'):
            return False

        content = p.read_text(encoding='utf-8')
        original_content = content

        # Pattern to find json tags without omitempty
        pattern = re.compile(r'`json:"([^,"]+)"')
        replacement = r'`json:"\1,omitempty"'
        
        # Perform the replacement
        new_content = pattern.sub(replacement, content)
        
        # Also handle cases where there might be other options after the field name
        pattern2 = re.compile(r'`json:"([^"]+)"')
        
        def add_omitempty_if_missing(match):
            tag_content = match.group(1)
            if "omitempty" not in tag_content:
                if "," in tag_content:
                    # Already has options, add omitempty
                    return f'`json:"{tag_content},omitempty"'
                else:
                    # No options yet, add omitempty
                    return f'`json:"{tag_content},omitempty"'
            return match.group(0)
        
        new_content = pattern2.sub(add_omitempty_if_missing, new_content)
        
        # Write back if content changed
        if new_content != original_content:
            _write_file(p, new_content, dry_run=dry_run, backup=backup, verbose=verbose)
            if verbose:
                print(f"  ✅ Added omitempty to JSON tags in {file_path}")
            return True
        return False
        
    except Exception as e:
        if verbose:
            print(f"  ❌ Error processing {file_path}: {e}")
        return False

def process_directory(directory_path, *, verbose=False, dry_run=False, backup=False):
    """Add omitempty to JSON tags across all model files in a directory."""
    directory = Path(directory_path)
    
    if not directory.exists():
        print(f"❌ Directory {directory_path} does not exist!")
        return False
    
    if verbose:
        print(f"🔍 Processing directory: {directory_path}")
    
    fixed_files = []
    total_files = 0
    
    # Find all model_*.go files
    for go_file in directory.rglob("model_*.go"):
        total_files += 1
        if verbose:
            print(f"Processing: {go_file.relative_to(directory)}")
        
        if add_omitempty_to_json_tags(str(go_file), verbose=verbose, dry_run=dry_run, backup=backup):
            fixed_files.append(str(go_file.relative_to(directory)))
            if verbose:
                print(f"  ✅ Fixed")
        elif verbose:
            print(f"  - No changes needed")
    
    print(f"\n📊 Result:")
    print(f"  Total model files: {total_files}")
    print(f"  Files modified: {len(fixed_files)}")
    
    if fixed_files and verbose:
        print(f"\nModified files:")
        for file in fixed_files:
            print(f"  - {file}")
    
    return len(fixed_files) > 0

def main():
    parser = argparse.ArgumentParser(description='Add omitempty tag to all JSON fields in generated Go models')
    parser.add_argument('directory', help='Path to the directory with generated code')
    parser.add_argument('--verbose', '-v', action='store_true', help='Verbose output')
    parser.add_argument('--dry-run', action='store_true', help='Show changes without writing')
    parser.add_argument('--backup', action='store_true', help='Create .bak backups for modified files')
    
    args = parser.parse_args()
    
    if not Path(args.directory).exists():
        print(f"❌ Directory {args.directory} does not exist!")
        sys.exit(1)
    
    success = process_directory(args.directory, verbose=args.verbose, dry_run=args.dry_run, backup=args.backup)
    
    if success:
        print("✅ omitempty tags added successfully!" if not args.dry_run else "📝 DRY-RUN completed")
        sys.exit(0)
    else:
        print("ℹ️  No changes were necessary")
        sys.exit(0)

if __name__ == "__main__":
    main()
