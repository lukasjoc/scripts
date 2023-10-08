package main

import (
	"fmt"
	"math/rand"
)

var hiragana = map[string]string{
	"a":   "あ",
	"i":   "い",
	"u":   "う",
	"e":   "え",
	"o":   "お",
	"ka":  "か",
	"ki":  "き",
	"ku":  "く",
	"ke":  "け",
	"ko":  "こ",
	"sa":  "さ",
	"shi": "し",
	"su":  "す",
	"se":  "せ",
	"so":  "そ",
}

func getRandomHiragana() (string, string) {
	keys := make([]string, 0, len(hiragana))
	for k := range hiragana {
		keys = append(keys, k)
	}
	randomIndex := rand.Intn(len(keys))
	k := keys[randomIndex]
	return k, hiragana[k]
}

const (
	reset = "\033[0m"
	bold  = "\033[1m"
)

func main() {
	streak := 0
	for {
		hirKey, hirVal := getRandomHiragana()
		fmt.Printf("%s%s%s (streak: %d)? ", bold, hirVal, reset, streak)
		var ans string
		fmt.Scanln(&ans)

		if ans == "exit" {
			break
		}

		correct, exists := hiragana[ans]
		if exists && correct == hirVal {
			streak++
		} else {
			streak = 0
			fmt.Printf("Not quite. The correct answer is '%s%s%s'.\n", bold, hirKey, reset)
		}
	}
}
