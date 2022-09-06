package utils

import "strings"

type StringUtils interface {
	ConvertStringToLowerCase(param string) string
}

type StringUtilsImpl struct {
}

func (u StringUtilsImpl) ConvertStringToLowerCase(param string) string {
	return strings.ToLower(param)
}
