package googlegkeonpremvmwareadmincluster


type GoogleGkeonpremVmwareAdminClusterNetworkConfig struct {
	// All pods in the cluster are assigned an RFC1918 IPv4 address from these ranges.
	//
	// Only a single range is supported. This field cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#pod_address_cidr_blocks GoogleGkeonpremVmwareAdminCluster#pod_address_cidr_blocks}
	PodAddressCidrBlocks *[]*string `field:"required" json:"podAddressCidrBlocks" yaml:"podAddressCidrBlocks"`
	// All services in the cluster are assigned an RFC1918 IPv4 address from these ranges.
	//
	// Only a single range is supported.. This field
	// cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#service_address_cidr_blocks GoogleGkeonpremVmwareAdminCluster#service_address_cidr_blocks}
	ServiceAddressCidrBlocks *[]*string `field:"required" json:"serviceAddressCidrBlocks" yaml:"serviceAddressCidrBlocks"`
	// dhcp_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#dhcp_ip_config GoogleGkeonpremVmwareAdminCluster#dhcp_ip_config}
	DhcpIpConfig *GoogleGkeonpremVmwareAdminClusterNetworkConfigDhcpIpConfig `field:"optional" json:"dhcpIpConfig" yaml:"dhcpIpConfig"`
	// ha_control_plane_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#ha_control_plane_config GoogleGkeonpremVmwareAdminCluster#ha_control_plane_config}
	HaControlPlaneConfig *GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfig `field:"optional" json:"haControlPlaneConfig" yaml:"haControlPlaneConfig"`
	// host_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#host_config GoogleGkeonpremVmwareAdminCluster#host_config}
	HostConfig *GoogleGkeonpremVmwareAdminClusterNetworkConfigHostConfig `field:"optional" json:"hostConfig" yaml:"hostConfig"`
	// static_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#static_ip_config GoogleGkeonpremVmwareAdminCluster#static_ip_config}
	StaticIpConfig *GoogleGkeonpremVmwareAdminClusterNetworkConfigStaticIpConfig `field:"optional" json:"staticIpConfig" yaml:"staticIpConfig"`
	// vcenter_network specifies vCenter network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gkeonprem_vmware_admin_cluster#vcenter_network GoogleGkeonpremVmwareAdminCluster#vcenter_network}
	VcenterNetwork *string `field:"optional" json:"vcenterNetwork" yaml:"vcenterNetwork"`
}

