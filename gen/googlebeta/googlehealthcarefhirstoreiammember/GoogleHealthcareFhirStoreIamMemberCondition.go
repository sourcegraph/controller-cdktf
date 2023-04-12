package googlehealthcarefhirstoreiammember


type GoogleHealthcareFhirStoreIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_fhir_store_iam_member#expression GoogleHealthcareFhirStoreIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_fhir_store_iam_member#title GoogleHealthcareFhirStoreIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_fhir_store_iam_member#description GoogleHealthcareFhirStoreIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

