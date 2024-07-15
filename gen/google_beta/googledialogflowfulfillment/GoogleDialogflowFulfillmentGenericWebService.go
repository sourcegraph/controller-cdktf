package googledialogflowfulfillment


type GoogleDialogflowFulfillmentGenericWebService struct {
	// The fulfillment URI for receiving POST requests. It must use https protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_fulfillment#uri GoogleDialogflowFulfillment#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The password for HTTP Basic authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_fulfillment#password GoogleDialogflowFulfillment#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The HTTP request headers to send together with fulfillment requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_fulfillment#request_headers GoogleDialogflowFulfillment#request_headers}
	RequestHeaders *map[string]*string `field:"optional" json:"requestHeaders" yaml:"requestHeaders"`
	// The user name for HTTP Basic authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_fulfillment#username GoogleDialogflowFulfillment#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

