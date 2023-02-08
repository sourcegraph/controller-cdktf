package notificationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#account_id NotificationPolicy#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#alert_type NotificationPolicy#alert_type}.
	AlertType *string `field:"required" json:"alertType" yaml:"alertType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#enabled NotificationPolicy#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#name NotificationPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#description NotificationPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// email_integration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#email_integration NotificationPolicy#email_integration}
	EmailIntegration interface{} `field:"optional" json:"emailIntegration" yaml:"emailIntegration"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#filters NotificationPolicy#filters}
	Filters *NotificationPolicyFilters `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#id NotificationPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// pagerduty_integration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#pagerduty_integration NotificationPolicy#pagerduty_integration}
	PagerdutyIntegration interface{} `field:"optional" json:"pagerdutyIntegration" yaml:"pagerdutyIntegration"`
	// webhooks_integration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#webhooks_integration NotificationPolicy#webhooks_integration}
	WebhooksIntegration interface{} `field:"optional" json:"webhooksIntegration" yaml:"webhooksIntegration"`
}

