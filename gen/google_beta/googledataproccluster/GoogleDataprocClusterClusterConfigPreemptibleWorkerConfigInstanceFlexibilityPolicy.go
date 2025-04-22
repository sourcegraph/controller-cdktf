package googledataproccluster


type GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicy struct {
	// instance_selection_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_cluster#instance_selection_list GoogleDataprocCluster#instance_selection_list}
	InstanceSelectionList interface{} `field:"optional" json:"instanceSelectionList" yaml:"instanceSelectionList"`
	// provisioning_model_mix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_cluster#provisioning_model_mix GoogleDataprocCluster#provisioning_model_mix}
	ProvisioningModelMix *GoogleDataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyProvisioningModelMix `field:"optional" json:"provisioningModelMix" yaml:"provisioningModelMix"`
}

