package googlegameservicesgameserverdeploymentrollout


type GoogleGameServicesGameServerDeploymentRolloutGameServerConfigOverrides struct {
	// Version of the configuration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_game_services_game_server_deployment_rollout#config_version GoogleGameServicesGameServerDeploymentRollout#config_version}
	ConfigVersion *string `field:"optional" json:"configVersion" yaml:"configVersion"`
	// realms_selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_game_services_game_server_deployment_rollout#realms_selector GoogleGameServicesGameServerDeploymentRollout#realms_selector}
	RealmsSelector *GoogleGameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelector `field:"optional" json:"realmsSelector" yaml:"realmsSelector"`
}

