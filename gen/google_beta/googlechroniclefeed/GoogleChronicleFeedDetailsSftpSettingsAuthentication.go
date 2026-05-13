package googlechroniclefeed


type GoogleChronicleFeedDetailsSftpSettingsAuthentication struct {
	// Password. Used for username and password authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#password GoogleChronicleFeed#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Private key. Used for private key authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#private_key GoogleChronicleFeed#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Private key passphrase. Used for private key authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#private_key_passphrase GoogleChronicleFeed#private_key_passphrase}
	PrivateKeyPassphrase *string `field:"optional" json:"privateKeyPassphrase" yaml:"privateKeyPassphrase"`
	// Username. Used for username and password authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#username GoogleChronicleFeed#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

