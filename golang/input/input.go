package input

import (
	"advent-of-code-2022/env"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func GetInput(link string) (string, error) {

	err := env.Load("../.env")
	if err != nil {
		return "", err
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

	client := http.Client{Timeout: time.Second * 30}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content), nil
}
