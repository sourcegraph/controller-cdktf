package backupdrrestoreworkload


type BackupDrRestoreWorkloadDiskRestoreProperties struct {
	// Required. Name of the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#name BackupDrRestoreWorkload#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Required. The size of the disk in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#size_gb BackupDrRestoreWorkload#size_gb}
	SizeGb *float64 `field:"required" json:"sizeGb" yaml:"sizeGb"`
	// Required. URL of the disk type resource describing which disk type to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#type BackupDrRestoreWorkload#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Optional. The access mode of the disk. Possible values: ["READ_WRITE_SINGLE", "READ_WRITE_MANY", "READ_ONLY_MANY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#access_mode BackupDrRestoreWorkload#access_mode}
	AccessMode *string `field:"optional" json:"accessMode" yaml:"accessMode"`
	// Optional. The architecture of the source disk. Possible values: ["ARCHITECTURE_UNSPECIFIED", "X86_64", "ARM64"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#architecture BackupDrRestoreWorkload#architecture}
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// Optional. An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#description BackupDrRestoreWorkload#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#disk_encryption_key BackupDrRestoreWorkload#disk_encryption_key}
	DiskEncryptionKey *BackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKey `field:"optional" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// Optional. Indicates whether this disk is using confidential compute mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#enable_confidential_compute BackupDrRestoreWorkload#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"optional" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
	// guest_os_feature block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#guest_os_feature BackupDrRestoreWorkload#guest_os_feature}
	GuestOsFeature interface{} `field:"optional" json:"guestOsFeature" yaml:"guestOsFeature"`
	// labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#labels BackupDrRestoreWorkload#labels}
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// Optional. A list of publicly available licenses that are applicable to this backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#licenses BackupDrRestoreWorkload#licenses}
	Licenses *[]*string `field:"optional" json:"licenses" yaml:"licenses"`
	// Optional. Physical block size of the persistent disk, in bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#physical_block_size_bytes BackupDrRestoreWorkload#physical_block_size_bytes}
	PhysicalBlockSizeBytes *float64 `field:"optional" json:"physicalBlockSizeBytes" yaml:"physicalBlockSizeBytes"`
	// Optional. Indicates how many IOPS to provision for the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#provisioned_iops BackupDrRestoreWorkload#provisioned_iops}
	ProvisionedIops *float64 `field:"optional" json:"provisionedIops" yaml:"provisionedIops"`
	// Optional. Indicates how much throughput to provision for the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#provisioned_throughput BackupDrRestoreWorkload#provisioned_throughput}
	ProvisionedThroughput *float64 `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// resource_manager_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#resource_manager_tags BackupDrRestoreWorkload#resource_manager_tags}
	ResourceManagerTags interface{} `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
	// Optional. Resource policies applied to this disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#resource_policy BackupDrRestoreWorkload#resource_policy}
	ResourcePolicy *[]*string `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// Optional. The storage pool in which the new disk is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#storage_pool BackupDrRestoreWorkload#storage_pool}
	StoragePool *string `field:"optional" json:"storagePool" yaml:"storagePool"`
}

