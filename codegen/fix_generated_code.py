#!/usr/bin/env python3
"""
Script to fix generated Go code.
Corrects specific data type issues in the generated sources.

Usage:
    python3 fix_generated_code.py <path_to_generated_code> [--verbose] [--dry-run] [--only fixes] [--backup]

Examples:
    python3 fix_generated_code.py ../sdk/generated/http/orkes --verbose --only users_in_group,signal_response
"""

import os
import re
import sys
import argparse
from pathlib import Path
from typing import List, Set

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


def fix_get_users_in_group(file_path, *, verbose=False, dry_run=False):
    """Fix GetUsersInGroup to return []interface{} instead of map[string]interface{} in a strictly scoped context."""
    try:
        p = Path(file_path)
        content = p.read_text(encoding='utf-8')
        original_content = content

        # Fast guard
        if 'GetUsersInGroupExecute' not in content or 'GroupResourceAPI' not in content:
            return False

        # Idempotency: skip if signature already returns an array
        if re.search(r'GetUsersInGroupExecute\(r GroupResourceAPIGetUsersInGroupRequest\) \(\[\]interface\{\}, \*http\.Response, error\)', content):
            return False

        # Fix function signatures
        content = re.sub(
            r'GetUsersInGroupExecute\(r GroupResourceAPIGetUsersInGroupRequest\) \(map\[string\]interface\{\}, \*http\.Response, error\)',
            'GetUsersInGroupExecute(r GroupResourceAPIGetUsersInGroupRequest) ([]interface{}, *http.Response, error)',
            content
        )
        content = re.sub(
            r'func \(r GroupResourceAPIGetUsersInGroupRequest\) Execute\(\) \(map\[string\]interface\{\}, \*http\.Response, error\)',
            'func (r GroupResourceAPIGetUsersInGroupRequest) Execute() ([]interface{}, *http.Response, error)',
            content
        )

        # Fix local variable type only inside Execute body
        def _fix_local_var(m: re.Match) -> str:
            block = m.group(0)
            block = block.replace('localVarReturnValue  map[string]interface{}', 'localVarReturnValue  []interface{}')
            block = block.replace('localVarReturnValue map[string]interface{}', 'localVarReturnValue []interface{}')
            return block

        content = re.sub(
            r'func \(a \*GroupResourceAPIService\) GetUsersInGroupExecute\([^)]*\) \([^)]*\) \{[\s\S]*?\n\}',
            _fix_local_var,
            content,
            flags=re.MULTILINE
        )

        # Update @return comment
        content = re.sub(
            r'//\s+@return map\[string\]interface\{\}',
            '//  @return []interface{}',
            content
        )

        if content != original_content:
            _write_file(p, content, dry_run=dry_run, backup=True, verbose=verbose)
            if verbose:
                print(f"  ✅ Fixed GetUsersInGroup in {file_path}")
            return True
        return False

    except Exception as e:
        if verbose:
            print(f"  ❌ Error processing {file_path}: {e}")
        return False


