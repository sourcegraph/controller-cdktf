package googlecloudrunv2jobiambinding


type GoogleCloudRunV2JobIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_job_iam_binding#expression GoogleCloudRunV2JobIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_job_iam_binding#title GoogleCloudRunV2JobIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_job_iam_binding#description GoogleCloudRunV2JobIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

