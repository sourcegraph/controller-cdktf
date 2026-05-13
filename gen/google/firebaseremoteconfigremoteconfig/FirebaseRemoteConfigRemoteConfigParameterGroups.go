package firebaseremoteconfigremoteconfig


type FirebaseRemoteConfigRemoteConfigParameterGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firebase_remote_config_remote_config#parameter_group_name FirebaseRemoteConfigRemoteConfig#parameter_group_name}.
	ParameterGroupName *string `field:"required" json:"parameterGroupName" yaml:"parameterGroupName"`
	// A description for the group.
	//
	// Its length must be less than or equal to 256
	// characters. A description may contain any Unicode characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firebase_remote_config_remote_config#description FirebaseRemoteConfigRemoteConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firebase_remote_config_remote_config#parameters FirebaseRemoteConfigRemoteConfig#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

