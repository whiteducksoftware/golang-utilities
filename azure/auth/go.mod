module github.com/whiteducksoftware/golang-utilities/azure/auth

go 1.17

require (
	github.com/Azure/go-autorest/autorest v0.11.20
	github.com/Azure/go-autorest/autorest/adal v0.9.14
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.8
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.3
)

require (
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/form3tech-oss/jwt-go v3.2.2+incompatible // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
)

// Temporary fix until https://github.com/Azure/go-autorest/pull/653 is merged.
replace github.com/Azure/go-autorest/autorest/azure/cli v0.4.3 => ../../libs/@azure/go-autorest/azure/cli
