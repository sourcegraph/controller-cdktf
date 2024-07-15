package googlefirebasehostingchannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirebaseHostingChannelConfig struct {
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
	// Required. Immutable. A unique ID within the site that identifies the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_channel#channel_id GoogleFirebaseHostingChannel#channel_id}
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// Required. The ID of the site in which to create this channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_channel#site_id GoogleFirebaseHostingChannel#site_id}
	SiteId *string `field:"required" json:"siteId" yaml:"siteId"`
	// The time at which the channel will be automatically deleted.
	//
	// If null, the channel
	// will not be automatically deleted. This field is present in the output whether it's
	// set directly or via the 'ttl' field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_channel#expire_time GoogleFirebaseHostingChannel#expire_time}
	ExpireTime *string `field:"optional" json:"expireTime" yaml:"expireTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_channel#id GoogleFirebaseHostingChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Text labels used for extra metadata and/or filtering.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_channel#labels GoogleFirebaseHostingChannel#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The number of previous releases to retain on the channel for rollback or other purposes.
	//
	// Must be a number between 1-100. Defaults to 10 for new channels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_channel#retained_release_count GoogleFirebaseHostingChannel#retained_release_count}
	RetainedReleaseCount *float64 `field:"optional" json:"retainedReleaseCount" yaml:"retainedReleaseCount"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_channel#timeouts GoogleFirebaseHostingChannel#timeouts}
	Timeouts *GoogleFirebaseHostingChannelTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Input only.
	//
	// A time-to-live for this channel. Sets 'expire_time' to the provided
	// duration past the time of the request. A duration in seconds with up to nine fractional
	// digits, terminated by 's'. Example: "86400s" (one day).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_channel#ttl GoogleFirebaseHostingChannel#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

