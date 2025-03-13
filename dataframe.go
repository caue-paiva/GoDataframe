package dataframe

import (
	"dataframe/dtypes"
	"encoding/csv"
	"fmt"
	"os"
)

type column struct {
	name    string
	colType dtypes.Dtype
}

type DataFrame struct {
	Shape [2]int
	Cols  []column
	data  [][]any
}

func readCsv(filePath string) (DataFrame, error) {
	reader, readErr := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)

	if readErr != nil {
		return DataFrame{}, fmt.Errorf("falha ao abrir arquivo csv, erro: %v", readErr)
	}

	var csvReader *csv.Reader = csv.NewReader(reader)
}
