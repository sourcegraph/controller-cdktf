package googlechroniclefeed


type GoogleChronicleFeedDetailsFoxItStixSettingsSsl struct {
	// The encoded private key.
	//
	// The string should be a private key in PEM format,
	// and should include the begin header and end footer lines. It may also
	// include newlines.
	//
	// Example:
	// -----BEGIN RSA PRIVATE KEY-----
	// Proc-Type: 4,ENCRYPTED
	// DEK-Info: DES-EDE3-CBC,F23074E02CF47304
	//
	//
	// -----END RSA PRIVATE KEY-----
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#encoded_private_key GoogleChronicleFeed#encoded_private_key}
	EncodedPrivateKey *string `field:"optional" json:"encodedPrivateKey" yaml:"encodedPrivateKey"`
	// The encoded SSL certificate.
	//
	// The string should be an SSL certificate in
	// PEM format, and should include the begin header and end footer lines. It
	// may also include newlines.
	//
	// Example:
	// -----BEGIN CERTIFICATE-----
	//
	// -----END CERTIFICATE-----
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#ssl_certificate GoogleChronicleFeed#ssl_certificate}
	SslCertificate *string `field:"optional" json:"sslCertificate" yaml:"sslCertificate"`
}

