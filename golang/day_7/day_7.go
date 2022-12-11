package main

import (
	"advent-of-code-2022/input"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Dir struct {
	name    string
	parent  *Dir
	subDirs []*Dir
	files   []File
	size    int
}

type File struct {
	name string
	size int
}

func main() {

	input, err := input.GetInput("7")
	if err != nil {
		log.Fatalf("failed to get input for day seven: %v", err)
	}

	var rootDir *Dir
	fileSystem := buildFileSystem(input)
	rootDir = fileSystem.parent

	for rootDir.parent != nil {
		rootDir = rootDir.parent
	}

	totalSize := totalDirs(0, *rootDir)

	fmt.Printf("Sum of the total of all directories is: %v", totalSize)
}

func totalDirs(total int, dir Dir) int {
	if dir.size <= 100000 {
		total += dir.size
	}

	if len(dir.subDirs) > 0 {
		for _, subDir := range dir.subDirs {
			total = totalDirs(total, *subDir)
		}
	}
	return total
}

func buildFileSystem(input string) *Dir {
	var currentDir *Dir
	commands := splitCommands(input)

	for len(commands) > 0 {

		cmd := commands[0]
		if cmd == "" {
			commands = commands[1:]
			continue
		}

		ok, currentName := isStepInto(cmd)
		if ok {
			if currentName == ".." {
				parentDir := currentDir.parent
				currentDir = parentDir
				commands = commands[1:]
				continue
			}

			currentDir = findDirectory(currentName, currentDir)
			commands = commands[1:]
			continue
		}

		if isLs(cmd) {
			commands = commands[1:]
			newCmds := fillDirectory(currentDir, commands)
			commands = newCmds
			continue
		}
	}
	return currentDir
}

func findDirectory(name string, fileSystem *Dir) *Dir {

	if fileSystem == nil {
		return &Dir{
			parent: nil,
			name:   name,
		}
	}

	for k, v := range fileSystem.subDirs {
		if v.name == name {
			return v
		}
		findDirectory(name, fileSystem.subDirs[k])
	}

	return nil
}

func fillDirectory(parentDir *Dir, commands []string) []string {

	for _, command := range commands {

		dir := regexp.MustCompile(`dir`)
		if dir.MatchString(command) {
			commands = commands[1:]
			sp := strings.Split(command, " ")
			d := &Dir{
				name:   sp[1],
				parent: parentDir,
			}
			parentDir.subDirs = append(parentDir.subDirs, d)
		}

		file := regexp.MustCompile(`(\d)`)
		if file.MatchString(command) {
			commands = commands[1:]
			sp := strings.Split(command, " ")
			size, _ := strconv.Atoi(sp[0])
			f := File{
				name: sp[1],
				size: size,
			}
			parentDir.files = append(parentDir.files, f)
			parentDir.size += f.size
			incrementParentDirs(parentDir, f.size)
		}

		chgDir := regexp.MustCompile(`(\$ cd)|(\$ cd)`)
		if chgDir.MatchString(command) {
			break
		}
	}
	return commands
}

func splitCommands(input string) []string {
	re := regexp.MustCompile("\n")
	return re.Split(input, -1)
}

func isStepInto(command string) (bool, string) {
	cd := regexp.MustCompile(`(\$ cd )`)
	if cd.MatchString(command) {
		return true, cd.ReplaceAllString(command, "")
	}
	return false, ""
}

func isLs(command string) bool {
	ls := regexp.MustCompile(`(\$ ls)`)
	return ls.MatchString(command)
}

func incrementParentDirs(dir *Dir, fileSize int) {
	loopDir := dir.parent

	if dir.parent == nil {
		return
	}

	for loopDir.parent != nil {
		loopDir.size += fileSize
		loopDir = loopDir.parent
	}
	loopDir.size += fileSize

}