def fix_signal_response(file_path, *, verbose=False, dry_run=False):
    """Add missing fields to SignalResponse (only if they are not present)."""
    try:
        p = Path(file_path)
        content = p.read_text(encoding='utf-8')
        original_content = content

        if 'type SignalResponse struct' not in content:
            return False

        # If Tasks already exist — no changes required
        if re.search(r'\n\s*Tasks\s+\[\]interface\{\}', content):
            return False

        # 1) Try to anchor by WorkflowId (fast path)
        pattern_anchor = re.compile(
            r'(WorkflowId\s+\*string\s+`json:"workflowId,omitempty"`\s*)(\})',
            re.DOTALL
        )

        def add_fields(match):
            before_brace = match.group(1)
            closing_brace = match.group(2)
            fields_to_add = '''

	// Additional fields for full response
	Tasks      []interface{} `json:"tasks,omitempty"`
	CreatedBy  string        `json:"createdBy,omitempty"`
	CreateTime int64         `json:"createTime,omitempty"`
	Status     string        `json:"status,omitempty"`
	UpdateTime int64         `json:"updateTime,omitempty"`

	// Task-specific fields for BLOCKING_TASK strategies
	TaskType          string `json:"taskType,omitempty"`
	TaskId            string `json:"taskId,omitempty"`
	ReferenceTaskName string `json:"referenceTaskName,omitempty"`
	RetryCount        int32  `json:"retryCount,omitempty"`
	TaskDefName       string `json:"taskDefName,omitempty"`
	WorkflowType      string `json:"workflowType,omitempty"`
'''
            return before_brace + fields_to_add + closing_brace

        new_content = pattern_anchor.sub(add_fields, content)

        # 2) If anchor replacement did not work — parse the entire struct block and insert before the closing brace
        if new_content == content:
            struct_pattern = re.compile(
                r'(type\s+SignalResponse\s+struct\s*\{)([\s\S]*?)(\})',
                re.MULTILINE
            )

            def insert_into_struct(m: re.Match) -> str:
                start = m.group(1)
                body = m.group(2)
                end = m.group(3)
                # If Tasks already exist — do nothing
                if re.search(r'\n\s*Tasks\s+\[\]interface\{\}', body):
                    return m.group(0)
                fields = '''

	// Additional fields for full response
	Tasks      []interface{} `json:"tasks,omitempty"`
	CreatedBy  string        `json:"createdBy,omitempty"`
	CreateTime int64         `json:"createTime,omitempty"`
	Status     string        `json:"status,omitempty"`
	UpdateTime int64         `json:"updateTime,omitempty"`

	// Task-specific fields for BLOCKING_TASK strategies
	TaskType          string `json:"taskType,omitempty"`
	TaskId            string `json:"taskId,omitempty"`
	ReferenceTaskName string `json:"referenceTaskName,omitempty"`
	RetryCount        int32  `json:"retryCount,omitempty"`
	TaskDefName       string `json:"taskDefName,omitempty"`
	WorkflowType      string `json:"workflowType,omitempty"`
'''
                return start + body + fields + end

            new_content = struct_pattern.sub(insert_into_struct, content)

        if new_content != original_content:
            _write_file(p, new_content, dry_run=dry_run, backup=True, verbose=verbose)
            if verbose:
                print(f"  ✅ Fixed SignalResponse in {file_path}")
            return True
        return False

    except Exception as e:
        if verbose:
            print(f"  ❌ Error processing {file_path}: {e}")
        return False


def fix_secret_primitive_response(file_path, *, verbose=False, dry_run=False):
    """In secret API, change map[string]interface{} responses to interface{} to support bool/string primitives."""
    try:
        p = Path(file_path)
        if p.name != 'api_secret_resource.go':
            return False
        content = p.read_text(encoding='utf-8')
        original = content

        # Replace request wrapper Execute() signatures to use interface{} instead of map[string]interface{}
        def _req_exec_sig(m: re.Match) -> str:
            s = m.group(0)
            return s.replace('map[string]interface{}', 'interface{}')
        content = re.sub(
            r'func \(r SecretResourceAPI(?:DeleteSecret|PutSecret|SecretExists)Request\) Execute\(\) \(map\[string\]interface\{\}, \*http\.Response, error\)',
            _req_exec_sig,
            content
        )

        # Service Execute signatures for specific methods
        def _svc_exec_sig(m: re.Match) -> str:
            s = m.group(0)
            return s.replace('map[string]interface{}', 'interface{}')
        content = re.sub(
            r'func \(a \*SecretResourceAPIService\) (?:DeleteSecret|PutSecret|SecretExists)Execute\([^)]*\) \(map\[string\]interface\{\}, \*http\.Response, error\)',
            _svc_exec_sig,
            content
        )

        # Patch interface method signatures and @return comments
        content = re.sub(
            r'(//\s+@return)\s+map\[string\]interface\{\}',
            r'\1 interface{}',
            content
        )
        content = re.sub(
            r'(DeleteSecretExecute\(r SecretResourceAPIDeleteSecretRequest\) )\(map\[string\]interface\{\}, \*http\.Response, error\)',
            r'\1(interface{}, *http.Response, error)',
            content
        )
        content = re.sub(
            r'(PutSecretExecute\(r SecretResourceAPIPutSecretRequest\) )\(map\[string\]interface\{\}, \*http\.Response, error\)',
            r'\1(interface{}, *http.Response, error)',
            content
        )
        content = re.sub(
            r'(SecretExistsExecute\(r SecretResourceAPISecretExistsRequest\) )\(map\[string\]interface\{\}, \*http\.Response, error\)',
            r'\1(interface{}, *http.Response, error)',
            content
        )

        # Change localVarReturnValue types inside those functions
        def _patch_block_local_var(match: re.Match) -> str:
            block = match.group(0)
            block = block.replace('localVarReturnValue  map[string]interface{}', 'localVarReturnValue  interface{}')
            block = block.replace('localVarReturnValue map[string]interface{}', 'localVarReturnValue interface{}')
            return block
        content = re.sub(
            r'func \(a \*SecretResourceAPIService\) (?:DeleteSecret|PutSecret|SecretExists)Execute\([^)]*\) \([^)]*\) \{[\s\S]*?\n\}',
            _patch_block_local_var,
            content,
            flags=re.MULTILINE
        )

        if content != original:
            _write_file(p, content, dry_run=dry_run, backup=True, verbose=verbose)
            if verbose:
                print(f"  ✅ Fixed primitive response in {file_path}")
            return True
        return False
    except Exception as e:
        if verbose:
            print(f"  ❌ Error processing {file_path}: {e}")
        return False


