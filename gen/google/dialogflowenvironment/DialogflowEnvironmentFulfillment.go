package dialogflowenvironment


type DialogflowEnvironmentFulfillment struct {
	// The human-readable name of the fulfillment, unique within the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_environment#display_name DialogflowEnvironment#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_environment#features DialogflowEnvironment#features}
	Features interface{} `field:"optional" json:"features" yaml:"features"`
	// generic_web_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_environment#generic_web_service DialogflowEnvironment#generic_web_service}
	GenericWebService *DialogflowEnvironmentFulfillmentGenericWebService `field:"optional" json:"genericWebService" yaml:"genericWebService"`
	// The unique identifier of the fulfillment. Supports the following formats: - projects/<Project ID>/agent/fulfillment - projects/<Project ID>/locations/<Location ID>/agent/fulfillment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_environment#name DialogflowEnvironment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

