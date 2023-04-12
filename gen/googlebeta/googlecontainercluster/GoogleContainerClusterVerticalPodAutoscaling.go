package googlecontainercluster


type GoogleContainerClusterVerticalPodAutoscaling struct {
	// Enables vertical pod autoscaling.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

