package googledialogflowcxwebhook


type GoogleDialogflowCxWebhookGenericWebService struct {
	// Whether to use speech adaptation for speech recognition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_webhook#uri GoogleDialogflowCxWebhook#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_webhook#allowed_ca_certs GoogleDialogflowCxWebhook#allowed_ca_certs}
	AllowedCaCerts *[]*string `field:"optional" json:"allowedCaCerts" yaml:"allowedCaCerts"`
	// The HTTP request headers to send together with webhook requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_webhook#request_headers GoogleDialogflowCxWebhook#request_headers}
	RequestHeaders *map[string]*string `field:"optional" json:"requestHeaders" yaml:"requestHeaders"`
}

