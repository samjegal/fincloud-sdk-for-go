package wms

// FINCLOUD_APACHE_NO_VERSION

// AlarmStatus enumerates the values for alarm status.
type AlarmStatus string

const (
	// OFF ...
	OFF AlarmStatus = "OFF"
	// ON ...
	ON AlarmStatus = "ON"
)

// PossibleAlarmStatusValues returns an array of possible values for the AlarmStatus const type.
func PossibleAlarmStatusValues() []AlarmStatus {
	return []AlarmStatus{OFF, ON}
}

// ErrorType enumerates the values for error type.
type ErrorType string

const (
	// JAVASCRIPT ...
	JAVASCRIPT ErrorType = "JAVASCRIPT"
	// URL ...
	URL ErrorType = "URL"
)

// PossibleErrorTypeValues returns an array of possible values for the ErrorType const type.
func PossibleErrorTypeValues() []ErrorType {
	return []ErrorType{JAVASCRIPT, URL}
}

// FilterType enumerates the values for filter type.
type FilterType string

const (
	// FilterTypeJs ...
	FilterTypeJs FilterType = "js"
	// FilterTypeJsprefix ...
	FilterTypeJsprefix FilterType = "jsprefix"
	// FilterTypeURL ...
	FilterTypeURL FilterType = "url"
	// FilterTypeUrlprefix ...
	FilterTypeUrlprefix FilterType = "urlprefix"
)

// PossibleFilterTypeValues returns an array of possible values for the FilterType const type.
func PossibleFilterTypeValues() []FilterType {
	return []FilterType{FilterTypeJs, FilterTypeJsprefix, FilterTypeURL, FilterTypeUrlprefix}
}

// LocationName enumerates the values for location name.
type LocationName string

const (
	// HK ...
	HK LocationName = "HK"
	// JP ...
	JP LocationName = "JP"
	// KR ...
	KR LocationName = "KR"
	// SG ...
	SG LocationName = "SG"
	// US ...
	US LocationName = "US"
)

// PossibleLocationNameValues returns an array of possible values for the LocationName const type.
func PossibleLocationNameValues() []LocationName {
	return []LocationName{HK, JP, KR, SG, US}
}

// MonitoringType enumerates the values for monitoring type.
type MonitoringType string

const (
	// MonitoringTypeSCENARIO ...
	MonitoringTypeSCENARIO MonitoringType = "SCENARIO"
	// MonitoringTypeURL ...
	MonitoringTypeURL MonitoringType = "URL"
)

// PossibleMonitoringTypeValues returns an array of possible values for the MonitoringType const type.
func PossibleMonitoringTypeValues() []MonitoringType {
	return []MonitoringType{MonitoringTypeSCENARIO, MonitoringTypeURL}
}

// ResultEnum enumerates the values for result enum.
type ResultEnum string

const (
	// FAIL ...
	FAIL ResultEnum = "FAIL"
	// SUCCESS ...
	SUCCESS ResultEnum = "SUCCESS"
)

// PossibleResultEnumValues returns an array of possible values for the ResultEnum const type.
func PossibleResultEnumValues() []ResultEnum {
	return []ResultEnum{FAIL, SUCCESS}
}

// ServiceType enumerates the values for service type.
type ServiceType string

const (
	// MOBILE ...
	MOBILE ServiceType = "MOBILE"
	// PC ...
	PC ServiceType = "PC"
)

// PossibleServiceTypeValues returns an array of possible values for the ServiceType const type.
func PossibleServiceTypeValues() []ServiceType {
	return []ServiceType{MOBILE, PC}
}

// StatusType enumerates the values for status type.
type StatusType string

const (
	// StatusTypeOFF ...
	StatusTypeOFF StatusType = "OFF"
	// StatusTypeON ...
	StatusTypeON StatusType = "ON"
)

// PossibleStatusTypeValues returns an array of possible values for the StatusType const type.
func PossibleStatusTypeValues() []StatusType {
	return []StatusType{StatusTypeOFF, StatusTypeON}
}

// StepEvent enumerates the values for step event.
type StepEvent string

const (
	// Click ...
	Click StepEvent = "click"
	// FindObject ...
	FindObject StepEvent = "find_object"
	// FindText ...
	FindText StepEvent = "find_text"
	// Text ...
	Text StepEvent = "text"
)

// PossibleStepEventValues returns an array of possible values for the StepEvent const type.
func PossibleStepEventValues() []StepEvent {
	return []StepEvent{Click, FindObject, FindText, Text}
}

// StepResult enumerates the values for step result.
type StepResult string

const (
	// StepResultFAIL ...
	StepResultFAIL StepResult = "FAIL"
	// StepResultSUCCESS ...
	StepResultSUCCESS StepResult = "SUCCESS"
)

// PossibleStepResultValues returns an array of possible values for the StepResult const type.
func PossibleStepResultValues() []StepResult {
	return []StepResult{StepResultFAIL, StepResultSUCCESS}
}
