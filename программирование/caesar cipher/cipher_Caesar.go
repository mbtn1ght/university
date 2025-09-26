package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var scale int
	var Flag string
	fmt.Println("Хотите зашифровать или расшифровать текст? (+/-)")
	fmt.Scan(&Flag)
	fmt.Print("Сдвиг: ")
	fmt.Scan(&scale)

	if Flag == "+" {
		encryptedText := encryptText(scale)
		fmt.Println("Зашифрованный текст:", encryptedText)
	} else if Flag == "-" {
		decryptedText := decryptText(scale)
		fmt.Println("Расшифрованный текст:", decryptedText)
	} else {
		fmt.Println("Неверно введён текст!")
	}
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

func decryptText(scale int) string {
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
	alphabetSize := len(language)
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
		newIndex := (index - scale + alphabetSize) % alphabetSize
		newChar := language[newIndex]
		if isUpper {
			newChar = unicode.ToUpper(newChar)
		}
		result.WriteRune(newChar)
	}
	return result.String()
}

func encryptText(scale int) string {
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
	alphabetSize := len(language)
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
		newIndex := (index + scale) % alphabetSize
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
