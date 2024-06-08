package explorer

import (
	"FJE/node"

	"encoding/json"
	"os"
	"log"
	"fmt"
)

// json 文件浏览器
type Explorer struct {
	FilePath         string
	ExampleContainer *node.Container
	ExampleLeaf      *node.Leaf

	rootContainer *node.Container
}

// 递归解析 JSON 数据，组织节点，并更新节点总数
func (e *Explorer)parseJSONData(key string, value interface{}) node.Node {
	node.NodeCount++
	switch v := value.(type) {
	case map[string]interface{}:
		container := e.ExampleContainer.Clone().(*node.Container)
		container.Name = key
		for k, val := range v {
			container.AddNode(e.parseJSONData(k, val))
		}
		return container
	case []interface{}:
		container := e.ExampleContainer.Clone().(*node.Container)
		container.Name = key
		for i, val := range v {
			container.AddNode(e.parseJSONData(fmt.Sprintf("%d", i), val))
		}
		return container
	default:
		leaf := e.ExampleLeaf.Clone().(*node.Leaf)
		if value == nil {
			leaf.Name = fmt.Sprintf("%v", key)
		} else {
		leaf.Name = fmt.Sprintf("%v:%v", key, value)
		}
		return leaf
	}
}

// 读取 JSON 文件并解析
func (e *Explorer) ReadJSONFile() {
	// 读取 JSON 文件
	jsonBytes, err := os.ReadFile(e.FilePath)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v\n", err)
		return
	}

	// 解析 JSON 数据到一个 map 中
	var data map[string]interface{}
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON data: %v\n", err)
		return
	}

	// 递归解析 JSON 数据，组织节点
	e.rootContainer = e.parseJSONData("root", data).(*node.Container)
}

// 绘制节点
func (e *Explorer) Show() {
	for i, n := range e.rootContainer.Children {
		n.Draw(os.Stdout, 0, 60, i == 0, false, true)
	}
}
