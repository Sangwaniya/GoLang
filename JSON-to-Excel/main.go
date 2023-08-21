package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	// Read JSON data from file
	jsonData, err := os.ReadFile("input.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal JSON data into a map
	var jsonDataMap map[string]interface{}
	err = json.Unmarshal(jsonData, &jsonDataMap)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Extract the "data" field as my data present int "data" field and i need that only
	dataArray, ok := jsonDataMap["data"].([]interface{})
	if !ok {
		fmt.Println("Error: 'data' field not found in JSON")
		return
	}

	// Create a new Excel file
	f := excelize.NewFile()

	// Create a new sheet
	sheetName := "Data"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Extract headers from the first data item
	headers := make([]string, 0)
	if len(dataArray) > 0 {
		firstData := dataArray[0].(map[string]interface{})
		for key := range firstData {
			headers = append(headers, key)
		}
	}

	// Add headers to the Excel sheet
	for col, header := range headers {
		f.SetCellValue(sheetName, fmt.Sprintf("%c1", 'A'+col), header)
	}

	// Add data to the Excel sheet
	for row, data := range dataArray {
		dataMap := data.(map[string]interface{})
		for col, header := range headers {
			value := dataMap[header]
			f.SetCellValue(sheetName, fmt.Sprintf("%c%d", 'A'+col, row+2), value)
		}
	}

	// Set active sheet
	f.SetActiveSheet(index)

	// Save the Excel file
	err = f.SaveAs("output.xlsx")
	if err != nil {
		fmt.Println("Error saving Excel file:", err)
		return
	}

	fmt.Println("Excel file saved as output.xlsx")
}
