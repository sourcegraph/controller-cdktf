package oauthapp


type OauthAppJwks struct {
	// Key ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/oauth_app#kid OauthApp#kid}
	Kid *string `field:"required" json:"kid" yaml:"kid"`
	// Key type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/oauth_app#kty OauthApp#kty}
	Kty *string `field:"required" json:"kty" yaml:"kty"`
	// RSA Exponent.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/oauth_app#e OauthApp#e}
	E *string `field:"optional" json:"e" yaml:"e"`
	// RSA Modulus.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/oauth_app#n OauthApp#n}
	N *string `field:"optional" json:"n" yaml:"n"`
}

