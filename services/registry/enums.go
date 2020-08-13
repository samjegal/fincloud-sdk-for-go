package registry

// FINCLOUD_APACHE_NO_VERSION

// RepositoryStatus enumerates the values for repository status.
type RepositoryStatus string

const (
	// Disconnected ...
	Disconnected RepositoryStatus = "disconnected"
	// Running ...
	Running RepositoryStatus = "running"
)

// PossibleRepositoryStatusValues returns an array of possible values for the RepositoryStatus const type.
func PossibleRepositoryStatusValues() []RepositoryStatus {
	return []RepositoryStatus{Disconnected, Running}
}
