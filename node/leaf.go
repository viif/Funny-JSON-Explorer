package node

import (
	"FJE/icon"
	"FJE/symbol"

	"log"
	"io"
	"unicode/utf8"
)

var _ = log.Println

// 叶节点
type Leaf struct {
	Name                           string        // 名称
	Icon                           icon.Icon     // 图标
	ElementPrefixFirstLineSymbol   symbol.Symbol // 元素前的结构符号（位于首行）
	ElementPrefixLastLineSymbol    symbol.Symbol // 元素前的结构符号（位于末行）
	ElementPrefixSiblingLastSymbol symbol.Symbol // 元素前的结构符号（位于同等级最后一个元素前）
	ElementPrefixSiblingSymbol     symbol.Symbol // 元素前的结构符号（位于同等级非最后一个元素前）
	ElementSuffixSymbol            symbol.Symbol // 元素后到行末的结构符号
	LineSuffixFirstLineSymbol      symbol.Symbol // 行末的结构符号（位于首行）
	LineSuffixLastLineSymbol       symbol.Symbol // 行末的结构符号（位于末行）
	LineSuffixMiddleLineSymbol     symbol.Symbol // 行末的结构符号（位于中间行）
	LinePrefixLastLineSymbol       symbol.Symbol // 行首的结构符号（位于末行）
	LinePrefixMiddleLineSymbol     symbol.Symbol // 行首的结构符号（位于中间行）
	ElementPrefixLinePrefixSymbol  symbol.Symbol // 元素前到行首的结构符号
}

func (l *Leaf) Clone() Node {
	return &Leaf{
		Name:                           l.Name,
		Icon:                           l.Icon,
		ElementPrefixFirstLineSymbol:   l.ElementPrefixFirstLineSymbol,
		ElementPrefixLastLineSymbol:    l.ElementPrefixLastLineSymbol,
		ElementPrefixSiblingLastSymbol: l.ElementPrefixSiblingLastSymbol,
		ElementPrefixSiblingSymbol:     l.ElementPrefixSiblingSymbol,
		ElementSuffixSymbol:            l.ElementSuffixSymbol,
		LineSuffixFirstLineSymbol:      l.LineSuffixFirstLineSymbol,
		LineSuffixLastLineSymbol:       l.LineSuffixLastLineSymbol,
		LineSuffixMiddleLineSymbol:     l.LineSuffixMiddleLineSymbol,
		LinePrefixLastLineSymbol:       l.LinePrefixLastLineSymbol,
		LinePrefixMiddleLineSymbol:     l.LinePrefixMiddleLineSymbol,
		ElementPrefixLinePrefixSymbol:  l.ElementPrefixLinePrefixSymbol,
	}
}

func (l *Leaf) Draw(writer io.Writer,
	indentLevel int, maxWidth int,
	isFirstLine bool, isLastLine bool,
	isLastSibling bool) {
	var writtenLen int

	// 绘制行首的结构符号
	if indentLevel > 0 {
		if isLastLine {
			symbolString := l.LinePrefixLastLineSymbol.RepeatSymbol(1)
			writtenLen += utf8.RuneCountInString(symbolString)
			data := []byte(symbolString)
			writer.Write(data)
		} else {
			symbolString := l.LinePrefixMiddleLineSymbol.RepeatSymbol(1)
			writtenLen += utf8.RuneCountInString(symbolString)
			data := []byte(symbolString)
			writer.Write(data)
		}
	}

	// 绘制元素前到行首的结构符号
	if indentLevel > 1 {
		symbolString := l.ElementPrefixLinePrefixSymbol.RepeatSymbol(1)
		writtenLen += utf8.RuneCountInString(symbolString)
		data := []byte(symbolString)
		writer.Write(data)
	}

	// 绘制元素前的结构符号
	switch {
	case isFirstLine:
		symbolString := l.ElementPrefixFirstLineSymbol.RepeatSymbol(1)
		writtenLen += utf8.RuneCountInString(symbolString)
		data := []byte(symbolString)
		writer.Write(data)
	case isLastLine:
		symbolString := l.ElementPrefixLastLineSymbol.RepeatSymbol(1)
		writtenLen += utf8.RuneCountInString(symbolString)
		data := []byte(symbolString)
		writer.Write(data)
	case isLastSibling:
		symbolString := l.ElementPrefixSiblingLastSymbol.RepeatSymbol(1)
		writtenLen += utf8.RuneCountInString(symbolString)
		data := []byte(symbolString)
		writer.Write(data)
	default:
		symbolString := l.ElementPrefixSiblingSymbol.RepeatSymbol(1)
		writtenLen += utf8.RuneCountInString(symbolString)
		data := []byte(symbolString)
		writer.Write(data)
	}

	// 绘制图标
	symbolString := l.Icon.GetSymbol()
	writtenLen += utf8.RuneCountInString(symbolString)
	data := []byte(symbolString)
	writer.Write(data)

	// 绘制名称
	symbolString = l.Name
	writtenLen += utf8.RuneCountInString(symbolString)
	data = []byte(symbolString)
	writer.Write(data)

	// 绘制元素后到行末的结构符号
	for writtenLen%3 != 0 {
		r, _ := utf8.DecodeRuneInString(l.ElementSuffixSymbol.RepeatSymbol(1))
		buf := make([]byte, 0, utf8.UTFMax)
		buf = append(buf, (byte)(r))
		writer.Write(buf)
		writtenLen++
	}


	// log.Println("writtenLen", writtenLen, "maxWidth", maxWidth)

	symbolString = l.ElementSuffixSymbol.RepeatSymbol((maxWidth - writtenLen)/3)
	writtenLen += utf8.RuneCountInString(symbolString)
	data = []byte(symbolString)
	writer.Write(data)

	// 绘制行末的结构符号
	switch {
	case isFirstLine:
		symbolString := l.LineSuffixFirstLineSymbol.RepeatSymbol(1)
		writtenLen += utf8.RuneCountInString(symbolString)
		data := []byte(symbolString)
		writer.Write(data)
	case isLastLine:
		symbolString := l.LineSuffixLastLineSymbol.RepeatSymbol(1)
		writtenLen += utf8.RuneCountInString(symbolString)
		data := []byte(symbolString)
		writer.Write(data)
	default:
		symbolString := l.LineSuffixMiddleLineSymbol.RepeatSymbol(1)
		writtenLen += utf8.RuneCountInString(symbolString)
		data := []byte(symbolString)
		writer.Write(data)
	}

	// log.Println("writtenLen", writtenLen, "maxWidth", maxWidth)

	// 绘制换行符
	writer.Write([]byte("\n"))

	// 更新节点编号
	curNumber++
}
