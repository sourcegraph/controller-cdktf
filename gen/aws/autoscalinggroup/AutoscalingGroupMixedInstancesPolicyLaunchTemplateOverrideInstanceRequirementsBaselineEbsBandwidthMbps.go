package autoscalinggroup


type AutoscalingGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsBaselineEbsBandwidthMbps struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/autoscaling_group#max AutoscalingGroup#max}.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/autoscaling_group#min AutoscalingGroup#min}.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

