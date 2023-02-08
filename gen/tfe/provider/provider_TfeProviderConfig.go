package provider


type TfeProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe#alias TfeProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The Terraform Enterprise hostname to connect to. Defaults to app.terraform.io.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe#hostname TfeProvider#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Whether or not to skip certificate verifications.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe#ssl_skip_verify TfeProvider#ssl_skip_verify}
	SslSkipVerify interface{} `field:"optional" json:"sslSkipVerify" yaml:"sslSkipVerify"`
	// The token used to authenticate with Terraform Enterprise.
	//
	// We recommend omitting
	// the token which can be set as credentials in the CLI config file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe#token TfeProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

