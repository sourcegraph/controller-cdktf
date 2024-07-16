package googlebeyondcorpappconnector


type GoogleBeyondcorpAppConnectorPrincipalInfo struct {
	// service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_beyondcorp_app_connector#service_account GoogleBeyondcorpAppConnector#service_account}
	ServiceAccount *GoogleBeyondcorpAppConnectorPrincipalInfoServiceAccount `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
}

