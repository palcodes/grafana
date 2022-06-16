package database

import (
	"context"
	"fmt"

	"github.com/grafana/grafana/pkg/services/dashboards"
	"github.com/grafana/grafana/pkg/services/sqlstore"
	"github.com/grafana/grafana/pkg/util"
)

// retrieves public dashboard configuration
func (d *DashboardStore) GetPublicDashboard(uid string) (*dashboards.PublicDashboard, *dashboards.Dashboard, error) {
	if uid == "" {
		return nil, nil, dashboards.ErrPublicDashboardIdentifierNotSet
	}

	// get public dashboard
	pdRes := &dashboards.PublicDashboard{Uid: uid}
	err := d.sqlStore.WithTransactionalDbSession(context.Background(), func(sess *sqlstore.DBSession) error {
		has, err := sess.Get(pdRes)
		if err != nil {
			return err
		}
		if !has {
			return dashboards.ErrPublicDashboardNotFound
		}
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	// find dashboard
	dashRes := &dashboards.Dashboard{OrgId: pdRes.OrgId, Uid: pdRes.DashboardUid}
	err = d.sqlStore.WithTransactionalDbSession(context.Background(), func(sess *sqlstore.DBSession) error {
		has, err := sess.Get(dashRes)
		if err != nil {
			return err
		}
		if !has {
			return dashboards.ErrPublicDashboardNotFound
		}
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return pdRes, dashRes, err
}

// generates a new unique uid to retrieve a public dashboard
func generateNewPublicDashboardUid(sess *sqlstore.DBSession) (string, error) {
	for i := 0; i < 3; i++ {
		uid := util.GenerateShortUID()

		exists, err := sess.Get(&dashboards.PublicDashboard{Uid: uid})
		if err != nil {
			return "", err
		}

		if !exists {
			return uid, nil
		}
	}

	return "", dashboards.ErrPublicDashboardFailedGenerateUniqueUid
}

// retrieves public dashboard configuration
func (d *DashboardStore) GetPublicDashboardConfig(orgId int64, dashboardUid string) (*dashboards.PublicDashboardConfig, error) {
	if dashboardUid == "" {
		return nil, dashboards.ErrDashboardIdentifierNotSet
	}

	// get dashboard and publicDashboard
	dashRes := &dashboards.Dashboard{OrgId: orgId, Uid: dashboardUid}
	pdRes := &dashboards.PublicDashboard{OrgId: orgId, DashboardUid: dashboardUid}
	err := d.sqlStore.WithTransactionalDbSession(context.Background(), func(sess *sqlstore.DBSession) error {
		// dashboard
		has, err := sess.Get(dashRes)
		if err != nil {
			return err
		}
		if !has {
			return dashboards.ErrDashboardNotFound
		}

		// publicDashboard
		_, err = sess.Get(pdRes)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	pdc := &dashboards.PublicDashboardConfig{
		IsPublic:        dashRes.IsPublic,
		PublicDashboard: *pdRes,
	}

	return pdc, err
}

// persists public dashboard configuration
func (d *DashboardStore) SavePublicDashboardConfig(cmd dashboards.SavePublicDashboardConfigCommand) (*dashboards.PublicDashboardConfig, error) {
	if len(cmd.PublicDashboardConfig.PublicDashboard.DashboardUid) == 0 {
		return nil, dashboards.ErrDashboardIdentifierNotSet
	}

	err := d.sqlStore.WithTransactionalDbSession(context.Background(), func(sess *sqlstore.DBSession) error {
		// update isPublic on dashboard entry
		affectedRowCount, err := sess.Table("dashboard").Where("org_id = ? AND uid = ?", cmd.OrgId, cmd.DashboardUid).Update(map[string]interface{}{"is_public": cmd.PublicDashboardConfig.IsPublic})
		if err != nil {
			return err
		}

		if affectedRowCount == 0 {
			return dashboards.ErrDashboardNotFound
		}

		// update dashboard_public_config
		// if we have a uid, public dashboard config exists. delete it otherwise generate a uid
		if cmd.PublicDashboardConfig.PublicDashboard.Uid != "" {
			if _, err = sess.Exec("DELETE FROM dashboard_public_config WHERE uid=?", cmd.PublicDashboardConfig.PublicDashboard.Uid); err != nil {
				return err
			}
		} else {
			uid, err := generateNewPublicDashboardUid(sess)
			if err != nil {
				return fmt.Errorf("failed to generate UID for public dashboard: %w", err)
			}
			cmd.PublicDashboardConfig.PublicDashboard.Uid = uid
		}

		_, err = sess.Insert(&cmd.PublicDashboardConfig.PublicDashboard)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &cmd.PublicDashboardConfig, nil
}
