package googlelookerinstance


type GoogleLookerInstancePeriodicExportConfig struct {
	// Cloud Storage bucket URI for periodic export. Format: gs://{bucket_name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_looker_instance#gcs_uri GoogleLookerInstance#gcs_uri}
	GcsUri *string `field:"required" json:"gcsUri" yaml:"gcsUri"`
	// Name of the CMEK key in KMS. Format: projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_looker_instance#kms_key GoogleLookerInstance#kms_key}
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_looker_instance#start_time GoogleLookerInstance#start_time}
	StartTime *GoogleLookerInstancePeriodicExportConfigStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

