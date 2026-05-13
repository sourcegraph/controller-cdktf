package dataplexdatascan


type DataplexDatascanExecutionIdentity struct {
	// dataplex_service_agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_datascan#dataplex_service_agent DataplexDatascan#dataplex_service_agent}
	DataplexServiceAgent *DataplexDatascanExecutionIdentityDataplexServiceAgent `field:"optional" json:"dataplexServiceAgent" yaml:"dataplexServiceAgent"`
	// service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_datascan#service_account DataplexDatascan#service_account}
	ServiceAccount *DataplexDatascanExecutionIdentityServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// user_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataplex_datascan#user_credential DataplexDatascan#user_credential}
	UserCredential *DataplexDatascanExecutionIdentityUserCredential `field:"optional" json:"userCredential" yaml:"userCredential"`
}

