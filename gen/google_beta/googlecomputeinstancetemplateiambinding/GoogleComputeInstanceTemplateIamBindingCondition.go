package googlecomputeinstancetemplateiambinding


type GoogleComputeInstanceTemplateIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_template_iam_binding#expression GoogleComputeInstanceTemplateIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_template_iam_binding#title GoogleComputeInstanceTemplateIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_template_iam_binding#description GoogleComputeInstanceTemplateIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

