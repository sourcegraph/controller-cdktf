package googleiapwebcloudrunserviceiammember


type GoogleIapWebCloudRunServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_web_cloud_run_service_iam_member#expression GoogleIapWebCloudRunServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_web_cloud_run_service_iam_member#title GoogleIapWebCloudRunServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_web_cloud_run_service_iam_member#description GoogleIapWebCloudRunServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

