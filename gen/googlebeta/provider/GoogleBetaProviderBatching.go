package provider


type GoogleBetaProviderBatching struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta#enable_batching GoogleBetaProvider#enable_batching}.
	EnableBatching interface{} `field:"optional" json:"enableBatching" yaml:"enableBatching"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta#send_after GoogleBetaProvider#send_after}.
	SendAfter *string `field:"optional" json:"sendAfter" yaml:"sendAfter"`
}

