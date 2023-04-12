package googlegkehubfeature


type GoogleGkeHubFeatureSpecMulticlusteringress struct {
	// Fully-qualified Membership name which hosts the MultiClusterIngress CRD. Example: `projects/foo-proj/locations/global/memberships/bar`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature#config_membership GoogleGkeHubFeature#config_membership}
	ConfigMembership *string `field:"required" json:"configMembership" yaml:"configMembership"`
}

