package nas

// FINCLOUD_APACHE_NO_VERSION

// EncryptedVolume enumerates the values for encrypted volume.
type EncryptedVolume string

const (
	// False ...
	False EncryptedVolume = "false"
	// True ...
	True EncryptedVolume = "true"
)

// PossibleEncryptedVolumeValues returns an array of possible values for the EncryptedVolume const type.
func PossibleEncryptedVolumeValues() []EncryptedVolume {
	return []EncryptedVolume{False, True}
}

// EventConfiguration enumerates the values for event configuration.
type EventConfiguration string

const (
	// EventConfigurationFalse ...
	EventConfigurationFalse EventConfiguration = "false"
	// EventConfigurationTrue ...
	EventConfigurationTrue EventConfiguration = "true"
)

// PossibleEventConfigurationValues returns an array of possible values for the EventConfiguration const type.
func PossibleEventConfigurationValues() []EventConfiguration {
	return []EventConfiguration{EventConfigurationFalse, EventConfigurationTrue}
}

// SnapshotConfiguration enumerates the values for snapshot configuration.
type SnapshotConfiguration string

const (
	// SnapshotConfigurationFalse ...
	SnapshotConfigurationFalse SnapshotConfiguration = "false"
	// SnapshotConfigurationTrue ...
	SnapshotConfigurationTrue SnapshotConfiguration = "true"
)

// PossibleSnapshotConfigurationValues returns an array of possible values for the SnapshotConfiguration const type.
func PossibleSnapshotConfigurationValues() []SnapshotConfiguration {
	return []SnapshotConfiguration{SnapshotConfigurationFalse, SnapshotConfigurationTrue}
}

// SortedBy enumerates the values for sorted by.
type SortedBy string

const (
	// NasVolumeInstanceNo ...
	NasVolumeInstanceNo SortedBy = "nasVolumeInstanceNo"
	// VolumeName ...
	VolumeName SortedBy = "volumeName"
	// VolumeTotalSize ...
	VolumeTotalSize SortedBy = "volumeTotalSize"
)

// PossibleSortedByValues returns an array of possible values for the SortedBy const type.
func PossibleSortedByValues() []SortedBy {
	return []SortedBy{NasVolumeInstanceNo, VolumeName, VolumeTotalSize}
}

// SortingOrder enumerates the values for sorting order.
type SortingOrder string

const (
	// ASC ...
	ASC SortingOrder = "ASC"
	// DESC ...
	DESC SortingOrder = "DESC"
)

// PossibleSortingOrderValues returns an array of possible values for the SortingOrder const type.
func PossibleSortingOrderValues() []SortingOrder {
	return []SortingOrder{ASC, DESC}
}

// VolumeAllotmentProtocolTypeCode enumerates the values for volume allotment protocol type code.
type VolumeAllotmentProtocolTypeCode string

const (
	// CIFS ...
	CIFS VolumeAllotmentProtocolTypeCode = "CIFS"
	// NFS ...
	NFS VolumeAllotmentProtocolTypeCode = "NFS"
)

// PossibleVolumeAllotmentProtocolTypeCodeValues returns an array of possible values for the VolumeAllotmentProtocolTypeCode const type.
func PossibleVolumeAllotmentProtocolTypeCodeValues() []VolumeAllotmentProtocolTypeCode {
	return []VolumeAllotmentProtocolTypeCode{CIFS, NFS}
}
