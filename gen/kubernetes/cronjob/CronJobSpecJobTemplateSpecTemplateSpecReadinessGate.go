package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecReadinessGate struct {
	// refers to a condition in the pod's condition list with matching type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job#condition_type CronJob#condition_type}
	ConditionType *string `field:"required" json:"conditionType" yaml:"conditionType"`
}

