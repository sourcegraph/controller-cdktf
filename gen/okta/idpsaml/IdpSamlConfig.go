package idpsaml

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdpSamlConfig struct {
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
	// URI that identifies the issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#issuer IdpSaml#issuer}
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// The ID of the signing key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#kid IdpSaml#kid}
	Kid *string `field:"required" json:"kid" yaml:"kid"`
	// Name of the IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#name IdpSaml#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// URL of binding-specific endpoint to send an AuthnRequest message to IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#sso_url IdpSaml#sso_url}
	SsoUrl *string `field:"required" json:"ssoUrl" yaml:"ssoUrl"`
	// Specifies the account linking action for an IdP user. Default: `AUTO`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#account_link_action IdpSaml#account_link_action}
	AccountLinkAction *string `field:"optional" json:"accountLinkAction" yaml:"accountLinkAction"`
	// Group memberships to determine link candidates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#account_link_group_include IdpSaml#account_link_group_include}
	AccountLinkGroupInclude *[]*string `field:"optional" json:"accountLinkGroupInclude" yaml:"accountLinkGroupInclude"`
	// The type of ACS. It can be `INSTANCE` or `ORG`. Default: `INSTANCE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#acs_type IdpSaml#acs_type}
	AcsType *string `field:"optional" json:"acsType" yaml:"acsType"`
	// Action for a previously deprovisioned IdP user during authentication. Can be `NONE` or `REACTIVATE`. Default: `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#deprovisioned_action IdpSaml#deprovisioned_action}
	DeprovisionedAction *string `field:"optional" json:"deprovisionedAction" yaml:"deprovisionedAction"`
	// Provisioning action for IdP user's group memberships. It can be `NONE`, `SYNC`, `APPEND`, or `ASSIGN`. Default: `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#groups_action IdpSaml#groups_action}
	GroupsAction *string `field:"optional" json:"groupsAction" yaml:"groupsAction"`
	// List of Okta Group IDs to add an IdP user as a member with the `ASSIGN` `groups_action`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#groups_assignment IdpSaml#groups_assignment}
	GroupsAssignment *[]*string `field:"optional" json:"groupsAssignment" yaml:"groupsAssignment"`
	// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#groups_attribute IdpSaml#groups_attribute}
	GroupsAttribute *string `field:"optional" json:"groupsAttribute" yaml:"groupsAttribute"`
	// Whitelist of Okta Group identifiers that are allowed for the `APPEND` or `SYNC` `groups_action`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#groups_filter IdpSaml#groups_filter}
	GroupsFilter *[]*string `field:"optional" json:"groupsFilter" yaml:"groupsFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#id IdpSaml#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#issuer_mode IdpSaml#issuer_mode}
	IssuerMode *string `field:"optional" json:"issuerMode" yaml:"issuerMode"`
	// Maximum allowable clock-skew when processing messages from the IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#max_clock_skew IdpSaml#max_clock_skew}
	MaxClockSkew *float64 `field:"optional" json:"maxClockSkew" yaml:"maxClockSkew"`
	// The name identifier format to use. By default `urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#name_format IdpSaml#name_format}
	NameFormat *string `field:"optional" json:"nameFormat" yaml:"nameFormat"`
	// Determines if the IdP should act as a source of truth for user profile attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#profile_master IdpSaml#profile_master}
	ProfileMaster interface{} `field:"optional" json:"profileMaster" yaml:"profileMaster"`
	// Provisioning action for an IdP user during authentication. Default: `AUTO`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#provisioning_action IdpSaml#provisioning_action}
	ProvisioningAction *string `field:"optional" json:"provisioningAction" yaml:"provisioningAction"`
	// The XML digital Signature Algorithm used when signing an `AuthnRequest` message. It can be `SHA-256` or `SHA-1`. Default: `SHA-256`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#request_signature_algorithm IdpSaml#request_signature_algorithm}
	RequestSignatureAlgorithm *string `field:"optional" json:"requestSignatureAlgorithm" yaml:"requestSignatureAlgorithm"`
	// Specifies whether to digitally sign an AuthnRequest messages to the IdP. It can be `REQUEST` or `NONE`. Default: `REQUEST`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#request_signature_scope IdpSaml#request_signature_scope}
	RequestSignatureScope *string `field:"optional" json:"requestSignatureScope" yaml:"requestSignatureScope"`
	// The minimum XML digital signature algorithm allowed when verifying a `SAMLResponse` message or Assertion element.
	//
	// It can be `SHA-256` or `SHA-1`. Default: `SHA-256`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#response_signature_algorithm IdpSaml#response_signature_algorithm}
	ResponseSignatureAlgorithm *string `field:"optional" json:"responseSignatureAlgorithm" yaml:"responseSignatureAlgorithm"`
	// Specifies whether to verify a `SAMLResponse` message or Assertion element XML digital signature.
	//
	// It can be `RESPONSE`, `ASSERTION`, or `ANY`. Default: `ANY`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#response_signature_scope IdpSaml#response_signature_scope}
	ResponseSignatureScope *string `field:"optional" json:"responseSignatureScope" yaml:"responseSignatureScope"`
	// The method of making an SSO request. It can be set to `HTTP-POST` or `HTTP-REDIRECT`. Default: `HTTP-POST`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#sso_binding IdpSaml#sso_binding}
	SsoBinding *string `field:"optional" json:"ssoBinding" yaml:"ssoBinding"`
	// URI reference indicating the address to which the AuthnRequest message is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#sso_destination IdpSaml#sso_destination}
	SsoDestination *string `field:"optional" json:"ssoDestination" yaml:"ssoDestination"`
	// Default to `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#status IdpSaml#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Optional regular expression pattern used to filter untrusted IdP usernames.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#subject_filter IdpSaml#subject_filter}
	SubjectFilter *string `field:"optional" json:"subjectFilter" yaml:"subjectFilter"`
	// The name format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#subject_format IdpSaml#subject_format}
	SubjectFormat *[]*string `field:"optional" json:"subjectFormat" yaml:"subjectFormat"`
	// Okta user profile attribute for matching transformed IdP username. Only for matchType `CUSTOM_ATTRIBUTE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#subject_match_attribute IdpSaml#subject_match_attribute}
	SubjectMatchAttribute *string `field:"optional" json:"subjectMatchAttribute" yaml:"subjectMatchAttribute"`
	// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username.
	//
	// By default, it is set to `USERNAME`. It can be set to `USERNAME`, `EMAIL`, `USERNAME_OR_EMAIL` or `CUSTOM_ATTRIBUTE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#subject_match_type IdpSaml#subject_match_type}
	SubjectMatchType *string `field:"optional" json:"subjectMatchType" yaml:"subjectMatchType"`
	// Action for a previously suspended IdP user during authentication. Can be `NONE` or `REACTIVATE`. Default: `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#suspended_action IdpSaml#suspended_action}
	SuspendedAction *string `field:"optional" json:"suspendedAction" yaml:"suspendedAction"`
	// Okta EL Expression to generate or transform a unique username for the IdP user. Default: `idpuser.email`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_saml#username_template IdpSaml#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

