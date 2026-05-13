package googlemanagedkafkacluster


type GoogleManagedKafkaClusterTlsConfig struct {
	// The rules for mapping mTLS certificate Distinguished Names (DNs) to shortened principal names for Kafka ACLs.
	//
	// This field corresponds exactly to the ssl.principal.mapping.rules broker config and matches the format and syntax defined in the Apache Kafka documentation. Setting or modifying this field will trigger a rolling restart of the Kafka brokers to apply the change. An empty string means that the default Kafka behavior is used. Example: 'RULE:^CN=(.?),OU=ServiceUsers.$/$1@example.com/,DEFAULT'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_managed_kafka_cluster#ssl_principal_mapping_rules GoogleManagedKafkaCluster#ssl_principal_mapping_rules}
	SslPrincipalMappingRules *string `field:"optional" json:"sslPrincipalMappingRules" yaml:"sslPrincipalMappingRules"`
	// trust_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_managed_kafka_cluster#trust_config GoogleManagedKafkaCluster#trust_config}
	TrustConfig *GoogleManagedKafkaClusterTlsConfigTrustConfig `field:"optional" json:"trustConfig" yaml:"trustConfig"`
}

