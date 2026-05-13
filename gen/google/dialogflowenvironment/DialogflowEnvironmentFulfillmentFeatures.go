package dialogflowenvironment


type DialogflowEnvironmentFulfillmentFeatures struct {
	// The type of the feature that enabled for fulfillment. Possible values: ["TYPE_UNSPECIFIED", "SMALLTALK"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_environment#type DialogflowEnvironment#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

