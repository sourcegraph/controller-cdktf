package googlehealthcarehl7v2storeiammember


type GoogleHealthcareHl7V2StoreIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_hl7_v2_store_iam_member#expression GoogleHealthcareHl7V2StoreIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_hl7_v2_store_iam_member#title GoogleHealthcareHl7V2StoreIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_hl7_v2_store_iam_member#description GoogleHealthcareHl7V2StoreIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

