package googlemanagedkafkacluster


type GoogleManagedKafkaClusterCapacityConfig struct {
	// The memory to provision for the cluster in bytes.
	//
	// The value must be between 1 GiB and 8 GiB per vCPU. Ex. 1024Mi, 4Gi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#memory_bytes GoogleManagedKafkaCluster#memory_bytes}
	MemoryBytes *string `field:"required" json:"memoryBytes" yaml:"memoryBytes"`
	// The number of vCPUs to provision for the cluster. The minimum is 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#vcpu_count GoogleManagedKafkaCluster#vcpu_count}
	VcpuCount *string `field:"required" json:"vcpuCount" yaml:"vcpuCount"`
}

