package googledialogflowenvironment


type GoogleDialogflowEnvironmentFulfillmentFeatures struct {
	// The type of the feature that enabled for fulfillment. Possible values: ["TYPE_UNSPECIFIED", "SMALLTALK"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#type GoogleDialogflowEnvironment#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

