package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfileMysqlProfileSslConfig struct {
	// PEM-encoded certificate of the CA that signed the source database server's certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_connection_profile#ca_certificate GoogleDatastreamConnectionProfile#ca_certificate}
	CaCertificate *string `field:"optional" json:"caCertificate" yaml:"caCertificate"`
	// PEM-encoded certificate that will be used by the replica to authenticate against the source database server.
	//
	// If this field
	// is used then the 'clientKey' and the 'caCertificate' fields are
	// mandatory.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_connection_profile#client_certificate GoogleDatastreamConnectionProfile#client_certificate}
	ClientCertificate *string `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// PEM-encoded private key associated with the Client Certificate.
	//
	// If this field is used then the 'client_certificate' and the
	// 'ca_certificate' fields are mandatory.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_connection_profile#client_key GoogleDatastreamConnectionProfile#client_key}
	ClientKey *string `field:"optional" json:"clientKey" yaml:"clientKey"`
}

