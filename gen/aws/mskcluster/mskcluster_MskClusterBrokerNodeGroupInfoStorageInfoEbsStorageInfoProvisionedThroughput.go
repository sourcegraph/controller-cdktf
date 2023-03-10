package mskcluster


type MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfoProvisionedThroughput struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_cluster#enabled MskCluster#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_cluster#volume_throughput MskCluster#volume_throughput}.
	VolumeThroughput *float64 `field:"optional" json:"volumeThroughput" yaml:"volumeThroughput"`
}

