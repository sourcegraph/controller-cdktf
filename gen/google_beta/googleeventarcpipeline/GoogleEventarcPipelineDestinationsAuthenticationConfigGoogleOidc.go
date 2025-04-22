package googleeventarcpipeline


type GoogleEventarcPipelineDestinationsAuthenticationConfigGoogleOidc struct {
	// Service account email used to generate the OIDC Token.
	//
	// The principal who calls this API must have
	// iam.serviceAccounts.actAs permission in the service account. See
	// https://cloud.google.com/iam/docs/understanding-service-accounts
	// for more information. Eventarc service agents must have
	// roles/roles/iam.serviceAccountTokenCreator role to allow the
	// Pipeline to create OpenID tokens for authenticated requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_pipeline#service_account GoogleEventarcPipeline#service_account}
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
	// Audience to be used to generate the OIDC Token.
	//
	// The audience claim
	// identifies the recipient that the JWT is intended for. If
	// unspecified, the destination URI will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_pipeline#audience GoogleEventarcPipeline#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
}

