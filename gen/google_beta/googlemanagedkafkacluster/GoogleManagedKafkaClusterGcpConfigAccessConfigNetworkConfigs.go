package googlemanagedkafkacluster


type GoogleManagedKafkaClusterGcpConfigAccessConfigNetworkConfigs struct {
	// Name of the VPC subnet from which the cluster is accessible.
	//
	// Both broker and bootstrap server IP addresses and DNS entries are automatically created in the subnet. The subnet must be located in the same region as the cluster. The project may differ. A minimum of 1 subnet is required. A maximum of 10 subnets can be specified. The name of the subnet must be in the format 'projects/PROJECT_ID/regions/REGION/subnetworks/SUBNET'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_managed_kafka_cluster#subnet GoogleManagedKafkaCluster#subnet}
	Subnet *string `field:"required" json:"subnet" yaml:"subnet"`
}

