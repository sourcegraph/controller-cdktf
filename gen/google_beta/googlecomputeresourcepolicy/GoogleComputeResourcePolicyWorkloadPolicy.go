package googlecomputeresourcepolicy


type GoogleComputeResourcePolicyWorkloadPolicy struct {
	// The type of workload policy. Possible values: ["HIGH_AVAILABILITY", "HIGH_THROUGHPUT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_resource_policy#type GoogleComputeResourcePolicy#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The accelerator topology.
	//
	// This field can be set only when the workload policy type is HIGH_THROUGHPUT
	// and cannot be set if max topology distance is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_resource_policy#accelerator_topology GoogleComputeResourcePolicy#accelerator_topology}
	AcceleratorTopology *string `field:"optional" json:"acceleratorTopology" yaml:"acceleratorTopology"`
	// The maximum topology distance.
	//
	// This field can be set only when the workload policy type is HIGH_THROUGHPUT
	// and cannot be set if accelerator topology is set. Possible values: ["BLOCK", "CLUSTER", "SUBBLOCK"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_resource_policy#max_topology_distance GoogleComputeResourcePolicy#max_topology_distance}
	MaxTopologyDistance *string `field:"optional" json:"maxTopologyDistance" yaml:"maxTopologyDistance"`
}

