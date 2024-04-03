package model

import "time"

// TableName 设置表名
func (RolePortlet) TableName() string {
	return "base_role_portlet"
}

type RolePortlet struct {
	Id        string `json:"id"`
	RoleId    string `json:"roleId"`
	PortletId string `json:"portletId"`
}

type RolePortletBase struct {
	RolePortlet
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
