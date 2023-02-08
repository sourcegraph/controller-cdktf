package idpoidc

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdpOidcConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#authorization_binding IdpOidc#authorization_binding}.
	AuthorizationBinding *string `field:"required" json:"authorizationBinding" yaml:"authorizationBinding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#authorization_url IdpOidc#authorization_url}.
	AuthorizationUrl *string `field:"required" json:"authorizationUrl" yaml:"authorizationUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#client_id IdpOidc#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#client_secret IdpOidc#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#issuer_url IdpOidc#issuer_url}.
	IssuerUrl *string `field:"required" json:"issuerUrl" yaml:"issuerUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#jwks_binding IdpOidc#jwks_binding}.
	JwksBinding *string `field:"required" json:"jwksBinding" yaml:"jwksBinding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#jwks_url IdpOidc#jwks_url}.
	JwksUrl *string `field:"required" json:"jwksUrl" yaml:"jwksUrl"`
	// Name of the IdP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#name IdpOidc#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#scopes IdpOidc#scopes}.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#token_binding IdpOidc#token_binding}.
	TokenBinding *string `field:"required" json:"tokenBinding" yaml:"tokenBinding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#token_url IdpOidc#token_url}.
	TokenUrl *string `field:"required" json:"tokenUrl" yaml:"tokenUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#account_link_action IdpOidc#account_link_action}.
	AccountLinkAction *string `field:"optional" json:"accountLinkAction" yaml:"accountLinkAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#account_link_group_include IdpOidc#account_link_group_include}.
	AccountLinkGroupInclude *[]*string `field:"optional" json:"accountLinkGroupInclude" yaml:"accountLinkGroupInclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#deprovisioned_action IdpOidc#deprovisioned_action}.
	DeprovisionedAction *string `field:"optional" json:"deprovisionedAction" yaml:"deprovisionedAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#groups_action IdpOidc#groups_action}.
	GroupsAction *string `field:"optional" json:"groupsAction" yaml:"groupsAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#groups_assignment IdpOidc#groups_assignment}.
	GroupsAssignment *[]*string `field:"optional" json:"groupsAssignment" yaml:"groupsAssignment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#groups_attribute IdpOidc#groups_attribute}.
	GroupsAttribute *string `field:"optional" json:"groupsAttribute" yaml:"groupsAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#groups_filter IdpOidc#groups_filter}.
	GroupsFilter *[]*string `field:"optional" json:"groupsFilter" yaml:"groupsFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#id IdpOidc#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates whether Okta uses the original Okta org domain URL, custom domain URL, or dynamic.
	//
	// See Identity Provider attributes - issuerMode - https://developer.okta.com/docs/reference/api/idps/#identity-provider-attributes
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#issuer_mode IdpOidc#issuer_mode}
	IssuerMode *string `field:"optional" json:"issuerMode" yaml:"issuerMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#max_clock_skew IdpOidc#max_clock_skew}.
	MaxClockSkew *float64 `field:"optional" json:"maxClockSkew" yaml:"maxClockSkew"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#profile_master IdpOidc#profile_master}.
	ProfileMaster interface{} `field:"optional" json:"profileMaster" yaml:"profileMaster"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#protocol_type IdpOidc#protocol_type}.
	ProtocolType *string `field:"optional" json:"protocolType" yaml:"protocolType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#provisioning_action IdpOidc#provisioning_action}.
	ProvisioningAction *string `field:"optional" json:"provisioningAction" yaml:"provisioningAction"`
	// The HMAC Signature Algorithm used when signing an authorization request.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#request_signature_algorithm IdpOidc#request_signature_algorithm}
	RequestSignatureAlgorithm *string `field:"optional" json:"requestSignatureAlgorithm" yaml:"requestSignatureAlgorithm"`
	// Specifies whether to digitally sign an authorization request to the IdP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#request_signature_scope IdpOidc#request_signature_scope}
	RequestSignatureScope *string `field:"optional" json:"requestSignatureScope" yaml:"requestSignatureScope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#status IdpOidc#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#subject_match_attribute IdpOidc#subject_match_attribute}.
	SubjectMatchAttribute *string `field:"optional" json:"subjectMatchAttribute" yaml:"subjectMatchAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#subject_match_type IdpOidc#subject_match_type}.
	SubjectMatchType *string `field:"optional" json:"subjectMatchType" yaml:"subjectMatchType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#suspended_action IdpOidc#suspended_action}.
	SuspendedAction *string `field:"optional" json:"suspendedAction" yaml:"suspendedAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#user_info_binding IdpOidc#user_info_binding}.
	UserInfoBinding *string `field:"optional" json:"userInfoBinding" yaml:"userInfoBinding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#user_info_url IdpOidc#user_info_url}.
	UserInfoUrl *string `field:"optional" json:"userInfoUrl" yaml:"userInfoUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_oidc#username_template IdpOidc#username_template}.
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

