// Code generated by mockery. DO NOT EDIT.

package mock

import (
	product "github.com/sebajax/go-vertical-slice-architecture/internal/product"
	mock "github.com/stretchr/testify/mock"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

type ProductRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ProductRepository) EXPECT() *ProductRepository_Expecter {
	return &ProductRepository_Expecter{mock: &_m.Mock}
}

// GetBySku provides a mock function with given fields: sku
func (_m *ProductRepository) GetBySku(sku string) (*product.Product, bool, error) {
	ret := _m.Called(sku)

	if len(ret) == 0 {
		panic("no return value specified for GetBySku")
	}

	var r0 *product.Product
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (*product.Product, bool, error)); ok {
		return rf(sku)
	}
	if rf, ok := ret.Get(0).(func(string) *product.Product); ok {
		r0 = rf(sku)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(sku)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(sku)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProductRepository_GetBySku_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBySku'
type ProductRepository_GetBySku_Call struct {
	*mock.Call
}

// GetBySku is a helper method to define mock.On call
//   - sku string
func (_e *ProductRepository_Expecter) GetBySku(sku interface{}) *ProductRepository_GetBySku_Call {
	return &ProductRepository_GetBySku_Call{Call: _e.mock.On("GetBySku", sku)}
}

func (_c *ProductRepository_GetBySku_Call) Run(run func(sku string)) *ProductRepository_GetBySku_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ProductRepository_GetBySku_Call) Return(_a0 *product.Product, _a1 bool, _a2 error) *ProductRepository_GetBySku_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProductRepository_GetBySku_Call) RunAndReturn(run func(string) (*product.Product, bool, error)) *ProductRepository_GetBySku_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: p
func (_m *ProductRepository) Save(p *product.Product) (int64, error) {
	ret := _m.Called(p)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*product.Product) (int64, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func(*product.Product) int64); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*product.Product) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProductRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type ProductRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - p *product.Product
func (_e *ProductRepository_Expecter) Save(p interface{}) *ProductRepository_Save_Call {
	return &ProductRepository_Save_Call{Call: _e.mock.On("Save", p)}
}

func (_c *ProductRepository_Save_Call) Run(run func(p *product.Product)) *ProductRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*product.Product))
	})
	return _c
}

func (_c *ProductRepository_Save_Call) Return(_a0 int64, _a1 error) *ProductRepository_Save_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProductRepository_Save_Call) RunAndReturn(run func(*product.Product) (int64, error)) *ProductRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
