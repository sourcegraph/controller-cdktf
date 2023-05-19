package ec2fleet


type Ec2FleetSpotOptionsMaintenanceStrategiesCapacityRebalance struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ec2_fleet#replacement_strategy Ec2Fleet#replacement_strategy}.
	ReplacementStrategy *string `field:"optional" json:"replacementStrategy" yaml:"replacementStrategy"`
}

