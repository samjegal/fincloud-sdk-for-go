package registryapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/samjegal/fincloud-sdk-for-go/services/registry"
)

// ContainerRegistryClientAPI contains the set of methods on the ContainerRegistryClient type.
type ContainerRegistryClientAPI interface {
	Create(ctx context.Context, registry string, parameters registry.RepositoryRequest) (result registry.MessageResponse, err error)
	Delete(ctx context.Context, registry string) (result registry.MessageResponse, err error)
	DeleteImage(ctx context.Context, registry string, imageName string) (result registry.MessageResponse, err error)
	DeleteTagReference(ctx context.Context, registry string, imageName string, reference string) (result registry.MessageResponse, err error)
	GetImageDetail(ctx context.Context, registry string, imageName string) (result registry.SetObject, err error)
	GetImageList(ctx context.Context, registry string, page *int32, pagesize *int32) (result registry.SetObject, err error)
	GetList(ctx context.Context, page *int32, pagesize *int32) (result registry.SetObject, err error)
	UpdateImage(ctx context.Context, registry string, imageName string, parameters registry.UpdateRepositoryImageRequest) (result registry.MessageResponse, err error)
}

var _ ContainerRegistryClientAPI = (*registry.ContainerRegistryClient)(nil)

// ContainerRegistryImageClientAPI contains the set of methods on the ContainerRegistryImageClient type.
type ContainerRegistryImageClientAPI interface {
	GetTagList(ctx context.Context, registry string, imageName string, page *int32, pagesize *int32) (result registry.SetObject, err error)
	GetTagReference(ctx context.Context, registry string, imageName string, reference string) (result registry.SetObject, err error)
}

var _ ContainerRegistryImageClientAPI = (*registry.ContainerRegistryImageClient)(nil)
