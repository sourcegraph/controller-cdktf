package googlefirebaseremoteconfigremoteconfig


type GoogleFirebaseRemoteConfigRemoteConfigConditions struct {
	// The logic of this condition.
	//
	// See the documentation regarding
	// [Condition
	// Expressions](https://firebase.google.com/docs/remote-config/condition-reference)
	// for the expected syntax of this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_remote_config_remote_config#expression GoogleFirebaseRemoteConfigRemoteConfig#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// A non-empty and unique name of this condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_remote_config_remote_config#name GoogleFirebaseRemoteConfigRemoteConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The color associated with this condition for display purposes in the Firebase Console.
	//
	// Not specifying this value results in the Console picking an arbitrary color to associate with the condition. Possible values: ["BLUE", "BROWN", "CYAN", "DEEP_ORANGE", "GREEN", "INDIGO", "LIME", "ORANGE", "PINK", "PURPLE", "TEAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_remote_config_remote_config#tag_color GoogleFirebaseRemoteConfigRemoteConfig#tag_color}
	TagColor *string `field:"optional" json:"tagColor" yaml:"tagColor"`
}

