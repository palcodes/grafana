package publicdashboards

import (
	"context"

	"github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/models"
	. "github.com/grafana/grafana/pkg/services/publicdashboards/models"
)

// These are the api contracts. The API should match the underlying service and
// store

//go:generate mockery --name PublicDashboardService --structname FakePublicDashboardService --inpackage --filename public_dashboard_service_mock.go
// DashboardProvisioningService is a service for operating on provisioned dashboards.
type Service interface {
	GetPublicDashboard(ctx context.Context, accessToken string) (*models.Dashboard, error)
	GetPublicDashboardConfig(ctx context.Context, orgId int64, dashboardUid string) (*PublicDashboard, error)
	SavePublicDashboardConfig(ctx context.Context, dto *SavePublicDashboardConfigDTO) (*PublicDashboard, error)
	BuildPublicDashboardMetricRequest(ctx context.Context, dashboard *models.Dashboard, publicDashboard *PublicDashboard, panelId int64) (dtos.MetricRequest, error)
}

//go:generate mockery --name Store --structname FakePublicDashboardStore --inpackage --filename store_mock.go
// Store is a dashboard store.
type Store interface {
	GetPublicDashboard(ctx context.Context, accessToken string) (*PublicDashboard, *models.Dashboard, error)
	GetPublicDashboardConfig(ctx context.Context, orgId int64, dashboardUid string) (*PublicDashboard, error)
	GenerateNewPublicDashboardUid(ctx context.Context) (string, error)
	SavePublicDashboardConfig(ctx context.Context, cmd SavePublicDashboardConfigCommand) (*PublicDashboard, error)
	UpdatePublicDashboardConfig(ctx context.Context, cmd SavePublicDashboardConfigCommand) error
}
