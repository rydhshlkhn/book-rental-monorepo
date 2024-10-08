// Code generated by candi v1.17.15. DO NOT EDIT.

package sdk

import (
	"sync"

	// @candi:serviceImport
	"monorepo/sdk/payment-service"
	"monorepo/sdk/library-service"
	"monorepo/sdk/book-service"
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

// SetPaymentservice option func
func SetPaymentservice(paymentservice paymentservice.Paymentservice) Option {
	return func(s *sdkInstance) {
		s.paymentservice = paymentservice
	}
}

// SetLibraryservice option func
func SetLibraryservice(libraryservice libraryservice.Libraryservice) Option {
	return func(s *sdkInstance) {
		s.libraryservice = libraryservice
	}
}

// SetBookservice option func
func SetBookservice(bookservice bookservice.Bookservice) Option {
	return func(s *sdkInstance) {
		s.bookservice = bookservice
	}
}

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
	Paymentservice() paymentservice.Paymentservice
	Libraryservice() libraryservice.Libraryservice
	Bookservice() bookservice.Bookservice
	Authservice() authservice.Authservice
	Userservice() userservice.Userservice
}

// sdkInstance implementation
type sdkInstance struct {
	// @candi:serviceField
	paymentservice	paymentservice.Paymentservice
	libraryservice	libraryservice.Libraryservice
	bookservice	bookservice.Bookservice
	authservice	authservice.Authservice
	userservice	userservice.Userservice
}

// @candi:instanceMethod
func (s *sdkInstance) Paymentservice() paymentservice.Paymentservice {
	return s.paymentservice
}
func (s *sdkInstance) Libraryservice() libraryservice.Libraryservice {
	return s.libraryservice
}
func (s *sdkInstance) Bookservice() bookservice.Bookservice {
	return s.bookservice
}
func (s *sdkInstance) Authservice() authservice.Authservice {
	return s.authservice
}
func (s *sdkInstance) Userservice() userservice.Userservice {
	return s.userservice
}
