package insight

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/insight"

// AggregationType enumerates the values for aggregation type.
type AggregationType string

const (
	// AVG ...
	AVG AggregationType = "AVG"
	// COUNT ...
	COUNT AggregationType = "COUNT"
	// MAX ...
	MAX AggregationType = "MAX"
	// MIN ...
	MIN AggregationType = "MIN"
	// SUM ...
	SUM AggregationType = "SUM"
)

// PossibleAggregationTypeValues returns an array of possible values for the AggregationType const type.
func PossibleAggregationTypeValues() []AggregationType {
	return []AggregationType{AVG, COUNT, MAX, MIN, SUM}
}

// QueryAggregationType enumerates the values for query aggregation type.
type QueryAggregationType string

const (
	// QueryAggregationTypeAVG ...
	QueryAggregationTypeAVG QueryAggregationType = "AVG"
	// QueryAggregationTypeCOUNT ...
	QueryAggregationTypeCOUNT QueryAggregationType = "COUNT"
	// QueryAggregationTypeFIRST ...
	QueryAggregationTypeFIRST QueryAggregationType = "FIRST"
	// QueryAggregationTypeLAST ...
	QueryAggregationTypeLAST QueryAggregationType = "LAST"
	// QueryAggregationTypeMAX ...
	QueryAggregationTypeMAX QueryAggregationType = "MAX"
	// QueryAggregationTypeMIN ...
	QueryAggregationTypeMIN QueryAggregationType = "MIN"
	// QueryAggregationTypeMULT ...
	QueryAggregationTypeMULT QueryAggregationType = "MULT"
	// QueryAggregationTypeNONE ...
	QueryAggregationTypeNONE QueryAggregationType = "NONE"
	// QueryAggregationTypeSUM ...
	QueryAggregationTypeSUM QueryAggregationType = "SUM"
)

// PossibleQueryAggregationTypeValues returns an array of possible values for the QueryAggregationType const type.
func PossibleQueryAggregationTypeValues() []QueryAggregationType {
	return []QueryAggregationType{QueryAggregationTypeAVG, QueryAggregationTypeCOUNT, QueryAggregationTypeFIRST, QueryAggregationTypeLAST, QueryAggregationTypeMAX, QueryAggregationTypeMIN, QueryAggregationTypeMULT, QueryAggregationTypeNONE, QueryAggregationTypeSUM}
}

// QueryIntervalTime enumerates the values for query interval time.
type QueryIntervalTime string

const (
	// DAY1 ...
	DAY1 QueryIntervalTime = "DAY1"
	// HOUR2 ...
	HOUR2 QueryIntervalTime = "HOUR2"
	// MIN1 ...
	MIN1 QueryIntervalTime = "MIN1"
	// MIN30 ...
	MIN30 QueryIntervalTime = "MIN30"
	// MIN5 ...
	MIN5 QueryIntervalTime = "MIN5"
)

// PossibleQueryIntervalTimeValues returns an array of possible values for the QueryIntervalTime const type.
func PossibleQueryIntervalTimeValues() []QueryIntervalTime {
	return []QueryIntervalTime{DAY1, HOUR2, MIN1, MIN30, MIN5}
}

// SchemaDataType enumerates the values for schema data type.
type SchemaDataType string

const (
	// FLOAT ...
	FLOAT SchemaDataType = "FLOAT"
	// INTEGER ...
	INTEGER SchemaDataType = "INTEGER"
	// LONG ...
	LONG SchemaDataType = "LONG"
	// STRING ...
	STRING SchemaDataType = "STRING"
)

// PossibleSchemaDataTypeValues returns an array of possible values for the SchemaDataType const type.
func PossibleSchemaDataTypeValues() []SchemaDataType {
	return []SchemaDataType{FLOAT, INTEGER, LONG, STRING}
}

// CloudInsightCollectorParameter ...
type CloudInsightCollectorParameter struct {
	// Data - 검색할 subnet 페이지 번호
	Data interface{} `json:"data,omitempty"`
	// CwKey - Schema 생성시 발급받은 키
	CwKey *string `json:"cw_key,omitempty"`
}

// CloudInsightQueryParameter ...
type CloudInsightQueryParameter struct {
	// TimeEnd - 마지막 조회 시간
	TimeEnd *float64 `json:"timeEnd,omitempty"`
	// TimeStart - 최초 조회 시간
	TimeStart *float64 `json:"timeStart,omitempty"`
	// Metric - Metric 이름
	Metric *string `json:"metric,omitempty"`
	// Interval - 조회 시간 간격. Possible values include: 'MIN1', 'MIN5', 'MIN30', 'HOUR2', 'DAY1'
	Interval QueryIntervalTime `json:"interval,omitempty"`
	// Aggregation - 일괄 처리. Possible values include: 'MIN', 'MAX', 'SUM', 'COUNT', 'AVG'
	Aggregation AggregationType `json:"aggregation,omitempty"`
	// QueryAggregation - Query 일괄 처리. Possible values include: 'QueryAggregationTypeAVG', 'QueryAggregationTypeCOUNT', 'QueryAggregationTypeMIN', 'QueryAggregationTypeMAX', 'QueryAggregationTypeNONE', 'QueryAggregationTypeSUM', 'QueryAggregationTypeFIRST', 'QueryAggregationTypeLAST', 'QueryAggregationTypeMULT'
	QueryAggregation QueryAggregationType `json:"queryAggregation,omitempty"`
	// CwKey - Schema 생성시 발급받은 키
	CwKey *string `json:"cw_key,omitempty"`
	// ProductName - Product 이름
	ProductName *string `json:"productName,omitempty"`
	// Dimensions - Query Dimension 데이터
	Dimensions interface{} `json:"dimensions,omitempty"`
}

// CloudInsightSchemaFieldsParameter ...
type CloudInsightSchemaFieldsParameter struct {
	// Metric - Metric 유무
	Metric *bool `json:"metric,omitempty"`
	// DataType - Metric 데이터 타입. Possible values include: 'STRING', 'INTEGER', 'LONG', 'FLOAT'
	DataType SchemaDataType `json:"dataType,omitempty"`
	// Name - Metric 이름
	Name *string `json:"name,omitempty"`
	// DefaultMetric - Metric 기본 유무
	DefaultMetric *bool `json:"defaultMetric,omitempty"`
	// Dimension - Dimension 설정
	Dimension *bool `json:"dimension,omitempty"`
	// Desc - Dimension 설명
	Desc *string `json:"desc,omitempty"`
}

// CloudInsightSchemaParameter ...
type CloudInsightSchemaParameter struct {
	// ProdName - Metric Product 이름
	ProdName *string `json:"prodName,omitempty"`
	// Fields - Metric 필드 정의
	Fields *[]CloudInsightSchemaFieldsParameter `json:"fields,omitempty"`
}

// ListListFloat64 ...
type ListListFloat64 struct {
	autorest.Response `json:"-"`
	Value             *[][]float64 `json:"value,omitempty"`
}
