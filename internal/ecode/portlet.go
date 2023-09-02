package ecode

import "tlc.platform/web-service/pkg/errno"

var (
	ErrPortletNotFound    = errno.NewError(20101, "未找到指定的门户组件")
	ErrInvalidPortletId   = errno.NewError(20102, "非法的门户组件ID")
	ErrDuplicatePortletId = errno.NewError(20103, "重复的门户组件ID")
	ErrPortletParams      = errno.NewError(20104, "非法的门户组件参数")
)
