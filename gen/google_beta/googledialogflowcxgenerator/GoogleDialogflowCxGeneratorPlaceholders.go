package googledialogflowcxgenerator


type GoogleDialogflowCxGeneratorPlaceholders struct {
	// Unique ID used to map custom placeholder to parameters in fulfillment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#id GoogleDialogflowCxGenerator#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Custom placeholder value in the prompt text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_generator#name GoogleDialogflowCxGenerator#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

