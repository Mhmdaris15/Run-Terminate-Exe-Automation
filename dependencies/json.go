package dependencies

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func ReadJSON() {
	file, err := os.Open("data/data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(data["pd"])

}
