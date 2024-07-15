package googleprivatecacapool


type GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesEllipticCurve struct {
	// The algorithm used. Possible values: ["ECDSA_P256", "ECDSA_P384", "EDDSA_25519"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_ca_pool#signature_algorithm GooglePrivatecaCaPool#signature_algorithm}
	SignatureAlgorithm *string `field:"required" json:"signatureAlgorithm" yaml:"signatureAlgorithm"`
}

