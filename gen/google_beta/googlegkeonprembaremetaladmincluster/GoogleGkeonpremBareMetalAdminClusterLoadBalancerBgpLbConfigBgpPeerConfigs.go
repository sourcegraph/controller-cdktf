package googlegkeonprembaremetaladmincluster


type GoogleGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigBgpPeerConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#asn GoogleGkeonpremBareMetalAdminCluster#asn}.
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// The IP address of the control plane node that connects to the external peer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#control_plane_nodes GoogleGkeonpremBareMetalAdminCluster#control_plane_nodes}
	ControlPlaneNodes *[]*string `field:"optional" json:"controlPlaneNodes" yaml:"controlPlaneNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#ip_address GoogleGkeonpremBareMetalAdminCluster#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}

