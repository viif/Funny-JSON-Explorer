package style

import (
	"FJE/node"
	"FJE/symbol"
	"errors"
)

// 生成器：风格生成器
type StyleBuilder interface {
	SetSymbolFactory() error
	SetElementPrefixFirstLineSymbol()
	SetElementPrefixLastLineSymbol()
	SetElementPrefixSiblingLastSymbol()
	SetElementPrefixSiblingSymbol()
	SetElementSuffixSymbol()
	SetLineSuffixFirstLineSymbol()
	SetLineSuffixLastLineSymbol()
	SetLineSuffixMiddleLineSymbol()
	SetLinePrefixLastLineSymbol()
	SetLinePrefixMiddleLineSymbol()
	SetElementPrefixLinePrefixSymbol()
	GetContainerNode() node.Container
	GetLeafNode() node.Leaf
}

func GetStyleBuilder(styleType string) (StyleBuilder, error) {
	switch styleType {
	case "tree":
		return &TreeStyleBuilder{}, nil
	case "rectangle":
		return &RectangleStyleBuilder{}, nil
	default:
		return nil, errors.New("unsupported style type")
	}
}

// 产品：具体风格的节点

// 具体生成器：树形风格生成器
type TreeStyleBuilder struct {
	treeSymbolFactory symbol.SymbolFactory
	ContainerNode     node.Container
	LeafNode          node.Leaf
}

func (f *TreeStyleBuilder) SetSymbolFactory() (err error) {
	f.treeSymbolFactory, err = symbol.GetSymbolFactory("tree")
	return err
}

func (f *TreeStyleBuilder) SetElementPrefixFirstLineSymbol() {
	f.ContainerNode.ElementPrefixFirstLineSymbol = f.treeSymbolFactory.CreateElementPrefixFirstLineSymbol()
	f.LeafNode.ElementPrefixFirstLineSymbol = f.treeSymbolFactory.CreateElementPrefixFirstLineSymbol()
}

func (f *TreeStyleBuilder) SetElementPrefixLastLineSymbol() {
	f.ContainerNode.ElementPrefixLastLineSymbol = f.treeSymbolFactory.CreateElementPrefixLastLineSymbol()
	f.LeafNode.ElementPrefixLastLineSymbol = f.treeSymbolFactory.CreateElementPrefixLastLineSymbol()
}

func (f *TreeStyleBuilder) SetElementPrefixSiblingLastSymbol() {
	f.ContainerNode.ElementPrefixSiblingLastSymbol = f.treeSymbolFactory.CreateElementPrefixSiblingLastSymbol()
	f.LeafNode.ElementPrefixSiblingLastSymbol = f.treeSymbolFactory.CreateElementPrefixSiblingLastSymbol()
}

func (f *TreeStyleBuilder) SetElementPrefixSiblingSymbol() {
	f.ContainerNode.ElementPrefixSiblingSymbol = f.treeSymbolFactory.CreateElementPrefixSiblingSymbol()
	f.LeafNode.ElementPrefixSiblingSymbol = f.treeSymbolFactory.CreateElementPrefixSiblingSymbol()
}

func (f *TreeStyleBuilder) SetElementSuffixSymbol() {
	f.ContainerNode.ElementSuffixSymbol = f.treeSymbolFactory.CreateElementSuffixSymbol()
	f.LeafNode.ElementSuffixSymbol = f.treeSymbolFactory.CreateElementSuffixSymbol()
}

func (f *TreeStyleBuilder) SetLineSuffixFirstLineSymbol() {
	f.ContainerNode.LineSuffixFirstLineSymbol = f.treeSymbolFactory.CreateLineSuffixFirstLineSymbol()
	f.LeafNode.LineSuffixFirstLineSymbol = f.treeSymbolFactory.CreateLineSuffixFirstLineSymbol()
}

func (f *TreeStyleBuilder) SetLineSuffixLastLineSymbol() {
	f.ContainerNode.LineSuffixLastLineSymbol = f.treeSymbolFactory.CreateLineSuffixLastLineSymbol()
	f.LeafNode.LineSuffixLastLineSymbol = f.treeSymbolFactory.CreateLineSuffixLastLineSymbol()
}

func (f *TreeStyleBuilder) SetLineSuffixMiddleLineSymbol() {
	f.ContainerNode.LineSuffixMiddleLineSymbol = f.treeSymbolFactory.CreateLineSuffixMiddleLineSymbol()
	f.LeafNode.LineSuffixMiddleLineSymbol = f.treeSymbolFactory.CreateLineSuffixMiddleLineSymbol()
}

func (f *TreeStyleBuilder) SetLinePrefixLastLineSymbol() {
	f.ContainerNode.LinePrefixLastLineSymbol = f.treeSymbolFactory.CreateLinePrefixLastLineSymbol()
	f.LeafNode.LinePrefixLastLineSymbol = f.treeSymbolFactory.CreateLinePrefixLastLineSymbol()
}

func (f *TreeStyleBuilder) SetLinePrefixMiddleLineSymbol() {
	f.ContainerNode.LinePrefixMiddleLineSymbol = f.treeSymbolFactory.CreateLinePrefixMiddleLineSymbol()
	f.LeafNode.LinePrefixMiddleLineSymbol = f.treeSymbolFactory.CreateLinePrefixMiddleLineSymbol()
}

