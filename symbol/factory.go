package symbol

import "errors"

// 抽象工厂：结构符号工厂
type SymbolFactory interface {
	CreateElementPrefixFirstLineSymbol() Symbol
	CreateElementPrefixLastLineSymbol() Symbol
	CreateElementPrefixSiblingLastSymbol() Symbol
	CreateElementPrefixSiblingSymbol() Symbol
	CreateElementSuffixSymbol() Symbol
	CreateLineSuffixFirstLineSymbol() Symbol
	CreateLineSuffixLastLineSymbol() Symbol
	CreateLineSuffixMiddleLineSymbol() Symbol
	CreateLinePrefixLastLineSymbol() Symbol
	CreateLinePrefixMiddleLineSymbol() Symbol
	CreateElementPrefixLinePrefixSymbol() Symbol
}

func GetSymbolFactory(symbolType string) (SymbolFactory, error) {
	switch symbolType {
	case "tree":
		return &TreeSymbolFactory{}, nil
	case "rectangle":
		return &RectangleSymbolFactory{}, nil
	default:
		return nil, errors.New("unsupported symbol type")
	}
}

// 具体工厂：树形结构符号工厂
type TreeSymbolFactory struct {
}

// 创建元素前的树形结构符号（位于首行）
func (f *TreeSymbolFactory) CreateElementPrefixFirstLineSymbol() Symbol {
	return &ElementPrefixFirstLineSymbol{symbol: "├─"}
}

// 创建元素前的树形结构符号（位于末行）
func (f *TreeSymbolFactory) CreateElementPrefixLastLineSymbol() Symbol {
	return &ElementPrefixLastLineSymbol{symbol: "└─"}
}

// 创建元素前的树形结构符号（位于同等级最后一个元素前）
func (f *TreeSymbolFactory) CreateElementPrefixSiblingLastSymbol() Symbol {
	return &ElementPrefixSiblingLastSymbol{symbol: "└─"}
}

// 创建元素前的树形结构符号（位于同等级非最后一个元素前）
func (f *TreeSymbolFactory) CreateElementPrefixSiblingSymbol() Symbol {
	return &ElementPrefixSiblingSymbol{symbol: "├─"}
}

// 创建元素后到行末的树形结构符号
func (f *TreeSymbolFactory) CreateElementSuffixSymbol() Symbol {
	return &ElementSuffixSymbol{symbol: "   "}
}

// 创建行末的树形结构符号（位于首行）
func (f *TreeSymbolFactory) CreateLineSuffixFirstLineSymbol() Symbol {
	return &LineSuffixFirstLineSymbol{symbol: "   "}
}

// 创建行末的树形结构符号（位于末行）
func (f *TreeSymbolFactory) CreateLineSuffixLastLineSymbol() Symbol {
	return &LineSuffixLastLineSymbol{symbol: "   "}
}

// 创建行末的树形结构符号（位于中间行）
func (f *TreeSymbolFactory) CreateLineSuffixMiddleLineSymbol() Symbol {
	return &LineSuffixMiddleLineSymbol{symbol: "   "}
}

// 创建行首的树形结构符号（位于末行）
func (f *TreeSymbolFactory) CreateLinePrefixLastLineSymbol() Symbol {
	return &LinePrefixLastLineSymbol{symbol: "   "}
}

// 创建行首的树形结构符号（位于中间行）
func (f *TreeSymbolFactory) CreateLinePrefixMiddleLineSymbol() Symbol {
	return &LinePrefixMiddleLineSymbol{symbol: "│  "}
}

// 创建元素前到行首的树形结构符号
func (f *TreeSymbolFactory) CreateElementPrefixLinePrefixSymbol() Symbol {
	return &ElementPrefixLinePrefixSymbol{symbol: "   "}
}

// 具体工厂：矩形结构符号工厂
type RectangleSymbolFactory struct {
}

// 创建元素前的矩形结构符号（位于首行）
func (f *RectangleSymbolFactory) CreateElementPrefixFirstLineSymbol() Symbol {
	return &ElementPrefixFirstLineSymbol{symbol: "┌─"}
}

// 创建元素前的矩形结构符号（位于末行）
func (f *RectangleSymbolFactory) CreateElementPrefixLastLineSymbol() Symbol {
	return &ElementPrefixLastLineSymbol{symbol: "└─"}
}

// 创建元素前的矩形结构符号（位于同等级最后一个元素前）
func (f *RectangleSymbolFactory) CreateElementPrefixSiblingLastSymbol() Symbol {
	return &ElementPrefixSiblingLastSymbol{symbol: "├─"}
}

// 创建元素前的矩形结构符号（位于同等级非最后一个元素前）
func (f *RectangleSymbolFactory) CreateElementPrefixSiblingSymbol() Symbol {
	return &ElementPrefixSiblingSymbol{symbol: "├─"}
}

// 创建元素后到行末的矩形结构符号
func (f *RectangleSymbolFactory) CreateElementSuffixSymbol() Symbol {
	return &ElementSuffixSymbol{symbol: "───"}
}

// 创建行末的矩形结构符号（位于首行）
func (f *RectangleSymbolFactory) CreateLineSuffixFirstLineSymbol() Symbol {
	return &LineSuffixFirstLineSymbol{symbol: "──┐"}
}

// 创建行末的矩形结构符号（位于末行）
func (f *RectangleSymbolFactory) CreateLineSuffixLastLineSymbol() Symbol {
	return &LineSuffixLastLineSymbol{symbol: "──┘"}
}

// 创建行末的矩形结构符号（位于中间行）
func (f *RectangleSymbolFactory) CreateLineSuffixMiddleLineSymbol() Symbol {
	return &LineSuffixMiddleLineSymbol{symbol: "──┤"}
}

// 创建行首的矩形结构符号（位于末行）
func (f *RectangleSymbolFactory) CreateLinePrefixLastLineSymbol() Symbol {
	return &LinePrefixLastLineSymbol{symbol: "└──"}
}

// 创建行首的矩形结构符号（位于中间行）
func (f *RectangleSymbolFactory) CreateLinePrefixMiddleLineSymbol() Symbol {
	return &LinePrefixMiddleLineSymbol{symbol: "│  "}
}

// 创建元素前到行首的矩形结构符号
func (f *RectangleSymbolFactory) CreateElementPrefixLinePrefixSymbol() Symbol {
	return &ElementPrefixLinePrefixSymbol{symbol: "│  "}
}
