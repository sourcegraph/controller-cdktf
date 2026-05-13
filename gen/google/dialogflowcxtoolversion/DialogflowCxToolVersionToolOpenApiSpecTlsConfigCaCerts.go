package dialogflowcxtoolversion


type DialogflowCxToolVersionToolOpenApiSpecTlsConfigCaCerts struct {
	// The allowed custom CA certificates (in DER format) for HTTPS verification.
	//
	// This overrides the default SSL trust store.
	// If this is empty or unspecified, Dialogflow will use Google's default trust store to verify certificates.
	// N.B. Make sure the HTTPS server certificates are signed with "subject alt name".
	// For instance a certificate can be self-signed using the following command:
	// ```
	//   openssl x509 -req -days 200 -in example.com.csr \
	//     -signkey example.com.key \
	//     -out example.com.crt \
	//     -extfile <(printf "\nsubjectAltName='DNS:www.example.com'")
	// ```
	// A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#cert DialogflowCxToolVersion#cert}
	Cert *string `field:"required" json:"cert" yaml:"cert"`
	// The name of the allowed custom CA certificates. This can be used to disambiguate the custom CA certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#display_name DialogflowCxToolVersion#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
}

