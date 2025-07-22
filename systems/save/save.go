package savesystem

import (
	"encoding/json"
	"io"
	"os"
)

func Save(data any, path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(json)
	if err != nil {
		panic(err)
	}
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
