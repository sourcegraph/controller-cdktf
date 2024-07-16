package googledatastreamstream


type GoogleDatastreamStreamDestinationConfig struct {
	// Destination connection profile resource. Format: projects/{project}/locations/{location}/connectionProfiles/{name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#destination_connection_profile GoogleDatastreamStream#destination_connection_profile}
	DestinationConnectionProfile *string `field:"required" json:"destinationConnectionProfile" yaml:"destinationConnectionProfile"`
	// bigquery_destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#bigquery_destination_config GoogleDatastreamStream#bigquery_destination_config}
	BigqueryDestinationConfig *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfig `field:"optional" json:"bigqueryDestinationConfig" yaml:"bigqueryDestinationConfig"`
	// gcs_destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#gcs_destination_config GoogleDatastreamStream#gcs_destination_config}
	GcsDestinationConfig *GoogleDatastreamStreamDestinationConfigGcsDestinationConfig `field:"optional" json:"gcsDestinationConfig" yaml:"gcsDestinationConfig"`
}

