package googledataproccluster


type GoogleDataprocClusterClusterConfigEndpointConfig struct {
	// The flag to enable http access to specific ports on the cluster from external sources (aka Component Gateway).
	//
	// Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#enable_http_port_access GoogleDataprocCluster#enable_http_port_access}
	EnableHttpPortAccess interface{} `field:"required" json:"enableHttpPortAccess" yaml:"enableHttpPortAccess"`
}

