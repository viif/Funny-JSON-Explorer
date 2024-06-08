package main

import (
	"FJE/icon"
	"FJE/style"
	"FJE/explorer"

	"flag"
	"fmt"
	"log"
)

func main() {
	// 解析命令行参数
	var jsonFile string
	var myStyle string
	var iconFamily string

	flag.StringVar(&jsonFile, "f", "", "JSON file path")
	flag.StringVar(&myStyle, "s", "tree", "Style (tree, rectangle)")
	flag.StringVar(&iconFamily, "i", "default", "Icon family (default, poker-face)")
	flag.Parse()

	if jsonFile == "" {
		log.Fatal("JSON file path is required")
		return
	}

	fmt.Printf("JSON File: %s\n", jsonFile)
	fmt.Printf("Style: %s\n", myStyle)
	fmt.Printf("Icon Family: %s\n", iconFamily)
	fmt.Println()

	// 设置风格
	treeStyleBuilder, err := style.GetStyleBuilder(myStyle)
	if err != nil {
		log.Fatalf("Error getting style builder: %v\n", err)
		return
	}
	director := style.NewDirector(treeStyleBuilder)
	containerNode, leafNode, err := director.Construct()
	if err != nil {
		log.Fatalf("Error constructing: %v\n", err)
		return
	}

	// 设置图标
	containerIconHander, err := icon.GetIconHandler("container")
	if err != nil {
		log.Fatalf("Error getting container icon handler: %v\n", err)
		return
	}
	containerIcon, err := containerIconHander.CreateIcon(iconFamily)
	if err != nil {
		log.Fatalf("Error creating container icon: %v\n", err)
		return
	}
	containerNode.Icon = containerIcon
	leafIconHander, err := icon.GetIconHandler("leaf")
	if err != nil {
		log.Fatalf("Error getting leaf icon handler: %v\n", err)
		return
	}
	leafIcon, err := leafIconHander.CreateIcon(iconFamily)
	if err != nil {
		log.Fatalf("Error creating leaf icon: %v\n", err)
		return
	}
	leafNode.Icon = leafIcon

	// 读取 JSON 文件并解析
	explorer := explorer.Explorer{
		FilePath:         jsonFile,
		ExampleContainer: &containerNode,
		ExampleLeaf:      &leafNode,
	}
	explorer.ReadJSONFile()

	// 可视化
	explorer.Show()
	fmt.Println()
}
