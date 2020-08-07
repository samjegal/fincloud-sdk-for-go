package partner

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/partner"

// ContractProductDemandList ...
type ContractProductDemandList struct {
	// PartnerMemberNo - 파트너 회원번호 (오너)
	PartnerMemberNo *string `json:"partnerMemberNo,omitempty"`
	// PartnerMemberLoginID - 파트너 회원 로그인 아이디 (오너)
	PartnerMemberLoginID *string `json:"partnerMemberLoginId,omitempty"`
	// PartnerMemberName - 파트너 회원명 (오너)
	PartnerMemberName *string `json:"partnerMemberName,omitempty"`
	// MemberNo - 파트너 회원번호 (멤버)
	MemberNo *string `json:"memberNo,omitempty"`
	// LoginID - 파트너 회원 로그인 아이디 (멤버)
	LoginID *string `json:"loginId,omitempty"`
	// MemberName - 파트너 회원명 (멤버)
	MemberName *string `json:"memberName,omitempty"`
	// DemandYm - 청구 연월
	DemandYm *string `json:"demandYm,omitempty"`
	// DemandTypeCode - 청구 유형 코드
	DemandTypeCode *string `json:"demandTypeCode,omitempty"`
	// DemandTypeCodeName - 청구 유형 코드명
	DemandTypeCodeName *string `json:"demandTypeCodeName,omitempty"`
	// DemandTypeDetailCode - 청구 유형 상세코드
	DemandTypeDetailCode *string `json:"demandTypeDetailCode,omitempty"`
	// DemandTypeDetailCodeName - 청구 유형 상세코드명
	DemandTypeDetailCodeName *string `json:"demandTypeDetailCodeName,omitempty"`
	// ContractNo - 계약번호
	ContractNo *string `json:"contractNo,omitempty"`
	// ContractProductSequence - 계약상품 일련 번호
	ContractProductSequence *string `json:"contractProductSequence,omitempty"`
	// InstanceNo - 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
	// ProductPrice - 상품가격 (정산기준)
	ProductPrice *string `json:"productPrice,omitempty"`
	// UseQuantity - 사용량
	UseQuantity *string `json:"useQuantity,omitempty"`
	// UnitUseQuantity - 단위 사용량
	UnitUseQuantity *string `json:"unitUseQuantity,omitempty"`
	// FreeUseQuantity - 무료 사용량
	FreeUseQuantity *string `json:"freeUseQuantity,omitempty"`
	// FreeUnitUseQuantity - 무료 단위 사용량
	FreeUnitUseQuantity *string `json:"freeUnitUseQuantity,omitempty"`
	// TotalUseQuantity - 총 사용량
	TotalUseQuantity *string `json:"totalUseQuantity,omitempty"`
	// TotalUnitUseQuantity - 총 단위 사용량
	TotalUnitUseQuantity *string `json:"totalUnitUseQuantity,omitempty"`
	// PromiseDiscountAmount - 약정 할인 금액
	PromiseDiscountAmount *string `json:"promiseDiscountAmount,omitempty"`
	// PromotionDiscountAmount - 프로모션 할인 금액
	PromotionDiscountAmount *string `json:"promotionDiscountAmount,omitempty"`
	// EtcDiscountAmount - 기타 할인 금액
	EtcDiscountAmount *string `json:"etcDiscountAmount,omitempty"`
	// DefaultAmount - 위약 할인 금액
	DefaultAmount *string `json:"defaultAmount,omitempty"`
	// UseAmount - 사용 금액
	UseAmount *string `json:"useAmount,omitempty"`
	// DemandAmount - 청구 금액
	DemandAmount *string `json:"demandAmount,omitempty"`
	// RegionNo - 리전 번호
	RegionNo *string `json:"regionNo,omitempty"`
	// RegionCode - 리전 코드
	RegionCode *string `json:"regionCode,omitempty"`
	// RegionName - 리전명
	RegionName *string `json:"regionName,omitempty"`
	// SingleProductContractTypeCode - 단품 계약 구분 코드
	SingleProductContractTypeCode *string `json:"singleProductContractTypeCode,omitempty"`
	// SingleProductContractTypeCodeName - 단품 계약 구분 코드명
	SingleProductContractTypeCodeName *string `json:"singleProductContractTypeCodeName,omitempty"`
	// InstanceName - 인스턴스명
	InstanceName *string `json:"instanceName,omitempty"`
	// ContractStatusCode - 계약 시작 일시
	ContractStatusCode *string `json:"contractStatusCode,omitempty"`
	// ContractStatusCodeName - 계약 종료 일시
	ContractStatusCodeName *string `json:"contractStatusCodeName,omitempty"`
	// PlatformTypeCode - 플랫폼 유형 코드
	PlatformTypeCode *string `json:"platformTypeCode,omitempty"`
	// ContractProductTypeCode - 계약 상태 코드
	ContractProductTypeCode *string `json:"contractProductTypeCode,omitempty"`
	// ContractProductTypeCodeName - 계약 상태 코드명
	ContractProductTypeCodeName *string `json:"contractProductTypeCodeName,omitempty"`
	// ProductCode - 상품 코드
	ProductCode *string `json:"productCode,omitempty"`
	// ProductName - 상품명
	ProductName *string `json:"productName,omitempty"`
	// ProductDescription - 상품 설명
	ProductDescription *string `json:"productDescription,omitempty"`
	// BlockStorageSize - 블록스토리지 사이즈
	BlockStorageSize *string `json:"blockStorageSize,omitempty"`
	// BlockStorageUsageTypeCode - 블록스토리지 용도 구분 코드
	BlockStorageUsageTypeCode *string `json:"blockStorageUsageTypeCode,omitempty"`
	// BlockStorageUsageTypeCodeName - 블록스토리지 용도 구분 코드명
	BlockStorageUsageTypeCodeName *string `json:"blockStorageUsageTypeCodeName,omitempty"`
	// ProductCount - 상품수
	ProductCount *string `json:"productCount,omitempty"`
	// ServiceStatusCode - 서비스 시작 일시
	ServiceStatusCode *string `json:"serviceStatusCode,omitempty"`
	// ServiceStatusCodeName - 서비스 종료 일시
	ServiceStatusCodeName *string `json:"serviceStatusCodeName,omitempty"`
	// ProductRatingTypeCode - 상품 과금 유형 코드
	ProductRatingTypeCode *string `json:"productRatingTypeCode,omitempty"`
	// ProductRatingTypeCodeName - 상품 과금 유형 코드명
	ProductRatingTypeCodeName *string `json:"productRatingTypeCodeName,omitempty"`
	// LastContractProductYn - 최종 계약 상품 여부
	LastContractProductYn *string `json:"lastContractProductYn,omitempty"`
	// ChangeTypeCode - 변경 유형 코드
	ChangeTypeCode *string `json:"changeTypeCode,omitempty"`
	// ChangeTypeCodeName - 변경 유형 코드명
	ChangeTypeCodeName *string `json:"changeTypeCodeName,omitempty"`
	// FeeSystemProductPrice - 상품 가격
	FeeSystemProductPrice *string `json:"feeSystemProductPrice,omitempty"`
	// FeeSystemBaseProductPrice - 정지 상품 가격
	FeeSystemBaseProductPrice *string `json:"feeSystemBaseProductPrice,omitempty"`
	// FeeSystemTypeCode - 요금제 구분 코드
	FeeSystemTypeCode *string `json:"feeSystemTypeCode,omitempty"`
	// FeeSystemTypeCodeName - 요금제 구분 코드명
	FeeSystemTypeCodeName *string `json:"feeSystemTypeCodeName,omitempty"`
	// RatingUnitTypeCode - 단위 구분 코드
	RatingUnitTypeCode *string `json:"ratingUnitTypeCode,omitempty"`
	// RatingUnitTypeCodeName - 단위 구분 코드명
	RatingUnitTypeCodeName *string `json:"ratingUnitTypeCodeName,omitempty"`
}

