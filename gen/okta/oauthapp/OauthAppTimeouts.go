package oauthapp


type OauthAppTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/oauth_app#create OauthApp#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/oauth_app#read OauthApp#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/oauth_app#update OauthApp#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

