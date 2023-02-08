package rolesubscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleSubscriptionConfig struct {
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
	// Type of the notification.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/role_subscription#notification_type RoleSubscription#notification_type}
	NotificationType *string `field:"required" json:"notificationType" yaml:"notificationType"`
	// Type of the role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/role_subscription#role_type RoleSubscription#role_type}
	RoleType *string `field:"required" json:"roleType" yaml:"roleType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/role_subscription#id RoleSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Status of subscription.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/role_subscription#status RoleSubscription#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

