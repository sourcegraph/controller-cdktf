package googlecloudrunv2jobiammember


type GoogleCloudRunV2JobIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_job_iam_member#expression GoogleCloudRunV2JobIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_job_iam_member#title GoogleCloudRunV2JobIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_job_iam_member#description GoogleCloudRunV2JobIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

