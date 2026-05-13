package googlefirebaseremoteconfigremoteconfig


type GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_remote_config_remote_config#parameter_name GoogleFirebaseRemoteConfigRemoteConfig#parameter_name}.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// conditional_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_remote_config_remote_config#conditional_values GoogleFirebaseRemoteConfigRemoteConfig#conditional_values}
	ConditionalValues interface{} `field:"optional" json:"conditionalValues" yaml:"conditionalValues"`
	// default_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_remote_config_remote_config#default_value GoogleFirebaseRemoteConfigRemoteConfig#default_value}
	DefaultValue *GoogleFirebaseRemoteConfigRemoteConfigParameterGroupsParametersDefaultValue `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// A description for this Parameter.
	//
	// Its length must be less than or equal to
	// 256 characters . A description may contain any Unicode characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_remote_config_remote_config#description GoogleFirebaseRemoteConfigRemoteConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The data type for all values of this parameter in the current version of the template.
	//
	// Default value: "STRING" Possible values: ["STRING", "BOOLEAN", "NUMBER", "JSON"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_remote_config_remote_config#value_type GoogleFirebaseRemoteConfigRemoteConfig#value_type}
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

