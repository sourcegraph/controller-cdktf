package job


type JobSpecTemplateSpecReadinessGate struct {
	// refers to a condition in the pod's condition list with matching type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/job#condition_type Job#condition_type}
	ConditionType *string `field:"required" json:"conditionType" yaml:"conditionType"`
}

