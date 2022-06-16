// Code generated by mockery v2.12.2. DO NOT EDIT.

package dashboards

import (
	context "context"

	models "github.com/grafana/grafana/pkg/models"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MockDashboardStore is an autogenerated mock type for the Store type
type MockDashboardStore struct {
	mock.Mock
}

// DeleteDashboard provides a mock function with given fields: ctx, cmd
func (_m *MockDashboardStore) DeleteDashboard(ctx context.Context, cmd *DeleteDashboardCommand) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *DeleteDashboardCommand) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOrphanedProvisionedDashboards provides a mock function with given fields: ctx, cmd
func (_m *MockDashboardStore) DeleteOrphanedProvisionedDashboards(ctx context.Context, cmd *DeleteOrphanedProvisionedDashboardsCommand) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *DeleteOrphanedProvisionedDashboardsCommand) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindDashboards provides a mock function with given fields: ctx, query
func (_m *MockDashboardStore) FindDashboards(ctx context.Context, query *models.FindPersistedDashboardsQuery) ([]DashboardSearchProjection, error) {
	ret := _m.Called(ctx, query)

	var r0 []DashboardSearchProjection
	if rf, ok := ret.Get(0).(func(context.Context, *models.FindPersistedDashboardsQuery) []DashboardSearchProjection); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]DashboardSearchProjection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.FindPersistedDashboardsQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboard provides a mock function with given fields: ctx, query
