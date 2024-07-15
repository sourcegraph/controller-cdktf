package googleprivatecacapool


type GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypes struct {
	// elliptic_curve block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_ca_pool#elliptic_curve GooglePrivatecaCaPool#elliptic_curve}
	EllipticCurve *GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve `field:"optional" json:"ellipticCurve" yaml:"ellipticCurve"`
	// rsa block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_ca_pool#rsa GooglePrivatecaCaPool#rsa}
	Rsa *GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa `field:"optional" json:"rsa" yaml:"rsa"`
}

