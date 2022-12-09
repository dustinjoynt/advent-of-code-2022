package main

import (
	"advent-of-code-2022/input"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Dir struct {
	name    string
	subDirs []Dir
	files   []File
}

type File struct {
	name string
	size int
}

var fileSystem Dir
var currentDir Dir
var currentName string
var commands []string

func main() {

	input, err := input.GetInput("7")
	if err != nil {
		log.Fatalf("failed to get input for day seven: %v", err)
	}

	commands = splitCommands(input)

	for _, cmd := range commands {

		ok, currentName := isStepInto(cmd)
		if ok {
			currentDir = findDirectory(currentName, fileSystem)
		}

		if isLs(cmd) {
			currentDir, currentName = fillDirectory(currentDir, commands)
		}

		addToFileSystem(currentName, currentDir)
	}

	// 	root := regexp.MustCompile(`(\$ cd /)`)

	// 	if root.MatchString(cmd) {
	// 		commands = commands[1:]
	// 		d := Dir{
	// 			name: "/",
	// 		}
	// 		currentDir = d
	// 	}

	// 	ls := regexp.MustCompile(`(\$ ls)`)
	// 	if ls.MatchString(cmd) {
	// 		commands = commands[1:]

	// 		cd, name := fillDirectory(currentDir, commands)
	// 		currentDir = cd
	// 		currentName = name
	// 	}
	// }
	// fmt.Printf("v", currentDir)
	// fmt.Printf("v", currentName)
}

func findDirectory(name string, fileSystem Dir) Dir {

	for k, v := range fileSystem.subDirs {
		if v.name == name {
			return v
		}
		findDirectory(name, fileSystem.subDirs[k])
	}
	return Dir{
		name: name,
	}
}

func addToFileSystem(name string, currentDir Dir) {
	fileSystem = currentDir
}

func fillDirectory(parentDir Dir, commands []string) (Dir, string) {

	for _, command := range commands {

		dir := regexp.MustCompile(`dir`)
		if dir.MatchString(command) {
			commands = commands[1:]
			sp := strings.Split(command, " ")
			d := Dir{
				name: sp[1],
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
		}

		cd := regexp.MustCompile(`(\$ cd [a-z]*)`)
		if cd.MatchString(command) {
			commands = commands[1:]
			newDir := regexp.MustCompile(`(\$ cd )`)
			name := newDir.ReplaceAllString(command, "")
			return parentDir, name
		}
	}
	return parentDir, ""
}

func splitCommands(input string) []string {
	re := regexp.MustCompile("\n")
	return re.Split(input, -1)
}

func isStepInto(command string) (bool, string) {
	cd := regexp.MustCompile(`(\$ cd )`)
	if cd.MatchString(command) {
		commands = commands[:1]
		return true, cd.ReplaceAllString(command, "")
	}
	return false, ""
}

func isStepOut(command string) bool {
	cd := regexp.MustCompile(`(\$ cd )`)
	if cd.MatchString(command) {
		dest := cd.ReplaceAllString(command, "")
		if dest == ".." {
			commands = commands[:1]
			return true
		}
	}
	return false
}

func isLs(command string) bool {
	ls := regexp.MustCompile(`(\$ ls)`)
	if ls.MatchString(command) {
		commands = commands[:1]
		return true
	}
	return false
}