def fix_get_access_keys(file_path, *, verbose=False, dry_run=False):
    """Fix ApplicationResourceAPI.GetAccessKeys to return []interface{} instead of map[string]interface{}.

    Applies only to generated files that contain ApplicationResourceAPI and GetAccessKeys.
    """
    try:
        p = Path(file_path)
        content = p.read_text(encoding='utf-8')
        original_content = content

        # Fast guard
        if 'ApplicationResourceAPI' not in content or 'GetAccessKeysExecute' not in content:
            return False

        # Idempotency: we still proceed to ensure local var type is corrected even if signature was already fixed

        # Update interface method signature comment and type (scoped later by signature changes)

        # Update interface method signature (Execute on service)
        content = re.sub(
            r'(GetAccessKeysExecute\(r ApplicationResourceAPIGetAccessKeysRequest\) )\(map\[string\]interface\{\}, \*http\.Response, error\)',
            r'\1([]interface{}, *http.Response, error)',
            content
        )

        # Update request wrapper Execute()
        content = re.sub(
            r'func \(r ApplicationResourceAPIGetAccessKeysRequest\) Execute\(\) \(map\[string\]interface\{\}, \*http\.Response, error\)',
            r'func (r ApplicationResourceAPIGetAccessKeysRequest) Execute() ([]interface{}, *http.Response, error)',
            content
        )

        # Update service method signature and local return variable type inside the function body
        def _fix_service_block(m: re.Match) -> str:
            block = m.group(0)
            block = block.replace('localVarReturnValue map[string]interface{}', 'localVarReturnValue []interface{}')
            block = block.replace('localVarReturnValue  map[string]interface{}', 'localVarReturnValue  []interface{}')
            return block

        content = re.sub(
            r'func \(a \*ApplicationResourceAPIService\) GetAccessKeysExecute\([^)]*\) \([^)]*\) \{[\s\S]*?\n\}',
            _fix_service_block,
            content,
            flags=re.MULTILINE
        )

        # Do NOT apply global fallbacks; keep changes limited to GetAccessKeys only

        if content != original_content:
            _write_file(p, content, dry_run=dry_run, backup=True, verbose=verbose)
            if verbose:
                print(f"  ✅ Fixed GetAccessKeys in {file_path}")
            return True
        return False

    except Exception as e:
        if verbose:
            print(f"  ❌ Error processing {file_path}: {e}")
        return False


