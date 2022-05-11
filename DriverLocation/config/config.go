package config

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

type Config struct {
	MongoDbUrl string `json:"mongoDbUrl"`
	MongoRecordSeparator string `json:"mongoRecordSeparator"`
	RecordBatchCount int `json:"recordBatchCount"`
	SecretMatchApiKey string `json:"secretMatchApiKey"`
	MatchApiHeader string `json:"matchApiHeader"`
	MongoDbName string `json:"mongoDbName"`
}

var c = &Config{}

func init() {
	mydir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if strings.Contains(mydir, "test") || strings.Contains(mydir, "pact") {
		os.Chdir("../..")
	}

	file, err := os.Open(".config/" + env + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	read, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(read, c)
	if err != nil {
		panic(err)
	}
}



func Get() *Config {
	return c
}
