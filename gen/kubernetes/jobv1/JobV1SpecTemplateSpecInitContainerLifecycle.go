package jobv1


type JobV1SpecTemplateSpecInitContainerLifecycle struct {
	// post_start block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/job_v1#post_start JobV1#post_start}
	PostStart interface{} `field:"optional" json:"postStart" yaml:"postStart"`
	// pre_stop block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/job_v1#pre_stop JobV1#pre_stop}
	PreStop interface{} `field:"optional" json:"preStop" yaml:"preStop"`
}

