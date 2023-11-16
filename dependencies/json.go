package dependencies

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func ReadWs(wsString string) (string, error) {
	file, err := os.Open("data/data.json")
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	log.Println(data[wsString])

	return data[wsString].(string), nil

}
