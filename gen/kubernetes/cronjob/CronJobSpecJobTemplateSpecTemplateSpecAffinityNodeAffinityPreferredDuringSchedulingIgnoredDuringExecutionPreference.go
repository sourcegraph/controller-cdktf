package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference struct {
	// match_expressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job#match_expressions CronJob#match_expressions}
	MatchExpressions interface{} `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
}

