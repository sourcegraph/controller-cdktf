package googleiapwebtypecomputeiammember


type GoogleIapWebTypeComputeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_web_type_compute_iam_member#expression GoogleIapWebTypeComputeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_web_type_compute_iam_member#title GoogleIapWebTypeComputeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_web_type_compute_iam_member#description GoogleIapWebTypeComputeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

