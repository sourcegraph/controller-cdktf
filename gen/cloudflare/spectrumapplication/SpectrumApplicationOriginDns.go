package spectrumapplication


type SpectrumApplicationOriginDns struct {
	// Fully qualified domain name of the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#name SpectrumApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

