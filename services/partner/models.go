package partner

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/partner"

// ContractProductDemandList ...
type ContractProductDemandList struct {
	PartnerMemberNo                   *string `json:"partnerMemberNo,omitempty"`
	PartnerMemberLoginID              *string `json:"partnerMemberLoginId,omitempty"`
	PartnerMemberName                 *string `json:"partnerMemberName,omitempty"`
	MemberNo                          *string `json:"memberNo,omitempty"`
	LoginID                           *string `json:"loginId,omitempty"`
	MemberName                        *string `json:"memberName,omitempty"`
	DemandYm                          *string `json:"demandYm,omitempty"`
	DemandTypeCode                    *string `json:"demandTypeCode,omitempty"`
	DemandTypeCodeName                *string `json:"demandTypeCodeName,omitempty"`
	DemandTypeDetailCode              *string `json:"demandTypeDetailCode,omitempty"`
	DemandTypeDetailCodeName          *string `json:"demandTypeDetailCodeName,omitempty"`
	ContractNo                        *string `json:"contractNo,omitempty"`
	ContractProductSequence           *string `json:"contractProductSequence,omitempty"`
	InstanceNo                        *string `json:"instanceNo,omitempty"`
	ProductPrice                      *string `json:"productPrice,omitempty"`
	UseQuantity                       *string `json:"useQuantity,omitempty"`
	UnitUseQuantity                   *string `json:"unitUseQuantity,omitempty"`
	FreeUseQuantity                   *string `json:"freeUseQuantity,omitempty"`
	FreeUnitUseQuantity               *string `json:"freeUnitUseQuantity,omitempty"`
	TotalUseQuantity                  *string `json:"totalUseQuantity,omitempty"`
	TotalUnitUseQuantity              *string `json:"totalUnitUseQuantity,omitempty"`
	PromiseDiscountAmount             *string `json:"promiseDiscountAmount,omitempty"`
	PromotionDiscountAmount           *string `json:"promotionDiscountAmount,omitempty"`
	EtcDiscountAmount                 *string `json:"etcDiscountAmount,omitempty"`
	DefaultAmount                     *string `json:"defaultAmount,omitempty"`
	UseAmount                         *string `json:"useAmount,omitempty"`
	DemandAmount                      *string `json:"demandAmount,omitempty"`
	RegionNo                          *string `json:"regionNo,omitempty"`
	RegionCode                        *string `json:"regionCode,omitempty"`
	RegionName                        *string `json:"regionName,omitempty"`
	SingleProductContractTypeCode     *string `json:"singleProductContractTypeCode,omitempty"`
	SingleProductContractTypeCodeName *string `json:"singleProductContractTypeCodeName,omitempty"`
	InstanceName                      *string `json:"instanceName,omitempty"`
	ContractStatusCode                *string `json:"contractStatusCode,omitempty"`
	ContractStatusCodeName            *string `json:"contractStatusCodeName,omitempty"`
	PlatformTypeCode                  *string `json:"platformTypeCode,omitempty"`
	ContractProductTypeCode           *string `json:"contractProductTypeCode,omitempty"`
	ContractProductTypeCodeName       *string `json:"contractProductTypeCodeName,omitempty"`
	ProductCode                       *string `json:"productCode,omitempty"`
	ProductName                       *string `json:"productName,omitempty"`
	ProductDescription                *string `json:"productDescription,omitempty"`
	BlockStorageSize                  *string `json:"blockStorageSize,omitempty"`
	BlockStorageUsageTypeCode         *string `json:"blockStorageUsageTypeCode,omitempty"`
	BlockStorageUsageTypeCodeName     *string `json:"blockStorageUsageTypeCodeName,omitempty"`
	ProductCount                      *string `json:"productCount,omitempty"`
	ServiceStatusCode                 *string `json:"serviceStatusCode,omitempty"`
	ServiceStatusCodeName             *string `json:"serviceStatusCodeName,omitempty"`
	ProductRatingTypeCode             *string `json:"productRatingTypeCode,omitempty"`
	ProductRatingTypeCodeName         *string `json:"productRatingTypeCodeName,omitempty"`
	LastContractProductYn             *string `json:"lastContractProductYn,omitempty"`
	ChangeTypeCode                    *string `json:"changeTypeCode,omitempty"`
	ChangeTypeCodeName                *string `json:"changeTypeCodeName,omitempty"`
	FeeSystemProductPrice             *string `json:"feeSystemProductPrice,omitempty"`
	FeeSystemBaseProductPrice         *string `json:"feeSystemBaseProductPrice,omitempty"`
	FeeSystemTypeCode                 *string `json:"feeSystemTypeCode,omitempty"`
	FeeSystemTypeCodeName             *string `json:"feeSystemTypeCodeName,omitempty"`
	RatingUnitTypeCode                *string `json:"ratingUnitTypeCode,omitempty"`
	RatingUnitTypeCodeName            *string `json:"ratingUnitTypeCodeName,omitempty"`
}

