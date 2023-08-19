# go-reloaded

## Project

In this project you will use some of your old functions made in your old repository. You will use them with the objective of making a simple text completion/editing/auto-correction tool.

## Easy Audit Guide

```bash
go run main.go sample.txt result.txt
```
```bash
go run main.go sample1.txt result.txt
```
```bash
go run main.go sample2.txt result.txt
```
```bash
go run main.go sample3.txt result.txt
```


## Usage

```bash
go run main.go file.txt result.txt
```

## Functions

```go 
func main

func ReadFileToString(s string) string {}

func StringToWriteFile(filename, myString string) {}

func indexOfStartBrackets(s string) []int {}

func indexOfEndBrackets(s string) []int {}

func hasComma(s string) bool {}

func returnSubStrAndNum(newStr string) (string, int) {}

func hexToDecimal(sentence string) string {}

func binaryToDecimal(sentence string) string {}

func up(s string) string {}

func low(s string) string {}

func cap(s string) string {}

func formatPunctuation(input string) string {}

func formatPunctuation2(input string) string {}

func shouldChangeAToAn(word string) bool {}

func transformAToAn(input string) string {}

func removeBracketsContent(input string) string {}

```

## Tests

It is recommended to have test files for unit testing.
https://go.dev/doc/tutorial/add-a-test

## License

[MIT](https://choosealicense.com/licenses/mit/)