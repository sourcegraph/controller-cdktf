package googlednsmanagedzoneiammember


type GoogleDnsManagedZoneIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone_iam_member#expression GoogleDnsManagedZoneIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone_iam_member#title GoogleDnsManagedZoneIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone_iam_member#description GoogleDnsManagedZoneIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

