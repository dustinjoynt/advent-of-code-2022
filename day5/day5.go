package main

import (
	"advent-of-code-2022/input"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Ship struct {
	position map[int][]string
}

func main() {

	input, err := input.GetInput("5")
	if err != nil {
		log.Fatalf("failed to get input for day five: %v", err)
	}

	split := splitStructure(input)
	ship := buildStartingStructure(split[0])

	newShip := followMoveCommands(ship, split[1])
	orderedCrates := getCratesOrdered(newShip)

	fmt.Printf("After rearrangement procuedure, the crates on the top of each stack are %v", orderedCrates)
}

func splitStructure(input string) []string {
	re := regexp.MustCompile(`\n\n`)
	return re.Split(input, -1)
}

func buildStartingStructure(s string) Ship {

	cleanedLayers := getLayersCleaned(s)
	positions := cleanedLayers[len(cleanedLayers)-1]
	cleanedLayers = cleanedLayers[:len(cleanedLayers)-1]

	positionMap := make(map[int][]string)

	for _, v := range positions {
		i, _ := strconv.Atoi(v)
		positionMap[i] = []string{}
	}

	for _, layer := range cleanedLayers {
		for i, poVal := range layer {
			if poVal != " " {
				positionMap[i+1] = append([]string{poVal}, positionMap[i+1]...)
			}
		}
	}

	return Ship{
		position: positionMap,
	}
}

func getLayersCleaned(ship string) [][]string {

	re := regexp.MustCompile(`\n`)
	layers := re.Split(ship, -1)

	layersSplit := [][]string{}
	for _, layer := range layers {

		layerSplit := []string{}

		s := strings.Split(layer, "")
		vIndex := 1
		for i, v := range s {

			if i == vIndex {
				layerSplit = append(layerSplit, v)
				vIndex = vIndex + 4
			}
		}
		layersSplit = append(layersSplit, layerSplit)
	}

	return layersSplit
}

func followMoveCommands(ship Ship, s string) Ship {

	cmdRe := regexp.MustCompile(`\n`)
	cmds := cmdRe.Split(s, -1)

	for _, cmd := range cmds {

		if cmd == "" {
			continue
		}

		moveRe := regexp.MustCompile(`[0-9]+`)
		moves := moveRe.FindAllString(cmd, -1)

		crateCount, _ := strconv.Atoi(moves[0])
		sP, _ := strconv.Atoi(moves[1])
		eP, _ := strconv.Atoi(moves[2])

		cratesToMove := ship.position[sP][len(ship.position[sP])-crateCount : len(ship.position[sP])]
		ship.position[sP] = ship.position[sP][:len(ship.position[sP])-crateCount]
		ship.position[eP] = append(ship.position[eP], cratesToMove...)
	}

	return ship
}

func getCratesOrdered(ship Ship) string {
	orderedCrates := ""
	for i := 1; i <= len(ship.position); i++ {
		orderedCrates += ship.position[i][len(ship.position[i])-1]
	}
	return orderedCrates
}
