package googlegkehubfleet


type GoogleGkeHubFleetDefaultClusterConfigBinaryAuthorizationConfig struct {
	// Mode of operation for binauthz policy evaluation. Possible values: ["DISABLED", "POLICY_BINDINGS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_fleet#evaluation_mode GoogleGkeHubFleet#evaluation_mode}
	EvaluationMode *string `field:"optional" json:"evaluationMode" yaml:"evaluationMode"`
	// policy_bindings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_fleet#policy_bindings GoogleGkeHubFleet#policy_bindings}
	PolicyBindings interface{} `field:"optional" json:"policyBindings" yaml:"policyBindings"`
}

