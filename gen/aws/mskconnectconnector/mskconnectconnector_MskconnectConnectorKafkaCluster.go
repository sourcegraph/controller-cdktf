package mskconnectconnector


type MskconnectConnectorKafkaCluster struct {
	// apache_kafka_cluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/mskconnect_connector#apache_kafka_cluster MskconnectConnector#apache_kafka_cluster}
	ApacheKafkaCluster *MskconnectConnectorKafkaClusterApacheKafkaCluster `field:"required" json:"apacheKafkaCluster" yaml:"apacheKafkaCluster"`
}