def fix_workflowdef_required_name(file_path, *, verbose=False, dry_run=False):
    """Remove 'name' from requiredProperties in WorkflowDef.UnmarshalJSON and ExtendedWorkflowDef.UnmarshalJSON."""
    try:
        p = Path(file_path)
        if p.name not in ('model_workflow_def.go', 'model_extended_workflow_def.go'):
            return False
        content = p.read_text(encoding='utf-8')
        original = content

        # Pattern to find requiredProperties slice and drop "name"
        def _drop_name_required(m: re.Match) -> str:
            block = m.group(0)
            # remove exact entry "\t\t\"name\",\n"
            block = block.replace('\n\t\t"name",\n', '\n')
            block = block.replace('\n\t\t"name"\n', '\n')
            block = block.replace('\t\t"name",\n', '')
            block = block.replace('\t\t"name"\n', '')
            return block

        content = re.sub(
            r'func \(o \*WorkflowDef\) UnmarshalJSON\([\s\S]*?requiredProperties := \[\]string\{[\s\S]*?\}[\s\S]*?for _, requiredProperty := range\(requiredProperties\)[\s\S]*?\}',
            _drop_name_required,
            content,
            flags=re.MULTILINE
        )

        content = re.sub(
            r'func \(o \*ExtendedWorkflowDef\) UnmarshalJSON\([\s\S]*?requiredProperties := \[\]string\{[\s\S]*?\}[\s\S]*?for _, requiredProperty := range\(requiredProperties\)[\s\S]*?\}',
            _drop_name_required,
            content,
            flags=re.MULTILINE
        )

        if content != original:
            _write_file(p, content, dry_run=dry_run, backup=True, verbose=verbose)
            if verbose:
                print(f"  ✅ Removed required 'name' in {file_path}")
            return True
        return False
    except Exception as e:
        if verbose:
            print(f"  ❌ Error processing {file_path}: {e}")
        return False


def fix_schema_default_value(file_path, *, verbose=False, dry_run=False):
    """Make Schema.defaultValue interface{} instead of map[string]interface{} and adjust accessors, scanning any model file that defines Schema."""
    try:
        p = Path(file_path)
        content = p.read_text(encoding='utf-8')
        original = content

        # Apply only if the file has defaultValue json tag or the DefaultValue field
        if 'defaultValue' not in content and 'DefaultValue' not in content:
            return False

        changed = False
        # Replace field type on Schema struct definitions
        new_content, n1 = re.subn(
            r'(DefaultValue\s+)map\[string\]interface\{\}(\s+`json:"defaultValue,?omitempty"`)?',
            r'\1interface{}\2',
            content
        )
        if n1:
            changed = True
        content = new_content

        # Replace getter signatures if present
        new_content, n2 = re.subn(
            r'func \(o \*Schema\) GetDefaultValue\(\) map\[string\]interface\{\}',
            r'func (o *Schema) GetDefaultValue() interface{}',
            content
        )
        if n2:
            changed = True
        content = new_content

        new_content, n3 = re.subn(
            r'func \(o \*Schema\) GetDefaultValueOk\(\) \(map\[string\]interface\{\}, bool\)',
            r'func (o *Schema) GetDefaultValueOk() (interface{}, bool)',
            content
        )
        if n3:
            changed = True
        content = new_content

        new_content, n4 = re.subn(
            r'func \(o \*Schema\) SetDefaultValue\(v map\[string\]interface\{\}\)',
            r'func (o *Schema) SetDefaultValue(v interface{})',
            content
        )
        if n4:
            changed = True
        content = new_content

        # Replace zero-value ret types inside getters
        content = re.sub(r'var ret map\[string\]interface\{\}', 'var ret interface{}', content)
        content = re.sub(r'return map\[string\]interface\{\}\{\}, false', 'return nil, false', content)

        if changed and content != original:
            _write_file(p, content, dry_run=dry_run, backup=True, verbose=verbose)
            if verbose:
                print(f"  ✅ Fixed Schema.defaultValue in {file_path}")
            return True
        return False
    except Exception as e:
        if verbose:
            print(f"  ❌ Error processing {file_path}: {e}")
        return False


