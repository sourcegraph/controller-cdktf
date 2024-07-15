package googlegkeonpremvmwarenodepool


type GoogleGkeonpremVmwareNodePoolNodePoolAutoscaling struct {
	// Maximum number of replicas in the NodePool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_node_pool#max_replicas GoogleGkeonpremVmwareNodePool#max_replicas}
	MaxReplicas *float64 `field:"required" json:"maxReplicas" yaml:"maxReplicas"`
	// Minimum number of replicas in the NodePool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_vmware_node_pool#min_replicas GoogleGkeonpremVmwareNodePool#min_replicas}
	MinReplicas *float64 `field:"required" json:"minReplicas" yaml:"minReplicas"`
}

