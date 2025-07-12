# Copilot Instructions

This document contains coding standards and rules for this Go project.

## Code Formatting and Structure

### Spacing and Line Breaks
- **Never add more than one empty line between declarations**
- **Always add an empty line at the end of every file**
- Ensure all files are properly formatted with `go fmt`

### Function Organization
- **Private helper functions should come after public functions**
- Extract duplicated logic to private helper functions when possible

## Testing Standards

### Test Framework
- Use testify and the `require` package for assertions instead of manual assertions
- Avoid assertion messages unless the assertion is unclear

### Test Structure
- **Avoid return statements in tests, use else blocks instead**
- Use `strings.ToUpper` for consistent capitalization of hexadecimal values in test cases

## Logic and Control Flow

### Conditional Logic
- Invert if logic to handle special cases first, use else blocks for normal cases
- Remove redundant validations when already using regex or other validation methods

## Dependencies Management

### go.mod Organization
- Run `go mod tidy` to organize dependencies properly
- Structure dependencies with:
  - Direct dependencies first
  - Indirect dependencies after in a separate require statement

## General Guidelines

- Maintain backward compatibility when making changes
- Focus on minimal, surgical changes rather than large refactors
- Follow Go naming conventions and best practices