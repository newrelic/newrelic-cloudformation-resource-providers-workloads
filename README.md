[sh![New Relic Experimental header](https://github.com/newrelic/opensource-website/raw/master/src/images/categories/Experimental.png)](https://opensource.newrelic.com/oss-category/#new-relic-experimental)

# NewRelic::Observability::Workloads

![GitHub forks](https://img.shields.io/github/forks/newrelic-experimental/newrelic-experimental-FIT-template?style=social)
![GitHub stars](https://img.shields.io/github/stars/newrelic-experimental/newrelic-experimental-FIT-template?style=social)
![GitHub watchers](https://img.shields.io/github/watchers/newrelic-experimental/newrelic-experimental-FIT-template?style=social)

![GitHub all releases](https://img.shields.io/github/downloads/newrelic-experimental/newrelic-experimental-FIT-template/total)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/newrelic-experimental/newrelic-experimental-FIT-template)
![GitHub last commit](https://img.shields.io/github/last-commit/newrelic-experimental/newrelic-experimental-FIT-template)
![GitHub Release Date](https://img.shields.io/github/release-date/newrelic-experimental/newrelic-experimental-FIT-template)


![GitHub issues](https://img.shields.io/github/issues/newrelic-experimental/newrelic-experimental-FIT-template)
![GitHub issues closed](https://img.shields.io/github/issues-closed/newrelic-experimental/newrelic-experimental-FIT-template)
![GitHub pull requests](https://img.shields.io/github/issues-pr/newrelic-experimental/newrelic-experimental-FIT-template)
![GitHub pull requests closed](https://img.shields.io/github/issues-pr-closed/newrelic-experimental/newrelic-experimental-FIT-template)

## Description
This Cloud Formation Custom Resource provides a CRUDL interface to the New Relic [NerdGraph (GraphQL) Workloads API](https://docs.newrelic.com/docs/apis/nerdgraph/examples/nerdgraph-workloads-api-tutorials/) for Cloud Formation stacks.

## Prerequisites
This document assumes familiarity with using CloudFormation Public extensions in CloudFormation templates. If you are not familiar with this [start here](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-public.html)

## [Extension Configuration](https://github.com/newrelic/newrelic-cloudformation-resource-providers-common/blob/main/EXTENSION_CONFIGURATION.md)

## Stack Configuration
| Field           | Type   | Default                | Create | Duplicate | Update | Delete | Read | Notes                                                                                                                                                                |
|-----------------|--------|------------------------|:------:|:---------:|:------:|:------:|:----:|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Guid            | string | none                   |        |     R     |   R    |   R    |      | [See Stack Common Configuration](https://github.com/newrelic/newrelic-cloudformation-resource-providers-common/blob/main/STACK_COMMON_CONFIGURATION.md) |
| ListQueryFilter | string | none                   |        |           |        |        |  R   | [See Stack Common Configuration](https://github.com/newrelic/newrelic-cloudformation-resource-providers-common/blob/main/STACK_COMMON_CONFIGURATION.md) |
| Variables       | Object | none                   |   O    |     O     |        |   O    |  O   | [See Stack Common Configuration](https://github.com/newrelic/newrelic-cloudformation-resource-providers-common/blob/main/STACK_COMMON_CONFIGURATION.md) |
| Tags            | Object | none                   |   O    |     O     |        |        |      | [See Stack Common Configuration](https://github.com/newrelic/newrelic-cloudformation-resource-providers-common/blob/main/STACK_COMMON_CONFIGURATION.md) |                                                                                                                             |
| Workload        | String | none                   |   R    |           |   R    |        |      |                                                                                                                                                                      |


Key:
- R- Requird
- O- Optional
- Blank- unused

### Workload
The entire `workload` fragment from a `workloadCreate` or `workloadUpdate` mutation.

This string is a valid GraphQL fragment representing a Workload, including the `workload: ` keyword. Your best bet is to use the
[GraphQL API Explorer](https://api.newrelic.com/graphiql?#query=mutation%20%7B%0A%20%20workloadCreate%28workload%3A%20%7B%7D%29%0A%7D%0A)
to create this and then copy and paste. Your fragment will be substituted in a create or update mutation like this:
```graphql
mutation {
    workloadCreate(accountId: <AccountID>, workload: {name: "A workload name"}) {
        guid
    }
}
```
_NOTE_: the `{{{` and `}}}` are for [Moustache](#Moustache) processing.

If you use a JSON CloudFormation template you will have to stringify the GraphQL fragment. YAML CloudFormation templates should follow [YAML multi-line input rules](https://yaml-multiline.info/) and avoid stringification.

## Example
An example stack configuration, this is valid and works:
```yaml
AWSTemplateFormatVersion: 2010-09-09
Description: Sample New Relic Workloads Template
Resources:
  Resource1:
    Type: 'NewRelic::Observability::Workloads'
    Properties:
      Workload: >-
        workload: {name: "CloudFormationTest-Create"}
Outputs:
  CustomResourceAttribute1:
    Value: !GetAtt  Resource1.Guid
```

## Troubleshooting
- Error log
- `Debug` log level for mustache substitution
- Validate mutation using the [Explorer](https://api.newrelic.com/graphiql)

## [Development](https://github.com/newrelic/newrelic-cloudformation-resource-providers-common/blob/main/DEVELOPMENT.md)

## Helpful links
- [CloudFormation CLI User Guide](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html)
- [New Relic GraphQL Explorer](https://api.newrelic.com/graphiql)

## Support
New Relic has open-sourced this project. This project is provided AS-IS WITHOUT WARRANTY OR DEDICATED SUPPORT. Issues and contributions should be reported to the project here on GitHub.

We encourage you to bring your experiences and questions to the [Explorers Hub](https://discuss.newrelic.com) where our community members collaborate on solutions and new ideas.

## Contributing
We encourage your contributions to improve New Relic NerdGraph CloudFormation Custom Resource! Keep in mind when you submit your pull request, you'll need to sign the CLA via the click-through using CLA-Assistant. You only have to sign the CLA one time per project. If you have any questions, or to execute our corporate CLA, required if your contribution is on behalf of a company, please drop us an email at opensource@newrelic.com.

**A note about vulnerabilities**

As noted in our [security policy](../../security/policy), New Relic is committed to the privacy and security of our customers and their data. We believe that providing coordinated disclosure by security researchers and engaging with the security community are important means to achieve our security goals.

If you believe you have found a security vulnerability in this project or any of New Relic's products or websites, we welcome and greatly appreciate you reporting it to New Relic through [HackerOne](https://hackerone.com/newrelic).

## License
New Relic NerdGraph CloudFormation Custom Resource is licensed under the [Apache 2.0](http://apache.org/licenses/LICENSE-2.0.txt) License.
