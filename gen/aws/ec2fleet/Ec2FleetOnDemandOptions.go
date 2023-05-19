package ec2fleet


type Ec2FleetOnDemandOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ec2_fleet#allocation_strategy Ec2Fleet#allocation_strategy}.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
}

