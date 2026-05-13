package firestoreindex


type FirestoreIndexFieldsSearchConfigTextSpec struct {
	// index_specs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firestore_index#index_specs FirestoreIndex#index_specs}
	IndexSpecs interface{} `field:"required" json:"indexSpecs" yaml:"indexSpecs"`
}

