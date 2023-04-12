package spectrumapplication


type SpectrumApplicationOriginPortRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#end SpectrumApplication#end}.
	End *float64 `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#start SpectrumApplication#start}.
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

