package main

import (
	"fmt"
	"os"
	"strings"
)

type aul2 struct {
}

func error(e any) { // why error isnt type
	if e != nil {
		panic(e)
	}
}

func ParseAul2(file string) map[string][][2]string {
	array := strings.Split(strings.ReplaceAll(file, "\r\n", "\n"), "\n") // i hate windows
	aul2 := make(map[string][][2]string)
	currentField := ""
	for _, v := range array {
		if len(v) == 0 {
			continue
		}
		switch v[0] {
		case ';':
			continue
		case '[':
			currentField = v[1 : len(v)-1]
			fmt.Println(currentField)
		default:
			if currentField == "" {
				return map[string][][2]string{} // error
			} else {
				original, translation, found := strings.Cut(v, "=")
				if !found {
					continue
				}
				aul2[currentField] = append(aul2[currentField], [2]string{original, translation})
			}
		}
	}
	return aul2
}

func main() {
	test, err := os.ReadFile("English.aul2")
	error(err)
	fmt.Println(ParseAul2(string(test)))
}
