package spotfleetrequest


type SpotFleetRequestLaunchTemplateConfigOverridesInstanceRequirementsAcceleratorTotalMemoryMib struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/spot_fleet_request#max SpotFleetRequest#max}.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/spot_fleet_request#min SpotFleetRequest#min}.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

