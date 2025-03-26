// Code generated by mockery v2.53.3. DO NOT EDIT.

package ports

import (
	context "context"
	courier "lisichkinuriy/delivery/internal/domain/courier"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ICourierRepositoryMock is an autogenerated mock type for the ICourierRepository type
type ICourierRepositoryMock struct {
	mock.Mock
}

type ICourierRepositoryMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ICourierRepositoryMock) EXPECT() *ICourierRepositoryMock_Expecter {
	return &ICourierRepositoryMock_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, aggregate
func (_m *ICourierRepositoryMock) Add(ctx context.Context, aggregate *courier.Courier) error {
	ret := _m.Called(ctx, aggregate)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *courier.Courier) error); ok {
		r0 = rf(ctx, aggregate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ICourierRepositoryMock_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type ICourierRepositoryMock_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx context.Context
//   - aggregate *courier.Courier
func (_e *ICourierRepositoryMock_Expecter) Add(ctx interface{}, aggregate interface{}) *ICourierRepositoryMock_Add_Call {
	return &ICourierRepositoryMock_Add_Call{Call: _e.mock.On("Add", ctx, aggregate)}
}

func (_c *ICourierRepositoryMock_Add_Call) Run(run func(ctx context.Context, aggregate *courier.Courier)) *ICourierRepositoryMock_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*courier.Courier))
	})
	return _c
}

func (_c *ICourierRepositoryMock_Add_Call) Return(_a0 error) *ICourierRepositoryMock_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ICourierRepositoryMock_Add_Call) RunAndReturn(run func(context.Context, *courier.Courier) error) *ICourierRepositoryMock_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *ICourierRepositoryMock) Get(ctx context.Context, id uuid.UUID) (*courier.Courier, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *courier.Courier
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*courier.Courier, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *courier.Courier); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*courier.Courier)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ICourierRepositoryMock_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ICourierRepositoryMock_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *ICourierRepositoryMock_Expecter) Get(ctx interface{}, id interface{}) *ICourierRepositoryMock_Get_Call {
	return &ICourierRepositoryMock_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *ICourierRepositoryMock_Get_Call) Run(run func(ctx context.Context, id uuid.UUID)) *ICourierRepositoryMock_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *ICourierRepositoryMock_Get_Call) Return(_a0 *courier.Courier, _a1 error) *ICourierRepositoryMock_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ICourierRepositoryMock_Get_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*courier.Courier, error)) *ICourierRepositoryMock_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllFreeCouriers provides a mock function with given fields: ctx
func (_m *ICourierRepositoryMock) GetAllFreeCouriers(ctx context.Context) ([]*courier.Courier, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllFreeCouriers")
	}

	var r0 []*courier.Courier
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*courier.Courier, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*courier.Courier); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*courier.Courier)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ICourierRepositoryMock_GetAllFreeCouriers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllFreeCouriers'
type ICourierRepositoryMock_GetAllFreeCouriers_Call struct {
	*mock.Call
}

// GetAllFreeCouriers is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ICourierRepositoryMock_Expecter) GetAllFreeCouriers(ctx interface{}) *ICourierRepositoryMock_GetAllFreeCouriers_Call {
	return &ICourierRepositoryMock_GetAllFreeCouriers_Call{Call: _e.mock.On("GetAllFreeCouriers", ctx)}
}

func (_c *ICourierRepositoryMock_GetAllFreeCouriers_Call) Run(run func(ctx context.Context)) *ICourierRepositoryMock_GetAllFreeCouriers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ICourierRepositoryMock_GetAllFreeCouriers_Call) Return(_a0 []*courier.Courier, _a1 error) *ICourierRepositoryMock_GetAllFreeCouriers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ICourierRepositoryMock_GetAllFreeCouriers_Call) RunAndReturn(run func(context.Context) ([]*courier.Courier, error)) *ICourierRepositoryMock_GetAllFreeCouriers_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, aggregate
func (_m *ICourierRepositoryMock) Update(ctx context.Context, aggregate *courier.Courier) error {
	ret := _m.Called(ctx, aggregate)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *courier.Courier) error); ok {
		r0 = rf(ctx, aggregate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ICourierRepositoryMock_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ICourierRepositoryMock_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - aggregate *courier.Courier
func (_e *ICourierRepositoryMock_Expecter) Update(ctx interface{}, aggregate interface{}) *ICourierRepositoryMock_Update_Call {
	return &ICourierRepositoryMock_Update_Call{Call: _e.mock.On("Update", ctx, aggregate)}
}

func (_c *ICourierRepositoryMock_Update_Call) Run(run func(ctx context.Context, aggregate *courier.Courier)) *ICourierRepositoryMock_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*courier.Courier))
	})
	return _c
}

func (_c *ICourierRepositoryMock_Update_Call) Return(_a0 error) *ICourierRepositoryMock_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ICourierRepositoryMock_Update_Call) RunAndReturn(run func(context.Context, *courier.Courier) error) *ICourierRepositoryMock_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewICourierRepositoryMock creates a new instance of ICourierRepositoryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICourierRepositoryMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ICourierRepositoryMock {
	mock := &ICourierRepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
