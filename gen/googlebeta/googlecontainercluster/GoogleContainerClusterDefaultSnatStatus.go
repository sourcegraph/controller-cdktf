package googlecontainercluster


type GoogleContainerClusterDefaultSnatStatus struct {
	// When disabled is set to false, default IP masquerade rules will be applied to the nodes to prevent sNAT on cluster internal traffic.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#disabled GoogleContainerCluster#disabled}
	Disabled interface{} `field:"required" json:"disabled" yaml:"disabled"`
}

