package networksecuritymirroringendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityMirroringEndpointConfig struct {
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
	// The cloud location of the endpoint, e.g. 'us-central1-a' or 'asia-south1-b'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_mirroring_endpoint#location NetworkSecurityMirroringEndpoint#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The endpoint group that this endpoint belongs to. Format is: 'projects/{project}/locations/{location}/mirroringEndpointGroups/{mirroringEndpointGroup}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_mirroring_endpoint#mirroring_endpoint_group NetworkSecurityMirroringEndpoint#mirroring_endpoint_group}
	MirroringEndpointGroup *string `field:"required" json:"mirroringEndpointGroup" yaml:"mirroringEndpointGroup"`
	// The ID to use for the new endpoint, which will become the final component of the endpoint's resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_mirroring_endpoint#mirroring_endpoint_id NetworkSecurityMirroringEndpoint#mirroring_endpoint_id}
	MirroringEndpointId *string `field:"required" json:"mirroringEndpointId" yaml:"mirroringEndpointId"`
	// User-provided description of the endpoint. Used as additional context for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_mirroring_endpoint#description NetworkSecurityMirroringEndpoint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_mirroring_endpoint#id NetworkSecurityMirroringEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels are key/value pairs that help to organize and filter resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_mirroring_endpoint#labels NetworkSecurityMirroringEndpoint#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_mirroring_endpoint#project NetworkSecurityMirroringEndpoint#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_mirroring_endpoint#timeouts NetworkSecurityMirroringEndpoint#timeouts}
	Timeouts *NetworkSecurityMirroringEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

