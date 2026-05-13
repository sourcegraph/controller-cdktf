package googlefirebaseremoteconfigremoteconfig


type GoogleFirebaseRemoteConfigRemoteConfigParametersDefaultValue struct {
	// If true, the parameter is omitted from the parameter values returned to a client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_remote_config_remote_config#use_in_app_default GoogleFirebaseRemoteConfigRemoteConfig#use_in_app_default}
	UseInAppDefault interface{} `field:"optional" json:"useInAppDefault" yaml:"useInAppDefault"`
	// The string value that the parameter is set to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_remote_config_remote_config#value GoogleFirebaseRemoteConfigRemoteConfig#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

