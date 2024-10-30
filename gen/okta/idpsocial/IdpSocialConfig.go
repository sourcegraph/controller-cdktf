package idpsocial

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdpSocialConfig struct {
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
	// Name of the IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#name IdpSocial#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The scopes of the IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#scopes IdpSocial#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Identity Provider Types: https://developer.okta.com/docs/reference/api/idps/#identity-provider-type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#type IdpSocial#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies the account linking action for an IdP user. Default: `AUTO`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#account_link_action IdpSocial#account_link_action}
	AccountLinkAction *string `field:"optional" json:"accountLinkAction" yaml:"accountLinkAction"`
	// Group memberships to determine link candidates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#account_link_group_include IdpSocial#account_link_group_include}
	AccountLinkGroupInclude *[]*string `field:"optional" json:"accountLinkGroupInclude" yaml:"accountLinkGroupInclude"`
	// The Key ID that you obtained from Apple when you created the private key for the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#apple_kid IdpSocial#apple_kid}
	AppleKid *string `field:"optional" json:"appleKid" yaml:"appleKid"`
	// The Key ID that you obtained from Apple when you created the private key for the client.
	//
	// PrivateKey is required when resource is first created. For all consecutive updates, it can be empty/omitted and keeps the existing value if it is empty/omitted. PrivateKey isn't returned when importing this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#apple_private_key IdpSocial#apple_private_key}
	ApplePrivateKey *string `field:"optional" json:"applePrivateKey" yaml:"applePrivateKey"`
	// The Team ID associated with your Apple developer account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#apple_team_id IdpSocial#apple_team_id}
	AppleTeamId *string `field:"optional" json:"appleTeamId" yaml:"appleTeamId"`
	// Unique identifier issued by AS for the Okta IdP instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#client_id IdpSocial#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Client secret issued by AS for the Okta IdP instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#client_secret IdpSocial#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Action for a previously deprovisioned IdP user during authentication. Can be `NONE` or `REACTIVATE`. Default: `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#deprovisioned_action IdpSocial#deprovisioned_action}
	DeprovisionedAction *string `field:"optional" json:"deprovisionedAction" yaml:"deprovisionedAction"`
	// Provisioning action for IdP user's group memberships. It can be `NONE`, `SYNC`, `APPEND`, or `ASSIGN`. Default: `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#groups_action IdpSocial#groups_action}
	GroupsAction *string `field:"optional" json:"groupsAction" yaml:"groupsAction"`
	// List of Okta Group IDs to add an IdP user as a member with the `ASSIGN` `groups_action`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#groups_assignment IdpSocial#groups_assignment}
	GroupsAssignment *[]*string `field:"optional" json:"groupsAssignment" yaml:"groupsAssignment"`
	// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#groups_attribute IdpSocial#groups_attribute}
	GroupsAttribute *string `field:"optional" json:"groupsAttribute" yaml:"groupsAttribute"`
	// Whitelist of Okta Group identifiers that are allowed for the `APPEND` or `SYNC` `groups_action`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#groups_filter IdpSocial#groups_filter}
	GroupsFilter *[]*string `field:"optional" json:"groupsFilter" yaml:"groupsFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#id IdpSocial#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL.
	//
	// It can be `ORG_URL` or `CUSTOM_URL`. Default: `ORG_URL`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#issuer_mode IdpSocial#issuer_mode}
	IssuerMode *string `field:"optional" json:"issuerMode" yaml:"issuerMode"`
	// Maximum allowable clock-skew when processing messages from the IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#max_clock_skew IdpSocial#max_clock_skew}
	MaxClockSkew *float64 `field:"optional" json:"maxClockSkew" yaml:"maxClockSkew"`
	// Determines if the IdP should act as a source of truth for user profile attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#profile_master IdpSocial#profile_master}
	ProfileMaster interface{} `field:"optional" json:"profileMaster" yaml:"profileMaster"`
	// The type of protocol to use. It can be `OIDC` or `OAUTH2`. Default: `OAUTH2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#protocol_type IdpSocial#protocol_type}
	ProtocolType *string `field:"optional" json:"protocolType" yaml:"protocolType"`
	// Provisioning action for an IdP user during authentication. Default: `AUTO`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#provisioning_action IdpSocial#provisioning_action}
	ProvisioningAction *string `field:"optional" json:"provisioningAction" yaml:"provisioningAction"`
	// Default to `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#status IdpSocial#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Okta user profile attribute for matching transformed IdP username. Only for matchType `CUSTOM_ATTRIBUTE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#subject_match_attribute IdpSocial#subject_match_attribute}
	SubjectMatchAttribute *string `field:"optional" json:"subjectMatchAttribute" yaml:"subjectMatchAttribute"`
	// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username.
	//
	// By default, it is set to `USERNAME`. It can be set to `USERNAME`, `EMAIL`, `USERNAME_OR_EMAIL` or `CUSTOM_ATTRIBUTE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#subject_match_type IdpSocial#subject_match_type}
	SubjectMatchType *string `field:"optional" json:"subjectMatchType" yaml:"subjectMatchType"`
	// Action for a previously suspended IdP user during authentication. Can be `NONE` or `REACTIVATE`. Default: `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#suspended_action IdpSocial#suspended_action}
	SuspendedAction *string `field:"optional" json:"suspendedAction" yaml:"suspendedAction"`
	// Okta EL Expression to generate or transform a unique username for the IdP user. Default: `idpuser.email`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/idp_social#username_template IdpSocial#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

