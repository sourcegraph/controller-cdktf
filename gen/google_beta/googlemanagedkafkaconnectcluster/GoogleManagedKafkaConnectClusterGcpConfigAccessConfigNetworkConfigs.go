package googlemanagedkafkaconnectcluster


type GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigs struct {
	// VPC subnet to make available to the Kafka Connect cluster.
	//
	// Structured like: projects/{project}/regions/{region}/subnetworks/{subnet_id}. It is used to create a Private Service Connect (PSC) interface for the Kafka Connect workers. It must be located in the same region as the Kafka Connect cluster. The CIDR range of the subnet must be within the IPv4 address ranges for private networks, as specified in RFC 1918. The primary subnet CIDR range must have a minimum size of /22 (1024 addresses).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_managed_kafka_connect_cluster#primary_subnet GoogleManagedKafkaConnectCluster#primary_subnet}
	PrimarySubnet *string `field:"required" json:"primarySubnet" yaml:"primarySubnet"`
	// Additional subnets may be specified.
	//
	// They may be in another region, but must be in the same VPC network. The Connect workers can communicate with network endpoints in either the primary or additional subnets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_managed_kafka_connect_cluster#additional_subnets GoogleManagedKafkaConnectCluster#additional_subnets}
	AdditionalSubnets *[]*string `field:"optional" json:"additionalSubnets" yaml:"additionalSubnets"`
	// Additional DNS domain names from the subnet's network to be made visible to the Connect Cluster.
	//
	// When using MirrorMaker2, it's necessary to add the bootstrap address's dns domain name of the target cluster to make it visible to the connector. For example: my-kafka-cluster.us-central1.managedkafka.my-project.cloud.goog
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_managed_kafka_connect_cluster#dns_domain_names GoogleManagedKafkaConnectCluster#dns_domain_names}
	DnsDomainNames *[]*string `field:"optional" json:"dnsDomainNames" yaml:"dnsDomainNames"`
}

