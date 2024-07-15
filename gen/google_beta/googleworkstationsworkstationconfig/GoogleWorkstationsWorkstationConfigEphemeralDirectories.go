package googleworkstationsworkstationconfig


type GoogleWorkstationsWorkstationConfigEphemeralDirectories struct {
	// gce_pd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#gce_pd GoogleWorkstationsWorkstationConfigA#gce_pd}
	GcePd *GoogleWorkstationsWorkstationConfigEphemeralDirectoriesGcePd `field:"optional" json:"gcePd" yaml:"gcePd"`
	// Location of this directory in the running workstation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_config#mount_path GoogleWorkstationsWorkstationConfigA#mount_path}
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

