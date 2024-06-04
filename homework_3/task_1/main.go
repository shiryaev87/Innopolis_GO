/*
MessageFormatter
Реализуйте интерфейс Formatter
С методом Format,
Который возвращает отформатированную строку.
Определите структуры, удовлетворяющие интерфейсу
Formatter: обычный текст(какесть), жирным шрифтом(****), код(“), курсив(__)
Опционально: иметь возможность задавать цепочку модификаторов
chainFormatter.AddFormatter(plainText) chainFormatter.AddFormatter(bold) chainFormatter.AddFormatter(code)
*/
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
	//var bold_1 = BoldText{}
	inputText := "Да пребудет с тобой сила."
	plainText := PlainText{}
	fmt.Printf("Неотформатированный текст: %s\n", plainText.Format(inputText))

	bold := BoldText{}
	fmt.Printf("Отформатированный жирный текст: %s\n", bold.Format(inputText))

	code := CodeText{}
	fmt.Printf("Отформатированный текст в виде кода: %s\n", code.Format(inputText))

	italic := ItalicText{}
	fmt.Printf("Отформатированный курсивный текст: %s\n", italic.Format(inputText))
	fmt.Println()
	// Пример использования ChainFormatter
	//объявляем цепочку форматирования

	fmt.Printf("Пример работы отформатированной цепочки\n")
	chainFormatter := ChainFormatter{}
	//добавляем в цепочку форматы
	chainFormatter.AddFormatter(plainText)
	chainFormatter.AddFormatter(bold)
	chainFormatter.AddFormatter(code)
	chainFormatter.AddFormatter(italic)

	//вызываем метод для обработки цепочки форматирования
	output := chainFormatter.Format(inputText)
	fmt.Println(output)
}
