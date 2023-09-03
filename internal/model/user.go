package model

import "time"

// TableName 设置表名
func (UserPortlet) TableName() string {
	return "base_user_portlet"
}

type UserPortlet struct {
	ID        string `json:"id"`
	UserId    string `json:"userId"`
	PortletId string `json:"portletId"`
}

type UserPortletBase struct {
	UserPortlet
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
