package googlecontainercluster


type GoogleContainerClusterFleet struct {
	// The Fleet host project of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#project GoogleContainerCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

