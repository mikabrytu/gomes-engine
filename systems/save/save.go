package savesystem

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Save(data any) {
	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("json data: %s\n", json)
}

func Load(path string, out any) error {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return json.Unmarshal([]byte(byteValue), out)
}
