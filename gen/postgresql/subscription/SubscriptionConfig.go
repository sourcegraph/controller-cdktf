package subscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SubscriptionConfig struct {
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
	// The connection string to the publisher. It should follow the keyword/value format (https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/subscription#conninfo Subscription#conninfo}
	Conninfo *string `field:"required" json:"conninfo" yaml:"conninfo"`
	// The name of the subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/subscription#name Subscription#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Names of the publications on the publisher to subscribe to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/subscription#publications Subscription#publications}
	Publications *[]*string `field:"required" json:"publications" yaml:"publications"`
	// Specifies whether the command should create the replication slot on the publisher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/subscription#create_slot Subscription#create_slot}
	CreateSlot interface{} `field:"optional" json:"createSlot" yaml:"createSlot"`
	// Sets the database to add the subscription for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/subscription#database Subscription#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/subscription#id Subscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of the replication slot to use.
	//
	// The default behavior is to use the name of the subscription for the slot name
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/subscription#slot_name Subscription#slot_name}
	SlotName *string `field:"optional" json:"slotName" yaml:"slotName"`
}

