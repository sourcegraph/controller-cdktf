package googleworkstationsworkstationconfig


type GoogleWorkstationsWorkstationConfigPersistentDirectoriesGcePd struct {
	// Type of the disk to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#disk_type GoogleWorkstationsWorkstationConfigA#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// Type of file system that the disk should be formatted with.
	//
	// The workstation image must support this file system type. Must be empty if sourceSnapshot is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#fs_type GoogleWorkstationsWorkstationConfigA#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// What should happen to the disk after the workstation is deleted. Defaults to DELETE. Possible values: ["DELETE", "RETAIN"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#reclaim_policy GoogleWorkstationsWorkstationConfigA#reclaim_policy}
	ReclaimPolicy *string `field:"optional" json:"reclaimPolicy" yaml:"reclaimPolicy"`
	// Size of the disk in GB. Must be empty if sourceSnapshot is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#size_gb GoogleWorkstationsWorkstationConfigA#size_gb}
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
}

