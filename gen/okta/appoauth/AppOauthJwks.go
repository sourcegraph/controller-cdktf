package appoauth


type AppOauthJwks struct {
	// Key ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#kid AppOauth#kid}
	Kid *string `field:"required" json:"kid" yaml:"kid"`
	// Key type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#kty AppOauth#kty}
	Kty *string `field:"required" json:"kty" yaml:"kty"`
	// RSA Exponent.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#e AppOauth#e}
	E *string `field:"optional" json:"e" yaml:"e"`
	// RSA Modulus.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#n AppOauth#n}
	N *string `field:"optional" json:"n" yaml:"n"`
}

