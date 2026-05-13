package networkservicesauthzextension

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesAuthzExtensionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#location NetworkServicesAuthzExtension#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Identifier. Name of the AuthzExtension resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#name NetworkServicesAuthzExtension#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The service that runs the extension.
	//
	// The following values and formats are accepted:
	// * 'iap.googleapis.com' when the policyProfile is set to REQUEST_AUTHZ
	// * 'modelarmor.{{region}}.rep.googleapis.com' when the policyProfile is set to CONTENT_AUTHZ
	// * A fully qualified domain name that can be resolved by the dataplane
	// * Backend service resource URI of the form 'https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/backendServices/{{name}}' or 'https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendServices/{{name}}}}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#service NetworkServicesAuthzExtension#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Specifies the timeout for each individual message on the stream. The timeout must be between 10-10000 milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#timeout NetworkServicesAuthzExtension#timeout}
	Timeout *string `field:"required" json:"timeout" yaml:"timeout"`
	// The :authority header in the gRPC request sent from Envoy to the extension service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#authority NetworkServicesAuthzExtension#authority}
	Authority *string `field:"optional" json:"authority" yaml:"authority"`
	// A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#description NetworkServicesAuthzExtension#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Determines how the proxy behaves if the call to the extension fails or times out.
	//
	// When set to TRUE, request or response processing continues without error. Any subsequent extensions in the extension chain are also executed. When set to FALSE or the default setting of FALSE is used, one of the following happens:
	// * If response headers have not been delivered to the downstream client, a generic 500 error is returned to the client. The error response can be tailored by configuring a custom error response in the load balancer.
	// * If response headers have been delivered, then the HTTP stream to the downstream client is reset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#fail_open NetworkServicesAuthzExtension#fail_open}
	FailOpen interface{} `field:"optional" json:"failOpen" yaml:"failOpen"`
	// List of the HTTP headers to forward to the extension (from the client).
	//
	// If omitted, all headers are sent. Each element is a string indicating the header name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#forward_headers NetworkServicesAuthzExtension#forward_headers}
	ForwardHeaders *[]*string `field:"optional" json:"forwardHeaders" yaml:"forwardHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#id NetworkServicesAuthzExtension#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of labels associated with the AuthzExtension resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#labels NetworkServicesAuthzExtension#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Required when the service points to a backend service.
	//
	// All backend services and forwarding rules referenced by
	// this extension must share the same load balancing scheme. For more information, refer to
	// [Backend services overview](https://cloud.google.com/load-balancing/docs/backend-service). Possible values: ["INTERNAL_MANAGED", "EXTERNAL_MANAGED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#load_balancing_scheme NetworkServicesAuthzExtension#load_balancing_scheme}
	LoadBalancingScheme *string `field:"optional" json:"loadBalancingScheme" yaml:"loadBalancingScheme"`
	// The metadata provided here is included as part of the metadata_context (of type google.protobuf.Struct) in the ProcessingRequest message sent to the extension server. The metadata is available under the namespace com.google.authz_extension.<resourceName>. The following variables are supported in the metadata Struct:.
	//
	// {forwarding_rule_id} - substituted with the forwarding rule's fully qualified resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#metadata NetworkServicesAuthzExtension#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#project NetworkServicesAuthzExtension#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#timeouts NetworkServicesAuthzExtension#timeouts}
	Timeouts *NetworkServicesAuthzExtensionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The format of communication supported by the callout extension.
	//
	// Applicable only when the policyProfile is REQUEST_AUTHZ.
	// This field is supported only for regional AuthzExtension resources. If not specified, the default value
	// EXT_PROC_GRPC is used. Global AuthzExtension resources use the EXT_PROC_GRPC wire format.
	//
	// Supported values:
	// - WIRE_FORMAT_UNSPECIFIED:
	//     No wire format is explicitly specified. The backend automatically
	//     defaults this value to EXT_PROC_GRPC.
	// - EXT_PROC_GRPC:
	//     Uses Envoy's External Processing (ext_proc) gRPC API over a single
	//     gRPC stream. The backend service must support HTTP/2 or H2C.
	//     All supported events for a client request are sent over the same
	//     gRPC stream. This is the default wire format.
	// - EXT_AUTHZ_GRPC:
	//     Uses Envoy's external authorization (ext_authz) gRPC API.
	//     The backend service must support HTTP/2 or H2C.
	//     This option is only supported for regional AuthzExtension resources. Possible values: ["WIRE_FORMAT_UNSPECIFIED", "EXT_PROC_GRPC", "EXT_AUTHZ_GRPC"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_authz_extension#wire_format NetworkServicesAuthzExtension#wire_format}
	WireFormat *string `field:"optional" json:"wireFormat" yaml:"wireFormat"`
}

