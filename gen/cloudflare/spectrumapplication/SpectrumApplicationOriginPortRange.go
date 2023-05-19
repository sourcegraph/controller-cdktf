package spectrumapplication


type SpectrumApplicationOriginPortRange struct {
	// Upper bound of the origin port range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#end SpectrumApplication#end}
	End *float64 `field:"required" json:"end" yaml:"end"`
	// Lower bound of the origin port range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#start SpectrumApplication#start}
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

