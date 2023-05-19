package googlenetworkservicesgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkServicesGatewayConfig struct {
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
	// Short name of the Gateway resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_network_services_gateway#name GoogleNetworkServicesGateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// One or more port numbers (1-65535), on which the Gateway will receive traffic.
	//
	// The proxy binds to the specified ports. Gateways of type 'SECURE_WEB_GATEWAY' are
	// limited to 1 port. Gateways of type 'OPEN_MESH' listen on 0.0.0.0 and support multiple ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_network_services_gateway#ports GoogleNetworkServicesGateway#ports}
	Ports *[]*float64 `field:"required" json:"ports" yaml:"ports"`
	// Immutable.
	//
	// Scope determines how configuration across multiple Gateway instances are merged.
	// The configuration for multiple Gateway instances with the same scope will be merged as presented as
	// a single coniguration to the proxy/load balancer.
	// Max length 64 characters. Scope should start with a letter and can only have letters, numbers, hyphens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_network_services_gateway#scope GoogleNetworkServicesGateway#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Immutable. The type of the customer-managed gateway. Possible values are: * OPEN_MESH * SECURE_WEB_GATEWAY. Possible values: ["TYPE_UNSPECIFIED", "OPEN_MESH", "SECURE_WEB_GATEWAY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_network_services_gateway#type GoogleNetworkServicesGateway#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_network_services_gateway#description GoogleNetworkServicesGateway#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_network_services_gateway#id GoogleNetworkServicesGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the Gateway resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_network_services_gateway#labels GoogleNetworkServicesGateway#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location of the gateway. The default value is 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_network_services_gateway#location GoogleNetworkServicesGateway#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_network_services_gateway#project GoogleNetworkServicesGateway#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A fully-qualified ServerTLSPolicy URL reference. Specifies how TLS traffic is terminated. If empty, TLS termination is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_network_services_gateway#server_tls_policy GoogleNetworkServicesGateway#server_tls_policy}
	ServerTlsPolicy *string `field:"optional" json:"serverTlsPolicy" yaml:"serverTlsPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_network_services_gateway#timeouts GoogleNetworkServicesGateway#timeouts}
	Timeouts *GoogleNetworkServicesGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

