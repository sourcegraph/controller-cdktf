package cloudwatcheventtarget


type CloudwatchEventTargetEcsTargetNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudwatch_event_target#subnets CloudwatchEventTarget#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudwatch_event_target#assign_public_ip CloudwatchEventTarget#assign_public_ip}.
	AssignPublicIp interface{} `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudwatch_event_target#security_groups CloudwatchEventTarget#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

