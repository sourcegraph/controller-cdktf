package googlecontainercluster


type GoogleContainerClusterNodeConfigGvnic struct {
	// Whether or not gvnic is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

