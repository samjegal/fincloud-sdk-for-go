package cloudinsight

// FINCLOUD_APACHE_NO_VERSION

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

// ChartDataWidgetInfoOption enumerates the values for chart data widget info option.
type ChartDataWidgetInfoOption string

const (
	// ChartDataWidgetInfoOptionAVG ...
	ChartDataWidgetInfoOptionAVG ChartDataWidgetInfoOption = "AVG"
	// ChartDataWidgetInfoOptionCOUNT ...
	ChartDataWidgetInfoOptionCOUNT ChartDataWidgetInfoOption = "COUNT"
	// ChartDataWidgetInfoOptionCOUNTER ...
	ChartDataWidgetInfoOptionCOUNTER ChartDataWidgetInfoOption = "COUNTER"
	// ChartDataWidgetInfoOptionMAX ...
	ChartDataWidgetInfoOptionMAX ChartDataWidgetInfoOption = "MAX"
	// ChartDataWidgetInfoOptionMIN ...
	ChartDataWidgetInfoOptionMIN ChartDataWidgetInfoOption = "MIN"
	// ChartDataWidgetInfoOptionSUM ...
	ChartDataWidgetInfoOptionSUM ChartDataWidgetInfoOption = "SUM"
)

// PossibleChartDataWidgetInfoOptionValues returns an array of possible values for the ChartDataWidgetInfoOption const type.
func PossibleChartDataWidgetInfoOptionValues() []ChartDataWidgetInfoOption {
	return []ChartDataWidgetInfoOption{ChartDataWidgetInfoOptionAVG, ChartDataWidgetInfoOptionCOUNT, ChartDataWidgetInfoOptionCOUNTER, ChartDataWidgetInfoOptionMAX, ChartDataWidgetInfoOptionMIN, ChartDataWidgetInfoOptionSUM}
}

// ChartDataWidgetInfoPeriod enumerates the values for chart data widget info period.
type ChartDataWidgetInfoPeriod string

const (
	// Day1 ...
	Day1 ChartDataWidgetInfoPeriod = "Day1"
	// Hour2 ...
	Hour2 ChartDataWidgetInfoPeriod = "Hour2"
	// Min1 ...
	Min1 ChartDataWidgetInfoPeriod = "Min1"
	// Min30 ...
	Min30 ChartDataWidgetInfoPeriod = "Min30"
	// Min5 ...
	Min5 ChartDataWidgetInfoPeriod = "Min5"
)

// PossibleChartDataWidgetInfoPeriodValues returns an array of possible values for the ChartDataWidgetInfoPeriod const type.
func PossibleChartDataWidgetInfoPeriodValues() []ChartDataWidgetInfoPeriod {
	return []ChartDataWidgetInfoPeriod{Day1, Hour2, Min1, Min30, Min5}
}

// ChartDataWidgetInfoStatistic enumerates the values for chart data widget info statistic.
type ChartDataWidgetInfoStatistic string

const (
	// ChartDataWidgetInfoStatisticAVG ...
	ChartDataWidgetInfoStatisticAVG ChartDataWidgetInfoStatistic = "AVG"
	// ChartDataWidgetInfoStatisticCOUNT ...
	ChartDataWidgetInfoStatisticCOUNT ChartDataWidgetInfoStatistic = "COUNT"
	// ChartDataWidgetInfoStatisticCOUNTER ...
	ChartDataWidgetInfoStatisticCOUNTER ChartDataWidgetInfoStatistic = "COUNTER"
	// ChartDataWidgetInfoStatisticMAX ...
	ChartDataWidgetInfoStatisticMAX ChartDataWidgetInfoStatistic = "MAX"
	// ChartDataWidgetInfoStatisticMIN ...
	ChartDataWidgetInfoStatisticMIN ChartDataWidgetInfoStatistic = "MIN"
	// ChartDataWidgetInfoStatisticSUM ...
	ChartDataWidgetInfoStatisticSUM ChartDataWidgetInfoStatistic = "SUM"
)

