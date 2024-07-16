package googlecertificatemanagertrustconfig


type GoogleCertificateManagerTrustConfigTrustStores struct {
	// intermediate_cas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_certificate_manager_trust_config#intermediate_cas GoogleCertificateManagerTrustConfig#intermediate_cas}
	IntermediateCas interface{} `field:"optional" json:"intermediateCas" yaml:"intermediateCas"`
	// trust_anchors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_certificate_manager_trust_config#trust_anchors GoogleCertificateManagerTrustConfig#trust_anchors}
	TrustAnchors interface{} `field:"optional" json:"trustAnchors" yaml:"trustAnchors"`
}

