package fileops

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func WriteCSV(filename string, records []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		if err := writer.Write([]string{record}); err != nil {
			return fmt.Errorf("could not write record to file: %v", err)
		}
	}
	return nil
}

func Read_csv(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var data []string

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("failed to read CSV data: %s", err)
			panic(err)
		}
		csvRecord := strings.Join(record, ",")
		data = append(data, csvRecord)
	}
	return data
}