def fix_bulk_response_success_list(file_path, *, verbose=False, dry_run=False):
    """Change BulkResponse.BulkSuccessfulResults to []string instead of []map[string]interface{} for orkes clients."""
    try:
        p = Path(file_path)
        if p.name != 'model_bulk_response.go':
            return False
        content = p.read_text(encoding='utf-8')
        original = content

        if 'type BulkResponse struct' not in content:
            return False

        # Field type
        content = re.sub(
            r'(BulkSuccessfulResults\s+)\[\]map\[string\]interface\{\}',
            r'\1[]string',
            content
        )
        # Getter return types
        content = re.sub(
            r'GetBulkSuccessfulResults\(\) \[\]map\[string\]interface\{\}',
            r'GetBulkSuccessfulResults() []string',
            content
        )
        content = re.sub(
            r'var ret \[\]map\[string\]interface\{\}',
            r'var ret []string',
            content
        )
        content = re.sub(
            r'GetBulkSuccessfulResultsOk\(\) \(\[\]map\[string\]interface\{\}, bool\)',
            r'GetBulkSuccessfulResultsOk() ([]string, bool)',
            content
        )
        content = re.sub(
            r'return nil, false',
            r'return nil, false',
            content
        )
        # Setter signature
        content = re.sub(
            r'SetBulkSuccessfulResults\(v \[\]map\[string\]interface\{\}\)',
            r'SetBulkSuccessfulResults(v []string)',
            content
        )

        if content != original:
            _write_file(p, content, dry_run=dry_run, backup=True, verbose=verbose)
            if verbose:
                print(f"  ✅ Fixed BulkResponse in {file_path}")
            return True
        return False
    except Exception as e:
        if verbose:
            print(f"  ❌ Error processing {file_path}: {e}")
        return False

def fix_remove_unused_fmt_import(file_path, *, verbose=False, dry_run=False):
    """Remove unused import \"fmt\" from Go files when no fmt. references remain."""
    try:
        p = Path(file_path)
        if p.suffix != '.go':
            return False
        content = p.read_text(encoding='utf-8')
        if '"fmt"' not in content:
            return False
        # If fmt is referenced, keep
        if re.search(r'\bfmt\.', content):
            return False
        # Remove fmt from single-line import or multi-line block
        new_content = content
        # Case 1: standalone import line
        new_content = re.sub(r'(?m)^\s*"fmt"\s*\n', '', new_content)
        # Case 2: inside block, handle trailing commas and empty lines
        new_content = re.sub(r'(?m)^\s*"fmt"\s*,?\s*\n', '', new_content)
        # Clean potential empty import blocks like: import ()
        new_content = re.sub(r'(?m)^import\s*\(\s*\)\s*\n', '', new_content)
        if new_content != content:
            _write_file(p, new_content, dry_run=dry_run, backup=True, verbose=verbose)
            if verbose:
                print(f"  ✓ Removed unused fmt in {file_path}")
            return True
        return False
    except Exception as e:
        if verbose:
            print(f"  ❌ Error processing {file_path}: {e}")
        return False

def fix_generated_code_in_file(file_path, *, verbose=False, dry_run=False, only: Set[str] = None):
    """Исправляет сгенерированный код в одном файле."""
    fixes_applied = 0

    run_users = (only is None) or ('users_in_group' in only)
    run_signal = (only is None) or ('signal_response' in only)
    run_secret = (only is None) or ('secret_primitive' in only)
    run_schema = (only is None) or ('schema_default' in only)
    run_bulk = (only is None) or ('bulk_response' in only)
    run_wfdef_req = (only is None) or ('workflowdef_required' in only)
    run_access_keys = (only is None) or ('get_access_keys' in only)
    # Allow relaxing required checks for specific models
    run_relax_required = (only is None) or ('relax_required' in only)
    # Global relax: replace any UnmarshalJSON that contains requiredProperties preamble
    run_relax_required_global = (only is None) or ('relax_required_global' in only)
    # Remove unused fmt imports
    run_fmt = (only is None) or ('rm_unused_fmt' in only)

    if run_users and fix_get_users_in_group(file_path, verbose=verbose, dry_run=dry_run):
        fixes_applied += 1

    if run_signal and fix_signal_response(file_path, verbose=verbose, dry_run=dry_run):
        fixes_applied += 1

    if run_secret and fix_secret_primitive_response(file_path, verbose=verbose, dry_run=dry_run):
        fixes_applied += 1

    if run_schema and fix_schema_default_value(file_path, verbose=verbose, dry_run=dry_run):
        fixes_applied += 1

    # Fix BulkResponse successful results type
    if run_bulk and fix_bulk_response_success_list(file_path, verbose=verbose, dry_run=dry_run):
        fixes_applied += 1

    # Make WorkflowDef.UnmarshalJSON not require "name" (it may be missing in history entries)
    if run_wfdef_req and fix_workflowdef_required_name(file_path, verbose=verbose, dry_run=dry_run):
        fixes_applied += 1

    # Fix GetAccessKeys array response type
    if run_access_keys and fix_get_access_keys(file_path, verbose=verbose, dry_run=dry_run):
        fixes_applied += 1

    # Remove requiredProperties validation blocks from UnmarshalJSON for selected models
    if run_relax_required and fix_remove_required_properties_check(file_path, verbose=verbose, dry_run=dry_run):
        fixes_applied += 1

    if run_relax_required_global and fix_remove_required_properties_check_global(file_path, verbose=verbose, dry_run=dry_run):
        fixes_applied += 1

    if run_fmt and fix_remove_unused_fmt_import(file_path, verbose=verbose, dry_run=dry_run):
        fixes_applied += 1

    return fixes_applied


