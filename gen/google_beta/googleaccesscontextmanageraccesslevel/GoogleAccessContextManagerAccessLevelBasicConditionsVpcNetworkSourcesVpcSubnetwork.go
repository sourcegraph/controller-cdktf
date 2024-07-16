package googleaccesscontextmanageraccesslevel


type GoogleAccessContextManagerAccessLevelBasicConditionsVpcNetworkSourcesVpcSubnetwork struct {
	// Required.
	//
	// Network name to be allowed by this Access Level. Networks of foreign organizations requires 'compute.network.get' permission to be granted to caller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#network GoogleAccessContextManagerAccessLevel#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// CIDR block IP subnetwork specification. Must be IPv4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#vpc_ip_subnetworks GoogleAccessContextManagerAccessLevel#vpc_ip_subnetworks}
	VpcIpSubnetworks *[]*string `field:"optional" json:"vpcIpSubnetworks" yaml:"vpcIpSubnetworks"`
}

