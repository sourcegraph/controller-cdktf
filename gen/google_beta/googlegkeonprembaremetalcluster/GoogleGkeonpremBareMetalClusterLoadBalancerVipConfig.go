package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterLoadBalancerVipConfig struct {
	// The VIP which you previously set aside for the Kubernetes API of this Bare Metal User Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#control_plane_vip GoogleGkeonpremBareMetalCluster#control_plane_vip}
	ControlPlaneVip *string `field:"required" json:"controlPlaneVip" yaml:"controlPlaneVip"`
	// The VIP which you previously set aside for ingress traffic into this Bare Metal User Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#ingress_vip GoogleGkeonpremBareMetalCluster#ingress_vip}
	IngressVip *string `field:"required" json:"ingressVip" yaml:"ingressVip"`
}

