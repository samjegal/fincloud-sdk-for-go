package partnerapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/samjegal/fincloud-sdk-for-go/services/partner"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	GetContractProductDemandList(ctx context.Context, responseFormatType string, pageNo *int32, pageSize *int32, startMonth string, endMonth string, loginID string) (result partner.ContractProductDemandListResponse, err error)
	GetDemandList(ctx context.Context, responseFormatType string, pageNo *int32, pageSize *int32, startMonth string, endMonth string, loginID string) (result partner.DemandListResponse, err error)
	GetDemandProductList(ctx context.Context, responseFormatType string, pageNo *int32, pageSize *int32, startMonth string, endMonth string, loginID string) (result partner.DemandProductListResponse, err error)
	GetPartnerDemandList(ctx context.Context, responseFormatType string, pageNo *int32, pageSize *int32, startMonth string, endMonth string, loginID string) (result partner.DemandListResponseType, err error)
}

var _ ClientAPI = (*partner.Client)(nil)
