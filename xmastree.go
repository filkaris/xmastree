package main

import (
	"fmt"
	"math/rand"
	"time"
)

const Clear = "\033[H\033[2J"

func main() {

	var newTree []rune
	for {
		newTree = []rune{}
		for _, c := range Tree {
			if c != '*' {
				newTree = append(newTree, c)
				continue
			}

			newTree = append(newTree, []rune(getColor())...)
			newTree = append(newTree, c)
			newTree = append(newTree, []rune("\033[0m")...)
		}

		fmt.Println(Clear)
		fmt.Println(string(newTree))

		time.Sleep(1 * time.Second)
	}

}

func getColor() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	colors := []string{
		"\033[0m",
		"\033[0;31m",
		"\033[0;32m",
		"\033[0;33m",
		"\033[0;34m",
		"\033[0;35m",
		"\033[0;36m",
	}

	return colors[r.Intn(7)]

}
