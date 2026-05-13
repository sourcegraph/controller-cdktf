package googledataproccluster


type GoogleDataprocClusterClusterConfigSecurityConfigIdentityConfig struct {
	// User to service account mappings for multi-tenancy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataproc_cluster#user_service_account_mapping GoogleDataprocCluster#user_service_account_mapping}
	UserServiceAccountMapping *map[string]*string `field:"required" json:"userServiceAccountMapping" yaml:"userServiceAccountMapping"`
}

