package googlefirestoreindex


type GoogleFirestoreIndexFields struct {
	// Indicates that this field supports operations on arrayValues.
	//
	// Only one of 'order', 'arrayConfig', and
	// 'vectorConfig' can be specified. Possible values: ["CONTAINS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_index#array_config GoogleFirestoreIndex#array_config}
	ArrayConfig *string `field:"optional" json:"arrayConfig" yaml:"arrayConfig"`
	// Name of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_index#field_path GoogleFirestoreIndex#field_path}
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
	// Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
	//
	// Only one of 'order', 'arrayConfig', and 'vectorConfig' can be specified. Possible values: ["ASCENDING", "DESCENDING"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_index#order GoogleFirestoreIndex#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// vector_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_index#vector_config GoogleFirestoreIndex#vector_config}
	VectorConfig *GoogleFirestoreIndexFieldsVectorConfig `field:"optional" json:"vectorConfig" yaml:"vectorConfig"`
}

