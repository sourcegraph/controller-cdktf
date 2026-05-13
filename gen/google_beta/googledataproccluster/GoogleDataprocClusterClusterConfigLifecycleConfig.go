package googledataproccluster


type GoogleDataprocClusterClusterConfigLifecycleConfig struct {
	// The time when cluster will be auto-deleted. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_cluster#auto_delete_time GoogleDataprocCluster#auto_delete_time}
	AutoDeleteTime *string `field:"optional" json:"autoDeleteTime" yaml:"autoDeleteTime"`
	// The time when cluster will be auto-stopped. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_cluster#auto_stop_time GoogleDataprocCluster#auto_stop_time}
	AutoStopTime *string `field:"optional" json:"autoStopTime" yaml:"autoStopTime"`
	// The duration to keep the cluster alive while idling (no jobs running).
	//
	// After this TTL, the cluster will be deleted. Valid range: [300s, 1209600s].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_cluster#idle_delete_ttl GoogleDataprocCluster#idle_delete_ttl}
	IdleDeleteTtl *string `field:"optional" json:"idleDeleteTtl" yaml:"idleDeleteTtl"`
	// The duration to keep the cluster started while idling (no jobs running).
	//
	// After this TTL, the cluster will be stopped. Valid range: [10m, 14d].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_cluster#idle_stop_ttl GoogleDataprocCluster#idle_stop_ttl}
	IdleStopTtl *string `field:"optional" json:"idleStopTtl" yaml:"idleStopTtl"`
}

