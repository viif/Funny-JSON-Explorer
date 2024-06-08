package icon

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// 创建者：IconHandler
type IconHandler interface {
	GetSymbol(string) (string, error) // 业务逻辑
	CreateIcon(string) (Icon, error)  // 工厂方法
}

type CreateIconError struct {
	IconFamilyName string
}

func (e *CreateIconError) Error() string {
	return "Icon family not found: " + e.IconFamilyName
}

func GetIconHandler(iconType string) (IconHandler, error) {
	switch iconType {
	case "container":
		return &ContainerIconHandler{}, nil
	case "leaf":
		return &LeafIconHandler{}, nil
	default:
		return nil, &CreateIconError{iconType}
	}
}

// 具体创建者：中间节点图标 handler
type ContainerIconHandler struct{}

type ContainerIconConfig map[string]struct {
	ContainerIcon string `json:"container-icon"`
}

func (cih *ContainerIconHandler) CreateIcon(iconFamilyName string) (Icon, error) {
	// 打开配置文件
	file, err := os.Open("icon_family.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 读取文件内容
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// 解析 json 文件
	var iconConfig ContainerIconConfig
	err = json.Unmarshal(contents, &iconConfig)
	if err != nil {
		return nil, err
	}

	// 根据图标族名称获取配置
	familyConfig, exists := iconConfig[iconFamilyName]
	if !exists {
		return nil, &CreateIconError{iconFamilyName}
	}

	// 创建中间节点图标
	return &ContainerIcon{familyConfig.ContainerIcon}, nil
}

func (cih *ContainerIconHandler) GetSymbol(iconFamilyName string) (string, error) {
	icon, err := cih.CreateIcon(iconFamilyName)
	if err != nil {
		return "", err
	}
	return icon.GetSymbol(), nil
}

// 具体创建者：叶子节点图标 handler
type LeafIconHandler struct{}

type LeafIconConfig map[string]struct {
	LeafIcon string `json:"leaf-icon"`
}

func (lih LeafIconHandler) CreateIcon(iconFamilyName string) (Icon, error) {
	// 打开配置文件
	file, err := os.Open("icon_family.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 读取文件内容
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// 解析 json 文件
	var iconConfig LeafIconConfig
	err = json.Unmarshal(contents, &iconConfig)
	if err != nil {
		return nil, err
	}

	// 根据图标族名称获取配置
	familyConfig, exists := iconConfig[iconFamilyName]
	if !exists {
		return nil, &CreateIconError{iconFamilyName}
	}

	// 创建叶子节点图标
	return &LeafIcon{familyConfig.LeafIcon}, nil
}

func (lih LeafIconHandler) GetSymbol(iconFamilyName string) (string, error) {
	icon, err := lih.CreateIcon(iconFamilyName)
	if err != nil {
		return "", err
	}
	return icon.GetSymbol(), nil
}
