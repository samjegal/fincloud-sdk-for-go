package cloudinsightapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/cloudinsight"
)

// ChartClientAPI contains the set of methods on the ChartClient type.
type ChartClientAPI interface {
	Preview(ctx context.Context, parameters cloudinsight.ChartDataWidgetRequest) (result cloudinsight.WidgetMetricInfoResponse, err error)
}

var _ ChartClientAPI = (*cloudinsight.ChartClient)(nil)

// DataClientAPI contains the set of methods on the DataClient type.
type DataClientAPI interface {
	Query(ctx context.Context, parameters cloudinsight.QueryRequest) (result cloudinsight.ListListFloat64, err error)
	QueryMultiple(ctx context.Context, parameters cloudinsight.QueryMultipleRequest) (result cloudinsight.ListMultipleDataParameter, err error)
}

var _ DataClientAPI = (*cloudinsight.DataClient)(nil)

// EventClientAPI contains the set of methods on the EventClient type.
type EventClientAPI interface {
	SearchByID(ctx context.Context, parameters cloudinsight.EventSearchRequest) (result cloudinsight.ListEventSearchResultParameter, err error)
	SearchEventCount(ctx context.Context, parameters cloudinsight.SearchEventCountConsoleRequest) (result cloudinsight.SearchEventCountConsoleResponse, err error)
}

var _ EventClientAPI = (*cloudinsight.EventClient)(nil)

// NotificationClientAPI contains the set of methods on the NotificationClient type.
type NotificationClientAPI interface {
	Get(ctx context.Context, grpID string) (result cloudinsight.EventRuleResponse, err error)
}

var _ NotificationClientAPI = (*cloudinsight.NotificationClient)(nil)

// MetricClientAPI contains the set of methods on the MetricClient type.
type MetricClientAPI interface {
	GetGroupItemsID(ctx context.Context, count int32) (result cloudinsight.ListString, err error)
	SearchList(ctx context.Context, parameters cloudinsight.MetricListRequest) (result autorest.Response, err error)
}

var _ MetricClientAPI = (*cloudinsight.MetricClient)(nil)

// RuleGroupClientAPI contains the set of methods on the RuleGroupClient type.
type RuleGroupClientAPI interface {
	Copy(ctx context.Context, ID string) (result cloudinsight.String, err error)
	CopyAsgGroup(ctx context.Context, parameters cloudinsight.RuleGroupCopyForAsgGroupRequest) (result cloudinsight.ListSetObject, err error)
	CopySettings(ctx context.Context, parameters cloudinsight.RuleGroupCopySettingRequest) (result cloudinsight.ListInt64, err error)
	Create(ctx context.Context, parameters cloudinsight.RuleGroupRequest) (result cloudinsight.Int64, err error)
	CreateDirectly(ctx context.Context, parameters cloudinsight.DirectRuleGroupCreateRequest) (result cloudinsight.String, err error)
	Delete(ctx context.Context, parameters cloudinsight.RuleGroupDeleteRequest) (result autorest.Response, err error)
	DeleteByProdKeyAndID(ctx context.Context, prodKey string, ID string) (result autorest.Response, err error)
	GetByMonitorGroupIds(ctx context.Context, prodKey string, parameters []string) (result cloudinsight.ListRuleGroupItemListParameter, err error)
	Query(ctx context.Context, parameters cloudinsight.RuleGroupListQueryRequest) (result cloudinsight.RuleGroupListQueryResponse, err error)
	QueryByProdKeyAndID(ctx context.Context, prodKey string, ID string) (result cloudinsight.RuleGroupParameter, err error)
	Update(ctx context.Context, parameters cloudinsight.RuleGroupRequest) (result autorest.Response, err error)
}

var _ RuleGroupClientAPI = (*cloudinsight.RuleGroupClient)(nil)

// MonitorGroupClientAPI contains the set of methods on the MonitorGroupClient type.
type MonitorGroupClientAPI interface {
	DeleteForce(ctx context.Context, prodKey string, parameters []cloudinsight.TypeGroupRelatedRuleParameter) (result autorest.Response, err error)
	DeleteForce1(ctx context.Context, prodKey string, parameters []cloudinsight.TypeGroupRelatedRuleParameter) (result cloudinsight.String, err error)
	Get(ctx context.Context, prodKey string, ID string) (result cloudinsight.MonitorGroupParameter, err error)
	GetByMonitorGroupIds(ctx context.Context, prodKey string, parameters []string) (result cloudinsight.String, err error)
	List(ctx context.Context, prodKey string) (result cloudinsight.ListMonitorGroupParameter, err error)
	RemoveResourceFromRules(ctx context.Context, parameters cloudinsight.RemoveResourceFromRulesParameter) (result autorest.Response, err error)
}

var _ MonitorGroupClientAPI = (*cloudinsight.MonitorGroupClient)(nil)

