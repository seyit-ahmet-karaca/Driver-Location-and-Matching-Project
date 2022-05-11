package utils

import (
	"driverlocation/config"
	"encoding/csv"
	"log"
	"os"
)

func CreateReader(filename string) (*csv.Reader, *os.File) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	csvReader := csv.NewReader(file)
	csvReader.Comma = rune(config.Get().MongoRecordSeparator[0])

	return csvReader, file
}
