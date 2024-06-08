package node

import (
	"FJE/icon"
	"FJE/symbol"

	"io"
	"log"
	"unicode/utf8"
)

var _ = log.Println

// 中间节点
type Container struct {
	Children                       []Node        // 子节点
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

func (c *Container) AddNode(node Node) {
	c.Children = append(c.Children, node)
}

func (c *Container) Clone() Node {
	return &Container{
		Children:                       c.Children,
		Name:                           c.Name,
		Icon:                           c.Icon,
		ElementPrefixFirstLineSymbol:   c.ElementPrefixFirstLineSymbol,
		ElementPrefixLastLineSymbol:    c.ElementPrefixLastLineSymbol,
		ElementPrefixSiblingLastSymbol: c.ElementPrefixSiblingLastSymbol,
		ElementPrefixSiblingSymbol:     c.ElementPrefixSiblingSymbol,
		ElementSuffixSymbol:            c.ElementSuffixSymbol,
		LineSuffixFirstLineSymbol:      c.LineSuffixFirstLineSymbol,
		LineSuffixLastLineSymbol:       c.LineSuffixLastLineSymbol,
		LineSuffixMiddleLineSymbol:     c.LineSuffixMiddleLineSymbol,
		LinePrefixLastLineSymbol:       c.LinePrefixLastLineSymbol,
		LinePrefixMiddleLineSymbol:     c.LinePrefixMiddleLineSymbol,
		ElementPrefixLinePrefixSymbol:  c.ElementPrefixLinePrefixSymbol,
	}
}

func (c *Container) Draw(writer io.Writer,
	indentLevel int, maxWidth int,
	isFirstLine bool, isLastLine bool,
	isLastSibling bool) {
	var writtenLen int

	// 绘制行首的结构符号
	if indentLevel > 0 {
		if isLastLine {
			symbolString := c.LinePrefixLastLineSymbol.RepeatSymbol(1)
			writtenLen += utf8.RuneCountInString(symbolString)
			data := []byte(symbolString)
			writer.Write(data)
		} else {
			symbolString := c.LinePrefixMiddleLineSymbol.RepeatSymbol(1)
			writtenLen += utf8.RuneCountInString(symbolString)
			data := []byte(symbolString)
			writer.Write(data)
		}
	}

	// 绘制元素前到行首的结构符号
	if indentLevel > 1 {
		symbolString := c.ElementPrefixLinePrefixSymbol.RepeatSymbol(1)
		writtenLen += utf8.RuneCountInString(symbolString)
		data := []byte(symbolString)
		writer.Write(data)
	}

	// 绘制元素前的结构符号
	switch {
	case isFirstLine:
		symbolString := c.ElementPrefixFirstLineSymbol.RepeatSymbol(1)
		writtenLen += utf8.RuneCountInString(symbolString)
		data := []byte(symbolString)
		writer.Write(data)
	case isLastLine:
		symbolString := c.ElementPrefixLastLineSymbol.RepeatSymbol(1)
		writtenLen += utf8.RuneCountInString(symbolString)
		data := []byte(symbolString)
		writer.Write(data)
	case isLastSibling:
		symbolString := c.ElementPrefixSiblingLastSymbol.RepeatSymbol(1)
		writtenLen += utf8.RuneCountInString(symbolString)
		data := []byte(symbolString)
		writer.Write(data)
	default:
		symbolString := c.ElementPrefixSiblingSymbol.RepeatSymbol(1)
		writtenLen += utf8.RuneCountInString(symbolString)
		data := []byte(symbolString)
		writer.Write(data)
	}

	// 绘制图标
	symbolString := c.Icon.GetSymbol()
	writtenLen += utf8.RuneCountInString(symbolString)
	data := []byte(symbolString)
	writer.Write(data)

	// 绘制名称
	symbolString = c.Name
	writtenLen += utf8.RuneCountInString(symbolString)
	data = []byte(symbolString)
	writer.Write(data)

	// 绘制元素后到行末的结构符号
	for writtenLen%3 != 0 {
		r, _ := utf8.DecodeRuneInString(c.ElementSuffixSymbol.RepeatSymbol(1))
		buf := make([]byte, 0, utf8.UTFMax)
		buf = append(buf, (byte)(r))
		writer.Write(buf)
		writtenLen++
	}

	// log.Println("writtenLen", writtenLen, "maxWidth", maxWidth)

	symbolString = c.ElementSuffixSymbol.RepeatSymbol((maxWidth - writtenLen) / 3)
	data = []byte(symbolString)
	writer.Write(data)

	// 绘制行末的结构符号
	switch {
	case isFirstLine:
		data = []byte(c.LineSuffixFirstLineSymbol.RepeatSymbol(1))
		writer.Write(data)
	case isLastLine:
		data = []byte(c.LineSuffixLastLineSymbol.RepeatSymbol(1))
		writer.Write(data)
	default:
		data = []byte(c.LineSuffixMiddleLineSymbol.RepeatSymbol(1))
		writer.Write(data)
	}

	// 绘制换行符
	writer.Write([]byte("\n"))

	// 递归绘制子节点
	curNumber++
	for i, child := range c.Children {
		child.Draw(writer, indentLevel+1, maxWidth,
			false, curNumber == NodeCount-2, i == len(c.Children)-1)
	}
}
