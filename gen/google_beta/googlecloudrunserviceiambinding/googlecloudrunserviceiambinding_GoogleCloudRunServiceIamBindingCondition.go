package googlecloudrunserviceiambinding


type GoogleCloudRunServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service_iam_binding#expression GoogleCloudRunServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service_iam_binding#title GoogleCloudRunServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service_iam_binding#description GoogleCloudRunServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

