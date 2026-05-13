package dialogflowcxagent


type DialogflowCxAgentClientCertificateSettings struct {
	// The name of the SecretManager secret version resource storing the private key encoded in PEM format. Format: **projects/{project}/secrets/{secret}/versions/{version}**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_agent#private_key DialogflowCxAgent#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_agent#ssl_certificate DialogflowCxAgent#ssl_certificate}
	SslCertificate *string `field:"required" json:"sslCertificate" yaml:"sslCertificate"`
	// The name of the SecretManager secret version resource storing the passphrase.
	//
	// 'passphrase' should be left unset if the private key is not encrypted. Format: **projects/{project}/secrets/{secret}/versions/{version}**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_agent#passphrase DialogflowCxAgent#passphrase}
	Passphrase *string `field:"optional" json:"passphrase" yaml:"passphrase"`
}

