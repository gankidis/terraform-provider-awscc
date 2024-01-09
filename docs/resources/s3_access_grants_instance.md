---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_s3_access_grants_instance Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::S3::AccessGrantsInstance resource is an Amazon S3 resource type that hosts Access Grants and their associated locations
---

# awscc_s3_access_grants_instance (Resource)

The AWS::S3::AccessGrantsInstance resource is an Amazon S3 resource type that hosts Access Grants and their associated locations



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `identity_center_arn` (String) The Amazon Resource Name (ARN) of the specified AWS Identity Center.
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `access_grants_instance_arn` (String) The Amazon Resource Name (ARN) of the specified Access Grants instance.
- `access_grants_instance_id` (String) A unique identifier for the specified access grants instance.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_s3_access_grants_instance.example <resource ID>
```