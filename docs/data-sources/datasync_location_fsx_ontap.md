---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_datasync_location_fsx_ontap Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::DataSync::LocationFSxONTAP
---

# awscc_datasync_location_fsx_ontap (Data Source)

Data Source schema for AWS::DataSync::LocationFSxONTAP



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `fsx_filesystem_arn` (String) The Amazon Resource Name (ARN) for the FSx ONAP file system.
- `location_arn` (String) The Amazon Resource Name (ARN) of the Amazon FSx ONTAP file system location that is created.
- `location_uri` (String) The URL of the FSx ONTAP file system that was described.
- `protocol` (Attributes) Configuration settings for NFS or SMB protocol. (see [below for nested schema](#nestedatt--protocol))
- `security_group_arns` (List of String) The ARNs of the security groups that are to use to configure the FSx ONTAP file system.
- `storage_virtual_machine_arn` (String) The Amazon Resource Name (ARN) for the FSx ONTAP SVM.
- `subdirectory` (String) A subdirectory in the location's path.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--protocol"></a>
### Nested Schema for `protocol`

Read-Only:

- `nfs` (Attributes) NFS protocol configuration for FSx ONTAP file system. (see [below for nested schema](#nestedatt--protocol--nfs))
- `smb` (Attributes) SMB protocol configuration for FSx ONTAP file system. (see [below for nested schema](#nestedatt--protocol--smb))

<a id="nestedatt--protocol--nfs"></a>
### Nested Schema for `protocol.nfs`

Read-Only:

- `mount_options` (Attributes) The NFS mount options that DataSync can use to mount your NFS share. (see [below for nested schema](#nestedatt--protocol--nfs--mount_options))

<a id="nestedatt--protocol--nfs--mount_options"></a>
### Nested Schema for `protocol.nfs.mount_options`

Read-Only:

- `version` (String) The specific NFS version that you want DataSync to use to mount your NFS share.



<a id="nestedatt--protocol--smb"></a>
### Nested Schema for `protocol.smb`

Read-Only:

- `domain` (String) The name of the Windows domain that the SMB server belongs to.
- `mount_options` (Attributes) The mount options used by DataSync to access the SMB server. (see [below for nested schema](#nestedatt--protocol--smb--mount_options))
- `password` (String) The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.
- `user` (String) The user who can mount the share, has the permissions to access files and folders in the SMB share.

<a id="nestedatt--protocol--smb--mount_options"></a>
### Nested Schema for `protocol.smb.mount_options`

Read-Only:

- `version` (String) The specific SMB version that you want DataSync to use to mount your SMB share.




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key for an AWS resource tag.
- `value` (String) The value for an AWS resource tag.

