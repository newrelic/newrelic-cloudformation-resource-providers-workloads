module newrelic-cloudformation-workloads

go 1.19

require (
	github.com/aws-cloudformation/cloudformation-cli-go-plugin v1.0.3
	github.com/newrelic-experimental/newrelic-cloudformation-resource-providers-common v0.3.1
	github.com/sirupsen/logrus v1.9.0
)

replace github.com/newrelic-experimental/newrelic-cloudformation-resource-providers-common => ../newrelic-cloudformation-resource-providers-common

require (
	github.com/aws/aws-lambda-go v1.13.3 // indirect
	github.com/aws/aws-sdk-go v1.44.111 // indirect
	github.com/cbroglie/mustache v1.4.0 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/graphql-go/graphql v0.8.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/segmentio/ksuid v1.0.2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	gopkg.in/validator.v2 v2.0.0-20191107172027-c3144fdedc21 // indirect
)
