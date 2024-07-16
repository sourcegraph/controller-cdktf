package googlefirebasehostingrelease

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirebaseHostingReleaseConfig struct {
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
	// Required. The ID of the site to which the release belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_release#site_id GoogleFirebaseHostingRelease#site_id}
	SiteId *string `field:"required" json:"siteId" yaml:"siteId"`
	// The ID of the channel to which the release belongs.
	//
	// If not provided, the release will
	// belong to the default "live" channel
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_release#channel_id GoogleFirebaseHostingRelease#channel_id}
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_release#id GoogleFirebaseHostingRelease#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The deploy description when the release was created. The value can be up to 512 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_release#message GoogleFirebaseHostingRelease#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_release#timeouts GoogleFirebaseHostingRelease#timeouts}
	Timeouts *GoogleFirebaseHostingReleaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The type of the release;
	//
	// indicates what happened to the content of the site. There is no need to specify
	// 'DEPLOY' or 'ROLLBACK' type if a 'version_name' is provided.
	// DEPLOY: A version was uploaded to Firebase Hosting and released. Output only.
	// ROLLBACK: The release points back to a previously deployed version. Output only.
	// SITE_DISABLE: The release prevents the site from serving content. Firebase Hosting acts as if the site never existed Possible values: ["DEPLOY", "ROLLBACK", "SITE_DISABLE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_release#type GoogleFirebaseHostingRelease#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The unique identifier for a version, in the format: sites/SITE_ID/versions/VERSION_ID.
	//
	// The content of the version specified will be actively displayed on the appropriate URL.
	// The Version must belong to the same site as in the 'site_id'.
	// This parameter must be empty if the 'type' of the release is 'SITE_DISABLE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_release#version_name GoogleFirebaseHostingRelease#version_name}
	VersionName *string `field:"optional" json:"versionName" yaml:"versionName"`
}

