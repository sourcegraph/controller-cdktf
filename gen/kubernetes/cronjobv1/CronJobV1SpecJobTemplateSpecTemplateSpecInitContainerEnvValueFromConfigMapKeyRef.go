package cronjobv1


type CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRef struct {
	// The key to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job_v1#key CronJobV1#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job_v1#name CronJobV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the ConfigMap or its key must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job_v1#optional CronJobV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

