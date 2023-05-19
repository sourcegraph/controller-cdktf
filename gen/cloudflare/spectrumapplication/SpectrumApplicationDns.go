package spectrumapplication


type SpectrumApplicationDns struct {
	// The name of the DNS record associated with the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#name SpectrumApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of DNS record associated with the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#type SpectrumApplication#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

