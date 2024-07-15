package googleedgecontainercluster


type GoogleEdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindow struct {
	// The time that the window ends. The end time must take place after the start time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_cluster#end_time GoogleEdgecontainerCluster#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The time that the window first starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_cluster#start_time GoogleEdgecontainerCluster#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

