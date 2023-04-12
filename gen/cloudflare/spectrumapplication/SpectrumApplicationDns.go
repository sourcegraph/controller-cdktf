package spectrumapplication


type SpectrumApplicationDns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#name SpectrumApplication#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#type SpectrumApplication#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

