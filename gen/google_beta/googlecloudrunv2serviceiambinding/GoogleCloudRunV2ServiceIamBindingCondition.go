package googlecloudrunv2serviceiambinding


type GoogleCloudRunV2ServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service_iam_binding#expression GoogleCloudRunV2ServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service_iam_binding#title GoogleCloudRunV2ServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_run_v2_service_iam_binding#description GoogleCloudRunV2ServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

