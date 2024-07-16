package googlebigtableauthorizedview


type GoogleBigtableAuthorizedViewSubsetView struct {
	// family_subsets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_authorized_view#family_subsets GoogleBigtableAuthorizedView#family_subsets}
	FamilySubsets interface{} `field:"optional" json:"familySubsets" yaml:"familySubsets"`
	// Base64-encoded row prefixes to be included in the authorized view.
	//
	// To provide access to all rows, include the empty string as a prefix ("").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_authorized_view#row_prefixes GoogleBigtableAuthorizedView#row_prefixes}
	RowPrefixes *[]*string `field:"optional" json:"rowPrefixes" yaml:"rowPrefixes"`
}

