package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterNodeConfig struct {
	// The available runtimes that can be used to run containers in a Bare Metal User Cluster.
	//
	// Possible values: ["CONTAINER_RUNTIME_UNSPECIFIED", "DOCKER", "CONTAINERD"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#container_runtime GoogleGkeonpremBareMetalCluster#container_runtime}
	ContainerRuntime *string `field:"optional" json:"containerRuntime" yaml:"containerRuntime"`
	// The maximum number of pods a node can run.
	//
	// The size of the CIDR range
	// assigned to the node will be derived from this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#max_pods_per_node GoogleGkeonpremBareMetalCluster#max_pods_per_node}
	MaxPodsPerNode *float64 `field:"optional" json:"maxPodsPerNode" yaml:"maxPodsPerNode"`
}

