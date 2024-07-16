package googlefirestoreindex


type GoogleFirestoreIndexFieldsVectorConfig struct {
	// The resulting index will only include vectors of this dimension, and can be used for vector search with the same dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_index#dimension GoogleFirestoreIndex#dimension}
	Dimension *float64 `field:"optional" json:"dimension" yaml:"dimension"`
	// flat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_index#flat GoogleFirestoreIndex#flat}
	Flat *GoogleFirestoreIndexFieldsVectorConfigFlat `field:"optional" json:"flat" yaml:"flat"`
}

