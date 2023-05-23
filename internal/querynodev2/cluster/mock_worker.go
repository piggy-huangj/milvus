// Code generated by mockery v2.21.1. DO NOT EDIT.

package cluster

import (
	context "context"

	internalpb "github.com/milvus-io/milvus/internal/proto/internalpb"
	mock "github.com/stretchr/testify/mock"

	querypb "github.com/milvus-io/milvus/internal/proto/querypb"
)

// MockWorker is an autogenerated mock type for the Worker type
type MockWorker struct {
	mock.Mock
}

type MockWorker_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWorker) EXPECT() *MockWorker_Expecter {
	return &MockWorker_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, req
func (_m *MockWorker) Delete(ctx context.Context, req *querypb.DeleteRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.DeleteRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorker_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockWorker_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - req *querypb.DeleteRequest
func (_e *MockWorker_Expecter) Delete(ctx interface{}, req interface{}) *MockWorker_Delete_Call {
	return &MockWorker_Delete_Call{Call: _e.mock.On("Delete", ctx, req)}
}

func (_c *MockWorker_Delete_Call) Run(run func(ctx context.Context, req *querypb.DeleteRequest)) *MockWorker_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.DeleteRequest))
	})
	return _c
}

func (_c *MockWorker_Delete_Call) Return(_a0 error) *MockWorker_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorker_Delete_Call) RunAndReturn(run func(context.Context, *querypb.DeleteRequest) error) *MockWorker_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatistics provides a mock function with given fields: ctx, req
func (_m *MockWorker) GetStatistics(ctx context.Context, req *querypb.GetStatisticsRequest) (*internalpb.GetStatisticsResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *internalpb.GetStatisticsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.GetStatisticsRequest) (*internalpb.GetStatisticsResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.GetStatisticsRequest) *internalpb.GetStatisticsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalpb.GetStatisticsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *querypb.GetStatisticsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorker_GetStatistics_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatistics'
type MockWorker_GetStatistics_Call struct {
	*mock.Call
}

// GetStatistics is a helper method to define mock.On call
//   - ctx context.Context
//   - req *querypb.GetStatisticsRequest
func (_e *MockWorker_Expecter) GetStatistics(ctx interface{}, req interface{}) *MockWorker_GetStatistics_Call {
	return &MockWorker_GetStatistics_Call{Call: _e.mock.On("GetStatistics", ctx, req)}
}

func (_c *MockWorker_GetStatistics_Call) Run(run func(ctx context.Context, req *querypb.GetStatisticsRequest)) *MockWorker_GetStatistics_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.GetStatisticsRequest))
	})
	return _c
}

func (_c *MockWorker_GetStatistics_Call) Return(_a0 *internalpb.GetStatisticsResponse, _a1 error) *MockWorker_GetStatistics_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorker_GetStatistics_Call) RunAndReturn(run func(context.Context, *querypb.GetStatisticsRequest) (*internalpb.GetStatisticsResponse, error)) *MockWorker_GetStatistics_Call {
	_c.Call.Return(run)
	return _c
}

// IsHealthy provides a mock function with given fields:
func (_m *MockWorker) IsHealthy() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockWorker_IsHealthy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsHealthy'
type MockWorker_IsHealthy_Call struct {
	*mock.Call
}

// IsHealthy is a helper method to define mock.On call
func (_e *MockWorker_Expecter) IsHealthy() *MockWorker_IsHealthy_Call {
	return &MockWorker_IsHealthy_Call{Call: _e.mock.On("IsHealthy")}
}

func (_c *MockWorker_IsHealthy_Call) Run(run func()) *MockWorker_IsHealthy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWorker_IsHealthy_Call) Return(_a0 bool) *MockWorker_IsHealthy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorker_IsHealthy_Call) RunAndReturn(run func() bool) *MockWorker_IsHealthy_Call {
	_c.Call.Return(run)
	return _c
}

// LoadSegments provides a mock function with given fields: _a0, _a1
func (_m *MockWorker) LoadSegments(_a0 context.Context, _a1 *querypb.LoadSegmentsRequest) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.LoadSegmentsRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorker_LoadSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadSegments'
type MockWorker_LoadSegments_Call struct {
	*mock.Call
}

// LoadSegments is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *querypb.LoadSegmentsRequest
func (_e *MockWorker_Expecter) LoadSegments(_a0 interface{}, _a1 interface{}) *MockWorker_LoadSegments_Call {
	return &MockWorker_LoadSegments_Call{Call: _e.mock.On("LoadSegments", _a0, _a1)}
}

func (_c *MockWorker_LoadSegments_Call) Run(run func(_a0 context.Context, _a1 *querypb.LoadSegmentsRequest)) *MockWorker_LoadSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.LoadSegmentsRequest))
	})
	return _c
}

func (_c *MockWorker_LoadSegments_Call) Return(_a0 error) *MockWorker_LoadSegments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorker_LoadSegments_Call) RunAndReturn(run func(context.Context, *querypb.LoadSegmentsRequest) error) *MockWorker_LoadSegments_Call {
	_c.Call.Return(run)
	return _c
}

