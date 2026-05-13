package managedkafkaconnectcluster


type ManagedKafkaConnectClusterCapacityConfig struct {
	// The memory to provision for the cluster in bytes.
	//
	// The CPU:memory ratio (vCPU:GiB) must be between 1:1 and 1:8. Minimum: 3221225472 (3 GiB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/managed_kafka_connect_cluster#memory_bytes ManagedKafkaConnectCluster#memory_bytes}
	MemoryBytes *string `field:"required" json:"memoryBytes" yaml:"memoryBytes"`
	// The number of vCPUs to provision for the cluster. The minimum is 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/managed_kafka_connect_cluster#vcpu_count ManagedKafkaConnectCluster#vcpu_count}
	VcpuCount *string `field:"required" json:"vcpuCount" yaml:"vcpuCount"`
}

