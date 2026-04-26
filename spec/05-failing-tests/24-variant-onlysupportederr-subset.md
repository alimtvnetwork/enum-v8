# Failing Test: Variant OnlySupportedErr / OnlySupportedMsgErr

## Root Cause

`OnlySupportedErr` checks ALL enum variant names against the provided supported-names list.
Passing a subset (e.g., "Owner", "Group") when other variants exist (Invalid, Other, etc.)
correctly returns an error listing the unsupported names.

The test expected `noErr: true`, but the method legitimately returns an error.

## Fix

Changed test expectations to assert `hasErr: true` since passing a subset of variant names
produces an unsupported-names error by design.
