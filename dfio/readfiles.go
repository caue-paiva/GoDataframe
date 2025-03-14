package dfio

import (
	"encoding/csv"
	"fmt"
	"os"
)

func __ReadCsv(filePath string) (DataFrame, error) {
	reader, readErr := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)

	if readErr != nil {
		return DataFrame{}, fmt.Errorf("falha ao abrir arquivo csv, erro: %v", readErr)
	}

	var csvReader *csv.Reader = csv.NewReader(reader)
}
