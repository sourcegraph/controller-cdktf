package samlidp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SamlIdpConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#issuer SamlIdp#issuer}.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#kid SamlIdp#kid}.
	Kid *string `field:"required" json:"kid" yaml:"kid"`
	// Name of the IdP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#name SamlIdp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#sso_url SamlIdp#sso_url}.
	SsoUrl *string `field:"required" json:"ssoUrl" yaml:"ssoUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#account_link_action SamlIdp#account_link_action}.
	AccountLinkAction *string `field:"optional" json:"accountLinkAction" yaml:"accountLinkAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#account_link_group_include SamlIdp#account_link_group_include}.
	AccountLinkGroupInclude *[]*string `field:"optional" json:"accountLinkGroupInclude" yaml:"accountLinkGroupInclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#acs_binding SamlIdp#acs_binding}.
	AcsBinding *string `field:"optional" json:"acsBinding" yaml:"acsBinding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#acs_type SamlIdp#acs_type}.
	AcsType *string `field:"optional" json:"acsType" yaml:"acsType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#deprovisioned_action SamlIdp#deprovisioned_action}.
	DeprovisionedAction *string `field:"optional" json:"deprovisionedAction" yaml:"deprovisionedAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#groups_action SamlIdp#groups_action}.
	GroupsAction *string `field:"optional" json:"groupsAction" yaml:"groupsAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#groups_assignment SamlIdp#groups_assignment}.
	GroupsAssignment *[]*string `field:"optional" json:"groupsAssignment" yaml:"groupsAssignment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#groups_attribute SamlIdp#groups_attribute}.
	GroupsAttribute *string `field:"optional" json:"groupsAttribute" yaml:"groupsAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#groups_filter SamlIdp#groups_filter}.
	GroupsFilter *[]*string `field:"optional" json:"groupsFilter" yaml:"groupsFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#id SamlIdp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#issuer_mode SamlIdp#issuer_mode}
	IssuerMode *string `field:"optional" json:"issuerMode" yaml:"issuerMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#max_clock_skew SamlIdp#max_clock_skew}.
	MaxClockSkew *float64 `field:"optional" json:"maxClockSkew" yaml:"maxClockSkew"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#name_format SamlIdp#name_format}.
	NameFormat *string `field:"optional" json:"nameFormat" yaml:"nameFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#profile_master SamlIdp#profile_master}.
	ProfileMaster interface{} `field:"optional" json:"profileMaster" yaml:"profileMaster"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#provisioning_action SamlIdp#provisioning_action}.
	ProvisioningAction *string `field:"optional" json:"provisioningAction" yaml:"provisioningAction"`
	// The XML digital Signature Algorithm used when signing an <AuthnRequest> message.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#request_signature_algorithm SamlIdp#request_signature_algorithm}
	RequestSignatureAlgorithm *string `field:"optional" json:"requestSignatureAlgorithm" yaml:"requestSignatureAlgorithm"`
	// Specifies whether to digitally sign <AuthnRequest> messages to the IdP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#request_signature_scope SamlIdp#request_signature_scope}
	RequestSignatureScope *string `field:"optional" json:"requestSignatureScope" yaml:"requestSignatureScope"`
	// The minimum XML digital Signature Algorithm allowed when verifying a <SAMLResponse> message or <Assertion> element.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#response_signature_algorithm SamlIdp#response_signature_algorithm}
	ResponseSignatureAlgorithm *string `field:"optional" json:"responseSignatureAlgorithm" yaml:"responseSignatureAlgorithm"`
	// Specifies whether to verify a <SAMLResponse> message or <Assertion> element XML digital signature.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#response_signature_scope SamlIdp#response_signature_scope}
	ResponseSignatureScope *string `field:"optional" json:"responseSignatureScope" yaml:"responseSignatureScope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#sso_binding SamlIdp#sso_binding}.
	SsoBinding *string `field:"optional" json:"ssoBinding" yaml:"ssoBinding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#sso_destination SamlIdp#sso_destination}.
	SsoDestination *string `field:"optional" json:"ssoDestination" yaml:"ssoDestination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#status SamlIdp#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#subject_filter SamlIdp#subject_filter}.
	SubjectFilter *string `field:"optional" json:"subjectFilter" yaml:"subjectFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#subject_format SamlIdp#subject_format}.
	SubjectFormat *[]*string `field:"optional" json:"subjectFormat" yaml:"subjectFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#subject_match_attribute SamlIdp#subject_match_attribute}.
	SubjectMatchAttribute *string `field:"optional" json:"subjectMatchAttribute" yaml:"subjectMatchAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#subject_match_type SamlIdp#subject_match_type}.
	SubjectMatchType *string `field:"optional" json:"subjectMatchType" yaml:"subjectMatchType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#suspended_action SamlIdp#suspended_action}.
	SuspendedAction *string `field:"optional" json:"suspendedAction" yaml:"suspendedAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/saml_idp#username_template SamlIdp#username_template}.
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

