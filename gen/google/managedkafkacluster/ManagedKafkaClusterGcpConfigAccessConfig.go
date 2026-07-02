package managedkafkacluster

type ManagedKafkaClusterGcpConfigAccessConfig struct {
	// network_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/managed_kafka_cluster#network_configs ManagedKafkaCluster#network_configs}
	NetworkConfigs any `field:"required" json:"networkConfigs" yaml:"networkConfigs"`
}
