package accessidentityprovider


type AccessIdentityProviderScimConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#enabled AccessIdentityProvider#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#group_member_deprovision AccessIdentityProvider#group_member_deprovision}.
	GroupMemberDeprovision interface{} `field:"optional" json:"groupMemberDeprovision" yaml:"groupMemberDeprovision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#seat_deprovision AccessIdentityProvider#seat_deprovision}.
	SeatDeprovision interface{} `field:"optional" json:"seatDeprovision" yaml:"seatDeprovision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#secret AccessIdentityProvider#secret}.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_identity_provider#user_deprovision AccessIdentityProvider#user_deprovision}.
	UserDeprovision interface{} `field:"optional" json:"userDeprovision" yaml:"userDeprovision"`
}