func (_m *MockDashboardStore) GetDashboard(ctx context.Context, query *GetDashboardQuery) (*Dashboard, error) {
	ret := _m.Called(ctx, query)

	var r0 *Dashboard
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardQuery) *Dashboard); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Dashboard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetDashboardQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboardAclInfoList provides a mock function with given fields: ctx, query
func (_m *MockDashboardStore) GetDashboardAclInfoList(ctx context.Context, query *models.GetDashboardAclInfoListQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetDashboardAclInfoListQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDashboardTags provides a mock function with given fields: ctx, query
func (_m *MockDashboardStore) GetDashboardTags(ctx context.Context, query *GetDashboardTagsQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardTagsQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDashboardUIDById provides a mock function with given fields: ctx, query
func (_m *MockDashboardStore) GetDashboardUIDById(ctx context.Context, query *GetDashboardRefByIdQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardRefByIdQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDashboards provides a mock function with given fields: ctx, query
func (_m *MockDashboardStore) GetDashboards(ctx context.Context, query *GetDashboardsQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardsQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDashboardsByPluginID provides a mock function with given fields: ctx, query
func (_m *MockDashboardStore) GetDashboardsByPluginID(ctx context.Context, query *GetDashboardsByPluginIdQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardsByPluginIdQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFolderByID provides a mock function with given fields: ctx, orgID, id
func (_m *MockDashboardStore) GetFolderByID(ctx context.Context, orgID int64, id int64) (*Folder, error) {
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
func (_m *MockDashboardStore) GetFolderByTitle(ctx context.Context, orgID int64, title string) (*Folder, error) {
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
func (_m *MockDashboardStore) GetFolderByUID(ctx context.Context, orgID int64, uid string) (*Folder, error) {
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

// GetProvisionedDashboardData provides a mock function with given fields: name
func (_m *MockDashboardStore) GetProvisionedDashboardData(name string) ([]*DashboardProvisioning, error) {
	ret := _m.Called(name)

	var r0 []*DashboardProvisioning
	if rf, ok := ret.Get(0).(func(string) []*DashboardProvisioning); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*DashboardProvisioning)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisionedDataByDashboardID provides a mock function with given fields: dashboardID
func (_m *MockDashboardStore) GetProvisionedDataByDashboardID(dashboardID int64) (*DashboardProvisioning, error) {
	ret := _m.Called(dashboardID)

	var r0 *DashboardProvisioning
	if rf, ok := ret.Get(0).(func(int64) *DashboardProvisioning); ok {
		r0 = rf(dashboardID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DashboardProvisioning)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(dashboardID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisionedDataByDashboardUID provides a mock function with given fields: orgID, dashboardUID
func (_m *MockDashboardStore) GetProvisionedDataByDashboardUID(orgID int64, dashboardUID string) (*DashboardProvisioning, error) {
	ret := _m.Called(orgID, dashboardUID)

	var r0 *DashboardProvisioning
	if rf, ok := ret.Get(0).(func(int64, string) *DashboardProvisioning); ok {
		r0 = rf(orgID, dashboardUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DashboardProvisioning)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string) error); ok {
		r1 = rf(orgID, dashboardUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPublicDashboard provides a mock function with given fields: uid
func (_m *MockDashboardStore) GetPublicDashboard(uid string) (*PublicDashboard, *Dashboard, error) {
	ret := _m.Called(uid)

	var r0 *PublicDashboard
	if rf, ok := ret.Get(0).(func(string) *PublicDashboard); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PublicDashboard)
		}
	}

	var r1 *Dashboard
	if rf, ok := ret.Get(1).(func(string) *Dashboard); ok {
		r1 = rf(uid)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*Dashboard)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(uid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetPublicDashboardConfig provides a mock function with given fields: orgId, dashboardUid
func (_m *MockDashboardStore) GetPublicDashboardConfig(orgId int64, dashboardUid string) (*PublicDashboardConfig, error) {
	ret := _m.Called(orgId, dashboardUid)

	var r0 *PublicDashboardConfig
	if rf, ok := ret.Get(0).(func(int64, string) *PublicDashboardConfig); ok {
		r0 = rf(orgId, dashboardUid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PublicDashboardConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string) error); ok {
		r1 = rf(orgId, dashboardUid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasAdminPermissionInFolders provides a mock function with given fields: ctx, query
func (_m *MockDashboardStore) HasAdminPermissionInFolders(ctx context.Context, query *HasAdminPermissionInFoldersQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *HasAdminPermissionInFoldersQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HasEditPermissionInFolders provides a mock function with given fields: ctx, query
func (_m *MockDashboardStore) HasEditPermissionInFolders(ctx context.Context, query *HasEditPermissionInFoldersQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *HasEditPermissionInFoldersQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveAlerts provides a mock function with given fields: ctx, dashID, alerts
func (_m *MockDashboardStore) SaveAlerts(ctx context.Context, dashID int64, alerts []*models.Alert) error {
	ret := _m.Called(ctx, dashID, alerts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []*models.Alert) error); ok {
		r0 = rf(ctx, dashID, alerts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveDashboard provides a mock function with given fields: cmd
func (_m *MockDashboardStore) SaveDashboard(cmd SaveDashboardCommand) (*Dashboard, error) {
	ret := _m.Called(cmd)

	var r0 *Dashboard
	if rf, ok := ret.Get(0).(func(SaveDashboardCommand) *Dashboard); ok {
		r0 = rf(cmd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Dashboard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(SaveDashboardCommand) error); ok {
		r1 = rf(cmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveProvisionedDashboard provides a mock function with given fields: cmd, provisioning
func (_m *MockDashboardStore) SaveProvisionedDashboard(cmd SaveDashboardCommand, provisioning *DashboardProvisioning) (*Dashboard, error) {
	ret := _m.Called(cmd, provisioning)

	var r0 *Dashboard
	if rf, ok := ret.Get(0).(func(SaveDashboardCommand, *DashboardProvisioning) *Dashboard); ok {
		r0 = rf(cmd, provisioning)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Dashboard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(SaveDashboardCommand, *DashboardProvisioning) error); ok {
		r1 = rf(cmd, provisioning)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SavePublicDashboardConfig provides a mock function with given fields: cmd
func (_m *MockDashboardStore) SavePublicDashboardConfig(cmd SavePublicDashboardConfigCommand) (*PublicDashboardConfig, error) {
	ret := _m.Called(cmd)

	var r0 *PublicDashboardConfig
	if rf, ok := ret.Get(0).(func(SavePublicDashboardConfigCommand) *PublicDashboardConfig); ok {
		r0 = rf(cmd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PublicDashboardConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(SavePublicDashboardConfigCommand) error); ok {
		r1 = rf(cmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnprovisionDashboard provides a mock function with given fields: ctx, id
func (_m *MockDashboardStore) UnprovisionDashboard(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDashboardACL provides a mock function with given fields: ctx, uid, items
func (_m *MockDashboardStore) UpdateDashboardACL(ctx context.Context, uid int64, items []*models.DashboardAcl) error {
	ret := _m.Called(ctx, uid, items)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []*models.DashboardAcl) error); ok {
		r0 = rf(ctx, uid, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateDashboardBeforeSave provides a mock function with given fields: dashboard, overwrite
func (_m *MockDashboardStore) ValidateDashboardBeforeSave(dashboard *Dashboard, overwrite bool) (bool, error) {
	ret := _m.Called(dashboard, overwrite)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*Dashboard, bool) bool); ok {
		r0 = rf(dashboard, overwrite)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Dashboard, bool) error); ok {
		r1 = rf(dashboard, overwrite)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockDashboardStore creates a new instance of MockDashboardStore. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDashboardStore(t testing.TB) *MockDashboardStore {
	mock := &MockDashboardStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
