package icon

// 产品：Icon 接口
type Icon interface {
	GetSymbol() string
}

// 具体产品：中间节点图标
type ContainerIcon struct {
	symbol string
}

func (ci *ContainerIcon) GetSymbol() string {
	return ci.symbol
}

// 具体产品：叶子节点图标
type LeafIcon struct {
	symbol string
}

func (li *LeafIcon) GetSymbol() string {
	return li.symbol
}
