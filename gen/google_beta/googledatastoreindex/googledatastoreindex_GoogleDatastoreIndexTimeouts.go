package googledatastoreindex


type GoogleDatastoreIndexTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastore_index#create GoogleDatastoreIndex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastore_index#delete GoogleDatastoreIndex#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

