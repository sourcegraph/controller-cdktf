package googlemanagedkafkacluster


type GoogleManagedKafkaClusterGcpConfig struct {
	// access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#access_config GoogleManagedKafkaCluster#access_config}
	AccessConfig *GoogleManagedKafkaClusterGcpConfigAccessConfig `field:"required" json:"accessConfig" yaml:"accessConfig"`
	// The Cloud KMS Key name to use for encryption.
	//
	// The key must be located in the same region as the cluster and cannot be changed. Must be in the format 'projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#kms_key GoogleManagedKafkaCluster#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

