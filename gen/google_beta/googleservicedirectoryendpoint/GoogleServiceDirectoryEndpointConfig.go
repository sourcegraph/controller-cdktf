package googleservicedirectoryendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleServiceDirectoryEndpointConfig struct {
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
	// The Resource ID must be 1-63 characters long, including digits, lowercase letters or the hyphen character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_endpoint#endpoint_id GoogleServiceDirectoryEndpoint#endpoint_id}
	EndpointId *string `field:"required" json:"endpointId" yaml:"endpointId"`
	// The resource name of the service that this endpoint provides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_endpoint#service GoogleServiceDirectoryEndpoint#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// IPv4 or IPv6 address of the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_endpoint#address GoogleServiceDirectoryEndpoint#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_endpoint#id GoogleServiceDirectoryEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Metadata for the endpoint.
	//
	// This data can be consumed
	// by service clients. The entire metadata dictionary may contain
	// up to 512 characters, spread across all key-value pairs.
	// Metadata that goes beyond any these limits will be rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_endpoint#metadata GoogleServiceDirectoryEndpoint#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// The URL to the network, such as projects/PROJECT_NUMBER/locations/global/networks/NETWORK_NAME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_endpoint#network GoogleServiceDirectoryEndpoint#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Port that the endpoint is running on, must be in the range of [0, 65535].
	//
	// If unspecified, the default is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_endpoint#port GoogleServiceDirectoryEndpoint#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_directory_endpoint#timeouts GoogleServiceDirectoryEndpoint#timeouts}
	Timeouts *GoogleServiceDirectoryEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

