package googlefirebaseremoteconfigremoteconfig


type GoogleFirebaseRemoteConfigRemoteConfigParameterGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_remote_config_remote_config#parameter_group_name GoogleFirebaseRemoteConfigRemoteConfig#parameter_group_name}.
	ParameterGroupName *string `field:"required" json:"parameterGroupName" yaml:"parameterGroupName"`
	// A description for the group.
	//
	// Its length must be less than or equal to 256
	// characters. A description may contain any Unicode characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_remote_config_remote_config#description GoogleFirebaseRemoteConfigRemoteConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_remote_config_remote_config#parameters GoogleFirebaseRemoteConfigRemoteConfig#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

