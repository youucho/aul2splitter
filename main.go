package main

import (
	"fmt"
	"maps"
	"os"
	"strings"
)

type aul2 map[string]map[string]string

func error(e any) { // why error isnt type
	if e != nil {
		panic(e)
	}
}

func ParseAul2(file string) aul2 {
	array := strings.Split(strings.ReplaceAll(file, "\r\n", "\n"), "\n") // i hate windows
	result := make(aul2)
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
			result[currentField] = make(map[string]string)
		default:
			if currentField == "" {
				return make(aul2) // error
			} else {
				original, translation, found := strings.Cut(v, "=")
				if !found {
					continue
				}
				result[currentField][original] = translation
			}
		}
	}
	return result
}

func MergeAul2(base aul2, sub aul2) aul2 {
	result := make(aul2)
	maps.Copy(result, base)
	for field, v := range sub {
		if result[field] == nil {
			result[field] = make(map[string]string)
		}
		for ori, tra := range v {
			if _, prs := base[field][ori]; prs {
				fmt.Println("skip cuz exist")
				continue
			} else {
				result[field][ori] = tra
			}
		}
	}
	return result
}

func main() {
	base, err := os.ReadFile("base.aul2")
	error(err)
	sub, err := os.ReadFile("sub.aul2")
	error(err)
	baseaul2 := ParseAul2(string(base))
	subaul2 := ParseAul2(string(sub))
	fmt.Println(baseaul2)
	fmt.Println(subaul2)
	result := MergeAul2(baseaul2, subaul2)
	fmt.Println(result)
}
