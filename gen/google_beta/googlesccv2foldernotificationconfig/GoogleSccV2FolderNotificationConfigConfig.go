package googlesccv2foldernotificationconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSccV2FolderNotificationConfigConfig struct {
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
	// This must be unique within the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_folder_notification_config#config_id GoogleSccV2FolderNotificationConfig#config_id}
	ConfigId *string `field:"required" json:"configId" yaml:"configId"`
	// Numerical ID of the parent folder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_folder_notification_config#folder GoogleSccV2FolderNotificationConfig#folder}
	Folder *string `field:"required" json:"folder" yaml:"folder"`
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_folder_notification_config#pubsub_topic GoogleSccV2FolderNotificationConfig#pubsub_topic}
	PubsubTopic *string `field:"required" json:"pubsubTopic" yaml:"pubsubTopic"`
	// streaming_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_folder_notification_config#streaming_config GoogleSccV2FolderNotificationConfig#streaming_config}
	StreamingConfig *GoogleSccV2FolderNotificationConfigStreamingConfig `field:"required" json:"streamingConfig" yaml:"streamingConfig"`
	// The description of the notification config (max of 1024 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_folder_notification_config#description GoogleSccV2FolderNotificationConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_folder_notification_config#id GoogleSccV2FolderNotificationConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Location ID of the parent organization. If not provided, 'global' will be used as the default location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_folder_notification_config#location GoogleSccV2FolderNotificationConfig#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_v2_folder_notification_config#timeouts GoogleSccV2FolderNotificationConfig#timeouts}
	Timeouts *GoogleSccV2FolderNotificationConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

