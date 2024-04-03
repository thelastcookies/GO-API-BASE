package ecode

import "thelastcookies/api-base/pkg/errno"

var (
	ErrRoleNotFound         = errno.NewError(20201, "未找到指定的角色")
	ErrInvalidRoleId        = errno.NewError(20202, "非法的角色ID")
	ErrDuplicateRoleId      = errno.NewError(20203, "重复的角色ID")
	ErrRoleParams           = errno.NewError(20204, "非法的角色信息参数")
	ErrRolePortletsNotFound = errno.NewError(20205, "未查询到当前角色的Portlet配置")
)
