---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_gamelift_fleet Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::GameLift::Fleet resource creates an Amazon GameLift (GameLift) fleet to host game servers.  A fleet is a set of EC2 instances, each of which can host multiple game sessions.
---

# awscc_gamelift_fleet (Resource)

The AWS::GameLift::Fleet resource creates an Amazon GameLift (GameLift) fleet to host game servers.  A fleet is a set of EC2 instances, each of which can host multiple game sessions.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **build_id** (String) A unique identifier for a build to be deployed on the new fleet. If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a READY status. This fleet setting cannot be changed once the fleet is created.
- **certificate_configuration** (Attributes) Information about the use of a TLS/SSL certificate for a fleet. TLS certificate generation is enabled at the fleet level, with one certificate generated for the fleet. When this feature is enabled, the certificate can be retrieved using the GameLift Server SDK call GetInstanceCertificate. All instances in a fleet share the same certificate. (see [below for nested schema](#nestedatt--certificate_configuration))
- **description** (String) A human-readable description of a fleet.
- **desired_ec2_instances** (Number) [DEPRECATED] The number of EC2 instances that you want this fleet to host. When creating a new fleet, GameLift automatically sets this value to "1" and initiates a single instance. Once the fleet is active, update this value to trigger GameLift to add or remove instances from the fleet.
- **ec2_inbound_permissions** (Attributes List) A range of IP addresses and port settings that allow inbound traffic to connect to server processes on an Amazon GameLift server. (see [below for nested schema](#nestedatt--ec2_inbound_permissions))
- **ec2_instance_type** (String) The name of an EC2 instance type that is supported in Amazon GameLift. A fleet instance type determines the computing resources of each instance in the fleet, including CPU, memory, storage, and networking capacity. Amazon GameLift supports the following EC2 instance types. See Amazon EC2 Instance Types for detailed descriptions.
- **fleet_type** (String) Indicates whether to use On-Demand instances or Spot instances for this fleet. If empty, the default is ON_DEMAND. Both categories of instances use identical hardware and configurations based on the instance type selected for this fleet.
- **instance_role_arn** (String) A unique identifier for an AWS IAM role that manages access to your AWS services. With an instance role ARN set, any application that runs on an instance in this fleet can assume the role, including install scripts, server processes, and daemons (background processes). Create a role or look up a role's ARN from the IAM dashboard in the AWS Management Console.
- **locations** (Attributes List) (see [below for nested schema](#nestedatt--locations))
- **log_paths** (List of String) This parameter is no longer used. When hosting a custom game build, specify where Amazon GameLift should store log files using the Amazon GameLift server API call ProcessReady()
- **max_size** (Number) [DEPRECATED] The maximum value that is allowed for the fleet's instance count. When creating a new fleet, GameLift automatically sets this value to "1". Once the fleet is active, you can change this value.
- **metric_groups** (List of String) The name of an Amazon CloudWatch metric group. A metric group aggregates the metrics for all fleets in the group. Specify a string containing the metric group name. You can use an existing name or use a new name to create a new metric group. Currently, this parameter can have only one string.
- **min_size** (Number) [DEPRECATED] The minimum value allowed for the fleet's instance count. When creating a new fleet, GameLift automatically sets this value to "0". After the fleet is active, you can change this value.
- **name** (String) A descriptive label that is associated with a fleet. Fleet names do not need to be unique.
- **new_game_session_protection_policy** (String) A game session protection policy to apply to all game sessions hosted on instances in this fleet. When protected, active game sessions cannot be terminated during a scale-down event. If this parameter is not set, instances in this fleet default to no protection. You can change a fleet's protection policy to affect future game sessions on the fleet. You can also set protection for individual game sessions.
- **peer_vpc_aws_account_id** (String) A unique identifier for the AWS account with the VPC that you want to peer your Amazon GameLift fleet with. You can find your account ID in the AWS Management Console under account settings.
- **peer_vpc_id** (String) A unique identifier for a VPC with resources to be accessed by your Amazon GameLift fleet. The VPC must be in the same Region as your fleet. To look up a VPC ID, use the VPC Dashboard in the AWS Management Console.
- **resource_creation_limit_policy** (Attributes) A policy that limits the number of game sessions a player can create on the same fleet. This optional policy gives game owners control over how players can consume available game server resources. A resource creation policy makes the following statement: "An individual player can create a maximum number of new game sessions within a specified time period".

The policy is evaluated when a player tries to create a new game session. For example, assume you have a policy of 10 new game sessions and a time period of 60 minutes. On receiving a CreateGameSession request, Amazon GameLift checks that the player (identified by CreatorId) has created fewer than 10 game sessions in the past 60 minutes. (see [below for nested schema](#nestedatt--resource_creation_limit_policy))
- **runtime_configuration** (Attributes) A collection of server process configurations that describe the processes to run on each instance in a fleet. All fleets must have a runtime configuration. Each instance in the fleet maintains server processes as specified in the runtime configuration, launching new ones as existing processes end. Each instance regularly checks for an updated runtime configuration makes adjustments as called for.

The runtime configuration enables the instances in a fleet to run multiple processes simultaneously. Potential scenarios are as follows: (1) Run multiple processes of a single game server executable to maximize usage of your hosting resources. (2) Run one or more processes of different executables, such as your game server and a metrics tracking program. (3) Run multiple processes of a single game server but with different launch parameters, for example to run one process on each instance in debug mode.

An Amazon GameLift instance is limited to 50 processes running simultaneously. A runtime configuration must specify fewer than this limit. To calculate the total number of processes specified in a runtime configuration, add the values of the ConcurrentExecutions parameter for each ServerProcess object in the runtime configuration. (see [below for nested schema](#nestedatt--runtime_configuration))
- **script_id** (String) A unique identifier for a Realtime script to be deployed on a new Realtime Servers fleet. The script must have been successfully uploaded to Amazon GameLift. This fleet setting cannot be changed once the fleet is created.

Note: It is not currently possible to use the !Ref command to reference a script created with a CloudFormation template for the fleet property ScriptId. Instead, use Fn::GetAtt Script.Arn or Fn::GetAtt Script.Id to retrieve either of these properties as input for ScriptId. Alternatively, enter a ScriptId string manually.
- **server_launch_parameters** (String) This parameter is no longer used but is retained for backward compatibility. Instead, specify server launch parameters in the RuntimeConfiguration parameter. A request must specify either a runtime configuration or values for both ServerLaunchParameters and ServerLaunchPath.
- **server_launch_path** (String) This parameter is no longer used. Instead, specify a server launch path using the RuntimeConfiguration parameter. Requests that specify a server launch path and launch parameters instead of a runtime configuration will continue to work.

### Read-Only

- **fleet_id** (String) Unique fleet ID
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--certificate_configuration"></a>
### Nested Schema for `certificate_configuration`

Optional:

- **certificate_type** (String)


<a id="nestedatt--ec2_inbound_permissions"></a>
### Nested Schema for `ec2_inbound_permissions`

Optional:

- **from_port** (Number) A starting value for a range of allowed port numbers.
- **ip_range** (String) A range of allowed IP addresses. This value must be expressed in CIDR notation. Example: "000.000.000.000/[subnet mask]" or optionally the shortened version "0.0.0.0/[subnet mask]".
- **protocol** (String) The network communication protocol used by the fleet.
- **to_port** (Number) An ending value for a range of allowed port numbers. Port numbers are end-inclusive. This value must be higher than FromPort.


<a id="nestedatt--locations"></a>
### Nested Schema for `locations`

Optional:

- **location** (String)
- **location_capacity** (Attributes) Current resource capacity settings in a specified fleet or location. The location value might refer to a fleet's remote location or its home Region. (see [below for nested schema](#nestedatt--locations--location_capacity))

<a id="nestedatt--locations--location_capacity"></a>
### Nested Schema for `locations.location_capacity`

Optional:

- **desired_ec2_instances** (Number) The number of EC2 instances you want to maintain in the specified fleet location. This value must fall between the minimum and maximum size limits.
- **max_size** (Number) The maximum value that is allowed for the fleet's instance count for a location. When creating a new fleet, GameLift automatically sets this value to "1". Once the fleet is active, you can change this value.
- **min_size** (Number) The minimum value allowed for the fleet's instance count for a location. When creating a new fleet, GameLift automatically sets this value to "0". After the fleet is active, you can change this value.



<a id="nestedatt--resource_creation_limit_policy"></a>
### Nested Schema for `resource_creation_limit_policy`

Optional:

- **new_game_sessions_per_creator** (Number) The maximum number of game sessions that an individual can create during the policy period.
- **policy_period_in_minutes** (Number) The time span used in evaluating the resource creation limit policy.


<a id="nestedatt--runtime_configuration"></a>
### Nested Schema for `runtime_configuration`

Optional:

- **game_session_activation_timeout_seconds** (Number) The maximum amount of time (in seconds) that a game session can remain in status ACTIVATING. If the game session is not active before the timeout, activation is terminated and the game session status is changed to TERMINATED.
- **max_concurrent_game_session_activations** (Number) The maximum number of game sessions with status ACTIVATING to allow on an instance simultaneously. This setting limits the amount of instance resources that can be used for new game activations at any one time.
- **server_processes** (Attributes List) A collection of server process configurations that describe which server processes to run on each instance in a fleet. (see [below for nested schema](#nestedatt--runtime_configuration--server_processes))

<a id="nestedatt--runtime_configuration--server_processes"></a>
### Nested Schema for `runtime_configuration.server_processes`

Optional:

- **concurrent_executions** (Number) The number of server processes that use this configuration to run concurrently on an instance.
- **launch_path** (String) The location of the server executable in a custom game build or the name of the Realtime script file that contains the Init() function. Game builds and Realtime scripts are installed on instances at the root:

Windows (for custom game builds only): C:\game. Example: "C:\game\MyGame\server.exe"

Linux: /local/game. Examples: "/local/game/MyGame/server.exe" or "/local/game/MyRealtimeScript.js"
- **parameters** (String) An optional list of parameters to pass to the server executable or Realtime script on launch.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_gamelift_fleet.example <resource ID>
```