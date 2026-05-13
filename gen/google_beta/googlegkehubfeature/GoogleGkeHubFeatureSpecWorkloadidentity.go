package googlegkehubfeature


type GoogleGkeHubFeatureSpecWorkloadidentity struct {
	// Pool to be used for Workload Identity.
	//
	// This pool in trust-domain mode is used with Fleet Tenancy, so that sameness can be enforced. ex: projects/example/locations/global/workloadidentitypools/custompool
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_feature#scope_tenancy_pool GoogleGkeHubFeature#scope_tenancy_pool}
	ScopeTenancyPool *string `field:"required" json:"scopeTenancyPool" yaml:"scopeTenancyPool"`
}