def fix_remove_required_properties_check(file_path, *, verbose=False, dry_run=False):
    """Remove requiredProperties validation inside UnmarshalJSON for specific models only.

    Targeted models:
      - WorkflowDef (model_workflow_def.go)
      - ExtendedWorkflowDef (model_extended_workflow_def.go)
      - StartWorkflowRequest (model_start_workflow_request.go)

    We carefully remove only the validation preamble that checks requiredProperties and leave
    the actual json.Unmarshal into the typed alias intact.
    """
    try:
        p = Path(file_path)
        fname = p.name
        if fname not in (
            'model_workflow_def.go',
            'model_extended_workflow_def.go',
            'model_start_workflow_request.go',
        ):
            return False

        content = p.read_text(encoding='utf-8')
        original = content

        # Decide alias type by file name
        if fname == 'model_workflow_def.go':
            struct_name = 'WorkflowDef'
            alias_name = '_WorkflowDef'
        elif fname == 'model_extended_workflow_def.go':
            struct_name = 'ExtendedWorkflowDef'
            alias_name = '_ExtendedWorkflowDef'
        else:  # model_start_workflow_request.go
            struct_name = 'StartWorkflowRequest'
            alias_name = '_StartWorkflowRequest'

        # Replace entire UnmarshalJSON function with a minimal implementation (no required checks)
        func_pattern = re.compile(
            rf"func \\(o \\*{struct_name}\\) UnmarshalJSON\\(bytes \\[\\]byte\\) \\(err error\\) \\{{[\\s\\S]*?\\}}\n",
            re.MULTILINE
        )

        replacement = (
            f"func (o *{struct_name}) UnmarshalJSON(bytes []byte) (err error) {{\n"
            f"\tvarObj := {alias_name}{{}}\n\n"
            f"\terr = json.Unmarshal(bytes, &varObj)\n\n"
            f"\tif err != nil {{\n"
            f"\t\treturn err\n"
            f"\t}}\n\n"
            f"\t*o = {struct_name}(varObj)\n\n"
            f"\treturn err\n"
            f"}}\n"
        )

        new_content, n = func_pattern.subn(replacement, content)

        if n > 0 and new_content != original:
            _write_file(p, new_content, dry_run=dry_run, backup=True, verbose=verbose)
            if verbose:
                print(f"  ✅ Replaced UnmarshalJSON (no required checks) in {file_path}")
            return True

        return False

    except Exception as e:
        if verbose:
            print(f"  ❌ Error processing {file_path}: {e}")
        return False


