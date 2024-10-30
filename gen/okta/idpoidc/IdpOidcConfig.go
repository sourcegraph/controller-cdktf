package idpoidc

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdpOidcConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// The method of making an authorization request. It can be set to `HTTP-POST` or `HTTP-REDIRECT`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#authorization_binding IdpOidc#authorization_binding}
	AuthorizationBinding *string `field:"required" json:"authorizationBinding" yaml:"authorizationBinding"`
	// IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#authorization_url IdpOidc#authorization_url}
	AuthorizationUrl *string `field:"required" json:"authorizationUrl" yaml:"authorizationUrl"`
	// Unique identifier issued by AS for the Okta IdP instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#client_id IdpOidc#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Client secret issued by AS for the Okta IdP instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#client_secret IdpOidc#client_secret}
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// URI that identifies the issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#issuer_url IdpOidc#issuer_url}
	IssuerUrl *string `field:"required" json:"issuerUrl" yaml:"issuerUrl"`
	// The method of making a request for the OIDC JWKS. It can be set to `HTTP-POST` or `HTTP-REDIRECT`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#jwks_binding IdpOidc#jwks_binding}
	JwksBinding *string `field:"required" json:"jwksBinding" yaml:"jwksBinding"`
	// Endpoint where the keys signer publishes its keys in a JWK Set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#jwks_url IdpOidc#jwks_url}
	JwksUrl *string `field:"required" json:"jwksUrl" yaml:"jwksUrl"`
	// Name of the IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#name IdpOidc#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The scopes of the IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#scopes IdpOidc#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The method of making a token request. It can be set to `HTTP-POST` or `HTTP-REDIRECT`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#token_binding IdpOidc#token_binding}
	TokenBinding *string `field:"required" json:"tokenBinding" yaml:"tokenBinding"`
	// IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#token_url IdpOidc#token_url}
	TokenUrl *string `field:"required" json:"tokenUrl" yaml:"tokenUrl"`
	// Specifies the account linking action for an IdP user. Default: `AUTO`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#account_link_action IdpOidc#account_link_action}
	AccountLinkAction *string `field:"optional" json:"accountLinkAction" yaml:"accountLinkAction"`
	// Group memberships to determine link candidates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#account_link_group_include IdpOidc#account_link_group_include}
	AccountLinkGroupInclude *[]*string `field:"optional" json:"accountLinkGroupInclude" yaml:"accountLinkGroupInclude"`
	// Action for a previously deprovisioned IdP user during authentication. Can be `NONE` or `REACTIVATE`. Default: `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#deprovisioned_action IdpOidc#deprovisioned_action}
	DeprovisionedAction *string `field:"optional" json:"deprovisionedAction" yaml:"deprovisionedAction"`
	// Provisioning action for IdP user's group memberships. It can be `NONE`, `SYNC`, `APPEND`, or `ASSIGN`. Default: `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#groups_action IdpOidc#groups_action}
	GroupsAction *string `field:"optional" json:"groupsAction" yaml:"groupsAction"`
	// List of Okta Group IDs to add an IdP user as a member with the `ASSIGN` `groups_action`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#groups_assignment IdpOidc#groups_assignment}
	GroupsAssignment *[]*string `field:"optional" json:"groupsAssignment" yaml:"groupsAssignment"`
	// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#groups_attribute IdpOidc#groups_attribute}
	GroupsAttribute *string `field:"optional" json:"groupsAttribute" yaml:"groupsAttribute"`
	// Whitelist of Okta Group identifiers that are allowed for the `APPEND` or `SYNC` `groups_action`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#groups_filter IdpOidc#groups_filter}
	GroupsFilter *[]*string `field:"optional" json:"groupsFilter" yaml:"groupsFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#id IdpOidc#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates whether Okta uses the original Okta org domain URL, a custom domain URL, or dynamic.
	//
	// It can be `ORG_URL`, `CUSTOM_URL`, or `DYNAMIC`. Default: `ORG_URL`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#issuer_mode IdpOidc#issuer_mode}
	IssuerMode *string `field:"optional" json:"issuerMode" yaml:"issuerMode"`
	// Maximum allowable clock-skew when processing messages from the IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#max_clock_skew IdpOidc#max_clock_skew}
	MaxClockSkew *float64 `field:"optional" json:"maxClockSkew" yaml:"maxClockSkew"`
	// Require Proof Key for Code Exchange (PKCE) for additional verification key rotation mode. See: https://developer.okta.com/docs/reference/api/idps/#oauth-2-0-and-openid-connect-client-object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#pkce_required IdpOidc#pkce_required}
	PkceRequired interface{} `field:"optional" json:"pkceRequired" yaml:"pkceRequired"`
	// Determines if the IdP should act as a source of truth for user profile attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#profile_master IdpOidc#profile_master}
	ProfileMaster interface{} `field:"optional" json:"profileMaster" yaml:"profileMaster"`
	// The type of protocol to use. It can be `OIDC` or `OAUTH2`. Default: `OIDC`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#protocol_type IdpOidc#protocol_type}
	ProtocolType *string `field:"optional" json:"protocolType" yaml:"protocolType"`
	// Provisioning action for an IdP user during authentication. Default: `AUTO`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#provisioning_action IdpOidc#provisioning_action}
	ProvisioningAction *string `field:"optional" json:"provisioningAction" yaml:"provisioningAction"`
	// The HMAC Signature Algorithm used when signing an authorization request.
	//
	// Defaults to `HS256`. It can be `HS256`, `HS384`, `HS512`, `SHA-256`. `RS256`, `RS384`, or `RS512`. NOTE: `SHA-256` an undocumented legacy value and not continue to be valid. See API docs https://developer.okta.com/docs/reference/api/idps/#oidc-request-signature-algorithm-object
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#request_signature_algorithm IdpOidc#request_signature_algorithm}
	RequestSignatureAlgorithm *string `field:"optional" json:"requestSignatureAlgorithm" yaml:"requestSignatureAlgorithm"`
	// Specifies whether to digitally sign an AuthnRequest messages to the IdP.
	//
	// Defaults to `REQUEST`. It can be `REQUEST` or `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#request_signature_scope IdpOidc#request_signature_scope}
	RequestSignatureScope *string `field:"optional" json:"requestSignatureScope" yaml:"requestSignatureScope"`
	// Default to `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#status IdpOidc#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Okta user profile attribute for matching transformed IdP username. Only for matchType `CUSTOM_ATTRIBUTE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#subject_match_attribute IdpOidc#subject_match_attribute}
	SubjectMatchAttribute *string `field:"optional" json:"subjectMatchAttribute" yaml:"subjectMatchAttribute"`
	// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username.
	//
	// By default, it is set to `USERNAME`. It can be set to `USERNAME`, `EMAIL`, `USERNAME_OR_EMAIL` or `CUSTOM_ATTRIBUTE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#subject_match_type IdpOidc#subject_match_type}
	SubjectMatchType *string `field:"optional" json:"subjectMatchType" yaml:"subjectMatchType"`
	// Action for a previously suspended IdP user during authentication. Can be `NONE` or `REACTIVATE`. Default: `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#suspended_action IdpOidc#suspended_action}
	SuspendedAction *string `field:"optional" json:"suspendedAction" yaml:"suspendedAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#user_info_binding IdpOidc#user_info_binding}.
	UserInfoBinding *string `field:"optional" json:"userInfoBinding" yaml:"userInfoBinding"`
	// Protected resource endpoint that returns claims about the authenticated user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#user_info_url IdpOidc#user_info_url}
	UserInfoUrl *string `field:"optional" json:"userInfoUrl" yaml:"userInfoUrl"`
	// Okta EL Expression to generate or transform a unique username for the IdP user. Default: `idpuser.email`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_oidc#username_template IdpOidc#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

