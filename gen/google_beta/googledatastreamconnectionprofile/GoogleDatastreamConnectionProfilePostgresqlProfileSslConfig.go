package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfilePostgresqlProfileSslConfig struct {
	// server_and_client_verification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_connection_profile#server_and_client_verification GoogleDatastreamConnectionProfile#server_and_client_verification}
	ServerAndClientVerification *GoogleDatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerification `field:"optional" json:"serverAndClientVerification" yaml:"serverAndClientVerification"`
	// server_verification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_connection_profile#server_verification GoogleDatastreamConnectionProfile#server_verification}
	ServerVerification *GoogleDatastreamConnectionProfilePostgresqlProfileSslConfigServerVerification `field:"optional" json:"serverVerification" yaml:"serverVerification"`
}

