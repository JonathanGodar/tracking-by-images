package main

import (
	"os"
	"regexp"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		panic("There has to be only one argument passed to the program")
	}

}

func ProcessFile(path string) {
	bytes, err := os.ReadFile(path)
	emptyBytes := make([]byte, 0)
	DieIf(err)

	// original := string(originalByte)
	modelsRe, err := regexp.Compile(`models\.`)
	DieIf(err)

	bytes = modelsRe.ReplaceAll(bytes, emptyBytes)
	// modelsRe.FindAllIndex()
	// modelsRe.FindAllIndex(bytes, )
	

	os.WriteFile(path, bytes, 0777)
}


func DieIf(err error){
	if err != nil {
		panic(err)
	}
}
