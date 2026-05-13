package googlefirestoreindex


type GoogleFirestoreIndexFieldsSearchConfigTextSpecIndexSpecs struct {
	// Ways to index the text field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firestore_index#index_type GoogleFirestoreIndex#index_type}
	IndexType *string `field:"optional" json:"indexType" yaml:"indexType"`
	// How to match the text field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firestore_index#match_type GoogleFirestoreIndex#match_type}
	MatchType *string `field:"optional" json:"matchType" yaml:"matchType"`
}

