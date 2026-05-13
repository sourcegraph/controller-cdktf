package datastreamconnectionprofile


type DatastreamConnectionProfilePostgresqlProfileSslConfigServerVerification struct {
	// PEM-encoded server root CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_connection_profile#ca_certificate DatastreamConnectionProfile#ca_certificate}
	CaCertificate *string `field:"required" json:"caCertificate" yaml:"caCertificate"`
}

