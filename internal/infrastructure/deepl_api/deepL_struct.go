package deeplapi

import (
	"fmt"
	"translate-remember-repeat/internal/domain"
)

type deepLStruct struct {
	defaultLink   string
	deepLAuthKey  string
	contentType   string
	requestMethod string
}

func New(
	deepLAuthKey string,
) *deepLStruct {
	defautLink := domain.DefaultLink
	contentType := domain.ContentType
	requestMethod := domain.RequestMethod
	defaultAuthKey := domain.HeaderKeyAuth
	return &deepLStruct{
		defaultLink:   defautLink,
		deepLAuthKey:  fmt.Sprintf("%s %s", defaultAuthKey, deepLAuthKey),
		contentType:   contentType,
		requestMethod: requestMethod,
	}
}
