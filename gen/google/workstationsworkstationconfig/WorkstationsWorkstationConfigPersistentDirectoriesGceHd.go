package workstationsworkstationconfig


type WorkstationsWorkstationConfigPersistentDirectoriesGceHd struct {
	// How long to wait before converting the disk into a snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/workstations_workstation_config#archive_timeout WorkstationsWorkstationConfigA#archive_timeout}
	ArchiveTimeout *string `field:"optional" json:"archiveTimeout" yaml:"archiveTimeout"`
	// Whether the persistent disk should be deleted when the workstation is deleted. Possible values: ["DELETE", "RETAIN"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/workstations_workstation_config#reclaim_policy WorkstationsWorkstationConfigA#reclaim_policy}
	ReclaimPolicy *string `field:"optional" json:"reclaimPolicy" yaml:"reclaimPolicy"`
	// The GB capacity of a persistent home directory. Defaults to '200'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/workstations_workstation_config#size_gb WorkstationsWorkstationConfigA#size_gb}
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
	// Name of the snapshot to use as the source for the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/workstations_workstation_config#source_snapshot WorkstationsWorkstationConfigA#source_snapshot}
	SourceSnapshot *string `field:"optional" json:"sourceSnapshot" yaml:"sourceSnapshot"`
}

