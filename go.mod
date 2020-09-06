module github.com/samjegal/fincloud-sdk-for-go

go 1.15

replace github.com/Azure/go-autorest/autorest => github.com/samjegal/go-autorest/autorest v0.11.5-0.20200906111652-5e286818fa7f

require (
	github.com/Azure/go-autorest/autorest v0.11.4
	github.com/Azure/go-autorest/autorest/date v0.3.0
	github.com/Azure/go-autorest/autorest/validation v0.3.0
	github.com/Azure/go-autorest/tracing v0.6.0
	github.com/Masterminds/semver v1.5.0
	github.com/samjegal/go-fincloud-helpers v0.2.4
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/tools v0.0.0-20200904185747-39188db58858
)