// ContractProductDemandListResponse ...
type ContractProductDemandListResponse struct {
	autorest.Response         `json:"-"`
	RequestID                 *string                      `json:"requestId,omitempty"`
	ReturnCode                *string                      `json:"returnCode,omitempty"`
	ReturnMessage             *string                      `json:"returnMessage,omitempty"`
	TotalRows                 *int32                       `json:"totalRows,omitempty"`
	ContractProductDemandList *[]ContractProductDemandList `json:"contractProductDemandList,omitempty"`
}

// DemandList ...
type DemandList struct {
	PartnerMemberNo              *string `json:"partnerMemberNo,omitempty"`
	PartnerMemberLoginID         *string `json:"partnerMemberLoginId,omitempty"`
	PartnerMemberName            *string `json:"partnerMemberName,omitempty"`
	MemberNo                     *string `json:"memberNo,omitempty"`
	LoginID                      *string `json:"loginId,omitempty"`
	MemberName                   *string `json:"memberName,omitempty"`
	DemandNo                     *string `json:"demandNo,omitempty"`
	IntegrationDemandNo          *string `json:"integrationDemandNo,omitempty"`
	DemandAttributeCode          *string `json:"demandAttributeCode,omitempty"`
	DemandAttributeCodeName      *string `json:"demandAttributeCodeName,omitempty"`
	DemandYm                     *string `json:"demandYm,omitempty"`
	BillingDemandYn              *string `json:"billingDemandYn,omitempty"`
	TotalUseAmount               *string `json:"totalUseAmount,omitempty"`
	TotalPromiseDiscountAmount   *string `json:"totalPromiseDiscountAmount,omitempty"`
	TotalPromotionDiscountAmount *string `json:"totalPromotionDiscountAmount,omitempty"`
	TotalEtcDiscountAmount       *string `json:"totalEtcDiscountAmount,omitempty"`
	TotalCustomerDiscountAmount  *string `json:"totalCustomerDiscountAmount,omitempty"`
	TotalProductDiscountAmount   *string `json:"totalProductDiscountAmount,omitempty"`
	TotalCreditDiscountAmount    *string `json:"totalCreditDiscountAmount,omitempty"`
	Under1000DiscountAmount      *string `json:"under1000DiscountAmount,omitempty"`
	Under100DiscountAmount       *string `json:"under100DiscountAmount,omitempty"`
	TotalDefaultAmount           *string `json:"totalDefaultAmount,omitempty"`
	ThisMonthDemandAmount        *string `json:"thisMonthDemandAmount,omitempty"`
	ThisMonthVatRatio            *string `json:"thisMonthVatRatio,omitempty"`
	ThisMonthVatAmount           *string `json:"thisMonthVatAmount,omitempty"`
	ThisMonthDemandVatSumAmount  *string `json:"thisMonthDemandVatSumAmount,omitempty"`
	TotalDemandAmount            *string `json:"totalDemandAmount,omitempty"`
	WriteYmdt                    *string `json:"writeYmdt,omitempty"`
	PaidUpYn                     *string `json:"paidUpYn,omitempty"`
	ThisMonthOverDueYn           *string `json:"thisMonthOverDueYn,omitempty"`
	ThisMonthOverDueDegreeCount  *string `json:"thisMonthOverDueDegreeCount,omitempty"`
	OverDuePlusAmount            *string `json:"overDuePlusAmount,omitempty"`
	OverDuePlusAmountRatio       *string `json:"overDuePlusAmountRatio,omitempty"`
	ThisMonthOverDueAmount       *string `json:"thisMonthOverDueAmount,omitempty"`
	BeforeMonthDemandNo          *string `json:"beforeMonthDemandNo,omitempty"`
	BeforeMonthDemandAmount      *string `json:"beforeMonthDemandAmount,omitempty"`
	BeforeMonthOverDueAmount     *string `json:"beforeMonthOverDueAmount,omitempty"`
	TotalOverDueAmount           *string `json:"totalOverDueAmount,omitempty"`
	TreatmentMeansCode           *string `json:"treatmentMeansCode,omitempty"`
	TreatmentMeansCodeName       *string `json:"treatmentMeansCodeName,omitempty"`
}

// DemandListResponse ...
type DemandListResponse struct {
	autorest.Response `json:"-"`
	RequestID         *string       `json:"requestId,omitempty"`
	ReturnCode        *string       `json:"returnCode,omitempty"`
	ReturnMessage     *string       `json:"returnMessage,omitempty"`
	TotalRows         *int32        `json:"totalRows,omitempty"`
	DemandList        *[]DemandList `json:"demandList,omitempty"`
}

