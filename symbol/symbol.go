package symbol

import (
	"strings"
)

// 抽象产品：可视化结构符号
type Symbol interface {
	RepeatSymbol(count int) string
}

// 具体产品：元素前的结构符号（位于首行）
type ElementPrefixFirstLineSymbol struct {
	symbol string
}

func (s *ElementPrefixFirstLineSymbol) RepeatSymbol(count int) string {
	return strings.Repeat(s.symbol, count)
}

// 具体产品：元素前的结构符号（位于末行）
type ElementPrefixLastLineSymbol struct {
	symbol string
}

func (s *ElementPrefixLastLineSymbol) RepeatSymbol(count int) string {
	return strings.Repeat(s.symbol, count)
}

// 具体产品：元素前的结构符号（位于同等级最后一个元素前）
type ElementPrefixSiblingLastSymbol struct {
	symbol string
}

func (s *ElementPrefixSiblingLastSymbol) RepeatSymbol(count int) string {
	return strings.Repeat(s.symbol, count)
}

// 具体产品：元素前的结构符号（位于同等级非最后一个元素前）
type ElementPrefixSiblingSymbol struct {
	symbol string
}

func (s *ElementPrefixSiblingSymbol) RepeatSymbol(count int) string {
	return strings.Repeat(s.symbol, count)
}

// 具体产品：元素后到行末的结构符号
type ElementSuffixSymbol struct {
	symbol string
}

func (s *ElementSuffixSymbol) RepeatSymbol(count int) string {
	return strings.Repeat(s.symbol, count)
}

// 具体产品：行末的结构符号（位于首行）
type LineSuffixFirstLineSymbol struct {
	symbol string
}

func (s *LineSuffixFirstLineSymbol) RepeatSymbol(count int) string {
	return strings.Repeat(s.symbol, count)
}

// 具体产品：行末的结构符号（位于末行）
type LineSuffixLastLineSymbol struct {
	symbol string
}

func (s *LineSuffixLastLineSymbol) RepeatSymbol(count int) string {
	return strings.Repeat(s.symbol, count)
}

// 具体产品：行末的结构符号（位于中间行）
type LineSuffixMiddleLineSymbol struct {
	symbol string
}

func (s *LineSuffixMiddleLineSymbol) RepeatSymbol(count int) string {
	return strings.Repeat(s.symbol, count)
}

// 具体产品：行首的结构符号（位于末行）
type LinePrefixLastLineSymbol struct {
	symbol string
}

func (s *LinePrefixLastLineSymbol) RepeatSymbol(count int) string {
	return strings.Repeat(s.symbol, count)
}

// 具体产品：行首的结构符号（位于中间行）
type LinePrefixMiddleLineSymbol struct {
	symbol string
}

func (s *LinePrefixMiddleLineSymbol) RepeatSymbol(count int) string {
	return strings.Repeat(s.symbol, count)
}

// 具体产品：元素前到行首的结构符号
type ElementPrefixLinePrefixSymbol struct {
	symbol string
}

func (s *ElementPrefixLinePrefixSymbol) RepeatSymbol(count int) string {
	return strings.Repeat(s.symbol, count)
}
