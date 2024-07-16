package googlebigtableauthorizedview


type GoogleBigtableAuthorizedViewSubsetViewFamilySubsets struct {
	// Name of the column family to be included in the authorized view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_authorized_view#family_name GoogleBigtableAuthorizedView#family_name}
	FamilyName *string `field:"required" json:"familyName" yaml:"familyName"`
	// Base64-encoded prefixes for qualifiers of the column family to be included in the authorized view.
	//
	// Every qualifier starting with one of these prefixes is included in the authorized view. To provide access to all qualifiers, include the empty string as a prefix ("").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_authorized_view#qualifier_prefixes GoogleBigtableAuthorizedView#qualifier_prefixes}
	QualifierPrefixes *[]*string `field:"optional" json:"qualifierPrefixes" yaml:"qualifierPrefixes"`
	// Base64-encoded individual exact column qualifiers of the column family to be included in the authorized view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_authorized_view#qualifiers GoogleBigtableAuthorizedView#qualifiers}
	Qualifiers *[]*string `field:"optional" json:"qualifiers" yaml:"qualifiers"`
}

