package spectrumapplication


type SpectrumApplicationOriginDns struct {
	// Fully qualified domain name of the origin.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#name SpectrumApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

