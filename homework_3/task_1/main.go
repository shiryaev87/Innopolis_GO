package main

import (
	"fmt"
)

// Formatter - интерфейс с методом Format
type Formatter interface {
	Format(string) string
}

// PlainText - структура для обычного текста
type PlainText struct{}

// Format - реализация метода Format для PlainText
func (p PlainText) Format(text string) string {
	return text
}

// BoldText - структура для жирного текста
type BoldText struct{}

// Format - реализация метода Format для BoldText
func (b BoldText) Format(text string) string {
	return "**" + text + "**"
}

// CodeText - структура для текста кода
type CodeText struct{}

// Format - реализация метода Format для CodeText
func (c CodeText) Format(text string) string {
	return "`" + text + "`"
}

// ItalicText - структура для курсивного текста
type ItalicText struct{}

// Format - реализация метода Format для ItalicText
func (i ItalicText) Format(text string) string {
	return "_" + text + "_"
}

// ChainFormatter - структура для цепочки форматировщиков
type ChainFormatter struct {
	formatters []Formatter
}

// AddFormatter - метод для добавления форматировщика в цепочку
func (cf *ChainFormatter) AddFormatter(f Formatter) {
	cf.formatters = append(cf.formatters, f)
}

// Format - реализация метода Format для ChainFormatter
func (cf *ChainFormatter) Format(text string) string {
	for _, formatter := range cf.formatters {
		text = formatter.Format(text)
	}
	return text
}

func main() {
	plainText := PlainText{}
	bold := BoldText{}
	code := CodeText{}
	italic := ItalicText{}

	// Пример использования ChainFormatter
	chainFormatter := ChainFormatter{}
	chainFormatter.AddFormatter(plainText)
	chainFormatter.AddFormatter(bold)
	chainFormatter.AddFormatter(code)
	chainFormatter.AddFormatter(italic)

	input := "Hello, world!"
	output := chainFormatter.Format(input)
	fmt.Println(output) // Ожидаемый результат: _`**Hello, world!**`_
}
