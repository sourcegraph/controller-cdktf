package appsaml

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppSamlConfig struct {
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
	// The Application's display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#label AppSaml#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Custom error page URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#accessibility_error_redirect_url AppSaml#accessibility_error_redirect_url}
	AccessibilityErrorRedirectUrl *string `field:"optional" json:"accessibilityErrorRedirectUrl" yaml:"accessibilityErrorRedirectUrl"`
	// Custom login page URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#accessibility_login_redirect_url AppSaml#accessibility_login_redirect_url}
	AccessibilityLoginRedirectUrl *string `field:"optional" json:"accessibilityLoginRedirectUrl" yaml:"accessibilityLoginRedirectUrl"`
	// Enable self service. Default is `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#accessibility_self_service AppSaml#accessibility_self_service}
	AccessibilitySelfService interface{} `field:"optional" json:"accessibilitySelfService" yaml:"accessibilitySelfService"`
	// An array of ACS endpoints. You can configure a maximum of 100 endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#acs_endpoints AppSaml#acs_endpoints}
	AcsEndpoints *[]*string `field:"optional" json:"acsEndpoints" yaml:"acsEndpoints"`
	// Application notes for admins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#admin_note AppSaml#admin_note}
	AdminNote *string `field:"optional" json:"adminNote" yaml:"adminNote"`
	// Displays specific appLinks for the app. The value for each application link should be boolean.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#app_links_json AppSaml#app_links_json}
	AppLinksJson *string `field:"optional" json:"appLinksJson" yaml:"appLinksJson"`
	// Application settings in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#app_settings_json AppSaml#app_settings_json}
	AppSettingsJson *string `field:"optional" json:"appSettingsJson" yaml:"appSettingsJson"`
	// Determines whether the SAML assertion is digitally signed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#assertion_signed AppSaml#assertion_signed}
	AssertionSigned interface{} `field:"optional" json:"assertionSigned" yaml:"assertionSigned"`
	// attribute_statements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#attribute_statements AppSaml#attribute_statements}
	AttributeStatements interface{} `field:"optional" json:"attributeStatements" yaml:"attributeStatements"`
	// Audience Restriction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#audience AppSaml#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// The ID of the associated `app_signon_policy`.
	//
	// If this property is removed from the application the `default` sign-on-policy will be associated with this application.y
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#authentication_policy AppSaml#authentication_policy}
	AuthenticationPolicy *string `field:"optional" json:"authenticationPolicy" yaml:"authenticationPolicy"`
	// Identifies the SAML authentication context class for the assertion’s authentication statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#authn_context_class_ref AppSaml#authn_context_class_ref}
	AuthnContextClassRef *string `field:"optional" json:"authnContextClassRef" yaml:"authnContextClassRef"`
	// Display auto submit toolbar. Default is: `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#auto_submit_toolbar AppSaml#auto_submit_toolbar}
	AutoSubmitToolbar interface{} `field:"optional" json:"autoSubmitToolbar" yaml:"autoSubmitToolbar"`
	// Identifies a specific application resource in an IDP initiated SSO scenario.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#default_relay_state AppSaml#default_relay_state}
	DefaultRelayState *string `field:"optional" json:"defaultRelayState" yaml:"defaultRelayState"`
	// Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#destination AppSaml#destination}
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Determines the digest algorithm used to digitally sign the SAML assertion and response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#digest_algorithm AppSaml#digest_algorithm}
	DigestAlgorithm *string `field:"optional" json:"digestAlgorithm" yaml:"digestAlgorithm"`
	// Application notes for end users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#enduser_note AppSaml#enduser_note}
	EnduserNote *string `field:"optional" json:"enduserNote" yaml:"enduserNote"`
	// Do not display application icon on mobile app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#hide_ios AppSaml#hide_ios}
	HideIos interface{} `field:"optional" json:"hideIos" yaml:"hideIos"`
	// Do not display application icon to users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#hide_web AppSaml#hide_web}
	HideWeb interface{} `field:"optional" json:"hideWeb" yaml:"hideWeb"`
	// Prompt user to re-authenticate if SP asks for it. Default is: `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#honor_force_authn AppSaml#honor_force_authn}
	HonorForceAuthn interface{} `field:"optional" json:"honorForceAuthn" yaml:"honorForceAuthn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#id AppSaml#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// SAML issuer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#idp_issuer AppSaml#idp_issuer}
	IdpIssuer *string `field:"optional" json:"idpIssuer" yaml:"idpIssuer"`
	// *Early Access Property*. Enable Federation Broker Mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#implicit_assignment AppSaml#implicit_assignment}
	ImplicitAssignment interface{} `field:"optional" json:"implicitAssignment" yaml:"implicitAssignment"`
	// Saml Inline Hook setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#inline_hook_id AppSaml#inline_hook_id}
	InlineHookId *string `field:"optional" json:"inlineHookId" yaml:"inlineHookId"`
	// Certificate name. This modulates the rotation of keys. New name == new key. Required to be set with `key_years_valid`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#key_name AppSaml#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Number of years the certificate is valid (2 - 10 years).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#key_years_valid AppSaml#key_years_valid}
	KeyYearsValid *float64 `field:"optional" json:"keyYearsValid" yaml:"keyYearsValid"`
	// Local file path to the logo.
	//
	// The file must be in PNG, JPG, or GIF format, and less than 1 MB in size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#logo AppSaml#logo}
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Name of application from the Okta Integration Network.
	//
	// For instance 'slack'. If not included a custom app will be created.  If not provided the following arguments are required:
	// 'sso_url'
	// 'recipient'
	// 'destination'
	// 'audience'
	// 'subject_name_id_template'
	// 'subject_name_id_format'
	// 'signature_algorithm'
	// 'digest_algorithm'
	// 'authn_context_class_ref'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#preconfigured_app AppSaml#preconfigured_app}
	PreconfiguredApp *string `field:"optional" json:"preconfiguredApp" yaml:"preconfiguredApp"`
	// The location where the app may present the SAML assertion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#recipient AppSaml#recipient}
	Recipient *string `field:"optional" json:"recipient" yaml:"recipient"`
	// Denotes whether the request is compressed or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#request_compressed AppSaml#request_compressed}
	RequestCompressed interface{} `field:"optional" json:"requestCompressed" yaml:"requestCompressed"`
	// Determines whether the SAML auth response message is digitally signed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#response_signed AppSaml#response_signed}
	ResponseSigned interface{} `field:"optional" json:"responseSigned" yaml:"responseSigned"`
	// SAML Signed Request enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#saml_signed_request_enabled AppSaml#saml_signed_request_enabled}
	SamlSignedRequestEnabled interface{} `field:"optional" json:"samlSignedRequestEnabled" yaml:"samlSignedRequestEnabled"`
	// SAML version for the app's sign-on mode. Valid values are: `2.0` or `1.1`. Default is `2.0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#saml_version AppSaml#saml_version}
	SamlVersion *string `field:"optional" json:"samlVersion" yaml:"samlVersion"`
	// Signature algorithm used to digitally sign the assertion and response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#signature_algorithm AppSaml#signature_algorithm}
	SignatureAlgorithm *string `field:"optional" json:"signatureAlgorithm" yaml:"signatureAlgorithm"`
	// x509 encoded certificate that the Service Provider uses to sign Single Logout requests.
	//
	// Note: should be provided without `-----BEGIN CERTIFICATE-----` and `-----END CERTIFICATE-----`, see [official documentation](https://developer.okta.com/docs/reference/api/apps/#service-provider-certificate).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#single_logout_certificate AppSaml#single_logout_certificate}
	SingleLogoutCertificate *string `field:"optional" json:"singleLogoutCertificate" yaml:"singleLogoutCertificate"`
	// The issuer of the Service Provider that generates the Single Logout request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#single_logout_issuer AppSaml#single_logout_issuer}
	SingleLogoutIssuer *string `field:"optional" json:"singleLogoutIssuer" yaml:"singleLogoutIssuer"`
	// The location where the logout response is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#single_logout_url AppSaml#single_logout_url}
	SingleLogoutUrl *string `field:"optional" json:"singleLogoutUrl" yaml:"singleLogoutUrl"`
	// SAML SP issuer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#sp_issuer AppSaml#sp_issuer}
	SpIssuer *string `field:"optional" json:"spIssuer" yaml:"spIssuer"`
	// Single Sign On URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#sso_url AppSaml#sso_url}
	SsoUrl *string `field:"optional" json:"ssoUrl" yaml:"ssoUrl"`
	// Status of application. By default, it is `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#status AppSaml#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Identifies the SAML processing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#subject_name_id_format AppSaml#subject_name_id_format}
	SubjectNameIdFormat *string `field:"optional" json:"subjectNameIdFormat" yaml:"subjectNameIdFormat"`
	// Template for app user's username when a user is assigned to the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#subject_name_id_template AppSaml#subject_name_id_template}
	SubjectNameIdTemplate *string `field:"optional" json:"subjectNameIdTemplate" yaml:"subjectNameIdTemplate"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#timeouts AppSaml#timeouts}
	Timeouts *AppSamlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Username template. Default: `${source.login}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#user_name_template AppSaml#user_name_template}
	UserNameTemplate *string `field:"optional" json:"userNameTemplate" yaml:"userNameTemplate"`
	// Push username on update. Valid values: `PUSH` and `DONT_PUSH`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#user_name_template_push_status AppSaml#user_name_template_push_status}
	UserNameTemplatePushStatus *string `field:"optional" json:"userNameTemplatePushStatus" yaml:"userNameTemplatePushStatus"`
	// Username template suffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#user_name_template_suffix AppSaml#user_name_template_suffix}
	UserNameTemplateSuffix *string `field:"optional" json:"userNameTemplateSuffix" yaml:"userNameTemplateSuffix"`
	// Username template type. Default: `BUILT_IN`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_saml#user_name_template_type AppSaml#user_name_template_type}
	UserNameTemplateType *string `field:"optional" json:"userNameTemplateType" yaml:"userNameTemplateType"`
}

