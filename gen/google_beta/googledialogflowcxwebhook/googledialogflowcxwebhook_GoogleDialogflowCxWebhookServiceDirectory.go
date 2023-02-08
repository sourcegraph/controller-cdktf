package googledialogflowcxwebhook


type GoogleDialogflowCxWebhookServiceDirectory struct {
	// generic_web_service block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dialogflow_cx_webhook#generic_web_service GoogleDialogflowCxWebhook#generic_web_service}
	GenericWebService *GoogleDialogflowCxWebhookServiceDirectoryGenericWebService `field:"required" json:"genericWebService" yaml:"genericWebService"`
	// The name of Service Directory service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dialogflow_cx_webhook#service GoogleDialogflowCxWebhook#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

