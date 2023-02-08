package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfileGcsProfile struct {
	// The Cloud Storage bucket name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_connection_profile#bucket GoogleDatastreamConnectionProfile#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The root path inside the Cloud Storage bucket.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_connection_profile#root_path GoogleDatastreamConnectionProfile#root_path}
	RootPath *string `field:"optional" json:"rootPath" yaml:"rootPath"`
}

