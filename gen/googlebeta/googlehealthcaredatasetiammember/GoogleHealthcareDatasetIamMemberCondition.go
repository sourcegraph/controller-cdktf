package googlehealthcaredatasetiammember


type GoogleHealthcareDatasetIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_dataset_iam_member#expression GoogleHealthcareDatasetIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_dataset_iam_member#title GoogleHealthcareDatasetIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_dataset_iam_member#description GoogleHealthcareDatasetIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

