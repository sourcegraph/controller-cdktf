package waitingroom

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WaitingRoomConfig struct {
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
	// Host name for which the waiting room will be applied (no wildcards).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#host WaitingRoom#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// A unique name to identify the waiting room. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#name WaitingRoom#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The number of new users that will be let into the route every minute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#new_users_per_minute WaitingRoom#new_users_per_minute}
	NewUsersPerMinute *float64 `field:"required" json:"newUsersPerMinute" yaml:"newUsersPerMinute"`
	// The total number of active user sessions on the route at a point in time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#total_active_users WaitingRoom#total_active_users}
	TotalActiveUsers *float64 `field:"required" json:"totalActiveUsers" yaml:"totalActiveUsers"`
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#zone_id WaitingRoom#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// This is a templated html file that will be rendered at the edge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#custom_page_html WaitingRoom#custom_page_html}
	CustomPageHtml *string `field:"optional" json:"customPageHtml" yaml:"customPageHtml"`
	// The language to use for the default waiting room page.
	//
	// Available values: `de-DE`, `es-ES`, `en-US`, `fr-FR`, `id-ID`, `it-IT`, `ja-JP`, `ko-KR`, `nl-NL`, `pl-PL`, `pt-BR`, `tr-TR`, `zh-CN`, `zh-TW`, `ru-RU`, `fa-IR`. Defaults to `en-US`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#default_template_language WaitingRoom#default_template_language}
	DefaultTemplateLanguage *string `field:"optional" json:"defaultTemplateLanguage" yaml:"defaultTemplateLanguage"`
	// A description to add more details about the waiting room.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#description WaitingRoom#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Disables automatic renewal of session cookies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#disable_session_renewal WaitingRoom#disable_session_renewal}
	DisableSessionRenewal interface{} `field:"optional" json:"disableSessionRenewal" yaml:"disableSessionRenewal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#id WaitingRoom#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, requests to the waiting room with the header `Accept: application/json` will receive a JSON response object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#json_response_enabled WaitingRoom#json_response_enabled}
	JsonResponseEnabled interface{} `field:"optional" json:"jsonResponseEnabled" yaml:"jsonResponseEnabled"`
	// The path within the host to enable the waiting room on. Defaults to `/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#path WaitingRoom#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// If queue_all is true, then all traffic will be sent to the waiting room.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#queue_all WaitingRoom#queue_all}
	QueueAll interface{} `field:"optional" json:"queueAll" yaml:"queueAll"`
	// The queueing method used by the waiting room. Available values: `fifo`, `random`, `passthrough`, `reject`. Defaults to `fifo`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#queueing_method WaitingRoom#queueing_method}
	QueueingMethod *string `field:"optional" json:"queueingMethod" yaml:"queueingMethod"`
	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to the origin.
	//
	// Defaults to `5`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#session_duration WaitingRoom#session_duration}
	SessionDuration *float64 `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// Suspends the waiting room.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#suspended WaitingRoom#suspended}
	Suspended interface{} `field:"optional" json:"suspended" yaml:"suspended"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room#timeouts WaitingRoom#timeouts}
	Timeouts *WaitingRoomTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

