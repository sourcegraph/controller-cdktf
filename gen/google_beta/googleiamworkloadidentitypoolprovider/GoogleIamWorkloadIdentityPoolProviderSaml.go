package googleiamworkloadidentitypoolprovider


type GoogleIamWorkloadIdentityPoolProviderSaml struct {
	// SAML Identity provider configuration metadata xml doc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iam_workload_identity_pool_provider#idp_metadata_xml GoogleIamWorkloadIdentityPoolProvider#idp_metadata_xml}
	IdpMetadataXml *string `field:"required" json:"idpMetadataXml" yaml:"idpMetadataXml"`
}

