package cesapp


type CesAppLoggingSettingsAudioRecordingConfig struct {
	// The [Cloud Storage](https://cloud.google.com/storage) bucket to store the session audio recordings. The URI must start with "gs://". Note: If the Cloud Storage bucket is in a different project from the app, you should grant 'storage.objects.create' permission to the CES service agent 'service-@gcp-sa-ces.iam.gserviceaccount.com'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#gcs_bucket CesApp#gcs_bucket}
	GcsBucket *string `field:"optional" json:"gcsBucket" yaml:"gcsBucket"`
	// The Cloud Storage path prefix for audio recordings.
	//
	// This prefix can include the following placeholders, which will be
	// dynamically substituted at serving time:
	// - $project:   project ID
	// - $location:  app location
	// - $app:       app ID
	// - $date:      session date in YYYY-MM-DD format
	// - $session:   session ID
	// If the path prefix is not specified, the default prefix
	// '$project/$location/$app/$date/$session/' will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#gcs_path_prefix CesApp#gcs_path_prefix}
	GcsPathPrefix *string `field:"optional" json:"gcsPathPrefix" yaml:"gcsPathPrefix"`
}

