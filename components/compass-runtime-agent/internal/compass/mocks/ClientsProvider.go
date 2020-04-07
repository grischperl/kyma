// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	connector "kyma-project.io/compass-runtime-agent/internal/compass/connector"
	config "kyma-project.io/compass-runtime-agent/internal/config"

	director "kyma-project.io/compass-runtime-agent/internal/compass/director"

	mock "github.com/stretchr/testify/mock"
)

// ClientsProvider is an autogenerated mock type for the ClientsProvider type
type ClientsProvider struct {
	mock.Mock
}

// GetConnectorCertSecuredClient provides a mock function with given fields:
func (_m *ClientsProvider) GetConnectorCertSecuredClient() (connector.Client, error) {
	ret := _m.Called()

	var r0 connector.Client
	if rf, ok := ret.Get(0).(func() connector.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(connector.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnectorTokensClient provides a mock function with given fields: url
func (_m *ClientsProvider) GetConnectorTokensClient(url string) (connector.Client, error) {
	ret := _m.Called(url)

	var r0 connector.Client
	if rf, ok := ret.Get(0).(func(string) connector.Client); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(connector.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDirectorClient provides a mock function with given fields: runtimeConfig
func (_m *ClientsProvider) GetDirectorClient(runtimeConfig config.RuntimeConfig) (director.DirectorClient, error) {
	ret := _m.Called(runtimeConfig)

	var r0 director.DirectorClient
	if rf, ok := ret.Get(0).(func(config.RuntimeConfig) director.DirectorClient); ok {
		r0 = rf(runtimeConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(director.DirectorClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(config.RuntimeConfig) error); ok {
		r1 = rf(runtimeConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}