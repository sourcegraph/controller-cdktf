package googledatafusioninstance


type GoogleDataFusionInstanceNetworkConfig struct {
	// Optional.
	//
	// Type of connection for establishing private IP connectivity between the Data Fusion customer project VPC and
	// the corresponding tenant project from a predefined list of available connection modes.
	// If this field is unspecified for a private instance, VPC peering is used. Possible values: ["VPC_PEERING", "PRIVATE_SERVICE_CONNECT_INTERFACES"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#connection_type GoogleDataFusionInstance#connection_type}
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// The IP range in CIDR notation to use for the managed Data Fusion instance nodes.
	//
	// This range must not overlap with any other ranges used in the Data Fusion instance network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#ip_allocation GoogleDataFusionInstance#ip_allocation}
	IpAllocation *string `field:"optional" json:"ipAllocation" yaml:"ipAllocation"`
	// Name of the network in the project with which the tenant project will be peered for executing pipelines.
	//
	// In case of shared VPC where the network resides in another host
	// project the network should specified in the form of projects/{host-project-id}/global/networks/{network}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#network GoogleDataFusionInstance#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// private_service_connect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#private_service_connect_config GoogleDataFusionInstance#private_service_connect_config}
	PrivateServiceConnectConfig *GoogleDataFusionInstanceNetworkConfigPrivateServiceConnectConfig `field:"optional" json:"privateServiceConnectConfig" yaml:"privateServiceConnectConfig"`
}

