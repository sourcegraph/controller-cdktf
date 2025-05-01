package googlemanagedkafkaconnectcluster


type GoogleManagedKafkaConnectClusterGcpConfig struct {
	// access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_managed_kafka_connect_cluster#access_config GoogleManagedKafkaConnectCluster#access_config}
	AccessConfig *GoogleManagedKafkaConnectClusterGcpConfigAccessConfig `field:"required" json:"accessConfig" yaml:"accessConfig"`
}

