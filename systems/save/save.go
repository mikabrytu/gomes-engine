package savesystem

import (
	"encoding/json"
	"io"
	"os"
)

func Save() {

}

func Load(path string) map[string]interface{} {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}
