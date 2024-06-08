package style

import (
	"FJE/node"
)

// 主管
type Director struct {
	builder StyleBuilder
}

func NewDirector(builder StyleBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) SetBuilder(builder StyleBuilder) {
	d.builder = builder
}

func (d *Director) Construct() (ContainerNode node.Container, LeafNode node.Leaf, err error) {
	if err = d.builder.SetSymbolFactory(); err != nil {
		return
	}
	d.builder.SetElementPrefixFirstLineSymbol()
	d.builder.SetElementPrefixLastLineSymbol()
	d.builder.SetElementPrefixSiblingLastSymbol()
	d.builder.SetElementPrefixSiblingSymbol()
	d.builder.SetElementSuffixSymbol()
	d.builder.SetLineSuffixFirstLineSymbol()
	d.builder.SetLineSuffixLastLineSymbol()
	d.builder.SetLineSuffixMiddleLineSymbol()
	d.builder.SetLinePrefixLastLineSymbol()
	d.builder.SetLinePrefixMiddleLineSymbol()
	d.builder.SetElementPrefixLinePrefixSymbol()
	return d.builder.GetContainerNode(), d.builder.GetLeafNode(), nil
}
