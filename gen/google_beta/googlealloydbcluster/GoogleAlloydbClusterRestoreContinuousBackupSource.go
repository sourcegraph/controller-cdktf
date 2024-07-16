package googlealloydbcluster


type GoogleAlloydbClusterRestoreContinuousBackupSource struct {
	// The name of the source cluster that this cluster is restored from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#cluster GoogleAlloydbCluster#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The point in time that this cluster is restored to, in RFC 3339 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#point_in_time GoogleAlloydbCluster#point_in_time}
	PointInTime *string `field:"required" json:"pointInTime" yaml:"pointInTime"`
}

