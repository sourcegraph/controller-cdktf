package googlemanagedkafkacluster


type GoogleManagedKafkaClusterTlsConfigTrustConfigCasConfigs struct {
	// The name of the CA pool to pull CA certificates from.
	//
	// The CA pool does not need to be in the same project or location as the Kafka cluster. Must be in the format 'projects/PROJECT_ID/locations/LOCATION/caPools/CA_POOL_ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_managed_kafka_cluster#ca_pool GoogleManagedKafkaCluster#ca_pool}
	CaPool *string `field:"required" json:"caPool" yaml:"caPool"`
}

