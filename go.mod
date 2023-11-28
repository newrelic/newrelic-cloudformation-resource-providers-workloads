module newrelic-cloudformation-workloads

go 1.19

require (
	github.com/aws-cloudformation/cloudformation-cli-go-plugin v1.2.0
	github.com/newrelic/newrelic-cloudformation-resource-providers-common v1.1.5
	github.com/sirupsen/logrus v1.9.3
)

// replace github.com/newrelic/newrelic-cloudformation-resource-providers-common => ../newrelic-cloudformation-resource-providers-common

require (
	github.com/aws/aws-lambda-go v1.41.0 // indirect
	github.com/aws/aws-sdk-go v1.48.6 // indirect
	github.com/cbroglie/mustache v1.4.0 // indirect
	github.com/go-resty/resty/v2 v2.10.0 // indirect
	github.com/graphql-go/graphql v0.8.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	gopkg.in/validator.v2 v2.0.1 // indirect
)
