package cestoolset


type CesToolsetOpenApiToolsetApiAuthenticationServiceAccountAuthConfig struct {
	// The email address of the service account used for authenticatation.
	//
	// CES
	// uses this service account to exchange an access token and the access token
	// is then sent in the 'Authorization' header of the request.
	//
	// The service account must have the
	// 'roles/iam.serviceAccountTokenCreator' role granted to the
	// CES service agent
	// 'service-@gcp-sa-ces.iam.gserviceaccount.com'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#service_account CesToolset#service_account}
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
	// The OAuth scopes to grant. If not specified, the default scope 'https://www.googleapis.com/auth/cloud-platform' is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#scopes CesToolset#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

