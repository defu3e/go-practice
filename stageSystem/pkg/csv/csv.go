package csv

import (
	"encoding/csv"
	"log"
	"os"
)

func GetCSVdata(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println("ошибка чтения файла:", err)
		return [][]string{}, err
	}

	return rows, nil
}

func WriteDataToCSV(filePath string, data [][]string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("ошибка создания файла:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(data)
	if err != nil {
		log.Println("ошибка записи в файл:", err)
	}
}