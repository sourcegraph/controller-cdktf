package googlemanagedkafkacluster


type GoogleManagedKafkaClusterBrokerCapacityConfig struct {
	// The disk to provision for each broker in Gibibytes. Minimum: 100 GiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_managed_kafka_cluster#disk_size_gib GoogleManagedKafkaCluster#disk_size_gib}
	DiskSizeGib *string `field:"optional" json:"diskSizeGib" yaml:"diskSizeGib"`
}

