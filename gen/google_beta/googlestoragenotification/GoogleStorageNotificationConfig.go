package googlestoragenotification

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleStorageNotificationConfig struct {
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
	// The name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_notification#bucket GoogleStorageNotification#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The desired content of the Payload. One of "JSON_API_V1" or "NONE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_notification#payload_format GoogleStorageNotification#payload_format}
	PayloadFormat *string `field:"required" json:"payloadFormat" yaml:"payloadFormat"`
	// The Cloud Pub/Sub topic to which this subscription publishes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_notification#topic GoogleStorageNotification#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// A set of key/value attribute pairs to attach to each Cloud Pub/Sub message published for this notification subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_notification#custom_attributes GoogleStorageNotification#custom_attributes}
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// List of event type filters for this notification config.
	//
	// If not specified, Cloud Storage will send notifications for all event types. The valid types are: "OBJECT_FINALIZE", "OBJECT_METADATA_UPDATE", "OBJECT_DELETE", "OBJECT_ARCHIVE"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_notification#event_types GoogleStorageNotification#event_types}
	EventTypes *[]*string `field:"optional" json:"eventTypes" yaml:"eventTypes"`
	// Specifies a prefix path filter for this notification config.
	//
	// Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_notification#object_name_prefix GoogleStorageNotification#object_name_prefix}
	ObjectNamePrefix *string `field:"optional" json:"objectNamePrefix" yaml:"objectNamePrefix"`
}

