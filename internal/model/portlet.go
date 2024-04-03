package model

import "time"

// TableName 设置表名
func (Portlet) TableName() string {
	return "base_portlet"
}

type Portlet struct {
	Id          string `json:"id"`
	PortletId   string `json:"portletId"`
	PortletName string `json:"portletName"`
	ParentId    string `json:"parentId"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Component   string `json:"component"`
}

type PortletBase struct {
	Portlet
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type PortletTreeNode struct {
	Portlet
	Children []*PortletTreeNode `json:"children"`
}
