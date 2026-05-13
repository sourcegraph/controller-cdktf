package containercluster


type ContainerClusterMaintenancePolicyMaintenanceExclusion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#exclusion_name ContainerCluster#exclusion_name}.
	ExclusionName *string `field:"required" json:"exclusionName" yaml:"exclusionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#start_time ContainerCluster#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#end_time ContainerCluster#end_time}.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// exclusion_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#exclusion_options ContainerCluster#exclusion_options}
	ExclusionOptions *ContainerClusterMaintenancePolicyMaintenanceExclusionExclusionOptions `field:"optional" json:"exclusionOptions" yaml:"exclusionOptions"`
}

