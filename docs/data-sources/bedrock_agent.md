---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_bedrock_agent Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Bedrock::Agent
---

# awscc_bedrock_agent (Data Source)

Data Source schema for AWS::Bedrock::Agent



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `action_groups` (Attributes List) List of ActionGroups (see [below for nested schema](#nestedatt--action_groups))
- `agent_arn` (String) Arn representation of the Agent.
- `agent_id` (String) Identifier for a resource.
- `agent_name` (String) Name for a resource.
- `agent_resource_role_arn` (String) ARN of a IAM role.
- `agent_status` (String) Schema Type for Action APIs.
- `agent_version` (String) Draft Agent Version.
- `auto_prepare` (Boolean) Specifies whether to automatically prepare after creating or updating the agent.
- `created_at` (String) Time Stamp.
- `customer_encryption_key_arn` (String) A KMS key ARN
- `description` (String) Description of the Resource.
- `failure_reasons` (List of String) Failure Reasons for Error.
- `foundation_model` (String) ARN or name of a Bedrock model.
- `idle_session_ttl_in_seconds` (Number) Max Session Time.
- `instruction` (String) Instruction for the agent.
- `knowledge_bases` (Attributes List) List of Agent Knowledge Bases (see [below for nested schema](#nestedatt--knowledge_bases))
- `prepared_at` (String) Time Stamp.
- `prompt_override_configuration` (Attributes) Configuration for prompt override. (see [below for nested schema](#nestedatt--prompt_override_configuration))
- `recommended_actions` (List of String) The recommended actions users can take to resolve an error in failureReasons.
- `skip_resource_in_use_check_on_delete` (Boolean) Specifies whether to allow deleting agent while it is in use.
- `tags` (Map of String) A map of tag keys and values
- `updated_at` (String) Time Stamp.

<a id="nestedatt--action_groups"></a>
### Nested Schema for `action_groups`

Read-Only:

- `action_group_executor` (Attributes) (see [below for nested schema](#nestedatt--action_groups--action_group_executor))
- `action_group_name` (String) Name of the action group
- `action_group_state` (String) State of the action group
- `api_schema` (Attributes) Contains information about the API Schema for the Action Group (see [below for nested schema](#nestedatt--action_groups--api_schema))
- `description` (String) Description of action group
- `parent_action_group_signature` (String) Action Group Signature for a BuiltIn Action
- `skip_resource_in_use_check_on_delete` (Boolean) Specifies whether to allow deleting action group while it is in use.

<a id="nestedatt--action_groups--action_group_executor"></a>
### Nested Schema for `action_groups.action_group_executor`

Read-Only:

- `lambda` (String) ARN of a Lambda.


<a id="nestedatt--action_groups--api_schema"></a>
### Nested Schema for `action_groups.api_schema`

Read-Only:

- `payload` (String) String OpenAPI Payload
- `s3` (Attributes) The identifier for the S3 resource. (see [below for nested schema](#nestedatt--action_groups--api_schema--s3))

<a id="nestedatt--action_groups--api_schema--s3"></a>
### Nested Schema for `action_groups.api_schema.s3`

Read-Only:

- `s3_bucket_name` (String) A bucket in S3.
- `s3_object_key` (String) A object key in S3.




<a id="nestedatt--knowledge_bases"></a>
### Nested Schema for `knowledge_bases`

Read-Only:

- `description` (String) Description of the Resource.
- `knowledge_base_id` (String) Identifier for a resource.
- `knowledge_base_state` (String) State of the knowledge base; whether it is enabled or disabled


<a id="nestedatt--prompt_override_configuration"></a>
### Nested Schema for `prompt_override_configuration`

Read-Only:

- `override_lambda` (String) ARN of a Lambda.
- `prompt_configurations` (Attributes List) List of BasePromptConfiguration (see [below for nested schema](#nestedatt--prompt_override_configuration--prompt_configurations))

<a id="nestedatt--prompt_override_configuration--prompt_configurations"></a>
### Nested Schema for `prompt_override_configuration.prompt_configurations`

Read-Only:

- `base_prompt_template` (String) Base Prompt Template.
- `inference_configuration` (Attributes) Configuration for inference in prompt configuration (see [below for nested schema](#nestedatt--prompt_override_configuration--prompt_configurations--inference_configuration))
- `parser_mode` (String) Creation Mode for Prompt Configuration.
- `prompt_creation_mode` (String) Creation Mode for Prompt Configuration.
- `prompt_state` (String) Prompt State.
- `prompt_type` (String) Prompt Type.

<a id="nestedatt--prompt_override_configuration--prompt_configurations--inference_configuration"></a>
### Nested Schema for `prompt_override_configuration.prompt_configurations.inference_configuration`

Read-Only:

- `maximum_length` (Number) Maximum length of output
- `stop_sequences` (List of String) List of stop sequences
- `temperature` (Number) Controls randomness, higher values increase diversity
- `top_k` (Number) Sample from the k most likely next tokens
- `top_p` (Number) Cumulative probability cutoff for token selection