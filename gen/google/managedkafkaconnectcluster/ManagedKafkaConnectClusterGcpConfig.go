package managedkafkaconnectcluster


type ManagedKafkaConnectClusterGcpConfig struct {
	// access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/managed_kafka_connect_cluster#access_config ManagedKafkaConnectCluster#access_config}
	AccessConfig *ManagedKafkaConnectClusterGcpConfigAccessConfig `field:"required" json:"accessConfig" yaml:"accessConfig"`
}

