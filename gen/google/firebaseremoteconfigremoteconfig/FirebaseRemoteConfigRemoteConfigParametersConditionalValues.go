package firebaseremoteconfigremoteconfig


type FirebaseRemoteConfigRemoteConfigParametersConditionalValues struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firebase_remote_config_remote_config#condition_name FirebaseRemoteConfigRemoteConfig#condition_name}.
	ConditionName *string `field:"required" json:"conditionName" yaml:"conditionName"`
	// If true, the parameter is omitted from the parameter values returned to a client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firebase_remote_config_remote_config#use_in_app_default FirebaseRemoteConfigRemoteConfig#use_in_app_default}
	UseInAppDefault interface{} `field:"optional" json:"useInAppDefault" yaml:"useInAppDefault"`
	// The string value that the parameter is set to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firebase_remote_config_remote_config#value FirebaseRemoteConfigRemoteConfig#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

