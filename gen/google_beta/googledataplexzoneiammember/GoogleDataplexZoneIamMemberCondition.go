package googledataplexzoneiammember


type GoogleDataplexZoneIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_zone_iam_member#expression GoogleDataplexZoneIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_zone_iam_member#title GoogleDataplexZoneIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_zone_iam_member#description GoogleDataplexZoneIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

