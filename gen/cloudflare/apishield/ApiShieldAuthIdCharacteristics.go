package apishield


type ApiShieldAuthIdCharacteristics struct {
	// The name of the characteristic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/api_shield#name ApiShield#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of characteristic. Available values: `header`, `cookie`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/api_shield#type ApiShield#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

