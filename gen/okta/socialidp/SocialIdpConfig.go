package socialidp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SocialIdpConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#name SocialIdp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#scopes SocialIdp#scopes}.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Identity Provider Types: https://developer.okta.com/docs/reference/api/idps/#identity-provider-type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#type SocialIdp#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#account_link_action SocialIdp#account_link_action}.
	AccountLinkAction *string `field:"optional" json:"accountLinkAction" yaml:"accountLinkAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#account_link_group_include SocialIdp#account_link_group_include}.
	AccountLinkGroupInclude *[]*string `field:"optional" json:"accountLinkGroupInclude" yaml:"accountLinkGroupInclude"`
	// The Key ID that you obtained from Apple when you created the private key for the client.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#apple_kid SocialIdp#apple_kid}
	AppleKid *string `field:"optional" json:"appleKid" yaml:"appleKid"`
	// The PKCS #8 encoded private key that you created for the client and downloaded from Apple.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#apple_private_key SocialIdp#apple_private_key}
	ApplePrivateKey *string `field:"optional" json:"applePrivateKey" yaml:"applePrivateKey"`
	// The Team ID associated with your Apple developer account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#apple_team_id SocialIdp#apple_team_id}
	AppleTeamId *string `field:"optional" json:"appleTeamId" yaml:"appleTeamId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#client_id SocialIdp#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#client_secret SocialIdp#client_secret}.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#deprovisioned_action SocialIdp#deprovisioned_action}.
	DeprovisionedAction *string `field:"optional" json:"deprovisionedAction" yaml:"deprovisionedAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#groups_action SocialIdp#groups_action}.
	GroupsAction *string `field:"optional" json:"groupsAction" yaml:"groupsAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#groups_assignment SocialIdp#groups_assignment}.
	GroupsAssignment *[]*string `field:"optional" json:"groupsAssignment" yaml:"groupsAssignment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#groups_attribute SocialIdp#groups_attribute}.
	GroupsAttribute *string `field:"optional" json:"groupsAttribute" yaml:"groupsAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#groups_filter SocialIdp#groups_filter}.
	GroupsFilter *[]*string `field:"optional" json:"groupsFilter" yaml:"groupsFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#id SocialIdp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#issuer_mode SocialIdp#issuer_mode}
	IssuerMode *string `field:"optional" json:"issuerMode" yaml:"issuerMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#match_attribute SocialIdp#match_attribute}.
	MatchAttribute *string `field:"optional" json:"matchAttribute" yaml:"matchAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#match_type SocialIdp#match_type}.
	MatchType *string `field:"optional" json:"matchType" yaml:"matchType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#max_clock_skew SocialIdp#max_clock_skew}.
	MaxClockSkew *float64 `field:"optional" json:"maxClockSkew" yaml:"maxClockSkew"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#profile_master SocialIdp#profile_master}.
	ProfileMaster interface{} `field:"optional" json:"profileMaster" yaml:"profileMaster"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#protocol_type SocialIdp#protocol_type}.
	ProtocolType *string `field:"optional" json:"protocolType" yaml:"protocolType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#provisioning_action SocialIdp#provisioning_action}.
	ProvisioningAction *string `field:"optional" json:"provisioningAction" yaml:"provisioningAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#status SocialIdp#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#subject_match_attribute SocialIdp#subject_match_attribute}.
	SubjectMatchAttribute *string `field:"optional" json:"subjectMatchAttribute" yaml:"subjectMatchAttribute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#subject_match_type SocialIdp#subject_match_type}.
	SubjectMatchType *string `field:"optional" json:"subjectMatchType" yaml:"subjectMatchType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#suspended_action SocialIdp#suspended_action}.
	SuspendedAction *string `field:"optional" json:"suspendedAction" yaml:"suspendedAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/social_idp#username_template SocialIdp#username_template}.
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

