package containercluster


type ContainerClusterMaintenancePolicyMaintenanceExclusionExclusionOptions struct {
	// The scope of automatic upgrades to restrict in the exclusion window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#scope ContainerCluster#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// The behavior of the exclusion end time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#end_time_behavior ContainerCluster#end_time_behavior}
	EndTimeBehavior *string `field:"optional" json:"endTimeBehavior" yaml:"endTimeBehavior"`
}

