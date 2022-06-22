package models

import (
	"errors"
	"time"

	"github.com/grafana/grafana/pkg/components/simplejson"
)

const (
	DS_GRAPHITE       = "graphite"
	DS_INFLUXDB       = "influxdb"
	DS_INFLUXDB_08    = "influxdb_08"
	DS_ES             = "elasticsearch"
	DS_PROMETHEUS     = "prometheus"
	DS_ALERTMANAGER   = "alertmanager"
	DS_JAEGER         = "jaeger"
	DS_LOKI           = "loki"
	DS_OPENTSDB       = "opentsdb"
	DS_TEMPO          = "tempo"
	DS_ZIPKIN         = "zipkin"
	DS_MYSQL          = "mysql"
	DS_POSTGRES       = "postgres"
	DS_MSSQL          = "mssql"
	DS_ACCESS_DIRECT  = "direct"
	DS_ACCESS_PROXY   = "proxy"
	DS_ES_OPEN_DISTRO = "grafana-es-open-distro-datasource"
	DS_ES_OPENSEARCH  = "grafana-opensearch-datasource"
)

var (
	ErrDataSourceNotFound                = errors.New("data source not found")
	ErrDataSourceNameExists              = errors.New("data source with the same name already exists")
	ErrDataSourceUidExists               = errors.New("data source with the same uid already exists")
	ErrDataSourceUpdatingOldVersion      = errors.New("trying to update old version of datasource")
	ErrDatasourceIsReadOnly              = errors.New("data source is readonly, can only be updated from configuration")
	ErrDataSourceAccessDenied            = errors.New("data source access denied")
	ErrDataSourceFailedGenerateUniqueUid = errors.New("failed to generate unique datasource ID")
	ErrDataSourceIdentifierNotSet        = errors.New("unique identifier and org id are needed to be able to get or delete a datasource")
	ErrCorrelationExists                 = errors.New("correlation to the same datasource already exists")
)

type DsAccess string

type Correlation struct {
	Target  string           `json:"target"`
	Options *simplejson.Json `json:"options,omitempty"`
}

