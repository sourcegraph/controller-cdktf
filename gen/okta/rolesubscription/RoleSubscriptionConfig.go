package rolesubscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleSubscriptionConfig struct {
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
	// Type of the notification.
	//
	// Valid values:
	// 	- 'CONNECTOR_AGENT' -  Disconnects and reconnects: On-prem provisioning, on-prem MFA agents, and RADIUS server agent.
	// 	- 'USER_LOCKED_OUT' - User lockouts.
	// 	- 'APP_IMPORT' - App user import status.
	// 	- 'LDAP_AGENT' - Disconnects and reconnects: LDAP agent.
	// 	- 'AD_AGENT' - Disconnects and reconnects: AD agent.
	// 	- 'OKTA_ANNOUNCEMENT' - Okta release notes and announcements.
	// 	- 'OKTA_ISSUE' - Trust incidents and updates.
	// 	- 'OKTA_UPDATE' - Scheduled system updates.
	// 	- 'IWA_AGENT' - Disconnects and reconnects: IWA agent.
	// 	- 'USER_DEPROVISION' - User deprovisions.
	// 	- 'REPORT_SUSPICIOUS_ACTIVITY' - User reporting of suspicious activity.
	// 	- 'RATELIMIT_NOTIFICATION' - Rate limit warning and violation.
	// 	- 'AGENT_AUTO_UPDATE_NOTIFICATION' - Agent auto-update notifications: AD Agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/role_subscription#notification_type RoleSubscription#notification_type}
	NotificationType *string `field:"required" json:"notificationType" yaml:"notificationType"`
	// Type of the role. Valid values: 	'API_ADMIN', 	'APP_ADMIN', 	'CUSTOM', 	'GROUP_MEMBERSHIP_ADMIN', 	'HELP_DESK_ADMIN', 	'MOBILE_ADMIN', 	'ORG_ADMIN', 	'READ_ONLY_ADMIN', 	'REPORT_ADMIN', 	'SUPER_ADMIN', 	'USER_ADMIN' 	. See [API docs](https://developer.okta.com/docs/reference/api/admin-notifications/#role-types).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/role_subscription#role_type RoleSubscription#role_type}
	RoleType *string `field:"required" json:"roleType" yaml:"roleType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/role_subscription#id RoleSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Subscription status. Valid values: `subscribed`, `unsubscribed`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/role_subscription#status RoleSubscription#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

