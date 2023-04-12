package googlehealthcaredatasetiambinding


type GoogleHealthcareDatasetIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_dataset_iam_binding#expression GoogleHealthcareDatasetIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_dataset_iam_binding#title GoogleHealthcareDatasetIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_dataset_iam_binding#description GoogleHealthcareDatasetIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

