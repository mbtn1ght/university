package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	decryptedText := decryptText()
	fmt.Println("Полученный текст: ", decryptedText)

}

func readFile() string {
	data, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла")
		return ""
	}
	return string(data)
}

func detectLanguage(text string) []rune {
	russian := []rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюя")
	english := []rune("abcdefghijklmnopqrstuvwxyz")

	for _, char := range text {
		if unicode.IsLetter(char) {
			if indexOfRune(russian, unicode.ToLower(char)) != -1 {
				return russian
			}
			if indexOfRune(english, unicode.ToLower(char)) != -1 {
				return english
			}
		}
	}
	return nil
}

func decryptText() string {
	text := readFile()
	if len(text) == 0 {
		fmt.Println("invalid file")
		return ""
	}
	language := detectLanguage(text)
	if language == nil {
		fmt.Println("Не удалось определить язык")
		return text
	}

	var result strings.Builder

	for _, char := range text {
		if !unicode.IsLetter(char) {
			result.WriteRune(char)
			continue
		}
		isUpper := unicode.IsUpper(char)
		lowerChar := unicode.ToLower(char)
		index := indexOfRune(language, lowerChar)
		if index == -1 {
			result.WriteRune(char)
			continue
		}
		newIndex := len(language) - index - 1
		newChar := language[newIndex]
		if isUpper {
			newChar = unicode.ToUpper(newChar)
		}
		result.WriteRune(newChar)
	}
	return result.String()
}

func indexOfRune(arr []rune, char rune) int {
	for i, r := range arr {
		if r == char {
			return i
		}
	}
	return -1
}
