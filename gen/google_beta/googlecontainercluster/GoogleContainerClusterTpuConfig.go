package googlecontainercluster


type GoogleContainerClusterTpuConfig struct {
	// Whether Cloud TPU integration is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Whether to use service networking for Cloud TPU or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#use_service_networking GoogleContainerCluster#use_service_networking}
	UseServiceNetworking interface{} `field:"optional" json:"useServiceNetworking" yaml:"useServiceNetworking"`
}

