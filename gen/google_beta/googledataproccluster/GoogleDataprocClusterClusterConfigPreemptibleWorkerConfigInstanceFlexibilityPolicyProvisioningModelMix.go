package googledataproccluster


type GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyProvisioningModelMix struct {
	// The base capacity that will always use Standard VMs to avoid risk of more preemption than the minimum capacity you need.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_cluster#standard_capacity_base GoogleDataprocCluster#standard_capacity_base}
	StandardCapacityBase *float64 `field:"optional" json:"standardCapacityBase" yaml:"standardCapacityBase"`
	// The percentage of target capacity that should use Standard VM. The remaining percentage will use Spot VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_cluster#standard_capacity_percent_above_base GoogleDataprocCluster#standard_capacity_percent_above_base}
	StandardCapacityPercentAboveBase *float64 `field:"optional" json:"standardCapacityPercentAboveBase" yaml:"standardCapacityPercentAboveBase"`
}

