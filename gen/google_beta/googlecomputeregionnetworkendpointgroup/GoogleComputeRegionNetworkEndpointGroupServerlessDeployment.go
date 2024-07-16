package googlecomputeregionnetworkendpointgroup


type GoogleComputeRegionNetworkEndpointGroupServerlessDeployment struct {
	// The platform of the NEG backend target(s). Possible values: API Gateway: apigateway.googleapis.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_network_endpoint_group#platform GoogleComputeRegionNetworkEndpointGroup#platform}
	Platform *string `field:"required" json:"platform" yaml:"platform"`
	// The user-defined name of the workload/instance.
	//
	// This value must be provided explicitly or in the urlMask.
	// The resource identified by this value is platform-specific and is as follows: API Gateway: The gateway ID, App Engine: The service name,
	// Cloud Functions: The function name, Cloud Run: The service name
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_network_endpoint_group#resource GoogleComputeRegionNetworkEndpointGroup#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
	// A template to parse platform-specific fields from a request URL.
	//
	// URL mask allows for routing to multiple resources
	// on the same serverless platform without having to create multiple Network Endpoint Groups and backend resources.
	// The fields parsed by this template are platform-specific and are as follows: API Gateway: The gateway ID,
	// App Engine: The service and version, Cloud Functions: The function name, Cloud Run: The service and tag
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_network_endpoint_group#url_mask GoogleComputeRegionNetworkEndpointGroup#url_mask}
	UrlMask *string `field:"optional" json:"urlMask" yaml:"urlMask"`
	// The optional resource version.
	//
	// The version identified by this value is platform-specific and is follows:
	// API Gateway: Unused, App Engine: The service version, Cloud Functions: Unused, Cloud Run: The service tag
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_network_endpoint_group#version GoogleComputeRegionNetworkEndpointGroup#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

