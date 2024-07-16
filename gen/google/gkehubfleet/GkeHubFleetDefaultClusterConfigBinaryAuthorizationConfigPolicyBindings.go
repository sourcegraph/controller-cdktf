package gkehubfleet


type GkeHubFleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBindings struct {
	// The relative resource name of the binauthz platform policy to audit. GKE platform policies have the following format: 'projects/{project_number}/platforms/gke/policies/{policy_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/gke_hub_fleet#name GkeHubFleet#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