// ContractProductDemandListResponse ...
type ContractProductDemandListResponse struct {
	autorest.Response `json:"-"`
	// RequestID - 요청 ID
	RequestID *string `json:"requestId,omitempty"`
	// ReturnCode - 반환 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 반환 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows                 *int32                       `json:"totalRows,omitempty"`
	ContractProductDemandList *[]ContractProductDemandList `json:"contractProductDemandList,omitempty"`
}

// DemandList ...
type DemandList struct {
	// PartnerMemberNo - 파트너 회원번호 (오너)
	PartnerMemberNo *string `json:"partnerMemberNo,omitempty"`
	// PartnerMemberLoginID - 파트너 회원 로그인 아이디 (오너)
	PartnerMemberLoginID *string `json:"partnerMemberLoginId,omitempty"`
	// PartnerMemberName - 파트너 회원명 (오너)
	PartnerMemberName *string `json:"partnerMemberName,omitempty"`
	// MemberNo - 파트너 회원번호 (멤버)
	MemberNo *string `json:"memberNo,omitempty"`
	// LoginID - 파트너 회원 로그인 아이디 (멤버)
	LoginID *string `json:"loginId,omitempty"`
	// MemberName - 파트너 회원명 (멤버)
	MemberName *string `json:"memberName,omitempty"`
	// DemandNo - 청구번호
	DemandNo *string `json:"demandNo,omitempty"`
	// IntegrationDemandNo - 통합 청구번호
	IntegrationDemandNo *string `json:"integrationDemandNo,omitempty"`
	// DemandAttributeCode - 청구 속성 코드
	DemandAttributeCode *string `json:"demandAttributeCode,omitempty"`
	// DemandAttributeCodeName - 청구 속성 코드명
	DemandAttributeCodeName *string `json:"demandAttributeCodeName,omitempty"`
	// DemandYm - 청구연월
	DemandYm *string `json:"demandYm,omitempty"`
	// BillingDemandYn - 빌링 청구 여부
	BillingDemandYn *string `json:"billingDemandYn,omitempty"`
	// TotalUseAmount - 총 사용 금액
	TotalUseAmount *string `json:"totalUseAmount,omitempty"`
	// TotalPromiseDiscountAmount - 총 약정 할인 금액
	TotalPromiseDiscountAmount *string `json:"totalPromiseDiscountAmount,omitempty"`
	// TotalPromotionDiscountAmount - 총 프로모션 할인 금액
	TotalPromotionDiscountAmount *string `json:"totalPromotionDiscountAmount,omitempty"`
	// TotalEtcDiscountAmount - 총 기타 할인 금액
	TotalEtcDiscountAmount *string `json:"totalEtcDiscountAmount,omitempty"`
	// TotalCustomerDiscountAmount - 총 고객 할인 금액
	TotalCustomerDiscountAmount *string `json:"totalCustomerDiscountAmount,omitempty"`
	// TotalProductDiscountAmount - 총 상품 할인 금액
	TotalProductDiscountAmount *string `json:"totalProductDiscountAmount,omitempty"`
	// TotalCreditDiscountAmount - 총 크레딧 할인 금액
	TotalCreditDiscountAmount *string `json:"totalCreditDiscountAmount,omitempty"`
	// Under1000DiscountAmount - 1000원 미만 할인 금액
	Under1000DiscountAmount *string `json:"under1000DiscountAmount,omitempty"`
	// Under100DiscountAmount - 100원 미만 할인 금액
	Under100DiscountAmount *string `json:"under100DiscountAmount,omitempty"`
	// TotalDefaultAmount - 총 위약 금액
	TotalDefaultAmount *string `json:"totalDefaultAmount,omitempty"`
	// ThisMonthDemandAmount - 당월 청구 금액
	ThisMonthDemandAmount *string `json:"thisMonthDemandAmount,omitempty"`
	// ThisMonthVatRatio - 당월 부가세 비율
	ThisMonthVatRatio *string `json:"thisMonthVatRatio,omitempty"`
	// ThisMonthVatAmount - 당월 부가가치세 금액
	ThisMonthVatAmount *string `json:"thisMonthVatAmount,omitempty"`
	// ThisMonthDemandVatSumAmount - 당월 청구 부가가치세 합계 금액
	ThisMonthDemandVatSumAmount *string `json:"thisMonthDemandVatSumAmount,omitempty"`
	// TotalDemandAmount - 총 청구 금액
	TotalDemandAmount *string `json:"totalDemandAmount,omitempty"`
	// WriteYmdt - 작성 일시
	WriteYmdt *string `json:"writeYmdt,omitempty"`
	// PaidUpYn - 납입 여부
	PaidUpYn *string `json:"paidUpYn,omitempty"`
	// PaidUpYmdt - 납입 일시
	PaidUpYmdt *string `json:"paidUpYmdt,omitempty"`
	// ThisMonthOverDueYn - 당월 연체 여부
	ThisMonthOverDueYn *string `json:"thisMonthOverDueYn,omitempty"`
	// ThisMonthOverDueDegreeCount - 당월 연체 누적 차수
	ThisMonthOverDueDegreeCount *string `json:"thisMonthOverDueDegreeCount,omitempty"`
	// OverDuePlusAmount - 연체 가산 금액
	OverDuePlusAmount *string `json:"overDuePlusAmount,omitempty"`
	// OverDuePlusAmountRatio - 연체 가산 금액 비율
	OverDuePlusAmountRatio *string `json:"overDuePlusAmountRatio,omitempty"`
	// ThisMonthOverDueAmount - 당월 가산 금액
	ThisMonthOverDueAmount *string `json:"thisMonthOverDueAmount,omitempty"`
	// BeforeMonthDemandNo - 전월 청구 번호
	BeforeMonthDemandNo *string `json:"beforeMonthDemandNo,omitempty"`
	// BeforeMonthDemandAmount - 전월 청구 금액
	BeforeMonthDemandAmount *string `json:"beforeMonthDemandAmount,omitempty"`
	// BeforeMonthOverDueAmount - 전월 연체 금액
	BeforeMonthOverDueAmount *string `json:"beforeMonthOverDueAmount,omitempty"`
	// TotalOverDueAmount - 총 연체 금액
	TotalOverDueAmount *string `json:"totalOverDueAmount,omitempty"`
	// TreatmentMeansCode - 처리 수단 코드
	TreatmentMeansCode *string `json:"treatmentMeansCode,omitempty"`
	// TreatmentMeansCodeName - 처리 수단 코드명
	TreatmentMeansCodeName *string `json:"treatmentMeansCodeName,omitempty"`
}

