package googleedgecontainercluster


type GoogleEdgecontainerClusterMaintenancePolicyWindowRecurringWindow struct {
	// An RRULE (https://tools.ietf.org/html/rfc5545#section-3.8.5.3) for how this window recurs. They go on for the span of time between the start and end time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_cluster#recurrence GoogleEdgecontainerCluster#recurrence}
	Recurrence *string `field:"optional" json:"recurrence" yaml:"recurrence"`
	// window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_cluster#window GoogleEdgecontainerCluster#window}
	Window *GoogleEdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindow `field:"optional" json:"window" yaml:"window"`
}

