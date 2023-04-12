package appswa


type AppSwaTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_swa#create AppSwa#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_swa#read AppSwa#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_swa#update AppSwa#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

