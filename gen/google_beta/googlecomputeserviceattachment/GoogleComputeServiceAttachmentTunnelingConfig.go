package googlecomputeserviceattachment


type GoogleComputeServiceAttachmentTunnelingConfig struct {
	// The encapsulation profile for tunneling traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_service_attachment#encapsulation_profile GoogleComputeServiceAttachment#encapsulation_profile}
	EncapsulationProfile *string `field:"optional" json:"encapsulationProfile" yaml:"encapsulationProfile"`
	// The routing mode for tunneling traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_service_attachment#routing_mode GoogleComputeServiceAttachment#routing_mode}
	RoutingMode *string `field:"optional" json:"routingMode" yaml:"routingMode"`
}

