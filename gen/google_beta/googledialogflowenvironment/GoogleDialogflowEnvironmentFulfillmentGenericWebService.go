package googledialogflowenvironment


type GoogleDialogflowEnvironmentFulfillmentGenericWebService struct {
	// The fulfillment URI for receiving POST requests. It must use https protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#uri GoogleDialogflowEnvironment#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The password for HTTP Basic authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#password GoogleDialogflowEnvironment#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The HTTP request headers to send together with fulfillment requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#request_headers GoogleDialogflowEnvironment#request_headers}
	RequestHeaders *map[string]*string `field:"optional" json:"requestHeaders" yaml:"requestHeaders"`
	// The user name for HTTP Basic authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#username GoogleDialogflowEnvironment#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