// MetricGroupClientAPI contains the set of methods on the MetricGroupClient type.
type MetricGroupClientAPI interface {
	Create(ctx context.Context, parameters cloudinsight.MetricsGroupRequest) (result cloudinsight.String, err error)
	Delete(ctx context.Context, prodKey string, parameters []string) (result autorest.Response, err error)
	DeleteByProdKeyAndID(ctx context.Context, prodKey string, ID string) (result autorest.Response, err error)
	Get(ctx context.Context, prodKey string, ID string) (result cloudinsight.MetricsGroupParameter, err error)
	List(ctx context.Context, prodKey string) (result cloudinsight.MetricGroupListResponse, err error)
	Update(ctx context.Context, parameters cloudinsight.MetricsGroupRequest) (result autorest.Response, err error)
}

var _ MetricGroupClientAPI = (*cloudinsight.MetricGroupClient)(nil)

// MonitorClientAPI contains the set of methods on the MonitorClient type.
type MonitorClientAPI interface {
	Create(ctx context.Context, parameters cloudinsight.MonitorGroupRequest) (result cloudinsight.String, err error)
}

var _ MonitorClientAPI = (*cloudinsight.MonitorClient)(nil)

// SchemaClientAPI contains the set of methods on the SchemaClient type.
type SchemaClientAPI interface {
	Delete(ctx context.Context, cwKey string, prodName string) (result autorest.Response, err error)
	Get(ctx context.Context, prodName string, cwKey string) (result cloudinsight.SchemaResponse, err error)
	Register(ctx context.Context, parameters cloudinsight.SchemaRequest) (result cloudinsight.SchemaRegisterResponse, err error)
	Update(ctx context.Context, parameters cloudinsight.SchemaRequest) (result cloudinsight.ScehmaUpdateResponse, err error)
}

var _ SchemaClientAPI = (*cloudinsight.SchemaClient)(nil)

// ExtendedClientAPI contains the set of methods on the ExtendedClient type.
type ExtendedClientAPI interface {
	Disable(ctx context.Context, cwKey string, instanceIds string) (result cloudinsight.SchemaExtendedDisableResponse, err error)
	Enable(ctx context.Context, cwKey string, instanceIds string) (result cloudinsight.SchemaExtendedEnableResponse, err error)
	Status(ctx context.Context, cwKey string, instanceIds string) (result cloudinsight.ListSchemaExtendedStatusParameter, err error)
}

var _ ExtendedClientAPI = (*cloudinsight.ExtendedClient)(nil)

// ServerClientAPI contains the set of methods on the ServerClient type.
type ServerClientAPI interface {
	GetTop(ctx context.Context, query cloudinsight.SeverTargetMetric) (result cloudinsight.ListServerTopMetricParameter, err error)
}

var _ ServerClientAPI = (*cloudinsight.ServerClient)(nil)

// CollectorClientAPI contains the set of methods on the CollectorClient type.
type CollectorClientAPI interface {
	SendMethod(ctx context.Context, parameters cloudinsight.CollectorRequest) (result cloudinsight.CollectorResponse, err error)
}

var _ CollectorClientAPI = (*cloudinsight.CollectorClient)(nil)

// ProcessPluginClientAPI contains the set of methods on the ProcessPluginClient type.
type ProcessPluginClientAPI interface {
	Create(ctx context.Context, parameters cloudinsight.ProcessPluginRequest) (result autorest.Response, err error)
	Get(ctx context.Context, instanceNo string) (result cloudinsight.ProcessPluginDetailResponse, err error)
	List(ctx context.Context) (result cloudinsight.ListProcessPluginParameter, err error)
}

var _ ProcessPluginClientAPI = (*cloudinsight.ProcessPluginClient)(nil)

// PortPluginClientAPI contains the set of methods on the PortPluginClient type.
type PortPluginClientAPI interface {
	Create(ctx context.Context, parameters cloudinsight.PortPluginRequest) (result autorest.Response, err error)
	Get(ctx context.Context, instanceNo string) (result cloudinsight.PortPluginDetailResponse, err error)
	List(ctx context.Context) (result cloudinsight.ListPortPluginParameter, err error)
}

var _ PortPluginClientAPI = (*cloudinsight.PortPluginClient)(nil)

// FilePluginClientAPI contains the set of methods on the FilePluginClient type.
type FilePluginClientAPI interface {
	Create(ctx context.Context, parameters cloudinsight.FilePluginRequest) (result autorest.Response, err error)
	Get(ctx context.Context, instanceNo string) (result cloudinsight.FilePluginDetailResponse, err error)
	List(ctx context.Context) (result cloudinsight.ListFilePluginParameter, err error)
}

var _ FilePluginClientAPI = (*cloudinsight.FilePluginClient)(nil)
