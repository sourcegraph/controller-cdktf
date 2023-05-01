package provider


type TlsProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls#alias TlsProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

