package autologinapp


type AutoLoginAppTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/auto_login_app#create AutoLoginApp#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/auto_login_app#read AutoLoginApp#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/auto_login_app#update AutoLoginApp#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

