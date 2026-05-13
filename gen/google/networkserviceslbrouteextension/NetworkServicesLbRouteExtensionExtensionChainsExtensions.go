package networkserviceslbrouteextension


type NetworkServicesLbRouteExtensionExtensionChainsExtensions struct {
	// The name for this extension.
	//
	// The name is logged as part of the HTTP request logs.
	// The name must conform with RFC-1034, is restricted to lower-cased letters, numbers and hyphens,
	// and can have a maximum length of 63 characters. Additionally, the first character must be a letter
	// and the last a letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_lb_route_extension#name NetworkServicesLbRouteExtension#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The reference to the service that runs the extension.
	//
	// * To configure a callout extension, service must be a fully-qualified reference to a backend service.
	// * To configure a plugin extension, service must be a reference to a WasmPlugin resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_lb_route_extension#service NetworkServicesLbRouteExtension#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// The :authority header in the gRPC request sent from Envoy to the extension service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_lb_route_extension#authority NetworkServicesLbRouteExtension#authority}
	Authority *string `field:"optional" json:"authority" yaml:"authority"`
	// Determines how the proxy behaves if the call to the extension fails or times out.
	//
	// When set to TRUE, request or response processing continues without error.
	// Any subsequent extensions in the extension chain are also executed.
	// When set to FALSE: * If response headers have not been delivered to the downstream client,
	// a generic 500 error is returned to the client. The error response can be tailored by
	// configuring a custom error response in the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_lb_route_extension#fail_open NetworkServicesLbRouteExtension#fail_open}
	FailOpen interface{} `field:"optional" json:"failOpen" yaml:"failOpen"`
	// List of the HTTP headers to forward to the extension (from the client or backend).
	//
	// If omitted, all headers are sent. Each element is a string indicating the header name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_lb_route_extension#forward_headers NetworkServicesLbRouteExtension#forward_headers}
	ForwardHeaders *[]*string `field:"optional" json:"forwardHeaders" yaml:"forwardHeaders"`
	// The metadata provided here is included as part of the 'metadata_context' (of type 'google.protobuf.Struct') in the 'ProcessingRequest' message sent to the extension server. The metadata is available under the namespace 'com.google.lb_route_extension.<resource_name>.<chain_name>.<extension_name>'. The following variables are supported in the metadata: '{forwarding_rule_id}' - substituted with the forwarding rule's fully qualified resource name. This field must not be set for plugin extensions. Setting it results in a validation error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_lb_route_extension#metadata NetworkServicesLbRouteExtension#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// When set to 'TRUE', enables 'observability_mode' on the 'ext_proc' filter.
	//
	// This makes 'ext_proc' calls asynchronous. Envoy doesn't check for the response from 'ext_proc' calls.
	// For more information about the filter, see: https://www.envoyproxy.io/docs/envoy/v1.32.3/api-v3/extensions/filters/http/ext_proc/v3/ext_proc.proto
	// This field is helpful when you want to try out the extension in async log-only mode.
	// Supported by regional 'LbTrafficExtension' and 'LbRouteExtension' resources.
	// Only 'STREAMED' (default) body processing mode is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_lb_route_extension#observability_mode NetworkServicesLbRouteExtension#observability_mode}
	ObservabilityMode interface{} `field:"optional" json:"observabilityMode" yaml:"observabilityMode"`
	// Configures the send mode for request body processing.
	//
	// The field can only be set if 'supported_events' includes 'REQUEST_BODY'.
	// If 'supported_events' includes 'REQUEST_BODY', but 'request_body_send_mode' is unset, the default value 'STREAMED' is used.
	// When this field is set to 'FULL_DUPLEX_STREAMED', 'supported_events' must include both 'REQUEST_BODY' and 'REQUEST_TRAILERS'.
	// This field can be set only when the 'service' field of the extension points to a 'BackendService'.
	// Only 'FULL_DUPLEX_STREAMED' mode is supported for 'LbRouteExtension' resources. Possible values: ["BODY_SEND_MODE_UNSPECIFIED", "BODY_SEND_MODE_STREAMED", "BODY_SEND_MODE_FULL_DUPLEX_STREAMED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_lb_route_extension#request_body_send_mode NetworkServicesLbRouteExtension#request_body_send_mode}
	RequestBodySendMode *string `field:"optional" json:"requestBodySendMode" yaml:"requestBodySendMode"`
	// A set of events during request or response processing for which this extension is called.
	//
	// This field is optional for the LbRouteExtension resource. If unspecified, 'REQUEST_HEADERS' event is assumed as supported.
	// Possible values: 'REQUEST_HEADERS', 'REQUEST_BODY', 'REQUEST_TRAILERS'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_lb_route_extension#supported_events NetworkServicesLbRouteExtension#supported_events}
	SupportedEvents *[]*string `field:"optional" json:"supportedEvents" yaml:"supportedEvents"`
	// Specifies the timeout for each individual message on the stream.
	//
	// The timeout must be between 10-1000 milliseconds.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_lb_route_extension#timeout NetworkServicesLbRouteExtension#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

