package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisks struct {
	// Optional. Specifies whether the disk will be auto-deleted when the instance is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#auto_delete GoogleBackupDrRestoreWorkload#auto_delete}
	AutoDelete interface{} `field:"optional" json:"autoDelete" yaml:"autoDelete"`
	// Optional. Indicates that this is a boot disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#boot GoogleBackupDrRestoreWorkload#boot}
	Boot interface{} `field:"optional" json:"boot" yaml:"boot"`
	// Optional. This is used as an identifier for the disks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#device_name GoogleBackupDrRestoreWorkload#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#disk_encryption_key GoogleBackupDrRestoreWorkload#disk_encryption_key}
	DiskEncryptionKey *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksDiskEncryptionKey `field:"optional" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// Optional. Specifies the disk interface to use for attaching this disk. Possible values: ["DISK_INTERFACE_UNSPECIFIED", "SCSI", "NVME", "NVDIMM", "ISCSI"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#disk_interface GoogleBackupDrRestoreWorkload#disk_interface}
	DiskInterface *string `field:"optional" json:"diskInterface" yaml:"diskInterface"`
	// Optional. The size of the disk in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#disk_size_gb GoogleBackupDrRestoreWorkload#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Output only. The URI of the disk type resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#disk_type GoogleBackupDrRestoreWorkload#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// guest_os_feature block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#guest_os_feature GoogleBackupDrRestoreWorkload#guest_os_feature}
	GuestOsFeature interface{} `field:"optional" json:"guestOsFeature" yaml:"guestOsFeature"`
	// Optional. A zero-based index to this disk, where 0 is reserved for the boot disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#index GoogleBackupDrRestoreWorkload#index}
	Index *float64 `field:"optional" json:"index" yaml:"index"`
	// initialize_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#initialize_params GoogleBackupDrRestoreWorkload#initialize_params}
	InitializeParams *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksInitializeParams `field:"optional" json:"initializeParams" yaml:"initializeParams"`
	// Optional. Type of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#kind GoogleBackupDrRestoreWorkload#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Optional. Any valid publicly visible licenses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#license GoogleBackupDrRestoreWorkload#license}
	License *[]*string `field:"optional" json:"license" yaml:"license"`
	// Optional. The mode in which to attach this disk. Possible values: ["DISK_MODE_UNSPECIFIED", "READ_WRITE", "READ_ONLY", "LOCKED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#mode GoogleBackupDrRestoreWorkload#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Optional. Specifies the saved state of the disk. Possible values: ["DISK_SAVED_STATE_UNSPECIFIED", "PRESERVED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#saved_state GoogleBackupDrRestoreWorkload#saved_state}
	SavedState *string `field:"optional" json:"savedState" yaml:"savedState"`
	// Optional. Specifies a valid partial or full URL to an existing Persistent Disk resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#source GoogleBackupDrRestoreWorkload#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Optional. Specifies the type of the disk. Possible values: ["DISK_TYPE_UNSPECIFIED", "SCRATCH", "PERSISTENT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#type GoogleBackupDrRestoreWorkload#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

