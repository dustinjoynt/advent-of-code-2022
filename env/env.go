package env

import (
	"os"
	"regexp"
)

func Load(path string) error {

	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	re := regexp.MustCompile("\n")
	split := re.Split(string(b), -1)

	for _, v := range split {
		re2 := regexp.MustCompile("=")
		split2 := re2.Split(v, -1)
		os.Setenv(split2[0], split2[1])
	}

	return nil
}
