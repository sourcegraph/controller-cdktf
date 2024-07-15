package googlemanagedkafkacluster


type GoogleManagedKafkaClusterRebalanceConfig struct {
	// The rebalance behavior for the cluster. When not specified, defaults to 'NO_REBALANCE'. Possible values: 'MODE_UNSPECIFIED', 'NO_REBALANCE', 'AUTO_REBALANCE_ON_SCALE_UP'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#mode GoogleManagedKafkaCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

