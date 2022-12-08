package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	name string
	parent *File
	files []*File
	size int
}

func (file *File) GetTotalSize() int {
	if file.files == nil {
		return file.size
	}

	totalSize := 0

	for _, f := range file.files {
		totalSize += f.GetTotalSize()
	}

	return totalSize
}

func SolveDay07() {
	content, err := os.ReadFile("input/day07.txt")
	if err != nil {
		log.Fatal(err)
	}

	contentStr := string(content)

	lines := strings.Split(contentStr, "\n")

	root := new(File)
	root.parent = nil
	root.name = "/"
	root.files = make([]*File, 0)
	root.size = 0

	currentDir := root

	dirs := make([]*File, 0)
	dirs = append(dirs, root)

	for _, line := range lines[1:] {
		switch line[:3] {
		case "$ c":
			fileName := line[5:]
			if fileName == ".." {
				currentDir = currentDir.parent
				break
			}

			newFile := new(File)
			newFile.name = fileName
			newFile.parent = currentDir
			newFile.files = make([]*File, 0)
			newFile.size = 0
			currentDir.files = append(currentDir.files, newFile)
			currentDir = newFile
			dirs = append(dirs, newFile)
		case "$ l":
			break
		case "dir":
			break
		default:
			info := strings.Split(line, " ")
			newFile := new(File)
			newFile.name = info[1]
			newFile.parent = currentDir
			newFile.size, _ = strconv.Atoi(info[0])
			newFile.files = nil
			currentDir.files = append(currentDir.files, newFile)
		}
	}

	sizes := make([]int, 0)

	for _, d := range dirs {
		sizes = append(sizes, d.GetTotalSize())
	}

	sort.Ints(sizes)

	unusedSpace := 70_000_000 - sizes[len(sizes)-1]

	for _, s := range sizes {
		if s + unusedSpace >= 30_000_000 {
			fmt.Println(s)
			return
		}
	}
}
