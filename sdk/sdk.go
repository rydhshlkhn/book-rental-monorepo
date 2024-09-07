// Code generated by candi v1.17.15. DO NOT EDIT.

package sdk

import (
	"sync"

	// @candi:serviceImport
	"monorepo/sdk/auth-service"
	"monorepo/sdk/user-service"
)

// Option func type
type Option func(*sdkInstance)

var (
	sdk  SDK
	once sync.Once
)

// SetGlobalSDK constructor with each sdk service option.
func SetGlobalSDK(opts ...Option) {
	s := new(sdkInstance)
	for _, o := range opts {
		o(s)
	}
	once.Do(func() {
		sdk = s
	})
}

// GetSDK get global sdk instance
func GetSDK() SDK {
	return sdk
}

// @candi:construct

// SetAuthservice option func
func SetAuthservice(authservice authservice.Authservice) Option {
	return func(s *sdkInstance) {
		s.authservice = authservice
	}
}

// SetUserservice option func
func SetUserservice(userservice userservice.Userservice) Option {
	return func(s *sdkInstance) {
		s.userservice = userservice
	}
}

// SDK instance abstraction
type SDK interface {
	// @candi:serviceMethod
	Authservice() authservice.Authservice
	Userservice() userservice.Userservice
}

// sdkInstance implementation
type sdkInstance struct {
	// @candi:serviceField
	authservice	authservice.Authservice
	userservice	userservice.Userservice
}

// @candi:instanceMethod
func (s *sdkInstance) Authservice() authservice.Authservice {
	return s.authservice
}
func (s *sdkInstance) Userservice() userservice.Userservice {
	return s.userservice
}
