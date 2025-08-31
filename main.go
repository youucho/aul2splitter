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
				continue
			} else {
				result[field][ori] = tra
			}
		}
	}
	return result
}

func BulkMergeAul2(directory string) aul2 {
	result := make(aul2)
	err := os.Chdir(directory)
	error(err)
	dir, err := os.ReadDir(".")
	error(err)
	var files []string
	for _, v := range dir {
		if v.IsDir() {
			continue
		}
		files = append(files, v.Name())
	}
	for _, name := range files {
		sub, err := os.ReadFile(name)
		error(err)
		subaul2 := ParseAul2(string(sub))
		result = MergeAul2(result, subaul2)
	}
	err = os.Chdir("../")
	error(err)
	return result
}

func WriteAul2(content aul2) {
	string := "; This file was written with aul2splitter\n"
	f, err := os.Create("Result.aul2")
	error(err)
	defer f.Close()
	for field, lines := range content {
		string += "[" + field + "]\n"
		for ori, tra := range lines {
			string += ori + "=" + tra + "\n"
		}
	}
	f.WriteString(string)
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Use 'aul2splitter <directory> to merge aul2 files'")
	} else {
		fmt.Printf("Merging %s...\n", args[1])
		result := BulkMergeAul2(args[1])
		WriteAul2(result)
	}
}
