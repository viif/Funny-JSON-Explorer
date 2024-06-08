package node

import (
	"io"
)

// 组件：节点
type Node interface {
	Draw(writer io.Writer,
		indentLevel int, maxWidth int,
		isFirstLine bool, isLastLine bool, 
		isLastSibling bool)
}

// 节点总数
var NodeCount int

// 当前节点编号
var curNumber int