// QuerySegments provides a mock function with given fields: ctx, req
func (_m *MockWorker) QuerySegments(ctx context.Context, req *querypb.QueryRequest) (*internalpb.RetrieveResults, error) {
	ret := _m.Called(ctx, req)

	var r0 *internalpb.RetrieveResults
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.QueryRequest) (*internalpb.RetrieveResults, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.QueryRequest) *internalpb.RetrieveResults); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalpb.RetrieveResults)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *querypb.QueryRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorker_QuerySegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QuerySegments'
type MockWorker_QuerySegments_Call struct {
	*mock.Call
}

// QuerySegments is a helper method to define mock.On call
//   - ctx context.Context
//   - req *querypb.QueryRequest
func (_e *MockWorker_Expecter) QuerySegments(ctx interface{}, req interface{}) *MockWorker_QuerySegments_Call {
	return &MockWorker_QuerySegments_Call{Call: _e.mock.On("QuerySegments", ctx, req)}
}

func (_c *MockWorker_QuerySegments_Call) Run(run func(ctx context.Context, req *querypb.QueryRequest)) *MockWorker_QuerySegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.QueryRequest))
	})
	return _c
}

func (_c *MockWorker_QuerySegments_Call) Return(_a0 *internalpb.RetrieveResults, _a1 error) *MockWorker_QuerySegments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorker_QuerySegments_Call) RunAndReturn(run func(context.Context, *querypb.QueryRequest) (*internalpb.RetrieveResults, error)) *MockWorker_QuerySegments_Call {
	_c.Call.Return(run)
	return _c
}

// ReleaseSegments provides a mock function with given fields: _a0, _a1
func (_m *MockWorker) ReleaseSegments(_a0 context.Context, _a1 *querypb.ReleaseSegmentsRequest) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.ReleaseSegmentsRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorker_ReleaseSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseSegments'
type MockWorker_ReleaseSegments_Call struct {
	*mock.Call
}

// ReleaseSegments is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *querypb.ReleaseSegmentsRequest
func (_e *MockWorker_Expecter) ReleaseSegments(_a0 interface{}, _a1 interface{}) *MockWorker_ReleaseSegments_Call {
	return &MockWorker_ReleaseSegments_Call{Call: _e.mock.On("ReleaseSegments", _a0, _a1)}
}

func (_c *MockWorker_ReleaseSegments_Call) Run(run func(_a0 context.Context, _a1 *querypb.ReleaseSegmentsRequest)) *MockWorker_ReleaseSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.ReleaseSegmentsRequest))
	})
	return _c
}

func (_c *MockWorker_ReleaseSegments_Call) Return(_a0 error) *MockWorker_ReleaseSegments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorker_ReleaseSegments_Call) RunAndReturn(run func(context.Context, *querypb.ReleaseSegmentsRequest) error) *MockWorker_ReleaseSegments_Call {
	_c.Call.Return(run)
	return _c
}

// SearchSegments provides a mock function with given fields: ctx, req
func (_m *MockWorker) SearchSegments(ctx context.Context, req *querypb.SearchRequest) (*internalpb.SearchResults, error) {
	ret := _m.Called(ctx, req)

	var r0 *internalpb.SearchResults
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.SearchRequest) (*internalpb.SearchResults, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.SearchRequest) *internalpb.SearchResults); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalpb.SearchResults)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *querypb.SearchRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorker_SearchSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchSegments'
type MockWorker_SearchSegments_Call struct {
	*mock.Call
}

// SearchSegments is a helper method to define mock.On call
//   - ctx context.Context
//   - req *querypb.SearchRequest
func (_e *MockWorker_Expecter) SearchSegments(ctx interface{}, req interface{}) *MockWorker_SearchSegments_Call {
	return &MockWorker_SearchSegments_Call{Call: _e.mock.On("SearchSegments", ctx, req)}
}

func (_c *MockWorker_SearchSegments_Call) Run(run func(ctx context.Context, req *querypb.SearchRequest)) *MockWorker_SearchSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.SearchRequest))
	})
	return _c
}

func (_c *MockWorker_SearchSegments_Call) Return(_a0 *internalpb.SearchResults, _a1 error) *MockWorker_SearchSegments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorker_SearchSegments_Call) RunAndReturn(run func(context.Context, *querypb.SearchRequest) (*internalpb.SearchResults, error)) *MockWorker_SearchSegments_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockWorker) Stop() {
	_m.Called()
}

// MockWorker_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockWorker_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockWorker_Expecter) Stop() *MockWorker_Stop_Call {
	return &MockWorker_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockWorker_Stop_Call) Run(run func()) *MockWorker_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWorker_Stop_Call) Return() *MockWorker_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockWorker_Stop_Call) RunAndReturn(run func()) *MockWorker_Stop_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockWorker interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockWorker creates a new instance of MockWorker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockWorker(t mockConstructorTestingTNewMockWorker) *MockWorker {
	mock := &MockWorker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