// PossibleChartDataWidgetInfoStatisticValues returns an array of possible values for the ChartDataWidgetInfoStatistic const type.
func PossibleChartDataWidgetInfoStatisticValues() []ChartDataWidgetInfoStatistic {
	return []ChartDataWidgetInfoStatistic{ChartDataWidgetInfoStatisticAVG, ChartDataWidgetInfoStatisticCOUNT, ChartDataWidgetInfoStatisticCOUNTER, ChartDataWidgetInfoStatisticMAX, ChartDataWidgetInfoStatisticMIN, ChartDataWidgetInfoStatisticSUM}
}

// ChartDataWidgetType enumerates the values for chart data widget type.
type ChartDataWidgetType string

const (
	// Area ...
	Area ChartDataWidgetType = "area"
	// Index ...
	Index ChartDataWidgetType = "index"
	// Line ...
	Line ChartDataWidgetType = "line"
	// Table ...
	Table ChartDataWidgetType = "table"
	// Text ...
	Text ChartDataWidgetType = "text"
)

// PossibleChartDataWidgetTypeValues returns an array of possible values for the ChartDataWidgetType const type.
func PossibleChartDataWidgetTypeValues() []ChartDataWidgetType {
	return []ChartDataWidgetType{Area, Index, Line, Table, Text}
}

// DimensionType enumerates the values for dimension type.
type DimensionType string

const (
	// Agent ...
	Agent DimensionType = "agent"
	// CPU ...
	CPU DimensionType = "cpu"
	// Disk ...
	Disk DimensionType = "disk"
	// Fs ...
	Fs DimensionType = "fs"
	// Memory ...
	Memory DimensionType = "memory"
	// Ntwk ...
	Ntwk DimensionType = "ntwk"
	// PluginFile ...
	PluginFile DimensionType = "plugin_file"
	// PluginPort ...
	PluginPort DimensionType = "plugin_port"
	// PluginProcess ...
	PluginProcess DimensionType = "plugin_process"
	// Process ...
	Process DimensionType = "process"
	// Svr ...
	Svr DimensionType = "svr"
)

