package route53recoveryreadinessresourceset


type Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResource struct {
	// nlb_resource block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/route53recoveryreadiness_resource_set#nlb_resource Route53RecoveryreadinessResourceSet#nlb_resource}
	NlbResource *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceNlbResource `field:"optional" json:"nlbResource" yaml:"nlbResource"`
	// r53_resource block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/route53recoveryreadiness_resource_set#r53_resource Route53RecoveryreadinessResourceSet#r53_resource}
	R53Resource *Route53RecoveryreadinessResourceSetResourcesDnsTargetResourceTargetResourceR53Resource `field:"optional" json:"r53Resource" yaml:"r53Resource"`
}

