package samlapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SamlAppConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#label SamlApp#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Custom error page URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#accessibility_error_redirect_url SamlApp#accessibility_error_redirect_url}
	AccessibilityErrorRedirectUrl *string `field:"optional" json:"accessibilityErrorRedirectUrl" yaml:"accessibilityErrorRedirectUrl"`
	// Custom login page URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#accessibility_login_redirect_url SamlApp#accessibility_login_redirect_url}
	AccessibilityLoginRedirectUrl *string `field:"optional" json:"accessibilityLoginRedirectUrl" yaml:"accessibilityLoginRedirectUrl"`
	// Enable self service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#accessibility_self_service SamlApp#accessibility_self_service}
	AccessibilitySelfService interface{} `field:"optional" json:"accessibilitySelfService" yaml:"accessibilitySelfService"`
	// List of ACS endpoints for this SAML application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#acs_endpoints SamlApp#acs_endpoints}
	AcsEndpoints *[]*string `field:"optional" json:"acsEndpoints" yaml:"acsEndpoints"`
	// Application notes for admins.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#admin_note SamlApp#admin_note}
	AdminNote *string `field:"optional" json:"adminNote" yaml:"adminNote"`
	// Displays specific appLinks for the app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#app_links_json SamlApp#app_links_json}
	AppLinksJson *string `field:"optional" json:"appLinksJson" yaml:"appLinksJson"`
	// Application settings in JSON format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#app_settings_json SamlApp#app_settings_json}
	AppSettingsJson *string `field:"optional" json:"appSettingsJson" yaml:"appSettingsJson"`
	// Determines whether the SAML assertion is digitally signed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#assertion_signed SamlApp#assertion_signed}
	AssertionSigned interface{} `field:"optional" json:"assertionSigned" yaml:"assertionSigned"`
	// attribute_statements block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#attribute_statements SamlApp#attribute_statements}
	AttributeStatements interface{} `field:"optional" json:"attributeStatements" yaml:"attributeStatements"`
	// Audience Restriction.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#audience SamlApp#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// Id of this apps authentication policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#authentication_policy SamlApp#authentication_policy}
	AuthenticationPolicy *string `field:"optional" json:"authenticationPolicy" yaml:"authenticationPolicy"`
	// Identifies the SAML authentication context class for the assertion’s authentication statement.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#authn_context_class_ref SamlApp#authn_context_class_ref}
	AuthnContextClassRef *string `field:"optional" json:"authnContextClassRef" yaml:"authnContextClassRef"`
	// Display auto submit toolbar.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#auto_submit_toolbar SamlApp#auto_submit_toolbar}
	AutoSubmitToolbar interface{} `field:"optional" json:"autoSubmitToolbar" yaml:"autoSubmitToolbar"`
	// Identifies a specific application resource in an IDP initiated SSO scenario.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#default_relay_state SamlApp#default_relay_state}
	DefaultRelayState *string `field:"optional" json:"defaultRelayState" yaml:"defaultRelayState"`
	// Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#destination SamlApp#destination}
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Determines the digest algorithm used to digitally sign the SAML assertion and response.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#digest_algorithm SamlApp#digest_algorithm}
	DigestAlgorithm *string `field:"optional" json:"digestAlgorithm" yaml:"digestAlgorithm"`
	// Application notes for end users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#enduser_note SamlApp#enduser_note}
	EnduserNote *string `field:"optional" json:"enduserNote" yaml:"enduserNote"`
	// Groups associated with the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#groups SamlApp#groups}
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Do not display application icon on mobile app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#hide_ios SamlApp#hide_ios}
	HideIos interface{} `field:"optional" json:"hideIos" yaml:"hideIos"`
	// Do not display application icon to users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#hide_web SamlApp#hide_web}
	HideWeb interface{} `field:"optional" json:"hideWeb" yaml:"hideWeb"`
	// Prompt user to re-authenticate if SP asks for it.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#honor_force_authn SamlApp#honor_force_authn}
	HonorForceAuthn interface{} `field:"optional" json:"honorForceAuthn" yaml:"honorForceAuthn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#id SamlApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// SAML issuer ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#idp_issuer SamlApp#idp_issuer}
	IdpIssuer *string `field:"optional" json:"idpIssuer" yaml:"idpIssuer"`
	// *Early Access Property*. Enable Federation Broker Mode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#implicit_assignment SamlApp#implicit_assignment}
	ImplicitAssignment interface{} `field:"optional" json:"implicitAssignment" yaml:"implicitAssignment"`
	// Saml Inline Hook setting.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#inline_hook_id SamlApp#inline_hook_id}
	InlineHookId *string `field:"optional" json:"inlineHookId" yaml:"inlineHookId"`
	// Certificate name. This modulates the rotation of keys. New name == new key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#key_name SamlApp#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Number of years the certificate is valid.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#key_years_valid SamlApp#key_years_valid}
	KeyYearsValid *float64 `field:"optional" json:"keyYearsValid" yaml:"keyYearsValid"`
	// Local path to logo of the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#logo SamlApp#logo}
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Name of preexisting SAML application. For instance 'slack'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#preconfigured_app SamlApp#preconfigured_app}
	PreconfiguredApp *string `field:"optional" json:"preconfiguredApp" yaml:"preconfiguredApp"`
	// The location where the app may present the SAML assertion.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#recipient SamlApp#recipient}
	Recipient *string `field:"optional" json:"recipient" yaml:"recipient"`
	// Denotes whether the request is compressed or not.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#request_compressed SamlApp#request_compressed}
	RequestCompressed interface{} `field:"optional" json:"requestCompressed" yaml:"requestCompressed"`
	// Determines whether the SAML auth response message is digitally signed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#response_signed SamlApp#response_signed}
	ResponseSigned interface{} `field:"optional" json:"responseSigned" yaml:"responseSigned"`
	// SAML Signed Request enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#saml_signed_request_enabled SamlApp#saml_signed_request_enabled}
	SamlSignedRequestEnabled interface{} `field:"optional" json:"samlSignedRequestEnabled" yaml:"samlSignedRequestEnabled"`
	// SAML version for the app's sign-on mode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#saml_version SamlApp#saml_version}
	SamlVersion *string `field:"optional" json:"samlVersion" yaml:"samlVersion"`
	// Signature algorithm used ot digitally sign the assertion and response.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#signature_algorithm SamlApp#signature_algorithm}
	SignatureAlgorithm *string `field:"optional" json:"signatureAlgorithm" yaml:"signatureAlgorithm"`
	// x509 encoded certificate that the Service Provider uses to sign Single Logout requests.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#single_logout_certificate SamlApp#single_logout_certificate}
	SingleLogoutCertificate *string `field:"optional" json:"singleLogoutCertificate" yaml:"singleLogoutCertificate"`
	// The issuer of the Service Provider that generates the Single Logout request.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#single_logout_issuer SamlApp#single_logout_issuer}
	SingleLogoutIssuer *string `field:"optional" json:"singleLogoutIssuer" yaml:"singleLogoutIssuer"`
	// The location where the logout response is sent.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#single_logout_url SamlApp#single_logout_url}
	SingleLogoutUrl *string `field:"optional" json:"singleLogoutUrl" yaml:"singleLogoutUrl"`
	// Ignore groups sync. This is a temporary solution until 'groups' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#skip_groups SamlApp#skip_groups}
	SkipGroups interface{} `field:"optional" json:"skipGroups" yaml:"skipGroups"`
	// Ignore users sync. This is a temporary solution until 'users' field is supported in all the app-like resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#skip_users SamlApp#skip_users}
	SkipUsers interface{} `field:"optional" json:"skipUsers" yaml:"skipUsers"`
	// SAML SP issuer ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#sp_issuer SamlApp#sp_issuer}
	SpIssuer *string `field:"optional" json:"spIssuer" yaml:"spIssuer"`
	// Single Sign On URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#sso_url SamlApp#sso_url}
	SsoUrl *string `field:"optional" json:"ssoUrl" yaml:"ssoUrl"`
	// Status of application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#status SamlApp#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Identifies the SAML processing rules.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#subject_name_id_format SamlApp#subject_name_id_format}
	SubjectNameIdFormat *string `field:"optional" json:"subjectNameIdFormat" yaml:"subjectNameIdFormat"`
	// Template for app user's username when a user is assigned to the app.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#subject_name_id_template SamlApp#subject_name_id_template}
	SubjectNameIdTemplate *string `field:"optional" json:"subjectNameIdTemplate" yaml:"subjectNameIdTemplate"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#timeouts SamlApp#timeouts}
	Timeouts *SamlAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Username template.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#user_name_template SamlApp#user_name_template}
	UserNameTemplate *string `field:"optional" json:"userNameTemplate" yaml:"userNameTemplate"`
	// Push username on update.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#user_name_template_push_status SamlApp#user_name_template_push_status}
	UserNameTemplatePushStatus *string `field:"optional" json:"userNameTemplatePushStatus" yaml:"userNameTemplatePushStatus"`
	// Username template suffix.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#user_name_template_suffix SamlApp#user_name_template_suffix}
	UserNameTemplateSuffix *string `field:"optional" json:"userNameTemplateSuffix" yaml:"userNameTemplateSuffix"`
	// Username template type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#user_name_template_type SamlApp#user_name_template_type}
	UserNameTemplateType *string `field:"optional" json:"userNameTemplateType" yaml:"userNameTemplateType"`
	// users block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_app#users SamlApp#users}
	Users interface{} `field:"optional" json:"users" yaml:"users"`
}

