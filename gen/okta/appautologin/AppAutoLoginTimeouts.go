package appautologin


type AppAutoLoginTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_auto_login#create AppAutoLogin#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_auto_login#read AppAutoLogin#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_auto_login#update AppAutoLogin#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

