package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	PrintDirTree(`D:/AIMP`)
}

func PrintDirTree(root string) {
	sb := new(strings.Builder)
	printDirTree(root, 0, sb)
	fmt.Println(sb.String())
}

func printDirTree(rootDir string, depth int, sb *strings.Builder) {
	if !IsDirExist(rootDir) {
		return
	}
	sb.WriteString(getIndent(depth))
	sb.WriteString("+--")
	if depth != 0 {
		fi, _ := os.Stat(rootDir)
		sb.WriteString(fi.Name())
	} else {
		sb.WriteString(rootDir)
	}
	sb.WriteString("/")
	sb.WriteString("\n")
	files, _ := ioutil.ReadDir(rootDir)
	for _, file := range files {
		if file.IsDir() {
			printDirTree(rootDir+"/"+file.Name(), depth+1, sb)
		} else {
			printFile(file, depth+1, sb)
		}
	}
}

func printFile(file os.FileInfo, indent int, sb *strings.Builder) {
	sb.WriteString(getIndent(indent))
	sb.WriteString("+--")
	sb.WriteString(file.Name())
	sb.WriteString("\n")
}

func getIndent(indent int) string {
	var sb strings.Builder
	for i := 0; i < indent; i++ {
		sb.WriteString("|  ")
	}
	return sb.String()
}

func IsDirExist(dir string) bool {
	fi, err := os.Stat(dir)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}
