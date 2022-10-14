# NewRelic::CloudFormation::Workloads

CRUD operations for New Relic Workloads via the NerdGraph API

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "NewRelic::CloudFormation::Workloads",
    "Properties" : {
        "<a href="#sourceguid" title="SourceGuid">SourceGuid</a>" : <i>String</i>,
        "<a href="#duplicatename" title="DuplicateName">DuplicateName</a>" : <i>String</i>,
        "<a href="#workload" title="Workload">Workload</a>" : <i>String</i>,
        "<a href="#listqueryfilter" title="ListQueryFilter">ListQueryFilter</a>" : <i>String</i>,
        "<a href="#variables" title="Variables">Variables</a>" : <i><a href="variables.md">Variables</a></i>
    }
}
</pre>

### YAML

<pre>
Type: NewRelic::CloudFormation::Workloads
Properties:
    <a href="#sourceguid" title="SourceGuid">SourceGuid</a>: <i>String</i>
    <a href="#duplicatename" title="DuplicateName">DuplicateName</a>: <i>String</i>
    <a href="#workload" title="Workload">Workload</a>: <i>String</i>
    <a href="#listqueryfilter" title="ListQueryFilter">ListQueryFilter</a>: <i>String</i>
    <a href="#variables" title="Variables">Variables</a>: <i><a href="variables.md">Variables</a></i>
</pre>

## Properties

#### SourceGuid

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### DuplicateName

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Workload

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ListQueryFilter

_Required_: No

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Variables

_Required_: No

_Type_: <a href="variables.md">Variables</a>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the Guid.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### Guid

Returns the <code>Guid</code> value.

