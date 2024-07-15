package googlehealthcareconsentstoreiammember


type GoogleHealthcareConsentStoreIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_consent_store_iam_member#expression GoogleHealthcareConsentStoreIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_consent_store_iam_member#title GoogleHealthcareConsentStoreIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_consent_store_iam_member#description GoogleHealthcareConsentStoreIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

