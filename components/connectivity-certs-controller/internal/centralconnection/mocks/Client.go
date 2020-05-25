// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	v1alpha1 "github.com/kyma-project/kyma/components/connectivity-certs-controller/pkg/apis/applicationconnector/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Upsert provides a mock function with given fields: name, spec
func (_m *Client) Upsert(name string, spec v1alpha1.CentralConnectionSpec) (*v1alpha1.CentralConnection, error) {
	ret := _m.Called(name, spec)

	var r0 *v1alpha1.CentralConnection
	if rf, ok := ret.Get(0).(func(string, v1alpha1.CentralConnectionSpec) *v1alpha1.CentralConnection); ok {
		r0 = rf(name, spec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.CentralConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, v1alpha1.CentralConnectionSpec) error); ok {
		r1 = rf(name, spec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}