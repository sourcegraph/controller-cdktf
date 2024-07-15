package googlesecretmanagersecret


type GoogleSecretManagerSecretTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secret_manager_secret#create GoogleSecretManagerSecret#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secret_manager_secret#delete GoogleSecretManagerSecret#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secret_manager_secret#update GoogleSecretManagerSecret#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

