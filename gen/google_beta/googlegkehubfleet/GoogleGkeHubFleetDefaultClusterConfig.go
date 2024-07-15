package googlegkehubfleet


type GoogleGkeHubFleetDefaultClusterConfig struct {
	// binary_authorization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_fleet#binary_authorization_config GoogleGkeHubFleet#binary_authorization_config}
	BinaryAuthorizationConfig *GoogleGkeHubFleetDefaultClusterConfigBinaryAuthorizationConfig `field:"optional" json:"binaryAuthorizationConfig" yaml:"binaryAuthorizationConfig"`
	// security_posture_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_fleet#security_posture_config GoogleGkeHubFleet#security_posture_config}
	SecurityPostureConfig *GoogleGkeHubFleetDefaultClusterConfigSecurityPostureConfig `field:"optional" json:"securityPostureConfig" yaml:"securityPostureConfig"`
}

