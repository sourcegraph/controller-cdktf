package googledialogflowcxagent


type GoogleDialogflowCxAgentClientCertificateSettings struct {
	// The name of the SecretManager secret version resource storing the private key encoded in PEM format. Format: **projects/{project}/secrets/{secret}/versions/{version}**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#private_key GoogleDialogflowCxAgent#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#ssl_certificate GoogleDialogflowCxAgent#ssl_certificate}
	SslCertificate *string `field:"required" json:"sslCertificate" yaml:"sslCertificate"`
	// The name of the SecretManager secret version resource storing the passphrase.
	//
	// 'passphrase' should be left unset if the private key is not encrypted. Format: **projects/{project}/secrets/{secret}/versions/{version}**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#passphrase GoogleDialogflowCxAgent#passphrase}
	Passphrase *string `field:"optional" json:"passphrase" yaml:"passphrase"`
}

