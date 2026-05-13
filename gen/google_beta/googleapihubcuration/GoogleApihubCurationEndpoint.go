package googleapihubcuration


type GoogleApihubCurationEndpoint struct {
	// application_integration_endpoint_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_curation#application_integration_endpoint_details GoogleApihubCuration#application_integration_endpoint_details}
	ApplicationIntegrationEndpointDetails *GoogleApihubCurationEndpointApplicationIntegrationEndpointDetails `field:"required" json:"applicationIntegrationEndpointDetails" yaml:"applicationIntegrationEndpointDetails"`
}

