// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	compute "google.golang.org/api/compute/v1"
)

// InstanceGetter is an autogenerated mock type for the InstanceGetter type
type InstanceGetter struct {
	mock.Mock
}

type InstanceGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *InstanceGetter) EXPECT() *InstanceGetter_Expecter {
	return &InstanceGetter_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: projectID, zone, instance
func (_m *InstanceGetter) Get(projectID string, zone string, instance string) (*compute.Instance, error) {
	ret := _m.Called(projectID, zone, instance)

	var r0 *compute.Instance
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (*compute.Instance, error)); ok {
		return rf(projectID, zone, instance)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) *compute.Instance); ok {
		r0 = rf(projectID, zone, instance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Instance)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(projectID, zone, instance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstanceGetter_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type InstanceGetter_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - projectID string
//   - zone string
//   - instance string
func (_e *InstanceGetter_Expecter) Get(projectID interface{}, zone interface{}, instance interface{}) *InstanceGetter_Get_Call {
	return &InstanceGetter_Get_Call{Call: _e.mock.On("Get", projectID, zone, instance)}
}

func (_c *InstanceGetter_Get_Call) Run(run func(projectID string, zone string, instance string)) *InstanceGetter_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *InstanceGetter_Get_Call) Return(_a0 *compute.Instance, _a1 error) *InstanceGetter_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *InstanceGetter_Get_Call) RunAndReturn(run func(string, string, string) (*compute.Instance, error)) *InstanceGetter_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewInstanceGetter creates a new instance of InstanceGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInstanceGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *InstanceGetter {
	mock := &InstanceGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}