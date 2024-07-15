package googlebigtableappprofile


type GoogleBigtableAppProfileSingleClusterRouting struct {
	// The cluster to which read/write requests should be routed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_app_profile#cluster_id GoogleBigtableAppProfile#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// If true, CheckAndMutateRow and ReadModifyWriteRow requests are allowed by this app profile.
	//
	// It is unsafe to send these requests to the same table/row/column in multiple clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_app_profile#allow_transactional_writes GoogleBigtableAppProfile#allow_transactional_writes}
	AllowTransactionalWrites interface{} `field:"optional" json:"allowTransactionalWrites" yaml:"allowTransactionalWrites"`
}

