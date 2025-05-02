# Workdays Calculator

A simple command-line tool to calculate workdays between two dates, with support for holidays.

## Features

- Calculate workdays between two dates
- Exclude holidays from workday calculation
- Read holidays from a comma-separated file
- Pretty terminal output using Charm libraries

## Usage

```bash
./workdays -start 2025-01-01 -end 2025-01-31
```

### Parameters

- `-start`: Start date in YYYY-MM-DD format (required)
- `-end`: End date in YYYY-MM-DD format (required)
- `-holidays`: Path to holidays file (default: holidays.txt)

## Holidays File Format

The holidays file must exist and contain at least one valid holiday date. Dates should be specified in YYYY-MM-DD format, separated by commas:

```
2025-01-01,2025-01-20,2025-02-17,2025-05-26
```

### Error Handling

The program will return an error if:

- The holidays file doesn't exist
- The holidays file is empty
- The holidays file contains no valid dates
- Any date in the file is in an invalid format

## Building

Using Go directly:

```bash
go build
```

Using Nix development environment:

```bash
# Enter development shell
nix develop

# Or with direnv (if installed)
direnv allow

# Then build
go build
```

## Example Output

```
  Workdays  

Date Range: Jan 1, 2025 to Jan 31, 2025

  1. Thursday, January 2, 2025
  2. Friday, January 3, 2025
  ...
  21. Friday, January 31, 2025

Total workdays: 21
```

## Dependencies

- github.com/charmbracelet/lipgloss - For terminal styling
- github.com/rickar/cal/v2 - For calendar and business day calculations