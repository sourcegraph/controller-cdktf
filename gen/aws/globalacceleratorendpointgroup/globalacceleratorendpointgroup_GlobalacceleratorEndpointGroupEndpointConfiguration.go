package globalacceleratorendpointgroup


type GlobalacceleratorEndpointGroupEndpointConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group#client_ip_preservation_enabled GlobalacceleratorEndpointGroup#client_ip_preservation_enabled}.
	ClientIpPreservationEnabled interface{} `field:"optional" json:"clientIpPreservationEnabled" yaml:"clientIpPreservationEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group#endpoint_id GlobalacceleratorEndpointGroup#endpoint_id}.
	EndpointId *string `field:"optional" json:"endpointId" yaml:"endpointId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/globalaccelerator_endpoint_group#weight GlobalacceleratorEndpointGroup#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

