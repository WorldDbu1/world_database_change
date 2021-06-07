// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	cid "github.com/ipfs/go-cid"
	car "github.com/ipld/go-car"

	context "context"

	io "io"

	ipld "github.com/ipld/go-ipld-prime"

	mock "github.com/stretchr/testify/mock"

	pieceio "github.com/filecoin-project/go-commp-utils/pieceio"
)

// CarIO is an autogenerated mock type for the CarIO type
type CarIO struct {
	mock.Mock
}

// LoadCar provides a mock function with given fields: bs, r
func (_m *CarIO) LoadCar(bs pieceio.WriteStore, r io.Reader) (cid.Cid, error) {
	ret := _m.Called(bs, r)

	var r0 cid.Cid
	if rf, ok := ret.Get(0).(func(pieceio.WriteStore, io.Reader) cid.Cid); ok {
		r0 = rf(bs, r)
	} else {
		r0 = ret.Get(0).(cid.Cid)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pieceio.WriteStore, io.Reader) error); ok {
		r1 = rf(bs, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrepareCar provides a mock function with given fields: ctx, bs, payloadCid, node, userOnNewCarBlocks
func (_m *CarIO) PrepareCar(ctx context.Context, bs pieceio.ReadStore, payloadCid cid.Cid, node ipld.Node, userOnNewCarBlocks ...car.OnNewCarBlockFunc) (pieceio.PreparedCar, error) {
	_va := make([]interface{}, len(userOnNewCarBlocks))
	for _i := range userOnNewCarBlocks {
		_va[_i] = userOnNewCarBlocks[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bs, payloadCid, node)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 pieceio.PreparedCar
	if rf, ok := ret.Get(0).(func(context.Context, pieceio.ReadStore, cid.Cid, ipld.Node, ...car.OnNewCarBlockFunc) pieceio.PreparedCar); ok {
		r0 = rf(ctx, bs, payloadCid, node, userOnNewCarBlocks...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pieceio.PreparedCar)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pieceio.ReadStore, cid.Cid, ipld.Node, ...car.OnNewCarBlockFunc) error); ok {
		r1 = rf(ctx, bs, payloadCid, node, userOnNewCarBlocks...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteCar provides a mock function with given fields: ctx, bs, payloadCid, node, w, userOnNewCarBlocks
func (_m *CarIO) WriteCar(ctx context.Context, bs pieceio.ReadStore, payloadCid cid.Cid, node ipld.Node, w io.Writer, userOnNewCarBlocks ...car.OnNewCarBlockFunc) error {
	_va := make([]interface{}, len(userOnNewCarBlocks))
	for _i := range userOnNewCarBlocks {
		_va[_i] = userOnNewCarBlocks[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, bs, payloadCid, node, w)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pieceio.ReadStore, cid.Cid, ipld.Node, io.Writer, ...car.OnNewCarBlockFunc) error); ok {
		r0 = rf(ctx, bs, payloadCid, node, w, userOnNewCarBlocks...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
