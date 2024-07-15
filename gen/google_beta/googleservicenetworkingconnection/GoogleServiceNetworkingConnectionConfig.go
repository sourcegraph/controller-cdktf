package googleservicenetworkingconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleServiceNetworkingConnectionConfig struct {
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
	// Name of VPC network connected with service producers using VPC peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_networking_connection#network GoogleServiceNetworkingConnection#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Named IP address range(s) of PEERING type reserved for this service provider.
	//
	// Note that invoking this method with a different range when connection is already established will not reallocate already provisioned service producer subnetworks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_networking_connection#reserved_peering_ranges GoogleServiceNetworkingConnection#reserved_peering_ranges}
	ReservedPeeringRanges *[]*string `field:"required" json:"reservedPeeringRanges" yaml:"reservedPeeringRanges"`
	// Provider peering service that is managing peering connectivity for a service provider organization.
	//
	// For Google services that support this functionality it is 'servicenetworking.googleapis.com'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_networking_connection#service GoogleServiceNetworkingConnection#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// When set to ABANDON, terraform will abandon management of the resource instead of deleting it.
	//
	// Prevents terraform apply failures with CloudSQL. Note: The resource will still exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_networking_connection#deletion_policy GoogleServiceNetworkingConnection#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_networking_connection#id GoogleServiceNetworkingConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_networking_connection#timeouts GoogleServiceNetworkingConnection#timeouts}
	Timeouts *GoogleServiceNetworkingConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

