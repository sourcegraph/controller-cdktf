package dataawssubnets


type DataAwsSubnetsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/subnets#name DataAwsSubnets#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/subnets#values DataAwsSubnets#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

