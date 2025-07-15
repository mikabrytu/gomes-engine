package savesystem

import (
	"container/list"
	"os"
)

var root string
var files *list.List

func Init(path string) {
	root = path
	files = list.New()
}

func Save() {

}

func Load(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	println("File opened!")

	defer file.Close()
}
