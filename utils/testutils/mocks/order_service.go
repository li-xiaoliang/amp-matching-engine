// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import bson "gopkg.in/mgo.v2/bson"
import common "github.com/ethereum/go-ethereum/common"
import engine "github.com/Proofsuite/amp-matching-engine/engine"
import mock "github.com/stretchr/testify/mock"
import rabbitmq "github.com/Proofsuite/amp-matching-engine/rabbitmq"

import types "github.com/Proofsuite/amp-matching-engine/types"

// OrderServiceInterface is an autogenerated mock type for the OrderServiceInterface type
type OrderServiceInterface struct {
	mock.Mock
}

// CancelOrder provides a mock function with given fields: oc
func (_m *OrderServiceInterface) CancelOrder(oc *types.OrderCancel) error {
	ret := _m.Called(oc)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.OrderCancel) error); ok {
		r0 = rf(oc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *OrderServiceInterface) GetByID(id bson.ObjectId) (*types.Order, error) {
	ret := _m.Called(id)

	var r0 *types.Order
	if rf, ok := ret.Get(0).(func(bson.ObjectId) *types.Order); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bson.ObjectId) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserAddress provides a mock function with given fields: addr
func (_m *OrderServiceInterface) GetByUserAddress(addr common.Address) ([]*types.Order, error) {
	ret := _m.Called(addr)

	var r0 []*types.Order
	if rf, ok := ret.Get(0).(func(common.Address) []*types.Order); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleEngineResponse provides a mock function with given fields: res
func (_m *OrderServiceInterface) HandleEngineResponse(res *engine.Response) error {
	ret := _m.Called(res)

	var r0 error
	if rf, ok := ret.Get(0).(func(*engine.Response) error); ok {
		r0 = rf(res)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOrder provides a mock function with given fields: o
func (_m *OrderServiceInterface) NewOrder(o *types.Order) error {
	ret := _m.Called(o)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Order) error); ok {
		r0 = rf(o)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishOrder provides a mock function with given fields: order
func (_m *OrderServiceInterface) PublishOrder(order *rabbitmq.Message) error {
	ret := _m.Called(order)

	var r0 error
	if rf, ok := ret.Get(0).(func(*rabbitmq.Message) error); ok {
		r0 = rf(order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RecoverOrders provides a mock function with given fields: res
func (_m *OrderServiceInterface) RecoverOrders(res *engine.Response) {
	_m.Called(res)
}

// RelayUpdateOverSocket provides a mock function with given fields: res
func (_m *OrderServiceInterface) RelayUpdateOverSocket(res *engine.Response) {
	_m.Called(res)
}

// SendMessage provides a mock function with given fields: msgType, hash, data
func (_m *OrderServiceInterface) SendMessage(msgType string, hash common.Hash, data interface{}) {
	_m.Called(msgType, hash, data)
}

// SubscribeQueue provides a mock function with given fields: fn
func (_m *OrderServiceInterface) SubscribeQueue(fn func(*rabbitmq.Message) error) error {
	ret := _m.Called(fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*rabbitmq.Message) error) error); ok {
		r0 = rf(fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
