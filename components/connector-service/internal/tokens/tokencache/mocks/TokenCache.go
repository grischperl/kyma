// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// TokenCache is an autogenerated mock type for the TokenCache type
type TokenCache struct {
	mock.Mock
}

// Delete provides a mock function with given fields: token
func (_m *TokenCache) Delete(token string) {
	_m.Called(token)
}

// Get provides a mock function with given fields: token
func (_m *TokenCache) Get(token string) (string, bool) {
	ret := _m.Called(token)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Put provides a mock function with given fields: token, data, ttl
func (_m *TokenCache) Put(token string, data string, ttl time.Duration) {
	_m.Called(token, data, ttl)
}