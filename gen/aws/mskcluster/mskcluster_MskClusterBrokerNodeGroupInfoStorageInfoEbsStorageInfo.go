package mskcluster


type MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfo struct {
	// provisioned_throughput block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_cluster#provisioned_throughput MskCluster#provisioned_throughput}
	ProvisionedThroughput *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughput `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_cluster#volume_size MskCluster#volume_size}.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
}
