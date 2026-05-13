package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfilePostgresqlProfileSslConfigServerVerification struct {
	// PEM-encoded server root CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_connection_profile#ca_certificate GoogleDatastreamConnectionProfile#ca_certificate}
	CaCertificate *string `field:"required" json:"caCertificate" yaml:"caCertificate"`
}

