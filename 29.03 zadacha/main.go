package main

import (
	"fmt"
	"strings"
	"unicode"
)

func findMissingLetters(letters string) []rune {
	var missingLetters []rune
	isRussian := false
	isEnglish := false
	for _, char := range letters {
		if unicode.Is(unicode.Cyrillic, char) {
			isRussian = true
		} else if unicode.Is(unicode.Latin, char) {
			isEnglish = true
		}
	}
	letterMap := make(map[rune]bool)
	for _, char := range letters {
		letterMap[char] = true
	}
	if isRussian {
		for char := 'а'; char <= 'я'; char++ {
			if !letterMap[char] {
				missingLetters = append(missingLetters, char)
			}
		}
	}
	if isEnglish {
		for char := 'a'; char <= 'z'; char++ {
			if !letterMap[char] {
				missingLetters = append(missingLetters, char)
			}
		}
	}
	return missingLetters
}

func main() {
	fmt.Println("Введи буквы:")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	letters := strings.Fields(input)

	missingLetters := findMissingLetters(strings.Join(letters, ""))

	if len(missingLetters) > 0 {
		fmt.Println("Недостающие буквы:")
		for _, char := range missingLetters {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	} else {
		fmt.Println("Недостающих букв нет.")
	}
}
