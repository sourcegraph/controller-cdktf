package googlefirestoreindex


type GoogleFirestoreIndexFieldsSearchConfig struct {
	// geo_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firestore_index#geo_spec GoogleFirestoreIndex#geo_spec}
	GeoSpec *GoogleFirestoreIndexFieldsSearchConfigGeoSpec `field:"optional" json:"geoSpec" yaml:"geoSpec"`
	// text_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firestore_index#text_spec GoogleFirestoreIndex#text_spec}
	TextSpec *GoogleFirestoreIndexFieldsSearchConfigTextSpec `field:"optional" json:"textSpec" yaml:"textSpec"`
}

