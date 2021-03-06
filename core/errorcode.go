/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package core

const (
	CodeOK int = iota
	CodeDbConnectError
	CodeSqlError
	CodeInvaldParams
	CodeSystemError
	CodeNoUser
	CodeErrorPassword
	CodeUserExist
	CodeNoBook
)

var errormap = map[int]string{
	CodeOK:             "OK",
	CodeDbConnectError: "db connect error",
	CodeSqlError:       "db sql error",
	CodeInvaldParams:   "error params",
	CodeSystemError:    "system error",
	CodeNoUser:         "no user",
	CodeErrorPassword:  "pass error",
	CodeUserExist:      "user exist",
	CodeNoBook:         "book no exist",
}

func GetMsg(errorcode int) string {
	return errormap[errorcode]
}
