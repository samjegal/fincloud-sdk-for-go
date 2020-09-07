package nasapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/samjegal/fincloud-sdk-for-go/services/nas"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	AddAccessControl(ctx context.Context, nasVolumeInstanceNo string, serverInstanceNoListN string) (result nas.VolumeAccessControlResponse, err error)
	ChangeSize(ctx context.Context, nasVolumeInstanceNo string, volumeSize string) (result nas.VolumeSizeResponse, err error)
	Create(ctx context.Context, volumeSize string, zoneCode string, volumeName string, volumeAllotmentProtocolTypeCode nas.VolumeAllotmentProtocolTypeCode, serverInstanceNoListN string, cifsUserName string, cifsUserPassword string, isEncryptedVolume nas.EncryptedVolume, nasVolumeDescription string) (result nas.VolumeInstancesResponse, err error)
	Delete(ctx context.Context, nasVolumeInstanceNoListN string) (result nas.VolumeInstancesResponse, err error)
	GetDetail(ctx context.Context, nasVolumeInstanceNo string) (result nas.VolumeInstanceDetailResponse, err error)
	GetList(ctx context.Context, volumeAllotmentProtocolTypeCode nas.VolumeAllotmentProtocolTypeCode, isEventConfiguration nas.EventConfiguration, isSnapshotConfiguration nas.SnapshotConfiguration, nasVolumeInstanceNoListN string, zoneCode string, pageNo string, pageSize string, volumeName string, sortedBy nas.SortedBy, sortingOrder nas.SortingOrder) (result nas.VolumeInstanceListResponse, err error)
	RemoveAccessControl(ctx context.Context, nasVolumeInstanceNo string, serverInstanceNoListN string) (result nas.VolumeAccessControlResponse, err error)
	SetAccessControl(ctx context.Context, nasVolumeInstanceNo string, serverInstanceNoListN string) (result nas.VolumeAccessControlResponse, err error)
}

var _ ClientAPI = (*nas.Client)(nil)
