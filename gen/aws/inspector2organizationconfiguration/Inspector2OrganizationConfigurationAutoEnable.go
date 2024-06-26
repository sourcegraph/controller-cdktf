package inspector2organizationconfiguration


type Inspector2OrganizationConfigurationAutoEnable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/inspector2_organization_configuration#ec2 Inspector2OrganizationConfiguration#ec2}.
	Ec2 interface{} `field:"required" json:"ec2" yaml:"ec2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/inspector2_organization_configuration#ecr Inspector2OrganizationConfiguration#ecr}.
	Ecr interface{} `field:"required" json:"ecr" yaml:"ecr"`
}

