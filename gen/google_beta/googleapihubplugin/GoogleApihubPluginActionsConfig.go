package googleapihubplugin


type GoogleApihubPluginActionsConfig struct {
	// The description of the operation performed by the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#description GoogleApihubPlugin#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The display name of the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#display_name GoogleApihubPlugin#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The id of the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#id GoogleApihubPlugin#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The trigger mode supported by the action. Possible values: TRIGGER_MODE_UNSPECIFIED API_HUB_ON_DEMAND_TRIGGER API_HUB_SCHEDULE_TRIGGER NON_API_HUB_MANAGED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#trigger_mode GoogleApihubPlugin#trigger_mode}
	TriggerMode *string `field:"required" json:"triggerMode" yaml:"triggerMode"`
}

