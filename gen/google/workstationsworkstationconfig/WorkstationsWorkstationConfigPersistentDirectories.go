package workstationsworkstationconfig


type WorkstationsWorkstationConfigPersistentDirectories struct {
	// gce_pd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#gce_pd WorkstationsWorkstationConfigA#gce_pd}
	GcePd *WorkstationsWorkstationConfigPersistentDirectoriesGcePd `field:"optional" json:"gcePd" yaml:"gcePd"`
	// Location of this directory in the running workstation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#mount_path WorkstationsWorkstationConfigA#mount_path}
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

