# JSON to Excel Converter

This repository contains a simple Go program that converts JSON data into an Excel file (XLSX format) using the `github.com/xuri/excelize/v2` and `encoding/json` package.

## Prerequisites

- Go (1.11 or later)
- [excelize](https://pkg.go.dev/github.com/xuri/excelize/v2) Go package (used for Excel file manipulation)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/sangwaniya/GOlang/json-to-excel.git

2. Go inside the folder 
    ```bash
   cd json-to-excel

3. Tide and vendor packages and their dependencies
    ```bash
   go mod tidy
   go mod vendor

4. Open in VS Code or GO land

## Usage
-----

1.  Place your JSON data in a file named `input.json`.

2.  Run the program to convert the JSON data into an Excel file:

    bashCopy code

    `go run main.go`

3.  The converted Excel file named `output.xlsx` will be generated in the same directory.