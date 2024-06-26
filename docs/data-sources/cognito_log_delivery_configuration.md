---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cognito_log_delivery_configuration Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Cognito::LogDeliveryConfiguration
---

# awscc_cognito_log_delivery_configuration (Data Source)

Data Source schema for AWS::Cognito::LogDeliveryConfiguration



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `log_configurations` (Attributes List) (see [below for nested schema](#nestedatt--log_configurations))
- `log_delivery_configuration_id` (String)
- `user_pool_id` (String)

<a id="nestedatt--log_configurations"></a>
### Nested Schema for `log_configurations`

Read-Only:

- `cloudwatch_logs_configuration` (Attributes) (see [below for nested schema](#nestedatt--log_configurations--cloudwatch_logs_configuration))
- `event_source` (String)
- `log_level` (String)

<a id="nestedatt--log_configurations--cloudwatch_logs_configuration"></a>
### Nested Schema for `log_configurations.cloudwatch_logs_configuration`

Read-Only:

- `log_group_arn` (String)
