package googlecontainercluster


type GoogleContainerClusterAddonsConfigCloudrunConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#disabled GoogleContainerCluster#disabled}.
	Disabled interface{} `field:"required" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#load_balancer_type GoogleContainerCluster#load_balancer_type}.
	LoadBalancerType *string `field:"optional" json:"loadBalancerType" yaml:"loadBalancerType"`
}

