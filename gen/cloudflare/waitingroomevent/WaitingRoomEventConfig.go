package waitingroomevent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WaitingRoomEventConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#event_end_time WaitingRoomEvent#event_end_time}.
	EventEndTime *string `field:"required" json:"eventEndTime" yaml:"eventEndTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#event_start_time WaitingRoomEvent#event_start_time}.
	EventStartTime *string `field:"required" json:"eventStartTime" yaml:"eventStartTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#name WaitingRoomEvent#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#waiting_room_id WaitingRoomEvent#waiting_room_id}.
	WaitingRoomId *string `field:"required" json:"waitingRoomId" yaml:"waitingRoomId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#zone_id WaitingRoomEvent#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#custom_page_html WaitingRoomEvent#custom_page_html}.
	CustomPageHtml *string `field:"optional" json:"customPageHtml" yaml:"customPageHtml"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#description WaitingRoomEvent#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#disable_session_renewal WaitingRoomEvent#disable_session_renewal}.
	DisableSessionRenewal interface{} `field:"optional" json:"disableSessionRenewal" yaml:"disableSessionRenewal"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#id WaitingRoomEvent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#new_users_per_minute WaitingRoomEvent#new_users_per_minute}.
	NewUsersPerMinute *float64 `field:"optional" json:"newUsersPerMinute" yaml:"newUsersPerMinute"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#prequeue_start_time WaitingRoomEvent#prequeue_start_time}.
	PrequeueStartTime *string `field:"optional" json:"prequeueStartTime" yaml:"prequeueStartTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#queueing_method WaitingRoomEvent#queueing_method}.
	QueueingMethod *string `field:"optional" json:"queueingMethod" yaml:"queueingMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#session_duration WaitingRoomEvent#session_duration}.
	SessionDuration *float64 `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#shuffle_at_event_start WaitingRoomEvent#shuffle_at_event_start}.
	ShuffleAtEventStart interface{} `field:"optional" json:"shuffleAtEventStart" yaml:"shuffleAtEventStart"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#suspended WaitingRoomEvent#suspended}.
	Suspended interface{} `field:"optional" json:"suspended" yaml:"suspended"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room_event#total_active_users WaitingRoomEvent#total_active_users}.
	TotalActiveUsers *float64 `field:"optional" json:"totalActiveUsers" yaml:"totalActiveUsers"`
}

