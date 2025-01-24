package googlesccv2organizationsourceiammember


type GoogleSccV2OrganizationSourceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_organization_source_iam_member#expression GoogleSccV2OrganizationSourceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_organization_source_iam_member#title GoogleSccV2OrganizationSourceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_organization_source_iam_member#description GoogleSccV2OrganizationSourceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

