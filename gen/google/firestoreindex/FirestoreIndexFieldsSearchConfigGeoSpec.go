package firestoreindex


type FirestoreIndexFieldsSearchConfigGeoSpec struct {
	// If true, disables GeoJSON indexing for the field.
	//
	// By default, GeoJSON points are indexed.
	// Firestore GeoPoints are indexed regardless of the value of this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firestore_index#geo_json_indexing_disabled FirestoreIndex#geo_json_indexing_disabled}
	GeoJsonIndexingDisabled interface{} `field:"required" json:"geoJsonIndexingDisabled" yaml:"geoJsonIndexingDisabled"`
}

