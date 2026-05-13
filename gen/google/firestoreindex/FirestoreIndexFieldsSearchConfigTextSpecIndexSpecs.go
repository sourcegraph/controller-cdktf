package firestoreindex


type FirestoreIndexFieldsSearchConfigTextSpecIndexSpecs struct {
	// Ways to index the text field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firestore_index#index_type FirestoreIndex#index_type}
	IndexType *string `field:"optional" json:"indexType" yaml:"indexType"`
	// How to match the text field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firestore_index#match_type FirestoreIndex#match_type}
	MatchType *string `field:"optional" json:"matchType" yaml:"matchType"`
}

