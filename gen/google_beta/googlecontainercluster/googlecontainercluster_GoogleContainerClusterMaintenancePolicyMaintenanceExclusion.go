package googlecontainercluster


type GoogleContainerClusterMaintenancePolicyMaintenanceExclusion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#end_time GoogleContainerCluster#end_time}.
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#exclusion_name GoogleContainerCluster#exclusion_name}.
	ExclusionName *string `field:"required" json:"exclusionName" yaml:"exclusionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#start_time GoogleContainerCluster#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// exclusion_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#exclusion_options GoogleContainerCluster#exclusion_options}
	ExclusionOptions *GoogleContainerClusterMaintenancePolicyMaintenanceExclusionExclusionOptions `field:"optional" json:"exclusionOptions" yaml:"exclusionOptions"`
}

