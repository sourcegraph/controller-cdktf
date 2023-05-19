package googleprivatecacertificateauthority


type GooglePrivatecaCertificateAuthoritySubordinateConfig struct {
	// This can refer to a CertificateAuthority that was used to create a subordinate CertificateAuthority.
	//
	// This field is used for information
	// and usability purposes only. The resource name is in the format
	// 'projects/*\/locations/*\/caPools/*\/certificateAuthorities/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_privateca_certificate_authority#certificate_authority GooglePrivatecaCertificateAuthority#certificate_authority}
	CertificateAuthority *string `field:"optional" json:"certificateAuthority" yaml:"certificateAuthority"`
	// pem_issuer_chain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_privateca_certificate_authority#pem_issuer_chain GooglePrivatecaCertificateAuthority#pem_issuer_chain}
	PemIssuerChain *GooglePrivatecaCertificateAuthoritySubordinateConfigPemIssuerChain `field:"optional" json:"pemIssuerChain" yaml:"pemIssuerChain"`
}

