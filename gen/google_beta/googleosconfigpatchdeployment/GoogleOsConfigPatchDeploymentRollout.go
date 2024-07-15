package googleosconfigpatchdeployment


type GoogleOsConfigPatchDeploymentRollout struct {
	// disruption_budget block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#disruption_budget GoogleOsConfigPatchDeployment#disruption_budget}
	DisruptionBudget *GoogleOsConfigPatchDeploymentRolloutDisruptionBudget `field:"required" json:"disruptionBudget" yaml:"disruptionBudget"`
	// Mode of the patch rollout. Possible values: ["ZONE_BY_ZONE", "CONCURRENT_ZONES"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_patch_deployment#mode GoogleOsConfigPatchDeployment#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

