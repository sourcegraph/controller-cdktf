package googlecomputefuturereservation


type GoogleComputeFutureReservationAggregateReservation struct {
	// reserved_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#reserved_resources GoogleComputeFutureReservation#reserved_resources}
	ReservedResources interface{} `field:"required" json:"reservedResources" yaml:"reservedResources"`
	// The VM family that all instances scheduled against this reservation must belong to.
	//
	// Possible values: ["VM_FAMILY_CLOUD_TPU_DEVICE_CT3", "VM_FAMILY_CLOUD_TPU_LITE_DEVICE_CT5L", "VM_FAMILY_CLOUD_TPU_LITE_POD_SLICE_CT5LP", "VM_FAMILY_CLOUD_TPU_LITE_POD_SLICE_CT6E", "VM_FAMILY_CLOUD_TPU_POD_SLICE_CT3P", "VM_FAMILY_CLOUD_TPU_POD_SLICE_CT4P", "VM_FAMILY_CLOUD_TPU_POD_SLICE_CT5P"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#vm_family GoogleComputeFutureReservation#vm_family}
	VmFamily *string `field:"optional" json:"vmFamily" yaml:"vmFamily"`
	// The workload type of the instances that will target this reservation. Possible values: ["BATCH", "SERVING", "UNSPECIFIED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_future_reservation#workload_type GoogleComputeFutureReservation#workload_type}
	WorkloadType *string `field:"optional" json:"workloadType" yaml:"workloadType"`
}

