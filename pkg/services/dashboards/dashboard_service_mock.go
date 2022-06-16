// Code generated by mockery v2.12.2. DO NOT EDIT.

package dashboards

import (
	context "context"

	dtos "github.com/grafana/grafana/pkg/api/dtos"
	mock "github.com/stretchr/testify/mock"

	models "github.com/grafana/grafana/pkg/models"

	testing "testing"
)

// MockDashboardService is an autogenerated mock type for the DashboardService type
type MockDashboardService struct {
	mock.Mock
}

// BuildPublicDashboardMetricRequest provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockDashboardService) BuildPublicDashboardMetricRequest(_a0 context.Context, _a1 string, _a2 int64) (dtos.MetricRequest, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 dtos.MetricRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) dtos.MetricRequest); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(dtos.MetricRequest)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildSaveDashboardCommand provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockDashboardService) BuildSaveDashboardCommand(_a0 context.Context, _a1 *SaveDashboardDTO, _a2 bool, _a3 bool) (*SaveDashboardCommand, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *SaveDashboardCommand
	if rf, ok := ret.Get(0).(func(context.Context, *SaveDashboardDTO, bool, bool) *SaveDashboardCommand); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SaveDashboardCommand)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *SaveDashboardDTO, bool, bool) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDashboard provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockDashboardService) DeleteDashboard(_a0 context.Context, _a1 int64, _a2 int64) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindDashboards provides a mock function with given fields: _a0, _a1
func (_m *MockDashboardService) FindDashboards(_a0 context.Context, _a1 *models.FindPersistedDashboardsQuery) ([]DashboardSearchProjection, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []DashboardSearchProjection
	if rf, ok := ret.Get(0).(func(context.Context, *models.FindPersistedDashboardsQuery) []DashboardSearchProjection); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]DashboardSearchProjection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.FindPersistedDashboardsQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboard provides a mock function with given fields: _a0, _a1
func (_m *MockDashboardService) GetDashboard(_a0 context.Context, _a1 *GetDashboardQuery) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardQuery) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDashboardAclInfoList provides a mock function with given fields: _a0, _a1
func (_m *MockDashboardService) GetDashboardAclInfoList(_a0 context.Context, _a1 *models.GetDashboardAclInfoListQuery) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetDashboardAclInfoListQuery) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDashboardTags provides a mock function with given fields: _a0, _a1
func (_m *MockDashboardService) GetDashboardTags(_a0 context.Context, _a1 *GetDashboardTagsQuery) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardTagsQuery) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDashboardUIDById provides a mock function with given fields: _a0, _a1
func (_m *MockDashboardService) GetDashboardUIDById(_a0 context.Context, _a1 *GetDashboardRefByIdQuery) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardRefByIdQuery) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDashboards provides a mock function with given fields: _a0, _a1
func (_m *MockDashboardService) GetDashboards(_a0 context.Context, _a1 *GetDashboardsQuery) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardsQuery) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPublicDashboard provides a mock function with given fields: _a0, _a1
func (_m *MockDashboardService) GetPublicDashboard(_a0 context.Context, _a1 string) (*Dashboard, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *Dashboard
	if rf, ok := ret.Get(0).(func(context.Context, string) *Dashboard); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Dashboard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPublicDashboardConfig provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockDashboardService) GetPublicDashboardConfig(_a0 context.Context, _a1 int64, _a2 string) (*PublicDashboardConfig, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *PublicDashboardConfig
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) *PublicDashboardConfig); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PublicDashboardConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasAdminPermissionInFolders provides a mock function with given fields: _a0, _a1
func (_m *MockDashboardService) HasAdminPermissionInFolders(_a0 context.Context, _a1 *HasAdminPermissionInFoldersQuery) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *HasAdminPermissionInFoldersQuery) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HasEditPermissionInFolders provides a mock function with given fields: _a0, _a1
func (_m *MockDashboardService) HasEditPermissionInFolders(_a0 context.Context, _a1 *HasEditPermissionInFoldersQuery) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *HasEditPermissionInFoldersQuery) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImportDashboard provides a mock function with given fields: _a0, _a1
func (_m *MockDashboardService) ImportDashboard(_a0 context.Context, _a1 *SaveDashboardDTO) (*Dashboard, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *Dashboard
	if rf, ok := ret.Get(0).(func(context.Context, *SaveDashboardDTO) *Dashboard); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Dashboard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *SaveDashboardDTO) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeUserAdmin provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *MockDashboardService) MakeUserAdmin(_a0 context.Context, _a1 int64, _a2 int64, _a3 int64, _a4 bool) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, int64, bool) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveDashboard provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockDashboardService) SaveDashboard(_a0 context.Context, _a1 *SaveDashboardDTO, _a2 bool) (*Dashboard, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *Dashboard
	if rf, ok := ret.Get(0).(func(context.Context, *SaveDashboardDTO, bool) *Dashboard); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Dashboard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *SaveDashboardDTO, bool) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SavePublicDashboardConfig provides a mock function with given fields: _a0, _a1
func (_m *MockDashboardService) SavePublicDashboardConfig(_a0 context.Context, _a1 *SavePublicDashboardConfigDTO) (*PublicDashboardConfig, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *PublicDashboardConfig
	if rf, ok := ret.Get(0).(func(context.Context, *SavePublicDashboardConfigDTO) *PublicDashboardConfig); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PublicDashboardConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *SavePublicDashboardConfigDTO) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchDashboards provides a mock function with given fields: _a0, _a1
func (_m *MockDashboardService) SearchDashboards(_a0 context.Context, _a1 *models.FindPersistedDashboardsQuery) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.FindPersistedDashboardsQuery) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDashboardACL provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockDashboardService) UpdateDashboardACL(_a0 context.Context, _a1 int64, _a2 []*models.DashboardAcl) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []*models.DashboardAcl) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockDashboardService creates a new instance of MockDashboardService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDashboardService(t testing.TB) *MockDashboardService {
	mock := &MockDashboardService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
