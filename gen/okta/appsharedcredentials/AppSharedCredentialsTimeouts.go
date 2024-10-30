package appsharedcredentials


type AppSharedCredentialsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_shared_credentials#create AppSharedCredentials#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_shared_credentials#read AppSharedCredentials#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_shared_credentials#update AppSharedCredentials#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

