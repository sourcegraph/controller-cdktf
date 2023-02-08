package idpsaml

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdpSamlConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#issuer IdpSaml#issuer}.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#kid IdpSaml#kid}.
	Kid *string `field:"required" json:"kid" yaml:"kid"`
	// Name of the IdP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#name IdpSaml#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#sso_url IdpSaml#sso_url}.
	SsoUrl *string `field:"required" json:"ssoUrl" yaml:"ssoUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#account_link_action IdpSaml#account_link_action}.
	AccountLinkAction *string `field:"optional" json:"accountLinkAction" yaml:"accountLinkAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#account_link_group_include IdpSaml#account_link_group_include}.
	AccountLinkGroupInclude *[]*string `field:"optional" json:"accountLinkGroupInclude" yaml:"accountLinkGroupInclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#acs_binding IdpSaml#acs_binding}.
	AcsBinding *string `field:"optional" json:"acsBinding" yaml:"acsBinding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#acs_type IdpSaml#acs_type}.
	AcsType *string `field:"optional" json:"acsType" yaml:"acsType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#deprovisioned_action IdpSaml#deprovisioned_action}.
	DeprovisionedAction *string `field:"optional" json:"deprovisionedAction" yaml:"deprovisionedAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#groups_action IdpSaml#groups_action}.
	GroupsAction *string `field:"optional" json:"groupsAction" yaml:"groupsAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#groups_assignment IdpSaml#groups_assignment}.
	GroupsAssignment *[]*string `field:"optional" json:"groupsAssignment" yaml:"groupsAssignment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#groups_attribute IdpSaml#groups_attribute}.
	GroupsAttribute *string `field:"optional" json:"groupsAttribute" yaml:"groupsAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#groups_filter IdpSaml#groups_filter}.
	GroupsFilter *[]*string `field:"optional" json:"groupsFilter" yaml:"groupsFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#id IdpSaml#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#issuer_mode IdpSaml#issuer_mode}
	IssuerMode *string `field:"optional" json:"issuerMode" yaml:"issuerMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#max_clock_skew IdpSaml#max_clock_skew}.
	MaxClockSkew *float64 `field:"optional" json:"maxClockSkew" yaml:"maxClockSkew"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#name_format IdpSaml#name_format}.
	NameFormat *string `field:"optional" json:"nameFormat" yaml:"nameFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#profile_master IdpSaml#profile_master}.
	ProfileMaster interface{} `field:"optional" json:"profileMaster" yaml:"profileMaster"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#provisioning_action IdpSaml#provisioning_action}.
	ProvisioningAction *string `field:"optional" json:"provisioningAction" yaml:"provisioningAction"`
	// The XML digital Signature Algorithm used when signing an <AuthnRequest> message.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#request_signature_algorithm IdpSaml#request_signature_algorithm}
	RequestSignatureAlgorithm *string `field:"optional" json:"requestSignatureAlgorithm" yaml:"requestSignatureAlgorithm"`
	// Specifies whether to digitally sign <AuthnRequest> messages to the IdP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#request_signature_scope IdpSaml#request_signature_scope}
	RequestSignatureScope *string `field:"optional" json:"requestSignatureScope" yaml:"requestSignatureScope"`
	// The minimum XML digital Signature Algorithm allowed when verifying a <SAMLResponse> message or <Assertion> element.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#response_signature_algorithm IdpSaml#response_signature_algorithm}
	ResponseSignatureAlgorithm *string `field:"optional" json:"responseSignatureAlgorithm" yaml:"responseSignatureAlgorithm"`
	// Specifies whether to verify a <SAMLResponse> message or <Assertion> element XML digital signature.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#response_signature_scope IdpSaml#response_signature_scope}
	ResponseSignatureScope *string `field:"optional" json:"responseSignatureScope" yaml:"responseSignatureScope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#sso_binding IdpSaml#sso_binding}.
	SsoBinding *string `field:"optional" json:"ssoBinding" yaml:"ssoBinding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#sso_destination IdpSaml#sso_destination}.
	SsoDestination *string `field:"optional" json:"ssoDestination" yaml:"ssoDestination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#status IdpSaml#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#subject_filter IdpSaml#subject_filter}.
	SubjectFilter *string `field:"optional" json:"subjectFilter" yaml:"subjectFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#subject_format IdpSaml#subject_format}.
	SubjectFormat *[]*string `field:"optional" json:"subjectFormat" yaml:"subjectFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#subject_match_attribute IdpSaml#subject_match_attribute}.
	SubjectMatchAttribute *string `field:"optional" json:"subjectMatchAttribute" yaml:"subjectMatchAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#subject_match_type IdpSaml#subject_match_type}.
	SubjectMatchType *string `field:"optional" json:"subjectMatchType" yaml:"subjectMatchType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#suspended_action IdpSaml#suspended_action}.
	SuspendedAction *string `field:"optional" json:"suspendedAction" yaml:"suspendedAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_saml#username_template IdpSaml#username_template}.
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

