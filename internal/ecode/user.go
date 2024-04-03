package ecode

import "thelastcookies/api-base/pkg/errno"

var (
	ErrUserNotFound         = errno.NewError(20301, "未找到指定的用户")
	ErrInvalidUserId        = errno.NewError(20302, "非法的用户ID")
	ErrDuplicateUserId      = errno.NewError(20303, "重复的用户ID")
	ErrUserParams           = errno.NewError(20304, "非法的用户信息参数")
	ErrUserPortletsNotFound = errno.NewError(20305, "未查询到当前用户的Portlet配置")
)
