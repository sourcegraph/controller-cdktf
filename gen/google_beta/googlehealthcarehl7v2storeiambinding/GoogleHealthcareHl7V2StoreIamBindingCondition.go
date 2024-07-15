package googlehealthcarehl7v2storeiambinding


type GoogleHealthcareHl7V2StoreIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_hl7_v2_store_iam_binding#expression GoogleHealthcareHl7V2StoreIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_hl7_v2_store_iam_binding#title GoogleHealthcareHl7V2StoreIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_hl7_v2_store_iam_binding#description GoogleHealthcareHl7V2StoreIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

