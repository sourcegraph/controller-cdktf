package route53zone


type Route53ZoneVpc struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/route53_zone#vpc_id Route53Zone#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/route53_zone#vpc_region Route53Zone#vpc_region}.
	VpcRegion *string `field:"optional" json:"vpcRegion" yaml:"vpcRegion"`
}

