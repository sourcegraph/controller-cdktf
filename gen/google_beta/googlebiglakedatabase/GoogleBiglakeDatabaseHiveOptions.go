package googlebiglakedatabase


type GoogleBiglakeDatabaseHiveOptions struct {
	// Cloud Storage folder URI where the database data is stored, starting with "gs://".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_biglake_database#location_uri GoogleBiglakeDatabase#location_uri}
	LocationUri *string `field:"optional" json:"locationUri" yaml:"locationUri"`
	// Stores user supplied Hive database parameters.
	//
	// An object containing a
	// list of"key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_biglake_database#parameters GoogleBiglakeDatabase#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

