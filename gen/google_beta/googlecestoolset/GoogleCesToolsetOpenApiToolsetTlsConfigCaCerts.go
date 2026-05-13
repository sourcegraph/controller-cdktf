package googlecestoolset


type GoogleCesToolsetOpenApiToolsetTlsConfigCaCerts struct {
	// The allowed custom CA certificates (in DER format) for HTTPS verification.
	//
	// This overrides the default SSL trust store. If this
	// is empty or unspecified, CES will use Google's default trust
	// store to verify certificates. N.B. Make sure the HTTPS server
	// certificates are signed with "subject alt name". For instance a
	// certificate can be self-signed using the following command,
	// openssl x509 -req -days 200 -in example.com.csr \
	// -signkey example.com.key \
	// -out example.com.crt \
	// -extfile <(printf "\nsubjectAltName='DNS:www.example.com'")
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#cert GoogleCesToolset#cert}
	Cert *string `field:"required" json:"cert" yaml:"cert"`
	// The name of the allowed custom CA certificates. This can be used to disambiguate the custom CA certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#display_name GoogleCesToolset#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
}

