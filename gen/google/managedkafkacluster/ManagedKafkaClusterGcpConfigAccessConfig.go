package managedkafkacluster


type ManagedKafkaClusterGcpConfigAccessConfig struct {
	// network_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/managed_kafka_cluster#network_configs ManagedKafkaCluster#network_configs}
	NetworkConfigs interface{} `field:"required" json:"networkConfigs" yaml:"networkConfigs"`
}

