// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import shared "github.com/kyma-project/kyma/components/console-backend-service/internal/domain/shared"

// ApplicationRetriever is an autogenerated mock type for the ApplicationRetriever type
type ApplicationRetriever struct {
	mock.Mock
}

// Application provides a mock function with given fields:
func (_m *ApplicationRetriever) Application() shared.ApplicationLister {
	ret := _m.Called()

	var r0 shared.ApplicationLister
	if rf, ok := ret.Get(0).(func() shared.ApplicationLister); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(shared.ApplicationLister)
		}
	}

	return r0
}