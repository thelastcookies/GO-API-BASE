package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"io"
	"math/rand"
	"os"
	"regexp"
	"time"
	"tlc.platform/web-service/internal/model"
)

// GetBytes interface 转 byte
func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Md5 字符串转 md5
func Md5(str string) (string, error) {
	h := md5.New()

	_, err := io.WriteString(h, str)
	if err != nil {
		return "", err
	}

	// 注意：这里不能使用string将[]byte转为字符串，否则会显示乱码
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// RandomStr 生成随机字符串
func RandomStr(n int) string {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	const pattern = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

	salt := make([]byte, 0, n)
	l := len(pattern)

	for i := 0; i < n; i++ {
		p := r.Intn(l)
		salt = append(salt, pattern[p])
	}

	return string(salt)
}

// RegexpReplace ...
func RegexpReplace(reg, src, temp string) string {
	result := []byte{}
	pattern := regexp.MustCompile(reg)
	for _, submatches := range pattern.FindAllStringSubmatchIndex(src, -1) {
		result = pattern.ExpandString(result, temp, src, submatches)
	}
	return string(result)
}

// GetHostname 获取主机名
func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		name = "unknown"
	}
	return name
}

// ListToTree 列表结构转树形结构
// 目前是 PortletTreeNode 特定
func ListToTree(list []*model.PortletTreeNode, parentId string) []*model.PortletTreeNode {
	if len(list) == 0 {
		return []*model.PortletTreeNode{}
	}
	var tree []*model.PortletTreeNode
	for _, node := range list {
		if node.Portlet.ParentId == parentId {
			children := ListToTree(list, node.Portlet.ID)
			if len(children) != 0 {
				node.Children = children
			}
			tree = append(tree, node)
		}
	}
	return tree
}

//func ListToTree(list []interface{}, parentId string) []interface{} {
//	if len(list) == 0 {
//		return []interface{}{}
//	}
//	nodeType := reflect.TypeOf(list[0])
//	tree := reflect.New(nodeType)
//	for _, node := range list {
//		if node.id == parentId {
//			children := ListToTree(list, node.id)
//			if len(children) != 0 {
//				node.children = children
//			}
//			nodeList = append(nodeList, node)
//		}
//	}
//	return tree
//}
