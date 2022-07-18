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


	// str := "{{asdf{sdf} {}}}"
	// bytes := []byte(str)

	ProcessFile(args[0])
}

func ProcessFile(path string) {
	bytes, err := os.ReadFile(path)
	emptyBytes := make([]byte, 0)
	DieIf(err)

	// str := string(bytes)
	// origStr := str 



	invalidFieldsRe, err := regexp.CompilePOSIX(`^[ \t]*-\??:.+;`)
	DieIf(err)

	bytes = invalidFieldsRe.ReplaceAll(bytes, emptyBytes)

	modelsRe, err := regexp.Compile(`models\.`)
	DieIf(err)

	bytes = modelsRe.ReplaceAll(bytes, emptyBytes)

	// str := string(bytes)
	badTypes := regexp.MustCompile(`[ \t]*export class \w+[LR] {`)



	// removeSegments := make([][][]int, 0)


	removedCount := 0
	indicies := badTypes.FindAllIndex(bytes, -1)
	if indicies != nil {
		for _, match := range(indicies) {
			// println(string(bytes[match[1] -1]))
			// println(string(bytes[match[1]-1]))
			match[0] -= removedCount
			match[1] -= removedCount

			pos, ok:= FindMatchingBrace(bytes, match[1] - 1)
			if !ok {
				panic("aaaa")
			}

			println(string(bytes[match[0]:pos+1]))
			println("--------------------------")
			bytes = append(bytes[:match[0]], bytes[pos+2:]...)

			removedCount += pos+2 - match[0]
		}
		// bytes = append(bytes[])
	}

	// runes := []rune(string(bytes))


	os.WriteFile(path, bytes, 0777)
}

func FindMatchingBrace(bytes []byte, startIdx int) (int, bool) {
	count := 1
	for idx := startIdx +1; idx < len(bytes); idx++ {
		if bytes[idx] == '{' {
			count++;
		} else if bytes[idx] == '}' {
			count--;
		}
		if count == 0 {
			return idx, true
		}
	}

	return 0, false
}


func DieIf(err error){
	if err != nil {
		panic(err)
	}
}