// DemandListResponse ...
type DemandListResponse struct {
	autorest.Response `json:"-"`
	// RequestID - 요청 ID
	RequestID *string `json:"requestId,omitempty"`
	// ReturnCode - 반환 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 반환 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows  *int32        `json:"totalRows,omitempty"`
	DemandList *[]DemandList `json:"demandList,omitempty"`
}

// DemandListResponseType ...
type DemandListResponseType struct {
	autorest.Response `json:"-"`
	// RequestID - 요청 ID
	RequestID *string `json:"requestId,omitempty"`
	// ReturnCode - 반환 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 반환 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows         *int32            `json:"totalRows,omitempty"`
	PartnerDemandList *[]DemandListType `json:"partnerDemandList,omitempty"`
}

// DemandListType ...
type DemandListType struct {
	// PartnerMemberNo - 파트너 회원번호 (오너)
	PartnerMemberNo *string `json:"partnerMemberNo,omitempty"`
	// PartnerMemberLoginID - 파트너 회원 로그인 아이디 (오너)
	PartnerMemberLoginID *string `json:"partnerMemberLoginId,omitempty"`
	// PartnerMemberName - 파트너 회원번호 (오너)
	PartnerMemberName *string `json:"partnerMemberName,omitempty"`
	// MemberNo - 파트너 회원번호 (멤버)
	MemberNo *string `json:"memberNo,omitempty"`
	// LoginID - 파트너 회원 로그인 아이디 (멤버)
	LoginID *string `json:"loginId,omitempty"`
	// MemberName - 파트너 회원명 (멤버)
	MemberName *string `json:"memberName,omitempty"`
	// DemandYm - 청구년월
	DemandYm *string `json:"demandYm,omitempty"`
	// TotalUseAmount - 총 사용금액
	TotalUseAmount *string `json:"totalUseAmount,omitempty"`
	// TotalPromiseDiscountAmount - 총 약정 할인 금액
	TotalPromiseDiscountAmount *string `json:"totalPromiseDiscountAmount,omitempty"`
	// TotalPromotionDiscountAmount - 총 프로모션 할인 금액
	TotalPromotionDiscountAmount *string `json:"totalPromotionDiscountAmount,omitempty"`
	// TotalEtcDiscountAmount - 총 기타 할인 금액
	TotalEtcDiscountAmount *string `json:"totalEtcDiscountAmount,omitempty"`
	// TotalCustomerDiscountAmount - 총 고객 할인 금액
	TotalCustomerDiscountAmount *string `json:"totalCustomerDiscountAmount,omitempty"`
	// TotalProductDiscountAmount - 총 상품 할인 금액
	TotalProductDiscountAmount *string `json:"totalProductDiscountAmount,omitempty"`
	// TotalCreditDiscountAmount - 총 크레딧 할인 금액
	TotalCreditDiscountAmount *string `json:"totalCreditDiscountAmount,omitempty"`
	// Under1000DiscountAmount - 1000원 미만 할인 금액
	Under1000DiscountAmount *string `json:"under1000DiscountAmount,omitempty"`
	// Under100DiscountAmount - 100원 미만 할인 금액
	Under100DiscountAmount *string `json:"under100DiscountAmount,omitempty"`
	// TotalDefaultAmount - 총 위약 금액
	TotalDefaultAmount *string `json:"totalDefaultAmount,omitempty"`
	// TotalDemandAmount - 총 청구 금액
	TotalDemandAmount *string `json:"totalDemandAmount,omitempty"`
	// TotalCommissionExceptionProductUseAmount - 총 커미션 예외 상품 사용 금액
	TotalCommissionExceptionProductUseAmount *string `json:"totalCommissionExceptionProductUseAmount,omitempty"`
	// TotalExceptionProductCommissionAmount - 총 예외 상품 커미션 금액
	TotalExceptionProductCommissionAmount *string `json:"totalExceptionProductCommissionAmount,omitempty"`
	// CommissionRatio - 커미션 비율
	CommissionRatio *string `json:"commissionRatio,omitempty"`
	// CommissionAmount - 커미션 금액
	CommissionAmount *string `json:"commissionAmount,omitempty"`
	// IncentiveAmount - 인센티브 금액
	IncentiveAmount *string `json:"incentiveAmount,omitempty"`
	// PartnerDiscountAmount - 파트너 할인 금액
	PartnerDiscountAmount *string `json:"partnerDiscountAmount,omitempty"`
	// PartnerCostAmount - 파트너 비용 금액
	PartnerCostAmount *string `json:"partnerCostAmount,omitempty"`
	// PartnerCustomerDiscountAmount - 파트너 고객 할인 금액
	PartnerCustomerDiscountAmount *string `json:"partnerCustomerDiscountAmount,omitempty"`
	// PartnerCustomerSellingAmount - 파트너 고객 매출 금액
	PartnerCustomerSellingAmount *string `json:"partnerCustomerSellingAmount,omitempty"`
	// PartnerTotalSellingAmount - 파트너 총 매출 금액
	PartnerTotalSellingAmount *string `json:"partnerTotalSellingAmount,omitempty"`
	// PartnerMarginalProfitAmount - 파트너 차익 금액 (손익 금액)
	PartnerMarginalProfitAmount *string `json:"partnerMarginalProfitAmount,omitempty"`
}

