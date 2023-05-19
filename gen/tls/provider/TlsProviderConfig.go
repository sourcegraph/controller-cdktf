package provider


type TlsProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs#alias TlsProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs#proxy TlsProvider#proxy}
	Proxy *TlsProviderProxy `field:"optional" json:"proxy" yaml:"proxy"`
}

