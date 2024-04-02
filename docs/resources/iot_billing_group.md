---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iot_billing_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::IoT::BillingGroup
---

# awscc_iot_billing_group (Resource)

Resource Type definition for AWS::IoT::BillingGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `billing_group_name` (String)
- `billing_group_properties` (Attributes) (see [below for nested schema](#nestedatt--billing_group_properties))
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `billing_group_id` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--billing_group_properties"></a>
### Nested Schema for `billing_group_properties`

Optional:

- `billing_group_description` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_iot_billing_group.example <resource ID>
```