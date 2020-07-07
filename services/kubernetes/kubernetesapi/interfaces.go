package kubernetesapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/kubernetes"
)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	Create(ctx context.Context, parameters kubernetes.ClusterRequestParameter) (result autorest.Response, err error)
	Delete(ctx context.Context, UUID string) (result autorest.Response, err error)
	Get(ctx context.Context, UUID string) (result kubernetes.ClusterResponseParameter, err error)
	List(ctx context.Context) (result kubernetes.ClustersListParameter, err error)
}

var _ ClustersClientAPI = (*kubernetes.ClustersClient)(nil)

// WorkerNodeClientAPI contains the set of methods on the WorkerNodeClient type.
type WorkerNodeClientAPI interface {
	Get(ctx context.Context, UUID string) (result kubernetes.WorkerNodeResponseParameter, err error)
}

var _ WorkerNodeClientAPI = (*kubernetes.WorkerNodeClient)(nil)

// NodePoolClientAPI contains the set of methods on the NodePoolClient type.
type NodePoolClientAPI interface {
	Create(ctx context.Context, UUID string, parameters kubernetes.NodePoolRequestParameter) (result autorest.Response, err error)
	Delete(ctx context.Context, UUID string, instanceNo string) (result autorest.Response, err error)
	Get(ctx context.Context, UUID string) (result kubernetes.NodePoolResponseParameter, err error)
	Update(ctx context.Context, UUID string, instanceNo string, parameters kubernetes.NodePoolUpdateParameter) (result autorest.Response, err error)
}

var _ NodePoolClientAPI = (*kubernetes.NodePoolClient)(nil)

// ConfigClientAPI contains the set of methods on the ConfigClient type.
type ConfigClientAPI interface {
	Get(ctx context.Context, UUID string) (result kubernetes.ConfigParameter, err error)
	Reset(ctx context.Context, UUID string) (result autorest.Response, err error)
}

var _ ConfigClientAPI = (*kubernetes.ConfigClient)(nil)
