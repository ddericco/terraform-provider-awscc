---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_rds_db_subnet_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::RDS::DBSubnetGroup resource creates a database subnet group. Subnet groups must contain at least two subnets in two different Availability Zones in the same region.
---

# awscc_rds_db_subnet_group (Resource)

The AWS::RDS::DBSubnetGroup resource creates a database subnet group. Subnet groups must contain at least two subnets in two different Availability Zones in the same region.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `db_subnet_group_description` (String)
- `subnet_ids` (List of String)

### Optional

- `db_subnet_group_name` (String)
- `tags` (Attributes List) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_rds_db_subnet_group.example <resource ID>
```