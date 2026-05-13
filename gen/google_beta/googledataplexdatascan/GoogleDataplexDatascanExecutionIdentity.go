package googledataplexdatascan


type GoogleDataplexDatascanExecutionIdentity struct {
	// dataplex_service_agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#dataplex_service_agent GoogleDataplexDatascan#dataplex_service_agent}
	DataplexServiceAgent *GoogleDataplexDatascanExecutionIdentityDataplexServiceAgent `field:"optional" json:"dataplexServiceAgent" yaml:"dataplexServiceAgent"`
	// service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#service_account GoogleDataplexDatascan#service_account}
	ServiceAccount *GoogleDataplexDatascanExecutionIdentityServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// user_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#user_credential GoogleDataplexDatascan#user_credential}
	UserCredential *GoogleDataplexDatascanExecutionIdentityUserCredential `field:"optional" json:"userCredential" yaml:"userCredential"`
}

