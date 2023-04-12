package appoauth

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppOauthConfig struct {
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
	// Pretty name of app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#label AppOauth#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// The type of client application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#type AppOauth#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Custom error page URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#accessibility_error_redirect_url AppOauth#accessibility_error_redirect_url}
	AccessibilityErrorRedirectUrl *string `field:"optional" json:"accessibilityErrorRedirectUrl" yaml:"accessibilityErrorRedirectUrl"`
	// Custom login page URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#accessibility_login_redirect_url AppOauth#accessibility_login_redirect_url}
	AccessibilityLoginRedirectUrl *string `field:"optional" json:"accessibilityLoginRedirectUrl" yaml:"accessibilityLoginRedirectUrl"`
	// Enable self service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#accessibility_self_service AppOauth#accessibility_self_service}
	AccessibilitySelfService interface{} `field:"optional" json:"accessibilitySelfService" yaml:"accessibilitySelfService"`
	// Application notes for admins.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#admin_note AppOauth#admin_note}
	AdminNote *string `field:"optional" json:"adminNote" yaml:"adminNote"`
	// Displays specific appLinks for the app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#app_links_json AppOauth#app_links_json}
	AppLinksJson *string `field:"optional" json:"appLinksJson" yaml:"appLinksJson"`
	// Application settings in JSON format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#app_settings_json AppOauth#app_settings_json}
	AppSettingsJson *string `field:"optional" json:"appSettingsJson" yaml:"appSettingsJson"`
	// Id of this apps authentication policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#authentication_policy AppOauth#authentication_policy}
	AuthenticationPolicy *string `field:"optional" json:"authenticationPolicy" yaml:"authenticationPolicy"`
	// Requested key rotation mode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#auto_key_rotation AppOauth#auto_key_rotation}
	AutoKeyRotation interface{} `field:"optional" json:"autoKeyRotation" yaml:"autoKeyRotation"`
	// Display auto submit toolbar.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#auto_submit_toolbar AppOauth#auto_submit_toolbar}
	AutoSubmitToolbar interface{} `field:"optional" json:"autoSubmitToolbar" yaml:"autoSubmitToolbar"`
	// OAuth client secret key, this can be set when token_endpoint_auth_method is client_secret_basic.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#client_basic_secret AppOauth#client_basic_secret}
	ClientBasicSecret *string `field:"optional" json:"clientBasicSecret" yaml:"clientBasicSecret"`
	// OAuth client ID. If set during creation, app is created with this id.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#client_id AppOauth#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// URI to a web page providing information about the client.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#client_uri AppOauth#client_uri}
	ClientUri *string `field:"optional" json:"clientUri" yaml:"clientUri"`
	// *Early Access Property*. Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#consent_method AppOauth#consent_method}
	ConsentMethod *string `field:"optional" json:"consentMethod" yaml:"consentMethod"`
	// **Deprecated** This property allows you to set your client_id during creation.
	//
	// NOTE: updating after creation will be a no-op, use client_id for that behavior instead.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#custom_client_id AppOauth#custom_client_id}
	CustomClientId *string `field:"optional" json:"customClientId" yaml:"customClientId"`
	// Application notes for end users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#enduser_note AppOauth#enduser_note}
	EnduserNote *string `field:"optional" json:"enduserNote" yaml:"enduserNote"`
	// List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#grant_types AppOauth#grant_types}
	GrantTypes *[]*string `field:"optional" json:"grantTypes" yaml:"grantTypes"`
	// Groups associated with the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#groups AppOauth#groups}
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// groups_claim block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#groups_claim AppOauth#groups_claim}
	GroupsClaim *AppOauthGroupsClaim `field:"optional" json:"groupsClaim" yaml:"groupsClaim"`
	// Do not display application icon on mobile app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#hide_ios AppOauth#hide_ios}
	HideIos interface{} `field:"optional" json:"hideIos" yaml:"hideIos"`
	// Do not display application icon to users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#hide_web AppOauth#hide_web}
	HideWeb interface{} `field:"optional" json:"hideWeb" yaml:"hideWeb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#id AppOauth#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// *Early Access Property*. Enable Federation Broker Mode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#implicit_assignment AppOauth#implicit_assignment}
	ImplicitAssignment interface{} `field:"optional" json:"implicitAssignment" yaml:"implicitAssignment"`
	// *Early Access Property*.
	//
	// Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#issuer_mode AppOauth#issuer_mode}
	IssuerMode *string `field:"optional" json:"issuerMode" yaml:"issuerMode"`
	// jwks block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#jwks AppOauth#jwks}
	Jwks interface{} `field:"optional" json:"jwks" yaml:"jwks"`
	// The type of Idp-Initiated login that the client supports, if any.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#login_mode AppOauth#login_mode}
	LoginMode *string `field:"optional" json:"loginMode" yaml:"loginMode"`
	// List of scopes to use for the request.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#login_scopes AppOauth#login_scopes}
	LoginScopes *[]*string `field:"optional" json:"loginScopes" yaml:"loginScopes"`
	// URI that initiates login.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#login_uri AppOauth#login_uri}
	LoginUri *string `field:"optional" json:"loginUri" yaml:"loginUri"`
	// Local path to logo of the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#logo AppOauth#logo}
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// URI that references a logo for the client.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#logo_uri AppOauth#logo_uri}
	LogoUri *string `field:"optional" json:"logoUri" yaml:"logoUri"`
	// This tells the provider not to persist the application's secret to state.
	//
	// If this is ever changes from true => false your app will be recreated.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#omit_secret AppOauth#omit_secret}
	OmitSecret interface{} `field:"optional" json:"omitSecret" yaml:"omitSecret"`
	// Require Proof Key for Code Exchange (PKCE) for additional verification key rotation mode. See: https://developer.okta.com/docs/reference/api/apps/#oauth-credential-object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#pkce_required AppOauth#pkce_required}
	PkceRequired interface{} `field:"optional" json:"pkceRequired" yaml:"pkceRequired"`
	// URI to web page providing client policy document.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#policy_uri AppOauth#policy_uri}
	PolicyUri *string `field:"optional" json:"policyUri" yaml:"policyUri"`
	// List of URIs for redirection after logout. Note: see okta_app_oauth_post_logout_redirect_uri for appending to this list in a decentralized way.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#post_logout_redirect_uris AppOauth#post_logout_redirect_uris}
	PostLogoutRedirectUris *[]*string `field:"optional" json:"postLogoutRedirectUris" yaml:"postLogoutRedirectUris"`
	// Custom JSON that represents an OAuth application's profile.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#profile AppOauth#profile}
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// List of URIs for use in the redirect-based flow.
	//
	// This is required for all application types except service. Note: see okta_app_oauth_redirect_uri for appending to this list in a decentralized way.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#redirect_uris AppOauth#redirect_uris}
	RedirectUris *[]*string `field:"optional" json:"redirectUris" yaml:"redirectUris"`
	// *Early Access Property* Grace period for token rotation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#refresh_token_leeway AppOauth#refresh_token_leeway}
	RefreshTokenLeeway *float64 `field:"optional" json:"refreshTokenLeeway" yaml:"refreshTokenLeeway"`
	// *Early Access Property* Refresh token rotation behavior.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#refresh_token_rotation AppOauth#refresh_token_rotation}
	RefreshTokenRotation *string `field:"optional" json:"refreshTokenRotation" yaml:"refreshTokenRotation"`
	// List of OAuth 2.0 response type strings.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#response_types AppOauth#response_types}
	ResponseTypes *[]*string `field:"optional" json:"responseTypes" yaml:"responseTypes"`
	// Ignore groups sync. This is a temporary solution until 'groups' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#skip_groups AppOauth#skip_groups}
	SkipGroups interface{} `field:"optional" json:"skipGroups" yaml:"skipGroups"`
	// Ignore users sync. This is a temporary solution until 'users' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#skip_users AppOauth#skip_users}
	SkipUsers interface{} `field:"optional" json:"skipUsers" yaml:"skipUsers"`
	// Status of application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#status AppOauth#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#timeouts AppOauth#timeouts}
	Timeouts *AppOauthTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Requested authentication method for the token endpoint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#token_endpoint_auth_method AppOauth#token_endpoint_auth_method}
	TokenEndpointAuthMethod *string `field:"optional" json:"tokenEndpointAuthMethod" yaml:"tokenEndpointAuthMethod"`
	// URI to web page providing client tos (terms of service).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#tos_uri AppOauth#tos_uri}
	TosUri *string `field:"optional" json:"tosUri" yaml:"tosUri"`
	// Username template.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#user_name_template AppOauth#user_name_template}
	UserNameTemplate *string `field:"optional" json:"userNameTemplate" yaml:"userNameTemplate"`
	// Push username on update.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#user_name_template_push_status AppOauth#user_name_template_push_status}
	UserNameTemplatePushStatus *string `field:"optional" json:"userNameTemplatePushStatus" yaml:"userNameTemplatePushStatus"`
	// Username template suffix.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#user_name_template_suffix AppOauth#user_name_template_suffix}
	UserNameTemplateSuffix *string `field:"optional" json:"userNameTemplateSuffix" yaml:"userNameTemplateSuffix"`
	// Username template type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#user_name_template_type AppOauth#user_name_template_type}
	UserNameTemplateType *string `field:"optional" json:"userNameTemplateType" yaml:"userNameTemplateType"`
	// users block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#users AppOauth#users}
	Users interface{} `field:"optional" json:"users" yaml:"users"`
	// *Early Access Property*. Indicates if the client is allowed to use wildcard matching of redirect_uris.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_oauth#wildcard_redirect AppOauth#wildcard_redirect}
	WildcardRedirect *string `field:"optional" json:"wildcardRedirect" yaml:"wildcardRedirect"`
}

