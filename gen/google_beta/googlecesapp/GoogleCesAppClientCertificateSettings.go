package googlecesapp


type GoogleCesAppClientCertificateSettings struct {
	// The name of the SecretManager secret version resource storing the private key encoded in PEM format. Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#private_key GoogleCesApp#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// The TLS certificate encoded in PEM format. This string must include the begin header and end footer lines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#tls_certificate GoogleCesApp#tls_certificate}
	TlsCertificate *string `field:"required" json:"tlsCertificate" yaml:"tlsCertificate"`
	// The passphrase to decrypt the private key. Should be left unset if the private key is not encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#passphrase GoogleCesApp#passphrase}
	Passphrase *string `field:"optional" json:"passphrase" yaml:"passphrase"`
}