// DemandListResponseType ...
type DemandListResponseType struct {
	autorest.Response `json:"-"`
	RequestID         *string           `json:"requestId,omitempty"`
	ReturnCode        *string           `json:"returnCode,omitempty"`
	ReturnMessage     *string           `json:"returnMessage,omitempty"`
	TotalRows         *int32            `json:"totalRows,omitempty"`
	PartnerDemandList *[]DemandListType `json:"partnerDemandList,omitempty"`
}

// DemandListType ...
type DemandListType struct {
	PartnerMemberNo                          *string `json:"partnerMemberNo,omitempty"`
	PartnerMemberLoginID                     *string `json:"partnerMemberLoginId,omitempty"`
	PartnerMemberName                        *string `json:"partnerMemberName,omitempty"`
	MemberNo                                 *string `json:"memberNo,omitempty"`
	LoginID                                  *string `json:"loginId,omitempty"`
	MemberName                               *string `json:"memberName,omitempty"`
	DemandYm                                 *string `json:"demandYm,omitempty"`
	TotalUseAmount                           *string `json:"totalUseAmount,omitempty"`
	TotalPromiseDiscountAmount               *string `json:"totalPromiseDiscountAmount,omitempty"`
	TotalPromotionDiscountAmount             *string `json:"totalPromotionDiscountAmount,omitempty"`
	TotalEtcDiscountAmount                   *string `json:"totalEtcDiscountAmount,omitempty"`
	TotalCustomerDiscountAmount              *string `json:"totalCustomerDiscountAmount,omitempty"`
	TotalProductDiscountAmount               *string `json:"totalProductDiscountAmount,omitempty"`
	TotalCreditDiscountAmount                *string `json:"totalCreditDiscountAmount,omitempty"`
	Under1000DiscountAmount                  *string `json:"under1000DiscountAmount,omitempty"`
	Under100DiscountAmount                   *string `json:"under100DiscountAmount,omitempty"`
	TotalDefaultAmount                       *string `json:"totalDefaultAmount,omitempty"`
	TotalDemandAmount                        *string `json:"totalDemandAmount,omitempty"`
	TotalCommissionExceptionProductUseAmount *string `json:"totalCommissionExceptionProductUseAmount,omitempty"`
	TotalExceptionProductCommissionAmount    *string `json:"totalExceptionProductCommissionAmount,omitempty"`
	CommissionRatio                          *string `json:"commissionRatio,omitempty"`
	CommissionAmount                         *string `json:"commissionAmount,omitempty"`
	IncentiveAmount                          *string `json:"incentiveAmount,omitempty"`
	PartnerDiscountAmount                    *string `json:"partnerDiscountAmount,omitempty"`
	PartnerCostAmount                        *string `json:"partnerCostAmount,omitempty"`
	PartnerCustomerDiscountAmount            *string `json:"partnerCustomerDiscountAmount,omitempty"`
	PartnerCustomerSellingAmount             *string `json:"partnerCustomerSellingAmount,omitempty"`
	PartnerTotalSellingAmount                *string `json:"partnerTotalSellingAmount,omitempty"`
	PartnerMarginalProfitAmount              *string `json:"partnerMarginalProfitAmount,omitempty"`
}

// DemandProductList ...
type DemandProductList struct {
	PartnerMemberNo                 *string `json:"partnerMemberNo,omitempty"`
	PartnerMemberLoginID            *string `json:"partnerMemberLoginId,omitempty"`
	PartnerMemberName               *string `json:"partnerMemberName,omitempty"`
	MemberNo                        *string `json:"memberNo,omitempty"`
	LoginID                         *string `json:"loginId,omitempty"`
	MemberName                      *string `json:"memberName,omitempty"`
	DemandYm                        *string `json:"demandYm,omitempty"`
	DemandTypeCreditProductCode     *string `json:"demandTypeCreditProductCode,omitempty"`
	DemandTypeCreditProductCodeName *string `json:"demandTypeCreditProductCodeName,omitempty"`
	PromiseDiscountAmount           *string `json:"promiseDiscountAmount,omitempty"`
	PromotionDiscountAmount         *string `json:"promotionDiscountAmount,omitempty"`
	EtcDiscountAmount               *string `json:"etcDiscountAmount,omitempty"`
	ProductDiscountAmount           *string `json:"productDiscountAmount,omitempty"`
	CreditDiscountAmount            *string `json:"creditDiscountAmount,omitempty"`
	DefaultAmount                   *string `json:"defaultAmount,omitempty"`
	UseAmount                       *string `json:"useAmount,omitempty"`
	DemandAmount                    *string `json:"demandAmount,omitempty"`
}

// DemandProductListResponse ...
type DemandProductListResponse struct {
	autorest.Response `json:"-"`
	RequestID         *string              `json:"requestId,omitempty"`
	ReturnCode        *string              `json:"returnCode,omitempty"`
	ReturnMessage     *string              `json:"returnMessage,omitempty"`
	TotalRows         *int32               `json:"totalRows,omitempty"`
	DemandProductList *[]DemandProductList `json:"demandProductList,omitempty"`
}
