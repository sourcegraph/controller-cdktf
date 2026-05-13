package workstationsworkstationconfig


type WorkstationsWorkstationConfigContainer struct {
	// Arguments passed to the entrypoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#args WorkstationsWorkstationConfigA#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// If set, overrides the default ENTRYPOINT specified by the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#command WorkstationsWorkstationConfigA#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Environment variables passed to the container.
	//
	// The elements are of the form "KEY=VALUE" for the environment variable "KEY" being given the value "VALUE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#env WorkstationsWorkstationConfigA#env}
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// Docker image defining the container. This image must be accessible by the config's service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#image WorkstationsWorkstationConfigA#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// If set, overrides the USER specified in the image with the given uid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#run_as_user WorkstationsWorkstationConfigA#run_as_user}
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
	// If set, overrides the default DIR specified by the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#working_dir WorkstationsWorkstationConfigA#working_dir}
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

