package datagooglefirebasehostingchannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleFirebaseHostingChannelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/data-sources/google_firebase_hosting_channel#channel_id DataGoogleFirebaseHostingChannel#channel_id}
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// Required. The ID of the site in which to create this channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/data-sources/google_firebase_hosting_channel#site_id DataGoogleFirebaseHostingChannel#site_id}
	SiteId *string `field:"required" json:"siteId" yaml:"siteId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/data-sources/google_firebase_hosting_channel#id DataGoogleFirebaseHostingChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

