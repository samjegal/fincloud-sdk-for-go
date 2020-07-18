package cloudinsightapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/cloudinsight"
)

// DataClientAPI contains the set of methods on the DataClient type.
type DataClientAPI interface {
	Preview(ctx context.Context) (result autorest.Response, err error)
	Query(ctx context.Context, parameters cloudinsight.QueryParameter) (result cloudinsight.ListListFloat64, err error)
	QueryMultiple(ctx context.Context, parameters cloudinsight.QueryMultipleParameter) (result cloudinsight.ListDataInfoParameter, err error)
}

var _ DataClientAPI = (*cloudinsight.DataClient)(nil)

// EventClientAPI contains the set of methods on the EventClient type.
type EventClientAPI interface {
	SearchEventCount(ctx context.Context) (result autorest.Response, err error)
}

var _ EventClientAPI = (*cloudinsight.EventClient)(nil)

// RuleClientAPI contains the set of methods on the RuleClient type.
type RuleClientAPI interface {
	Create(ctx context.Context) (result autorest.Response, err error)
	Delete(ctx context.Context, ruleID string) (result autorest.Response, err error)
	Get(ctx context.Context, ruleID string) (result autorest.Response, err error)
	GetList(ctx context.Context, pageNum string, pageSize string) (result autorest.Response, err error)
	Update(ctx context.Context, ruleID string) (result autorest.Response, err error)
}

var _ RuleClientAPI = (*cloudinsight.RuleClient)(nil)

// SchemaClientAPI contains the set of methods on the SchemaClient type.
type SchemaClientAPI interface {
	Create(ctx context.Context, parameters cloudinsight.SchemaParameter, prodName string) (result autorest.Response, err error)
	Delete(ctx context.Context, prodName string) (result autorest.Response, err error)
	Disable(ctx context.Context, cwKey string, instanceIds string) (result autorest.Response, err error)
	Enable(ctx context.Context, cwKey string, instanceIds string) (result autorest.Response, err error)
	Get(ctx context.Context, prodName string) (result autorest.Response, err error)
	GetList(ctx context.Context) (result autorest.Response, err error)
	QueryStatus(ctx context.Context, cwKey string, instanceIds string) (result autorest.Response, err error)
	Update(ctx context.Context, prodName string) (result autorest.Response, err error)
}

var _ SchemaClientAPI = (*cloudinsight.SchemaClient)(nil)

// ServerClientAPI contains the set of methods on the ServerClient type.
type ServerClientAPI interface {
	CreatTop(ctx context.Context, query string) (result autorest.Response, err error)
}

var _ ServerClientAPI = (*cloudinsight.ServerClient)(nil)

// CollectorClientAPI contains the set of methods on the CollectorClient type.
type CollectorClientAPI interface {
	Push(ctx context.Context, parameters cloudinsight.CollectorParameter) (result autorest.Response, err error)
}

var _ CollectorClientAPI = (*cloudinsight.CollectorClient)(nil)

// ProcessPluginClientAPI contains the set of methods on the ProcessPluginClient type.
type ProcessPluginClientAPI interface {
	Create(ctx context.Context, parameters cloudinsight.ProcessPluginParameter) (result autorest.Response, err error)
	Get(ctx context.Context, instanceNo string) (result cloudinsight.ProcessPluginParameter, err error)
	List(ctx context.Context) (result cloudinsight.ListProcessPluginParameter, err error)
}

var _ ProcessPluginClientAPI = (*cloudinsight.ProcessPluginClient)(nil)

// PortPluginClientAPI contains the set of methods on the PortPluginClient type.
type PortPluginClientAPI interface {
	Create(ctx context.Context, parameters cloudinsight.PortPluginParameter) (result autorest.Response, err error)
	Get(ctx context.Context, instanceNo string) (result cloudinsight.PortPluginParameter, err error)
	List(ctx context.Context) (result cloudinsight.ListPortPluginParameter, err error)
}

var _ PortPluginClientAPI = (*cloudinsight.PortPluginClient)(nil)
