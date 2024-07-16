package googlesecretmanagersecret


type GoogleSecretManagerSecretReplication struct {
	// auto block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secret_manager_secret#auto GoogleSecretManagerSecret#auto}
	Auto *GoogleSecretManagerSecretReplicationAuto `field:"optional" json:"auto" yaml:"auto"`
	// user_managed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secret_manager_secret#user_managed GoogleSecretManagerSecret#user_managed}
	UserManaged *GoogleSecretManagerSecretReplicationUserManaged `field:"optional" json:"userManaged" yaml:"userManaged"`
}

