package googlegameservicesgameserverdeploymentrollout


type GoogleGameServicesGameServerDeploymentRolloutGameServerConfigOverrides struct {
	// Version of the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_game_services_game_server_deployment_rollout#config_version GoogleGameServicesGameServerDeploymentRollout#config_version}
	ConfigVersion *string `field:"optional" json:"configVersion" yaml:"configVersion"`
	// realms_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_game_services_game_server_deployment_rollout#realms_selector GoogleGameServicesGameServerDeploymentRollout#realms_selector}
	RealmsSelector *GoogleGameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelector `field:"optional" json:"realmsSelector" yaml:"realmsSelector"`
}

