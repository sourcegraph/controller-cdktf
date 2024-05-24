package googleaccesscontextmanagerserviceperimeters


type GoogleAccessContextManagerServicePerimetersServicePerimetersStatusIngressPoliciesIngressFrom struct {
	// 'A list of identities that are allowed access through this 'IngressPolicy'.
	//
	// To specify an identity or identity group, use the IAM v1 format
	// specified [here](https://cloud.google.com/iam/docs/principal-identifiers.md#v1).
	// The following prefixes are supprted: user, group, serviceAccount, principal, and principalSet.'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_access_context_manager_service_perimeters#identities GoogleAccessContextManagerServicePerimeters#identities}
	Identities *[]*string `field:"optional" json:"identities" yaml:"identities"`
	// Specifies the type of identities that are allowed access from outside the perimeter.
	//
	// If left unspecified, then members of 'identities' field will be
	// allowed access. Possible values: ["IDENTITY_TYPE_UNSPECIFIED", "ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_access_context_manager_service_perimeters#identity_type GoogleAccessContextManagerServicePerimeters#identity_type}
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
	// sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_access_context_manager_service_perimeters#sources GoogleAccessContextManagerServicePerimeters#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

