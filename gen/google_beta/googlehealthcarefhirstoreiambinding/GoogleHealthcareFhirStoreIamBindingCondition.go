package googlehealthcarefhirstoreiambinding


type GoogleHealthcareFhirStoreIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_fhir_store_iam_binding#expression GoogleHealthcareFhirStoreIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_fhir_store_iam_binding#title GoogleHealthcareFhirStoreIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_fhir_store_iam_binding#description GoogleHealthcareFhirStoreIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

