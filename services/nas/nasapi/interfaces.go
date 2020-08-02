package nasapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/nas"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	AddAccessControl(ctx context.Context, responseFormatType string, nasVolumeInstanceNo string, serverInstanceNoListN string, regionCode string) (result autorest.Response, err error)
	ChangeSize(ctx context.Context, responseFormatType string, nasVolumeInstanceNo string, volumeSize string, regionCode string) (result autorest.Response, err error)
	Create(ctx context.Context, responseFormatType string, volumeSize string, regionCode string, zoneCode string, volumeName string, volumeAllotmentProtocolTypeCode nas.VolumeAllotmentProtocolTypeCode, serverInstanceNoListN string, cifsUserName string, cifsUserPassword string, isEncryptedVolume nas.EncryptedVolume, nasVolumeDescription string) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, nasVolumeInstanceNoListN string, regionCode string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, nasVolumeInstanceNo string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, volumeAllotmentProtocolTypeCode nas.VolumeAllotmentProtocolTypeCode, isEventConfiguration nas.EventConfiguration, isSnapshotConfiguration nas.SnapshotConfiguration, nasVolumeInstanceNoListN string, zoneCode string, pageNo string, pageSize string, volumeName string, sortedBy nas.SortedBy, sortingOrder nas.SortingOrder) (result autorest.Response, err error)
	RemoveAccessControl(ctx context.Context, responseFormatType string, nasVolumeInstanceNo string, serverInstanceNoListN string, regionCode string) (result autorest.Response, err error)
	SetAccessControl(ctx context.Context, responseFormatType string, nasVolumeInstanceNo string, regionCode string, serverInstanceNoListN string) (result autorest.Response, err error)
}

var _ ClientAPI = (*nas.Client)(nil)
