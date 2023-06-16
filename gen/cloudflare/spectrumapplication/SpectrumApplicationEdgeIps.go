package spectrumapplication


type SpectrumApplicationEdgeIps struct {
	// The type of edge IP configuration specified. Available values: `dynamic`, `static`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#type SpectrumApplication#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	//
	// Required when `type` is not `static`. Available values: `all`, `ipv4`, `ipv6`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#connectivity SpectrumApplication#connectivity}
	Connectivity *string `field:"optional" json:"connectivity" yaml:"connectivity"`
	// The collection of customer owned IPs to broadcast via anycast for this hostname and application.
	//
	// Requires [Bring Your Own IP](https://developers.cloudflare.com/spectrum/getting-started/byoip/) provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#ips SpectrumApplication#ips}
	Ips *[]*string `field:"optional" json:"ips" yaml:"ips"`
}

