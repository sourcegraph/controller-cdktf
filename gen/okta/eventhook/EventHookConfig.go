package eventhook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventHookConfig struct {
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
	// Details of the endpoint the event hook will hit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/event_hook#channel EventHook#channel}
	Channel *map[string]*string `field:"required" json:"channel" yaml:"channel"`
	// The events that will be delivered to this hook. [See here for a list of supported events](https://developer.okta.com/docs/reference/api/event-types/?q=event-hook-eligible).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/event_hook#events EventHook#events}
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// The event hook display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/event_hook#name EventHook#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Details of the endpoint the event hook will hit.
	//
	// - 'version' - (Required) The version of the channel. The currently-supported version is '1.0.0'.
	// 	- 'uri' - (Required) The URI the hook will hit.
	// 	- 'type' - (Optional) The type of hook to trigger. Currently, the only supported type is 'HTTP'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/event_hook#auth EventHook#auth}
	Auth *map[string]*string `field:"optional" json:"auth" yaml:"auth"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/event_hook#headers EventHook#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/event_hook#id EventHook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Default to `ACTIVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/event_hook#status EventHook#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

