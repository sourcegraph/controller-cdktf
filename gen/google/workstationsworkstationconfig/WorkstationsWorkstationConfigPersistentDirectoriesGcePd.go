package workstationsworkstationconfig


type WorkstationsWorkstationConfigPersistentDirectoriesGcePd struct {
	// The type of the persistent disk for the home directory. Defaults to 'pd-standard'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#disk_type WorkstationsWorkstationConfigA#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// Type of file system that the disk should be formatted with.
	//
	// The workstation image must support this file system type. Must be empty if 'sourceSnapshot' is set. Defaults to 'ext4'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#fs_type WorkstationsWorkstationConfigA#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Whether the persistent disk should be deleted when the workstation is deleted.
	//
	// Valid values are 'DELETE' and 'RETAIN'. Defaults to 'DELETE'. Possible values: ["DELETE", "RETAIN"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#reclaim_policy WorkstationsWorkstationConfigA#reclaim_policy}
	ReclaimPolicy *string `field:"optional" json:"reclaimPolicy" yaml:"reclaimPolicy"`
	// The GB capacity of a persistent home directory for each workstation created with this configuration.
	//
	// Must be empty if 'sourceSnapshot' is set.
	// Valid values are '10', '50', '100', '200', '500', or '1000'. Defaults to '200'. If less than '200' GB, the 'diskType' must be 'pd-balanced' or 'pd-ssd'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#size_gb WorkstationsWorkstationConfigA#size_gb}
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
	// Name of the snapshot to use as the source for the disk.
	//
	// This can be the snapshot's 'self_link', 'id', or a string in the format of 'projects/{project}/global/snapshots/{snapshot}'. If set, 'sizeGb' and 'fsType' must be empty. Can only be updated if it has an existing value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#source_snapshot WorkstationsWorkstationConfigA#source_snapshot}
	SourceSnapshot *string `field:"optional" json:"sourceSnapshot" yaml:"sourceSnapshot"`
}

