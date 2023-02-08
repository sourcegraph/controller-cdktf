package googlehealthcarefhirstoreiambinding


type GoogleHealthcareFhirStoreIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_fhir_store_iam_binding#expression GoogleHealthcareFhirStoreIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_fhir_store_iam_binding#title GoogleHealthcareFhirStoreIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_fhir_store_iam_binding#description GoogleHealthcareFhirStoreIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