// PossibleDimensionTypeValues returns an array of possible values for the DimensionType const type.
func PossibleDimensionTypeValues() []DimensionType {
	return []DimensionType{Agent, CPU, Disk, Fs, Memory, Ntwk, PluginFile, PluginPort, PluginProcess, Process, Svr}
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

// RecipientNotificationNotifyType enumerates the values for recipient notification notify type.
type RecipientNotificationNotifyType string

const (
	// EMAIL ...
	EMAIL RecipientNotificationNotifyType = "EMAIL"
	// SMS ...
	SMS RecipientNotificationNotifyType = "SMS"
)

// PossibleRecipientNotificationNotifyTypeValues returns an array of possible values for the RecipientNotificationNotifyType const type.
func PossibleRecipientNotificationNotifyTypeValues() []RecipientNotificationNotifyType {
	return []RecipientNotificationNotifyType{EMAIL, SMS}
}

// RuleGroupProductDataStatus enumerates the values for rule group product data status.
type RuleGroupProductDataStatus string

const (
	// CREATED ...
	CREATED RuleGroupProductDataStatus = "CREATED"
	// DELETED ...
	DELETED RuleGroupProductDataStatus = "DELETED"
	// UPDATED ...
	UPDATED RuleGroupProductDataStatus = "UPDATED"
)

// PossibleRuleGroupProductDataStatusValues returns an array of possible values for the RuleGroupProductDataStatus const type.
func PossibleRuleGroupProductDataStatusValues() []RuleGroupProductDataStatus {
	return []RuleGroupProductDataStatus{CREATED, DELETED, UPDATED}
}

// RuleGroupStatus enumerates the values for rule group status.
type RuleGroupStatus string

const (
	// INSUFFICIENT ...
	INSUFFICIENT RuleGroupStatus = "INSUFFICIENT"
	// OK ...
	OK RuleGroupStatus = "OK"
	// VIOLATED ...
	VIOLATED RuleGroupStatus = "VIOLATED"
)

// PossibleRuleGroupStatusValues returns an array of possible values for the RuleGroupStatus const type.
func PossibleRuleGroupStatusValues() []RuleGroupStatus {
	return []RuleGroupStatus{INSUFFICIENT, OK, VIOLATED}
}

// RuleGroupVersion enumerates the values for rule group version.
type RuleGroupVersion string

const (
	// V1 ...
	V1 RuleGroupVersion = "V1"
	// V2 ...
	V2 RuleGroupVersion = "V2"
)

// PossibleRuleGroupVersionValues returns an array of possible values for the RuleGroupVersion const type.
func PossibleRuleGroupVersionValues() []RuleGroupVersion {
	return []RuleGroupVersion{V1, V2}
}

// SchemaFieldaAggregation enumerates the values for schema fielda aggregation.
type SchemaFieldaAggregation string

const (
	// SchemaFieldaAggregationAVG ...
	SchemaFieldaAggregationAVG SchemaFieldaAggregation = "AVG"
	// SchemaFieldaAggregationCOUNT ...
	SchemaFieldaAggregationCOUNT SchemaFieldaAggregation = "COUNT"
	// SchemaFieldaAggregationCOUNTER ...
	SchemaFieldaAggregationCOUNTER SchemaFieldaAggregation = "COUNTER"
	// SchemaFieldaAggregationMAX ...
	SchemaFieldaAggregationMAX SchemaFieldaAggregation = "MAX"
	// SchemaFieldaAggregationMIN ...
	SchemaFieldaAggregationMIN SchemaFieldaAggregation = "MIN"
	// SchemaFieldaAggregationSUM ...
	SchemaFieldaAggregationSUM SchemaFieldaAggregation = "SUM"
)

// PossibleSchemaFieldaAggregationValues returns an array of possible values for the SchemaFieldaAggregation const type.
func PossibleSchemaFieldaAggregationValues() []SchemaFieldaAggregation {
	return []SchemaFieldaAggregation{SchemaFieldaAggregationAVG, SchemaFieldaAggregationCOUNT, SchemaFieldaAggregationCOUNTER, SchemaFieldaAggregationMAX, SchemaFieldaAggregationMIN, SchemaFieldaAggregationSUM}
}

// SchemaFieldDataType enumerates the values for schema field data type.
type SchemaFieldDataType string

const (
	// ARRAY ...
	ARRAY SchemaFieldDataType = "ARRAY"
	// FLOAT ...
	FLOAT SchemaFieldDataType = "FLOAT"
	// INTEGER ...
	INTEGER SchemaFieldDataType = "INTEGER"
	// LONG ...
	LONG SchemaFieldDataType = "LONG"
	// STRING ...
	STRING SchemaFieldDataType = "STRING"
	// TIMESTAMP ...
	TIMESTAMP SchemaFieldDataType = "TIMESTAMP"
)

// PossibleSchemaFieldDataTypeValues returns an array of possible values for the SchemaFieldDataType const type.
func PossibleSchemaFieldDataTypeValues() []SchemaFieldDataType {
	return []SchemaFieldDataType{ARRAY, FLOAT, INTEGER, LONG, STRING, TIMESTAMP}
}

// SchemaFieldDimensionType enumerates the values for schema field dimension type.
type SchemaFieldDimensionType string

const (
	// CUSTOM ...
	CUSTOM SchemaFieldDimensionType = "CUSTOM"
	// DEFAULT ...
	DEFAULT SchemaFieldDimensionType = "DEFAULT"
)

// PossibleSchemaFieldDimensionTypeValues returns an array of possible values for the SchemaFieldDimensionType const type.
func PossibleSchemaFieldDimensionTypeValues() []SchemaFieldDimensionType {
	return []SchemaFieldDimensionType{CUSTOM, DEFAULT}
}

// SchemaFieldMetricType enumerates the values for schema field metric type.
type SchemaFieldMetricType string

const (
	// SchemaFieldMetricTypeBASIC ...
	SchemaFieldMetricTypeBASIC SchemaFieldMetricType = "BASIC"
	// SchemaFieldMetricTypeCUSTOM ...
	SchemaFieldMetricTypeCUSTOM SchemaFieldMetricType = "CUSTOM"
	// SchemaFieldMetricTypeEXTENDED ...
	SchemaFieldMetricTypeEXTENDED SchemaFieldMetricType = "EXTENDED"
)

// PossibleSchemaFieldMetricTypeValues returns an array of possible values for the SchemaFieldMetricType const type.
func PossibleSchemaFieldMetricTypeValues() []SchemaFieldMetricType {
	return []SchemaFieldMetricType{SchemaFieldMetricTypeBASIC, SchemaFieldMetricTypeCUSTOM, SchemaFieldMetricTypeEXTENDED}
}

// SeverTargetMetric enumerates the values for sever target metric.
type SeverTargetMetric string

const (
	// AvgCPUUsedRto ...
	AvgCPUUsedRto SeverTargetMetric = "avg_cpu_used_rto"
	// AvgFsUsert ...
	AvgFsUsert SeverTargetMetric = "avg_fs_usert"
	// MemUsert ...
	MemUsert SeverTargetMetric = "mem_usert"
)

// PossibleSeverTargetMetricValues returns an array of possible values for the SeverTargetMetric const type.
func PossibleSeverTargetMetricValues() []SeverTargetMetric {
	return []SeverTargetMetric{AvgCPUUsedRto, AvgFsUsert, MemUsert}
}

// WidgetMetricInfoPeriod enumerates the values for widget metric info period.
type WidgetMetricInfoPeriod string

const (
	// WidgetMetricInfoPeriodDay1 ...
	WidgetMetricInfoPeriodDay1 WidgetMetricInfoPeriod = "Day1"
	// WidgetMetricInfoPeriodHour2 ...
	WidgetMetricInfoPeriodHour2 WidgetMetricInfoPeriod = "Hour2"
	// WidgetMetricInfoPeriodMin1 ...
	WidgetMetricInfoPeriodMin1 WidgetMetricInfoPeriod = "Min1"
	// WidgetMetricInfoPeriodMin30 ...
	WidgetMetricInfoPeriodMin30 WidgetMetricInfoPeriod = "Min30"
	// WidgetMetricInfoPeriodMin5 ...
	WidgetMetricInfoPeriodMin5 WidgetMetricInfoPeriod = "Min5"
)

// PossibleWidgetMetricInfoPeriodValues returns an array of possible values for the WidgetMetricInfoPeriod const type.
func PossibleWidgetMetricInfoPeriodValues() []WidgetMetricInfoPeriod {
	return []WidgetMetricInfoPeriod{WidgetMetricInfoPeriodDay1, WidgetMetricInfoPeriodHour2, WidgetMetricInfoPeriodMin1, WidgetMetricInfoPeriodMin30, WidgetMetricInfoPeriodMin5}
}

// WidgetMetricInfoQueryAggregation enumerates the values for widget metric info query aggregation.
type WidgetMetricInfoQueryAggregation string

const (
	// WidgetMetricInfoQueryAggregationAVG ...
	WidgetMetricInfoQueryAggregationAVG WidgetMetricInfoQueryAggregation = "AVG"
	// WidgetMetricInfoQueryAggregationCOUNT ...
	WidgetMetricInfoQueryAggregationCOUNT WidgetMetricInfoQueryAggregation = "COUNT"
	// WidgetMetricInfoQueryAggregationFIRST ...
	WidgetMetricInfoQueryAggregationFIRST WidgetMetricInfoQueryAggregation = "FIRST"
	// WidgetMetricInfoQueryAggregationLAST ...
	WidgetMetricInfoQueryAggregationLAST WidgetMetricInfoQueryAggregation = "LAST"
	// WidgetMetricInfoQueryAggregationMAX ...
	WidgetMetricInfoQueryAggregationMAX WidgetMetricInfoQueryAggregation = "MAX"
	// WidgetMetricInfoQueryAggregationMIMMAX ...
	WidgetMetricInfoQueryAggregationMIMMAX WidgetMetricInfoQueryAggregation = "MIMMAX"
	// WidgetMetricInfoQueryAggregationMIMMIN ...
	WidgetMetricInfoQueryAggregationMIMMIN WidgetMetricInfoQueryAggregation = "MIMMIN"
	// WidgetMetricInfoQueryAggregationMIN ...
	WidgetMetricInfoQueryAggregationMIN WidgetMetricInfoQueryAggregation = "MIN"
	// WidgetMetricInfoQueryAggregationMULT ...
	WidgetMetricInfoQueryAggregationMULT WidgetMetricInfoQueryAggregation = "MULT"
	// WidgetMetricInfoQueryAggregationNONE ...
	WidgetMetricInfoQueryAggregationNONE WidgetMetricInfoQueryAggregation = "NONE"
	// WidgetMetricInfoQueryAggregationSUM ...
	WidgetMetricInfoQueryAggregationSUM WidgetMetricInfoQueryAggregation = "SUM"
	// WidgetMetricInfoQueryAggregationZIMSUM ...
	WidgetMetricInfoQueryAggregationZIMSUM WidgetMetricInfoQueryAggregation = "ZIMSUM"
)

// PossibleWidgetMetricInfoQueryAggregationValues returns an array of possible values for the WidgetMetricInfoQueryAggregation const type.
func PossibleWidgetMetricInfoQueryAggregationValues() []WidgetMetricInfoQueryAggregation {
	return []WidgetMetricInfoQueryAggregation{WidgetMetricInfoQueryAggregationAVG, WidgetMetricInfoQueryAggregationCOUNT, WidgetMetricInfoQueryAggregationFIRST, WidgetMetricInfoQueryAggregationLAST, WidgetMetricInfoQueryAggregationMAX, WidgetMetricInfoQueryAggregationMIMMAX, WidgetMetricInfoQueryAggregationMIMMIN, WidgetMetricInfoQueryAggregationMIN, WidgetMetricInfoQueryAggregationMULT, WidgetMetricInfoQueryAggregationNONE, WidgetMetricInfoQueryAggregationSUM, WidgetMetricInfoQueryAggregationZIMSUM}
}

// WidgetMetricInfoStatistic enumerates the values for widget metric info statistic.
type WidgetMetricInfoStatistic string

const (
	// WidgetMetricInfoStatisticAVG ...
	WidgetMetricInfoStatisticAVG WidgetMetricInfoStatistic = "AVG"
	// WidgetMetricInfoStatisticCOUNT ...
	WidgetMetricInfoStatisticCOUNT WidgetMetricInfoStatistic = "COUNT"
	// WidgetMetricInfoStatisticCOUNTER ...
	WidgetMetricInfoStatisticCOUNTER WidgetMetricInfoStatistic = "COUNTER"
	// WidgetMetricInfoStatisticMAX ...
	WidgetMetricInfoStatisticMAX WidgetMetricInfoStatistic = "MAX"
	// WidgetMetricInfoStatisticMIN ...
	WidgetMetricInfoStatisticMIN WidgetMetricInfoStatistic = "MIN"
	// WidgetMetricInfoStatisticSUM ...
	WidgetMetricInfoStatisticSUM WidgetMetricInfoStatistic = "SUM"
)

// PossibleWidgetMetricInfoStatisticValues returns an array of possible values for the WidgetMetricInfoStatistic const type.
func PossibleWidgetMetricInfoStatisticValues() []WidgetMetricInfoStatistic {
	return []WidgetMetricInfoStatistic{WidgetMetricInfoStatisticAVG, WidgetMetricInfoStatisticCOUNT, WidgetMetricInfoStatisticCOUNTER, WidgetMetricInfoStatisticMAX, WidgetMetricInfoStatisticMIN, WidgetMetricInfoStatisticSUM}
}
