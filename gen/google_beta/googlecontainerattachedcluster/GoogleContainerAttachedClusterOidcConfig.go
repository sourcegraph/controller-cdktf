package googlecontainerattachedcluster


type GoogleContainerAttachedClusterOidcConfig struct {
	// A JSON Web Token (JWT) issuer URI. 'issuer' must start with 'https://'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#issuer_url GoogleContainerAttachedCluster#issuer_url}
	IssuerUrl *string `field:"required" json:"issuerUrl" yaml:"issuerUrl"`
	// OIDC verification keys in JWKS format (RFC 7517).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#jwks GoogleContainerAttachedCluster#jwks}
	Jwks *string `field:"optional" json:"jwks" yaml:"jwks"`
}

