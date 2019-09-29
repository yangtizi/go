package sysutils

import (
	uuid "github.com/satori/go.uuid"
)

// GenUUID 获取uuid
func GenUUID() string {
	guid := uuid.NewV4()
	strGUID := guid.String()
	strGUID = StringReplaceAll(strGUID, "-", "")
	strGUID = UpperCase(strGUID)
	return strGUID
}

// GenGUID 生成GUID
func GenGUID() string {
	guid := uuid.NewV4()
	strGUID := guid.String()
	strGUID = UpperCase(strGUID)
	return strGUID
}
