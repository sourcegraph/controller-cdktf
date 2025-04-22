package googlesccv2organizationsourceiambinding


type GoogleSccV2OrganizationSourceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_scc_v2_organization_source_iam_binding#expression GoogleSccV2OrganizationSourceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_scc_v2_organization_source_iam_binding#title GoogleSccV2OrganizationSourceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_scc_v2_organization_source_iam_binding#description GoogleSccV2OrganizationSourceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