def fix_remove_required_properties_check_global(file_path, *, verbose=False, dry_run=False):
    """Globally replace any UnmarshalJSON function that contains requiredProperties validation
    with a minimal implementation that just unmarshals into the alias type and assigns back.

    This is safe because openapi-generator always uses the pattern:
      type _X X
      func (o *X) UnmarshalJSON(bytes []byte) (err error) {
          // requiredProperties preamble...
          varX := _X{}
          err = json.Unmarshal(bytes, &varX)
          ...
          *o = X(varX)
          return err
      }
    We detect functions that include 'requiredProperties :=' and reconstruct them.
    """
    try:
        p = Path(file_path)
        if p.suffix != '.go':
            return False

        content = p.read_text(encoding='utf-8')
        original = content

        # Find each UnmarshalJSON function block
        func_re = re.compile(
            r'(func\s*\(o\s*\*([A-Za-z_][\w]*)\)\s*UnmarshalJSON\(bytes\s*\[\]\s*byte\)\s*\(err\s*error\)\s*\{)([\s\S]*?)(\n\})',
            re.MULTILINE
        )

        changed = False

        def repl(m: re.Match) -> str:
            nonlocal changed
            header = m.group(1)
            body = m.group(3)
            footer = m.group(4)
            struct_name = m.group(2)
            if 'requiredProperties :=' not in body:
                return m.group(0)
            # Attempt to detect alias name used later (e.g., _WorkflowDef)
            alias_match = re.search(r'var\s*[A-Za-z_][\w]*\s*:=\s*(_[A-Za-z_][\w]*)\{\}', body)
            if alias_match:
                alias_name = alias_match.group(1)
            else:
                # Fallback alias based on struct name
                alias_name = f'_{struct_name}'
            minimal = (
                f"{header}\n"
                f"\tvarObj := {alias_name}{{}}\n\n"
                f"\terr = json.Unmarshal(bytes, &varObj)\n\n"
                f"\tif err != nil {{\n\t\treturn err\n\t}}\n\n"
                f"\t*o = {struct_name}(varObj)\n\n"
                f"\treturn err\n"
                f"{footer}"
            )
            changed = True
            return minimal

        new_content = func_re.sub(repl, content)

        if changed and new_content != original:
            _write_file(p, new_content, dry_run=dry_run, backup=True, verbose=verbose)
            if verbose:
                print(f"  ✅ Globally relaxed required checks in {file_path}")
            return True
        return False

    except Exception as e:
        if verbose:
            print(f"  ❌ Error processing {file_path}: {e}")
        return False


def fix_generated_code_in_directory(directory_path, verbose=False, dry_run=False, only: Set[str] = None):
    """Fix generated code across all .go files in a directory."""
    directory = Path(directory_path)

    if not directory.exists():
        print(f"❌ Directory {directory_path} does not exist!")
        return False

    if verbose:
        print(f"🔍 Processing directory: {directory_path}")

    fixed_files = []
    total_files = 0
    total_fixes = 0

    for go_file in directory.rglob("*.go"):
        total_files += 1
        if verbose:
            print(f"Processing: {go_file.relative_to(directory)}")

        fixes = fix_generated_code_in_file(str(go_file), verbose=verbose, dry_run=dry_run, only=only)
        if fixes > 0:
            fixed_files.append(str(go_file.relative_to(directory)))
            total_fixes += fixes
            if verbose:
                print(f"  ✅ Fixes applied: {fixes}")
        elif verbose:
            print(f"  - No changes")

    print(f"\n📊 Result:")
    print(f"  Total files: {total_files}")
    print(f"  Files modified: {len(fixed_files)}")
    print(f"  Total fixes: {total_fixes}")

    if fixed_files and verbose:
        print(f"\nModified files:")
        for file in fixed_files:
            print(f"  - {file}")

    return len(fixed_files) > 0


def parse_only(value: str) -> Set[str]:
    if not value:
        return set()
    return {v.strip() for v in value.split(',') if v.strip()}


def main():
    parser = argparse.ArgumentParser(description='Fixes issues in generated Go code')
    parser.add_argument('directory', help='Path to the directory with generated code')
    parser.add_argument('--verbose', '-v', action='store_true', help='Verbose output')
    parser.add_argument('--dry-run', action='store_true', help='Show changes without writing')
    parser.add_argument('--only', type=str, default='', help='Comma-separated list of fixes: users_in_group,signal_response')
    parser.add_argument('--backup', action='store_true', help='Create .bak backups for modified files')

    args = parser.parse_args()

    if not Path(args.directory).exists():
        print(f"❌ Directory {args.directory} does not exist!")
        sys.exit(1)

    only = parse_only(args.only)
    # If --only is omitted or empty, run all fixes
    if not only:
        only = None
    success = fix_generated_code_in_directory(args.directory, verbose=args.verbose, dry_run=args.dry_run, only=only)

    if success:
        print("✅ Fixes applied successfully!" if not args.dry_run else "📝 DRY-RUN completed")
        sys.exit(0)
    else:
        print("ℹ️  No fixes were necessary")
        sys.exit(0)

if __name__ == "__main__":
    main()
 