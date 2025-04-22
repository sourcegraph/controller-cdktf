package googlemanagedkafkaconnectcluster


type GoogleManagedKafkaConnectClusterGcpConfigAccessConfig struct {
	// network_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_managed_kafka_connect_cluster#network_configs GoogleManagedKafkaConnectCluster#network_configs}
	NetworkConfigs interface{} `field:"required" json:"networkConfigs" yaml:"networkConfigs"`
}

