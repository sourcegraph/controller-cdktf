package googlecontainerattachedcluster


type GoogleContainerAttachedClusterLoggingConfigComponentConfig struct {
	// The components to be enabled. Possible values: ["SYSTEM_COMPONENTS", "WORKLOADS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#enable_components GoogleContainerAttachedCluster#enable_components}
	EnableComponents *[]*string `field:"optional" json:"enableComponents" yaml:"enableComponents"`
}

