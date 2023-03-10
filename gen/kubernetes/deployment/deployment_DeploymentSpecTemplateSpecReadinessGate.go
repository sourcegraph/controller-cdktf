package deployment


type DeploymentSpecTemplateSpecReadinessGate struct {
	// refers to a condition in the pod's condition list with matching type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#condition_type Deployment#condition_type}
	ConditionType *string `field:"required" json:"conditionType" yaml:"conditionType"`
}

