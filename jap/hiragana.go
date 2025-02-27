package main

import "fmt"

const (
	reset = "\033[0m"
	bold  = "\033[1m"
)

var hiragana = map[string]string{
	"a": "あ", "i": "い", "u": "う", "e": "え", "o": "お",
	"ka": "か", "ki": "き", "ku": "く", "ke": "け", "ko": "こ",
	"sa": "さ", "shi": "し", "su": "す", "se": "せ", "so": "そ",
	"ta": "た", "chi": "ち", "tsu": "つ", "te": "て", "to": "と",
	// TODO: NA,HA,MA,YA,RA,WA ...
}

func getRandomHiragana(next *int) (string, string) {
	keys := make([]string, 0, len(hiragana))
	for k := range hiragana {
		keys = append(keys, k)
	}
	k := keys[*next]
	if *next == len(keys)-1 {
		*next = 0
	} else {
		*next++
	}
	return k, hiragana[k]
}

func main() {
	streak := 0
	next := 0
	for {
		hirKey, hirVal := getRandomHiragana(&next)
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
			fmt.Printf("Incorrect! The correct answer is '%s%s%s'.\n", bold, hirKey, reset)
		}
	}
}
