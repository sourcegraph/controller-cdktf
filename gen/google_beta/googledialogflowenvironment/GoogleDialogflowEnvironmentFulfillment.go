package googledialogflowenvironment


type GoogleDialogflowEnvironmentFulfillment struct {
	// The human-readable name of the fulfillment, unique within the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#display_name GoogleDialogflowEnvironment#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#features GoogleDialogflowEnvironment#features}
	Features interface{} `field:"optional" json:"features" yaml:"features"`
	// generic_web_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#generic_web_service GoogleDialogflowEnvironment#generic_web_service}
	GenericWebService *GoogleDialogflowEnvironmentFulfillmentGenericWebService `field:"optional" json:"genericWebService" yaml:"genericWebService"`
	// The unique identifier of the fulfillment. Supports the following formats: - projects/<Project ID>/agent/fulfillment - projects/<Project ID>/locations/<Location ID>/agent/fulfillment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#name GoogleDialogflowEnvironment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

