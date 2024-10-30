package appoauth


type AppOauthJwks struct {
	// Key ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_oauth#kid AppOauth#kid}
	Kid *string `field:"required" json:"kid" yaml:"kid"`
	// Key type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_oauth#kty AppOauth#kty}
	Kty *string `field:"required" json:"kty" yaml:"kty"`
	// RSA Exponent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_oauth#e AppOauth#e}
	E *string `field:"optional" json:"e" yaml:"e"`
	// RSA Modulus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_oauth#n AppOauth#n}
	N *string `field:"optional" json:"n" yaml:"n"`
	// X coordinate of the elliptic curve point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_oauth#x AppOauth#x}
	X *string `field:"optional" json:"x" yaml:"x"`
	// Y coordinate of the elliptic curve point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_oauth#y AppOauth#y}
	Y *string `field:"optional" json:"y" yaml:"y"`
}

