#!/usr/bin/env python3
"""
Comprehensive fix for receiver/types.go - handles all missing type definitions
Run from the onvif project root directory
"""

import re
import shutil
from pathlib import Path

PROJECT_ROOT = Path("/Users/nk/Documents/workspace/OpenSource/onvif")
TYPES_FILE = PROJECT_ROOT / "receiver" / "types.go"

def main():
    print("🔧 Comprehensive fix for receiver/types.go")
    print("=" * 60)
    print()
    
    # Backup original
    backup_file = TYPES_FILE.with_suffix('.go.backup')
    print("📦 Creating backup...")
    shutil.copy2(TYPES_FILE, backup_file)
    print(f"   ✅ Backup saved to {backup_file}")
    print()
    
    # Read the file
    with open(TYPES_FILE, 'r') as f:
        content = f.read()
    
    # Fix 1: Add ALL missing type definitions after "var _ xml.Name"
    print("🔧 Fix 1: Adding all missing type definitions...")
    
    type_definitions = '''
// QName represents an XML qualified name
type QName string

// AnySimpleType represents any simple XML type
type AnySimpleType string

// Duration represents an XML duration type
type Duration string

// NonNegativeInteger represents a non-negative integer
type NonNegativeInteger int
'''
    
    # Find where to insert (after "var _ xml.Name")
    pattern = r'(var _ xml\.Name)'
    replacement = r'\1' + type_definitions
    content = re.sub(pattern, replacement, content, count=1)
    print("   ✅ Added QName, AnySimpleType, Duration, NonNegativeInteger")
    print()
    
    # Fix 2: Remove duplicate QueryExpressionType (keep first, remove second)
    print("🔧 Fix 2: Removing duplicate QueryExpressionType...")
    
    # Find all occurrences of QueryExpressionType struct
    pattern = r'type QueryExpressionType struct\s*{[^}]*}'
    matches = list(re.finditer(pattern, content, re.DOTALL))
    
    if len(matches) > 1:
        # Remove the second occurrence
        second_match = matches[1]
        content = content[:second_match.start()] + content[second_match.end():]
        print(f"   ✅ Removed duplicate QueryExpressionType at position {second_match.start()}")
    else:
        print("   ⚠️  Only one QueryExpressionType found (already fixed?)")
    print()
    
    # Fix 3: Remove duplicate Capabilities (keep first, remove second)
    print("🔧 Fix 3: Removing duplicate Capabilities...")
    
    # Find all occurrences of Capabilities struct
    # This is trickier because it's a larger struct with nested fields
    pattern = r'type Capabilities struct\s*{.*?(?=\ntype\s|\Z)'
    matches = list(re.finditer(pattern, content, re.DOTALL))
    
    if len(matches) > 1:
        # Remove the second occurrence
        second_match = matches[1]
        content = content[:second_match.start()] + content[second_match.end():]
        print(f"   ✅ Removed duplicate Capabilities at position {second_match.start()}")
    else:
        print("   ⚠️  Only one Capabilities found (already fixed?)")
    print()
    
    # Write the fixed file
    with open(TYPES_FILE, 'w') as f:
        f.write(content)
    
    print("=" * 60)
    print("✅ All fixes applied successfully!")
    print("=" * 60)
    print()
    print(f"Original file backed up to: {backup_file}")
    print(f"Fixed file: {TYPES_FILE}")
    print()
    print("Next steps:")
    print("  1. Run: go fmt ./receiver/")
    print("  2. Run: go build ./receiver/")
    print("  3. If successful, continue with SDK generation")

if __name__ == '__main__':
    main()
