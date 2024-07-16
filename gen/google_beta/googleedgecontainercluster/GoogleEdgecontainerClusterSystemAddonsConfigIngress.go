package googleedgecontainercluster


type GoogleEdgecontainerClusterSystemAddonsConfigIngress struct {
	// Whether Ingress is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_cluster#disabled GoogleEdgecontainerCluster#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Ingress VIP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_cluster#ipv4_vip GoogleEdgecontainerCluster#ipv4_vip}
	Ipv4Vip *string `field:"optional" json:"ipv4Vip" yaml:"ipv4Vip"`
}

