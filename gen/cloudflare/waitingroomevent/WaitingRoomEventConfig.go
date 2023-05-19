package waitingroomevent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WaitingRoomEventConfig struct {
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
	// ISO 8601 timestamp that marks the end of the event.
	//
	// **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#event_end_time WaitingRoomEvent#event_end_time}
	EventEndTime *string `field:"required" json:"eventEndTime" yaml:"eventEndTime"`
	// ISO 8601 timestamp that marks the start of the event.
	//
	// Must occur at least 1 minute before `event_end_time`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#event_start_time WaitingRoomEvent#event_start_time}
	EventStartTime *string `field:"required" json:"eventStartTime" yaml:"eventStartTime"`
	// A unique name to identify the event.
	//
	// Only alphanumeric characters, hyphens, and underscores are allowed. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#name WaitingRoomEvent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Waiting Room ID the event should apply to. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#waiting_room_id WaitingRoomEvent#waiting_room_id}
	WaitingRoomId *string `field:"required" json:"waitingRoomId" yaml:"waitingRoomId"`
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#zone_id WaitingRoomEvent#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// This is a templated html file that will be rendered at the edge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#custom_page_html WaitingRoomEvent#custom_page_html}
	CustomPageHtml *string `field:"optional" json:"customPageHtml" yaml:"customPageHtml"`
	// A description to let users add more details about the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#description WaitingRoomEvent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Disables automatic renewal of session cookies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#disable_session_renewal WaitingRoomEvent#disable_session_renewal}
	DisableSessionRenewal interface{} `field:"optional" json:"disableSessionRenewal" yaml:"disableSessionRenewal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#id WaitingRoomEvent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The number of new users that will be let into the route every minute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#new_users_per_minute WaitingRoomEvent#new_users_per_minute}
	NewUsersPerMinute *float64 `field:"optional" json:"newUsersPerMinute" yaml:"newUsersPerMinute"`
	// ISO 8601 timestamp that marks when to begin queueing all users before the event starts.
	//
	// Must occur at least 5 minutes before `event_start_time`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#prequeue_start_time WaitingRoomEvent#prequeue_start_time}
	PrequeueStartTime *string `field:"optional" json:"prequeueStartTime" yaml:"prequeueStartTime"`
	// The queueing method used by the waiting room. Available values: `fifo`, `random`, `passthrough`, `reject`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#queueing_method WaitingRoomEvent#queueing_method}
	QueueingMethod *string `field:"optional" json:"queueingMethod" yaml:"queueingMethod"`
	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#session_duration WaitingRoomEvent#session_duration}
	SessionDuration *float64 `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// Users in the prequeue will be shuffled randomly at the `event_start_time`.
	//
	// Requires that `prequeue_start_time` is not null. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#shuffle_at_event_start WaitingRoomEvent#shuffle_at_event_start}
	ShuffleAtEventStart interface{} `field:"optional" json:"shuffleAtEventStart" yaml:"shuffleAtEventStart"`
	// If suspended, the event is ignored and traffic will be handled based on the waiting room configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#suspended WaitingRoomEvent#suspended}
	Suspended interface{} `field:"optional" json:"suspended" yaml:"suspended"`
	// The total number of active user sessions on the route at a point in time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event#total_active_users WaitingRoomEvent#total_active_users}
	TotalActiveUsers *float64 `field:"optional" json:"totalActiveUsers" yaml:"totalActiveUsers"`
}

