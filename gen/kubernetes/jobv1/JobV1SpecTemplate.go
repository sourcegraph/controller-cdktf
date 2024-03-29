package jobv1


type JobV1SpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/job_v1#metadata JobV1#metadata}
	Metadata *JobV1SpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/job_v1#spec JobV1#spec}
	Spec *JobV1SpecTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

