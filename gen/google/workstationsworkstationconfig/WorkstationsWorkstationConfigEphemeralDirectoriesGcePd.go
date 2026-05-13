package workstationsworkstationconfig


type WorkstationsWorkstationConfigEphemeralDirectoriesGcePd struct {
	// Type of the disk to use. Defaults to '"pd-standard"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#disk_type WorkstationsWorkstationConfigA#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// Whether the disk is read only.
	//
	// If true, the disk may be shared by multiple VMs and 'sourceSnapshot' must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#read_only WorkstationsWorkstationConfigA#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Name of the disk image to use as the source for the disk.
	//
	// Must be empty 'sourceSnapshot' is set.
	// Updating 'sourceImage' will update content in the ephemeral directory after the workstation is restarted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#source_image WorkstationsWorkstationConfigA#source_image}
	SourceImage *string `field:"optional" json:"sourceImage" yaml:"sourceImage"`
	// Name of the snapshot to use as the source for the disk.
	//
	// Must be empty if 'sourceImage' is set.
	// Must be empty if 'read_only' is false.
	// Updating 'source_snapshot' will update content in the ephemeral directory after the workstation is restarted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#source_snapshot WorkstationsWorkstationConfigA#source_snapshot}
	SourceSnapshot *string `field:"optional" json:"sourceSnapshot" yaml:"sourceSnapshot"`
}

