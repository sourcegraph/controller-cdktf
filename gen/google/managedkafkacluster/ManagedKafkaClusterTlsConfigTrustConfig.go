package managedkafkacluster


type ManagedKafkaClusterTlsConfigTrustConfig struct {
	// cas_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/managed_kafka_cluster#cas_configs ManagedKafkaCluster#cas_configs}
	CasConfigs interface{} `field:"optional" json:"casConfigs" yaml:"casConfigs"`
}

