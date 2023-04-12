package appsecurepasswordstore


type AppSecurePasswordStoreTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_secure_password_store#create AppSecurePasswordStore#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_secure_password_store#read AppSecurePasswordStore#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_secure_password_store#update AppSecurePasswordStore#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