func (f *TreeStyleBuilder) SetElementPrefixLinePrefixSymbol() {
	f.ContainerNode.ElementPrefixLinePrefixSymbol = f.treeSymbolFactory.CreateElementPrefixLinePrefixSymbol()
	f.LeafNode.ElementPrefixLinePrefixSymbol = f.treeSymbolFactory.CreateElementPrefixLinePrefixSymbol()
}

func (f *TreeStyleBuilder) GetContainerNode() node.Container {
	return f.ContainerNode
}

func (f *TreeStyleBuilder) GetLeafNode() node.Leaf {
	return f.LeafNode
}

// 具体生成器：矩形风格生成器
type RectangleStyleBuilder struct {
	rectangleSymbolFactory symbol.SymbolFactory
	ContainerNode          node.Container
	LeafNode               node.Leaf
}

func (f *RectangleStyleBuilder) SetSymbolFactory() (err error) {
	f.rectangleSymbolFactory, err = symbol.GetSymbolFactory("rectangle")
	return err
}

func (f *RectangleStyleBuilder) SetElementPrefixFirstLineSymbol() {
	f.ContainerNode.ElementPrefixFirstLineSymbol = f.rectangleSymbolFactory.CreateElementPrefixFirstLineSymbol()
	f.LeafNode.ElementPrefixFirstLineSymbol = f.rectangleSymbolFactory.CreateElementPrefixFirstLineSymbol()
}

func (f *RectangleStyleBuilder) SetElementPrefixLastLineSymbol() {
	f.ContainerNode.ElementPrefixLastLineSymbol = f.rectangleSymbolFactory.CreateElementPrefixLastLineSymbol()
	f.LeafNode.ElementPrefixLastLineSymbol = f.rectangleSymbolFactory.CreateElementPrefixLastLineSymbol()
}

func (f *RectangleStyleBuilder) SetElementPrefixSiblingLastSymbol() {
	f.ContainerNode.ElementPrefixSiblingLastSymbol = f.rectangleSymbolFactory.CreateElementPrefixSiblingLastSymbol()
	f.LeafNode.ElementPrefixSiblingLastSymbol = f.rectangleSymbolFactory.CreateElementPrefixSiblingLastSymbol()
}

func (f *RectangleStyleBuilder) SetElementPrefixSiblingSymbol() {
	f.ContainerNode.ElementPrefixSiblingSymbol = f.rectangleSymbolFactory.CreateElementPrefixSiblingSymbol()
	f.LeafNode.ElementPrefixSiblingSymbol = f.rectangleSymbolFactory.CreateElementPrefixSiblingSymbol()
}

func (f *RectangleStyleBuilder) SetElementSuffixSymbol() {
	f.ContainerNode.ElementSuffixSymbol = f.rectangleSymbolFactory.CreateElementSuffixSymbol()
	f.LeafNode.ElementSuffixSymbol = f.rectangleSymbolFactory.CreateElementSuffixSymbol()
}

func (f *RectangleStyleBuilder) SetLineSuffixFirstLineSymbol() {
	f.ContainerNode.LineSuffixFirstLineSymbol = f.rectangleSymbolFactory.CreateLineSuffixFirstLineSymbol()
	f.LeafNode.LineSuffixFirstLineSymbol = f.rectangleSymbolFactory.CreateLineSuffixFirstLineSymbol()
}

func (f *RectangleStyleBuilder) SetLineSuffixLastLineSymbol() {
	f.ContainerNode.LineSuffixLastLineSymbol = f.rectangleSymbolFactory.CreateLineSuffixLastLineSymbol()
	f.LeafNode.LineSuffixLastLineSymbol = f.rectangleSymbolFactory.CreateLineSuffixLastLineSymbol()
}

func (f *RectangleStyleBuilder) SetLineSuffixMiddleLineSymbol() {
	f.ContainerNode.LineSuffixMiddleLineSymbol = f.rectangleSymbolFactory.CreateLineSuffixMiddleLineSymbol()
	f.LeafNode.LineSuffixMiddleLineSymbol = f.rectangleSymbolFactory.CreateLineSuffixMiddleLineSymbol()
}

func (f *RectangleStyleBuilder) SetLinePrefixLastLineSymbol() {
	f.ContainerNode.LinePrefixLastLineSymbol = f.rectangleSymbolFactory.CreateLinePrefixLastLineSymbol()
	f.LeafNode.LinePrefixLastLineSymbol = f.rectangleSymbolFactory.CreateLinePrefixLastLineSymbol()
}

func (f *RectangleStyleBuilder) SetLinePrefixMiddleLineSymbol() {
	f.ContainerNode.LinePrefixMiddleLineSymbol = f.rectangleSymbolFactory.CreateLinePrefixMiddleLineSymbol()
	f.LeafNode.LinePrefixMiddleLineSymbol = f.rectangleSymbolFactory.CreateLinePrefixMiddleLineSymbol()
}

func (f *RectangleStyleBuilder) SetElementPrefixLinePrefixSymbol() {
	f.ContainerNode.ElementPrefixLinePrefixSymbol = f.rectangleSymbolFactory.CreateElementPrefixLinePrefixSymbol()
	f.LeafNode.ElementPrefixLinePrefixSymbol = f.rectangleSymbolFactory.CreateElementPrefixLinePrefixSymbol()
}

func (f *RectangleStyleBuilder) GetContainerNode() node.Container {
	return f.ContainerNode
}

func (f *RectangleStyleBuilder) GetLeafNode() node.Leaf {
	return f.LeafNode
}
