package codedeploydeploymentgroup


type CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group#action_on_timeout CodedeployDeploymentGroup#action_on_timeout}.
	ActionOnTimeout *string `field:"optional" json:"actionOnTimeout" yaml:"actionOnTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codedeploy_deployment_group#wait_time_in_minutes CodedeployDeploymentGroup#wait_time_in_minutes}.
	WaitTimeInMinutes *float64 `field:"optional" json:"waitTimeInMinutes" yaml:"waitTimeInMinutes"`
}

