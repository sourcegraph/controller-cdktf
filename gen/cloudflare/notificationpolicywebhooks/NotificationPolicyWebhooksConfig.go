package notificationpolicywebhooks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationPolicyWebhooksConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy_webhooks#account_id NotificationPolicyWebhooks#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The name of the webhook destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy_webhooks#name NotificationPolicyWebhooks#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy_webhooks#id NotificationPolicyWebhooks#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// An optional secret can be provided that will be passed in the `cf-webhook-auth` header when dispatching a webhook notification.
	//
	// Secrets are not returned in any API response body. Refer to the [documentation](https://api.cloudflare.com/#notification-webhooks-create-webhook) for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy_webhooks#secret NotificationPolicyWebhooks#secret}
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// The URL of the webhook destinations. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/notification_policy_webhooks#url NotificationPolicyWebhooks#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

