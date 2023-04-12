package idpsocial

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdpSocialConfig struct {
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
	// Name of the IdP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#name IdpSocial#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#scopes IdpSocial#scopes}.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Identity Provider Types: https://developer.okta.com/docs/reference/api/idps/#identity-provider-type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#type IdpSocial#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#account_link_action IdpSocial#account_link_action}.
	AccountLinkAction *string `field:"optional" json:"accountLinkAction" yaml:"accountLinkAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#account_link_group_include IdpSocial#account_link_group_include}.
	AccountLinkGroupInclude *[]*string `field:"optional" json:"accountLinkGroupInclude" yaml:"accountLinkGroupInclude"`
	// The Key ID that you obtained from Apple when you created the private key for the client.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#apple_kid IdpSocial#apple_kid}
	AppleKid *string `field:"optional" json:"appleKid" yaml:"appleKid"`
	// The PKCS #8 encoded private key that you created for the client and downloaded from Apple.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#apple_private_key IdpSocial#apple_private_key}
	ApplePrivateKey *string `field:"optional" json:"applePrivateKey" yaml:"applePrivateKey"`
	// The Team ID associated with your Apple developer account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#apple_team_id IdpSocial#apple_team_id}
	AppleTeamId *string `field:"optional" json:"appleTeamId" yaml:"appleTeamId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#client_id IdpSocial#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#client_secret IdpSocial#client_secret}.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#deprovisioned_action IdpSocial#deprovisioned_action}.
	DeprovisionedAction *string `field:"optional" json:"deprovisionedAction" yaml:"deprovisionedAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#groups_action IdpSocial#groups_action}.
	GroupsAction *string `field:"optional" json:"groupsAction" yaml:"groupsAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#groups_assignment IdpSocial#groups_assignment}.
	GroupsAssignment *[]*string `field:"optional" json:"groupsAssignment" yaml:"groupsAssignment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#groups_attribute IdpSocial#groups_attribute}.
	GroupsAttribute *string `field:"optional" json:"groupsAttribute" yaml:"groupsAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#groups_filter IdpSocial#groups_filter}.
	GroupsFilter *[]*string `field:"optional" json:"groupsFilter" yaml:"groupsFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#id IdpSocial#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#issuer_mode IdpSocial#issuer_mode}
	IssuerMode *string `field:"optional" json:"issuerMode" yaml:"issuerMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#match_attribute IdpSocial#match_attribute}.
	MatchAttribute *string `field:"optional" json:"matchAttribute" yaml:"matchAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#match_type IdpSocial#match_type}.
	MatchType *string `field:"optional" json:"matchType" yaml:"matchType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#max_clock_skew IdpSocial#max_clock_skew}.
	MaxClockSkew *float64 `field:"optional" json:"maxClockSkew" yaml:"maxClockSkew"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#profile_master IdpSocial#profile_master}.
	ProfileMaster interface{} `field:"optional" json:"profileMaster" yaml:"profileMaster"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#protocol_type IdpSocial#protocol_type}.
	ProtocolType *string `field:"optional" json:"protocolType" yaml:"protocolType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#provisioning_action IdpSocial#provisioning_action}.
	ProvisioningAction *string `field:"optional" json:"provisioningAction" yaml:"provisioningAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#status IdpSocial#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#subject_match_attribute IdpSocial#subject_match_attribute}.
	SubjectMatchAttribute *string `field:"optional" json:"subjectMatchAttribute" yaml:"subjectMatchAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#subject_match_type IdpSocial#subject_match_type}.
	SubjectMatchType *string `field:"optional" json:"subjectMatchType" yaml:"subjectMatchType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#suspended_action IdpSocial#suspended_action}.
	SuspendedAction *string `field:"optional" json:"suspendedAction" yaml:"suspendedAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/idp_social#username_template IdpSocial#username_template}.
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

