package appbasicauth


type AppBasicAuthTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_basic_auth#create AppBasicAuth#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_basic_auth#read AppBasicAuth#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_basic_auth#update AppBasicAuth#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

