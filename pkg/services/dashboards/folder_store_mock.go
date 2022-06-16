// Code generated by mockery v2.12.2. DO NOT EDIT.

package dashboards

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockFolderStore is an autogenerated mock type for the FolderStore type
type MockFolderStore struct {
	mock.Mock
}

// GetFolderByID provides a mock function with given fields: ctx, orgID, id
func (_m *MockFolderStore) GetFolderByID(ctx context.Context, orgID int64, id int64) (*Folder, error) {
	ret := _m.Called(ctx, orgID, id)

	var r0 *Folder
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) *Folder); ok {
		r0 = rf(ctx, orgID, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Folder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, orgID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFolderByTitle provides a mock function with given fields: ctx, orgID, title
func (_m *MockFolderStore) GetFolderByTitle(ctx context.Context, orgID int64, title string) (*Folder, error) {
	ret := _m.Called(ctx, orgID, title)

	var r0 *Folder
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) *Folder); ok {
		r0 = rf(ctx, orgID, title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Folder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, orgID, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFolderByUID provides a mock function with given fields: ctx, orgID, uid
func (_m *MockFolderStore) GetFolderByUID(ctx context.Context, orgID int64, uid string) (*Folder, error) {
	ret := _m.Called(ctx, orgID, uid)

	var r0 *Folder
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) *Folder); ok {
		r0 = rf(ctx, orgID, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Folder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, orgID, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockFolderStore creates a new instance of MockFolderStore. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockFolderStore(t testing.TB) *MockFolderStore {
	mock := &MockFolderStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
