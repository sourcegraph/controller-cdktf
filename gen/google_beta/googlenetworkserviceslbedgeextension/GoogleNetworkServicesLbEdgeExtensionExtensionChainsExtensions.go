package googlenetworkserviceslbedgeextension


type GoogleNetworkServicesLbEdgeExtensionExtensionChainsExtensions struct {
	// The name for this extension.
	//
	// The name is logged as part of the HTTP request logs.
	// The name must conform with RFC-1034, is restricted to lower-cased letters, numbers and hyphens,
	// and can have a maximum length of 63 characters. Additionally, the first character must be a letter
	// and the last a letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#name GoogleNetworkServicesLbEdgeExtension#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The reference to the service that runs the extension.
	//
	// * To configure a callout extension, service must be a fully-qualified reference to a backend service.
	// * To configure a plugin extension, service must be a reference to a WasmPlugin resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#service GoogleNetworkServicesLbEdgeExtension#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Determines how the proxy behaves if the call to the extension fails or times out.
	//
	// When set to TRUE, request or response processing continues without error.
	// Any subsequent extensions in the extension chain are also executed.
	// When set to FALSE: * If response headers have not been delivered to the downstream client,
	// a generic 500 error is returned to the client. The error response can be tailored by
	// configuring a custom error response in the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#fail_open GoogleNetworkServicesLbEdgeExtension#fail_open}
	FailOpen interface{} `field:"optional" json:"failOpen" yaml:"failOpen"`
	// List of the HTTP headers to forward to the extension (from the client or backend).
	//
	// If omitted, all headers are sent. Each element is a string indicating the header name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#forward_headers GoogleNetworkServicesLbEdgeExtension#forward_headers}
	ForwardHeaders *[]*string `field:"optional" json:"forwardHeaders" yaml:"forwardHeaders"`
	// A set of events during request or response processing for which this extension is called.
	//
	// This field is required for the LbEdgeExtension resource and only supports the value 'REQUEST_HEADERS'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#supported_events GoogleNetworkServicesLbEdgeExtension#supported_events}
	SupportedEvents *[]*string `field:"optional" json:"supportedEvents" yaml:"supportedEvents"`
}

