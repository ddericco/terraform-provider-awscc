---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_imagebuilder_distribution_configuration Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::ImageBuilder::DistributionConfiguration
---

# awscc_imagebuilder_distribution_configuration (Resource)

Resource schema for AWS::ImageBuilder::DistributionConfiguration



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **distributions** (Attributes List) The distributions of the distribution configuration. (see [below for nested schema](#nestedatt--distributions))
- **name** (String) The name of the distribution configuration.

### Optional

- **description** (String) The description of the distribution configuration.
- **tags** (Map of String) The tags associated with the component.

### Read-Only

- **arn** (String) The Amazon Resource Name (ARN) of the distribution configuration.
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--distributions"></a>
### Nested Schema for `distributions`

Required:

- **ami_distribution_configuration** (Attributes) The specific AMI settings (for example, launch permissions, AMI tags). (see [below for nested schema](#nestedatt--distributions--ami_distribution_configuration))
- **container_distribution_configuration** (Attributes) Container distribution settings for encryption, licensing, and sharing in a specific Region. (see [below for nested schema](#nestedatt--distributions--container_distribution_configuration))
- **launch_template_configurations** (Attributes List) A group of launchTemplateConfiguration settings that apply to image distribution. (see [below for nested schema](#nestedatt--distributions--launch_template_configurations))
- **license_configuration_arns** (List of String) The License Manager Configuration to associate with the AMI in the specified Region.
- **region** (String) region

<a id="nestedatt--distributions--ami_distribution_configuration"></a>
### Nested Schema for `distributions.ami_distribution_configuration`

Required:

- **ami_tags** (Map of String) The tags to apply to AMIs distributed to this Region.
- **description** (String) The description of the AMI distribution configuration.
- **kms_key_id** (String) The KMS key identifier used to encrypt the distributed image.
- **launch_permission_configuration** (Attributes) Launch permissions can be used to configure which AWS accounts can use the AMI to launch instances. (see [below for nested schema](#nestedatt--distributions--ami_distribution_configuration--launch_permission_configuration))
- **name** (String) The name of the AMI distribution configuration.
- **target_account_ids** (List of String) The ID of accounts to which you want to distribute an image.

<a id="nestedatt--distributions--ami_distribution_configuration--launch_permission_configuration"></a>
### Nested Schema for `distributions.ami_distribution_configuration.launch_permission_configuration`

Required:

- **user_groups** (List of String) The name of the group.
- **user_ids** (List of String) The AWS account ID.



<a id="nestedatt--distributions--container_distribution_configuration"></a>
### Nested Schema for `distributions.container_distribution_configuration`

Required:

- **container_tags** (List of String) Tags that are attached to the container distribution configuration.
- **description** (String) The description of the container distribution configuration.
- **target_repository** (Attributes) The destination repository for the container image. (see [below for nested schema](#nestedatt--distributions--container_distribution_configuration--target_repository))

<a id="nestedatt--distributions--container_distribution_configuration--target_repository"></a>
### Nested Schema for `distributions.container_distribution_configuration.target_repository`

Required:

- **repository_name** (String) The repository name of target container repository.
- **service** (String) The service of target container repository.



<a id="nestedatt--distributions--launch_template_configurations"></a>
### Nested Schema for `distributions.launch_template_configurations`

Required:

- **account_id** (String) The account ID that this configuration applies to.
- **launch_template_id** (String) Identifies the EC2 launch template to use.
- **set_default_version** (Boolean) Set the specified EC2 launch template as the default launch template for the specified account.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_imagebuilder_distribution_configuration.example <resource ID>
```