// DemandProductList ...
type DemandProductList struct {
	// PartnerMemberNo - 파트너 회원번호 (오너)
	PartnerMemberNo *string `json:"partnerMemberNo,omitempty"`
	// PartnerMemberLoginID - 파트너 회원 로그인 아이디 (오너)
	PartnerMemberLoginID *string `json:"partnerMemberLoginId,omitempty"`
	// PartnerMemberName - 파트너 회원명 (오너)
	PartnerMemberName *string `json:"partnerMemberName,omitempty"`
	// MemberNo - 파트너 회원번호 (멤버)
	MemberNo *string `json:"memberNo,omitempty"`
	// LoginID - 파트너 회원 로그인 아이디 (멤버)
	LoginID *string `json:"loginId,omitempty"`
	// MemberName - 파트너 회원명 (멤버)
	MemberName *string `json:"memberName,omitempty"`
	// DemandYm - 청구연월
	DemandYm *string `json:"demandYm,omitempty"`
	// DemandTypeCreditProductCode - 청구 유형 크레딧 상품 코드
	DemandTypeCreditProductCode *string `json:"demandTypeCreditProductCode,omitempty"`
	// DemandTypeCreditProductCodeName - 청구 유형 크레딧 상품 코드명
	DemandTypeCreditProductCodeName *string `json:"demandTypeCreditProductCodeName,omitempty"`
	// PromiseDiscountAmount - 약정 할인 금액
	PromiseDiscountAmount *string `json:"promiseDiscountAmount,omitempty"`
	// PromotionDiscountAmount - 프로모션 할인 금액
	PromotionDiscountAmount *string `json:"promotionDiscountAmount,omitempty"`
	// EtcDiscountAmount - 기타 할인 금액
	EtcDiscountAmount *string `json:"etcDiscountAmount,omitempty"`
	// ProductDiscountAmount - 상품 할인 금액
	ProductDiscountAmount *string `json:"productDiscountAmount,omitempty"`
	// CreditDiscountAmount - 크레딧 할인 금액
	CreditDiscountAmount *string `json:"creditDiscountAmount,omitempty"`
	// DefaultAmount - 위약 금액
	DefaultAmount *string `json:"defaultAmount,omitempty"`
	// UseAmount - 사용 금액
	UseAmount *string `json:"useAmount,omitempty"`
	// DemandAmount - 청구 금액
	DemandAmount *string `json:"demandAmount,omitempty"`
}

// DemandProductListResponse ...
type DemandProductListResponse struct {
	autorest.Response `json:"-"`
	// RequestID - 요청 ID
	RequestID *string `json:"requestId,omitempty"`
	// ReturnCode - 반환 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 반환 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows         *int32               `json:"totalRows,omitempty"`
	DemandProductList *[]DemandProductList `json:"demandProductList,omitempty"`
}