type DataSource struct {
	Id      int64 `json:"id"`
	OrgId   int64 `json:"orgId"`
	Version int   `json:"version"`

	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Access DsAccess `json:"access"`
	Url    string   `json:"url"`
	// swagger:ignore
	Password      string `json:"-"`
	User          string `json:"user"`
	Database      string `json:"database"`
	BasicAuth     bool   `json:"basicAuth"`
	BasicAuthUser string `json:"basicAuthUser"`
	// swagger:ignore
	BasicAuthPassword string            `json:"-"`
	WithCredentials   bool              `json:"withCredentials"`
	IsDefault         bool              `json:"isDefault"`
	Correlations      []Correlation     `json:"correlations"`
	JsonData          *simplejson.Json  `json:"jsonData"`
	SecureJsonData    map[string][]byte `json:"secureJsonData"`
	ReadOnly          bool              `json:"readOnly"`
	Uid               string            `json:"uid"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// AllowedCookies parses the jsondata.keepCookies and returns a list of
// allowed cookies, otherwise an empty list.
func (ds DataSource) AllowedCookies() []string {
	if ds.JsonData != nil {
		if keepCookies := ds.JsonData.Get("keepCookies"); keepCookies != nil {
			return keepCookies.MustStringArray()
		}
	}

	return []string{}
}

// Specific error type for grpc secrets management so that we can show more detailed plugin errors to users
type ErrDatasourceSecretsPluginUserFriendly struct {
	Err string
}

func (e ErrDatasourceSecretsPluginUserFriendly) Error() string {
	return e.Err
}

// ----------------------
// COMMANDS

// Also acts as api DTO
type AddDataSourceCommand struct {
	Name            string            `json:"name" binding:"Required"`
	Type            string            `json:"type" binding:"Required"`
	Access          DsAccess          `json:"access" binding:"Required"`
	Url             string            `json:"url"`
	Database        string            `json:"database"`
	User            string            `json:"user"`
	BasicAuth       bool              `json:"basicAuth"`
	BasicAuthUser   string            `json:"basicAuthUser"`
	WithCredentials bool              `json:"withCredentials"`
	IsDefault       bool              `json:"isDefault"`
	JsonData        *simplejson.Json  `json:"jsonData"`
	SecureJsonData  map[string]string `json:"secureJsonData"`
	Uid             string            `json:"uid"`

	OrgId                   int64             `json:"-"`
	UserId                  int64             `json:"-"`
	ReadOnly                bool              `json:"-"`
	EncryptedSecureJsonData map[string][]byte `json:"-"`
	UpdateSecretFn          UpdateSecretFn    `json:"-"`

	Result *DataSource `json:"-"`
}

// Also acts as api DTO
type UpdateDataSourceCommand struct {
	Name            string            `json:"name" binding:"Required"`
	Type            string            `json:"type" binding:"Required"`
	Access          DsAccess          `json:"access" binding:"Required"`
	Url             string            `json:"url"`
	User            string            `json:"user"`
	Database        string            `json:"database"`
	BasicAuth       bool              `json:"basicAuth"`
	BasicAuthUser   string            `json:"basicAuthUser"`
	WithCredentials bool              `json:"withCredentials"`
	IsDefault       bool              `json:"isDefault"`
	JsonData        *simplejson.Json  `json:"jsonData"`
	SecureJsonData  map[string]string `json:"secureJsonData"`
	Version         int               `json:"version"`
	Uid             string            `json:"uid"`

	OrgId                   int64             `json:"-"`
	Id                      int64             `json:"-"`
	ReadOnly                bool              `json:"-"`
	EncryptedSecureJsonData map[string][]byte `json:"-"`
	UpdateSecretFn          UpdateSecretFn    `json:"-"`

	Result *DataSource `json:"-"`
}

// DeleteDataSourceCommand will delete a DataSource based on OrgID as well as the UID (preferred), ID, or Name.
// At least one of the UID, ID, or Name properties must be set in addition to OrgID.
type DeleteDataSourceCommand struct {
	ID   int64
	UID  string
	Name string

	OrgID int64

	DeletedDatasourcesCount int64

	UpdateSecretFn UpdateSecretFn
}

// Correlations are uniquely identified by a SourceUid+TargetUid pair.

// AddCorrelationCommand adds a correlation
type AddCorrelationCommand struct {
	OrgID     int64
	SourceUID string
	TargetUID string `json:"targetUid" binding:"Required"`
	// TODO: Options

	Result Correlation
}

// UpdateCorrelationsCommand updates a correlation
type UpdateCorrelationsCommand struct {
	OrgId     int64  `json:"-"`
	Id        int64  `json:"-"`
	SourceUID string `json:"sourceUid" binding:"Required"`
	TargetUID string `json:"targetUid" binding:"Required"`
	// TODO: Options

	Correlations []Correlation

	Result []Correlation
}

// Function for updating secrets along with datasources, to ensure atomicity
type UpdateSecretFn func() error

// ---------------------
// QUERIES

type GetDataSourcesQuery struct {
	OrgId           int64
	DataSourceLimit int
	User            *SignedInUser
	Result          []*DataSource
}

type GetDataSourcesByTypeQuery struct {
	Type   string
	Result []*DataSource
}

type GetDefaultDataSourceQuery struct {
	OrgId  int64
	User   *SignedInUser
	Result *DataSource
}

// GetDataSourceQuery will get a DataSource based on OrgID as well as the UID (preferred), ID, or Name.
// At least one of the UID, ID, or Name properties must be set in addition to OrgID.
type GetDataSourceQuery struct {
	Id   int64
	Uid  string
	Name string

	OrgId int64

	Result *DataSource
}

// ---------------------
//  Permissions
// ---------------------

// Datasource permission
// Description:
// * `0` - No Access
// * `1` - Query
// Enum: 0,1
// swagger:model
type DsPermissionType int

const (
	DsPermissionNoAccess DsPermissionType = iota
	DsPermissionQuery
)

func (p DsPermissionType) String() string {
	names := map[int]string{
		int(DsPermissionQuery):    "Query",
		int(DsPermissionNoAccess): "No Access",
	}
	return names[int(p)]
}

type DatasourcesPermissionFilterQuery struct {
	User        *SignedInUser
	Datasources []*DataSource
	Result      []*DataSource
}
