package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) == 3 {
		str := ReadFileToString(os.Args[1]) // convert content to a 'string'
		// fmt.Println(str) // print the content as a 'string'
		startArrBrac := indexOfStartBrackets(str)
		endArrBrac := indexOfEndBrackets(str)
		lenstartArrBrac := len(startArrBrac)
		// (low), (up), (cap), (hex), (bin)
		for i := 0; i < lenstartArrBrac; i++ {
			newStr := str[startArrBrac[i]+1 : endArrBrac[i]]
			// panic: runtime error: slice bounds out of range [:584] with length 583 when string ends with round brackets ()
			newStr2, _ := returnSubStrAndNum(newStr)
			// fmt.Println(newStr)
			if newStr2 == "up" {
				str = up(str)
				// fmt.Println(integer)
			} else if newStr2 == "cap" {
				str = cap(str)
			} else if newStr2 == "low" {
				str = low(str)
			} else if newStr2 == "hex" {
				str = hexToDecimal(str)
			} else if newStr2 == "bin" {
				str = binaryToDecimal(str)
			}
			startArrBrac = indexOfStartBrackets(str)
			endArrBrac = indexOfEndBrackets(str)
			// fmt.Println(newStr2, integer)
		}

		// removeBracketsContent
		str = removeBracketsContent(str)

		// transform A To An
		str = transformAToAn(str)

		// formatPunctuation
		str = formatPunctuation(str)
		str = formatPunctuation2(str)

		fmt.Println("\n", str)
		StringToWriteFile(os.Args[2], str)

	}
}

func removeBracketsContent(input string) string {
	var result string
	inBracket := false

	for _, char := range input {
		if char == '(' {
			inBracket = true
		} else if char == ')' {
			inBracket = false
		} else if !inBracket {
			result += string(char)
		}
	}

	return result
}

func binaryToDecimal(sentence string) string {
	words := strings.Fields(sentence)
	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" && i > 0 {
			binaryStr := words[i-1]
			decimalNum, err := strconv.ParseInt(binaryStr, 2, 64)
			if err == nil {
				words[i-1] = fmt.Sprintf("%d", decimalNum)
			}
		}
	}

	return strings.Join(words, " ")
}

func hexToDecimal(sentence string) string {
	words := strings.Fields(sentence)
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			hexStr := words[i-1]
			decimalNum, err := strconv.ParseInt(hexStr, 16, 64)
			if err == nil {
				words[i-1] = fmt.Sprintf("%d", decimalNum)
			}
		}
	}

	return strings.Join(words, " ")
}

func formatPunctuation2(input string) string {
	var output []rune
	var prevRune rune

	for _, r := range input {

		if r == '\'' || r == ' ' {
			if prevRune == ' ' && r == '\'' {
				// Remove the space before punctuation
				output = output[:len(output)-1]
				output = append(output, r)
			} else if prevRune == '\'' && r == ' ' {
			} else {
				output = append(output, r)
			}
		} else {
			output = append(output, r)
		}

		prevRune = r
	}

	return string(output)
}

func formatPunctuation(input string) string {
	var output []rune
	var prevRune rune

	for index, r := range input {

		if r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';' {
			if prevRune == ' ' {
				// Remove the space before punctuation
				output = output[:len(output)-1]
			}
			output = append(output, r)
			if index != len(input)-1 {
				if !unicode.IsPunct(rune(input[index+1])) && rune(input[index+1]) != ' ' {
					output = append(output, ' ') // Add a space after punctuation
				}
			}

		} else {
			output = append(output, r)
		}

		prevRune = r
	}

	return string(output)
}

func returnSubStrAndNum(newStr string) (string, int) {
	if hasComma(newStr) {
		tempStr := strings.Split(newStr, " ")
		split, err := strconv.Atoi(tempStr[1])
		if err != nil {
			fmt.Println(err)
		}
		returnString := strings.Split(tempStr[0], ",")
		return returnString[0], split
	} else {
		return newStr, 1
	}
}

func ReadFileToString(s string) string {
	b, err := os.ReadFile(os.Args[1]) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	return string(b) // convert content to a 'string'
}

func StringToWriteFile(filename, myString string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(myString)
	if err2 != nil {
		fmt.Println(err2)
	}
}

func indexOfStartBrackets(s string) []int {
	var mArr []int
	for index, v := range s {
		if v == '(' {
			mArr = append(mArr, index)
		}
	}
	return mArr
}

func indexOfEndBrackets(s string) []int {
	var mArr []int
	for index, v := range s {
		if v == ')' {
			mArr = append(mArr, index)
		}
	}
	return mArr
}

func hasComma(s string) bool {
	for _, v := range s {
		if v == ',' {
			return true
		}
	}
	return false
}

func shouldChangeAToAn(word string) bool {
	if len(word) == 0 {
		return false
	}
	firstChar := rune(word[0])
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, firstChar) || firstChar == 'h' || firstChar == 'H'
}

func transformAToAn(input string) string {
	words := strings.Fields(input)
	result := make([]string, 0, len(words))

	for i := 0; i < len(words); i++ {
		currentWord := words[i]
		if i+1 < len(words) && shouldChangeAToAn(currentWord) {
			nextWord := words[i+1]
			if shouldChangeAToAn(nextWord) {
				if strings.Contains(currentWord, "m") || strings.Contains(currentWord, "s") {
				} else {
					currentWord = strings.Replace(currentWord, "a", "an", 1)
				}
			}
		}
		result = append(result, currentWord)
	}

	return strings.Join(result, " ")
}

func up(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(up)") {
			startIndex := i - 1
			for j := startIndex; j < i; j++ {
				words[j] = strings.ToUpper(words[j])
			}
		}
		if strings.HasPrefix(words[i], "(up,") {
			parts := strings.Split(words[i+1], ")")
			if len(parts) >= 2 {
				number, err := strconv.Atoi(strings.TrimSuffix(parts[0], " "))
				if err == nil {
					startIndex := i - number
					for j := startIndex; j < i; j++ {
						words[j] = strings.ToUpper(words[j])
					}
				}
			}
		}
	}
	return strings.Join(words, " ")
}

func low(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(low)") {
			startIndex := i - 1
			for j := startIndex; j < i; j++ {
				words[j] = strings.ToLower(words[j])
			}
		}
		if strings.HasPrefix(words[i], "(low,") {
			parts := strings.Split(words[i+1], ")")
			if len(parts) >= 2 {
				number, err := strconv.Atoi(strings.TrimSuffix(parts[0], " "))
				if err == nil {
					startIndex := i - number
					for j := startIndex; j < i; j++ {
						words[j] = strings.ToLower(words[j])
					}
				}
			}
		}
	}
	return strings.Join(words, " ")
}

func cap(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(cap)") {
			startIndex := i - 1
			for j := startIndex; j < i; j++ {
				words[j] = strings.Title(words[j])
			}
		}
		if strings.HasPrefix(words[i], "(cap,") {
			parts := strings.Split(words[i+1], ")")
			if len(parts) >= 2 {
				number, err := strconv.Atoi(strings.TrimSuffix(parts[0], " "))
				if err == nil {
					startIndex := i - number
					for j := startIndex; j < i; j++ {
						words[j] = strings.Title(words[j])
					}
				}
			}
		}
	}
	return strings.Join(words, " ")
}
