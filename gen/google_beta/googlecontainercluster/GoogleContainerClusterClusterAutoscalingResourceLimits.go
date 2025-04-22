package googlecontainercluster


type GoogleContainerClusterClusterAutoscalingResourceLimits struct {
	// Maximum amount of the resource in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_container_cluster#maximum GoogleContainerCluster#maximum}
	Maximum *float64 `field:"required" json:"maximum" yaml:"maximum"`
	// The type of the resource.
	//
	// For example, cpu and memory. See the guide to using Node Auto-Provisioning for a list of types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_container_cluster#resource_type GoogleContainerCluster#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Minimum amount of the resource in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_container_cluster#minimum GoogleContainerCluster#minimum}
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
}

