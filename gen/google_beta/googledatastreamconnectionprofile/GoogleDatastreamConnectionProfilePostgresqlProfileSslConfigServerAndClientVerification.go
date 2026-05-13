package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfilePostgresqlProfileSslConfigServerAndClientVerification struct {
	// PEM-encoded server root CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_connection_profile#ca_certificate GoogleDatastreamConnectionProfile#ca_certificate}
	CaCertificate *string `field:"required" json:"caCertificate" yaml:"caCertificate"`
	// PEM-encoded certificate used by the source database to authenticate the client identity (i.e., the Datastream's identity). This certificate is signed by either a root certificate trusted by the server or one or more intermediate certificates (which is stored with the leaf certificate) to link to this certificate to the trusted root certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_connection_profile#client_certificate GoogleDatastreamConnectionProfile#client_certificate}
	ClientCertificate *string `field:"required" json:"clientCertificate" yaml:"clientCertificate"`
	// PEM-encoded private key associated with the client certificate.
	//
	// This value will be used during the SSL/TLS handshake, allowing
	// the PostgreSQL server to authenticate the client's identity,
	// i.e. identity of the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_connection_profile#client_key GoogleDatastreamConnectionProfile#client_key}
	ClientKey *string `field:"required" json:"clientKey" yaml:"clientKey"`
}

