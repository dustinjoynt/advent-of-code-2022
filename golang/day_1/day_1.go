package main

import (
	"advent-of-code-2022/env"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	input := getInput("https://adventofcode.com/2022/day/1/input")

	re := regexp.MustCompile("\n\n")
	elfSplit := re.Split(input, -1)

	elfLoads := []int{}
	for _, v := range elfSplit {
		re := regexp.MustCompile("\n")
		loadSplit := re.Split(v, -1)

		loadSum := 0
		for _, v := range loadSplit {

			intV, _ := strconv.Atoi(v)
			loadSum += intV
		}

		elfLoads = append(elfLoads, loadSum)
		loadSum = 0
	}

	topThree := []int{}
	for _, v := range elfLoads {

		if len(topThree) < 3 {
			topThree = append(topThree, v)
		} else {
			if v > topThree[0] {
				topThree[2] = topThree[1]
				topThree[1] = topThree[0]
				topThree[0] = v
			} else if v > topThree[1] {
				topThree[2] = topThree[1]
				topThree[1] = v
			} else if v > topThree[2] {
				topThree[2] = v
			}
		}
	}

	topThreeSum := 0
	for _, v := range topThree {
		topThreeSum += v
	}

	fmt.Printf("total calories carried by the elf with the most calories is: %v \n", topThree[0])
	fmt.Printf("total calories carried by the top three elves is: %v", topThreeSum)
}

// todo::dynamically auth this
func getInput(link string) string {

	err := env.Load("../.env")
	if err != nil {
		panic(err)
	}
	token := os.Getenv("AUTH_TOKEN")

	req, _ := http.NewRequest(http.MethodGet, link, nil)

	c := &http.Cookie{
		Name:    "session",
		Value:   token,
		Domain:  ".adventofcode.com",
		Path:    "/",
		Expires: time.Now().Add(48 * time.Hour),
	}
	req.AddCookie(c)

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
