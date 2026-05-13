package googleprivilegedaccessmanagersettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GooglePrivilegedAccessManagerSettingsConfig struct {
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
	// The region of the PAM settings resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#location GooglePrivilegedAccessManagerSettings#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Format: projects/{project-id|project-number} or organizations/{organization-number} or folders/{folder-number}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#parent GooglePrivilegedAccessManagerSettings#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// email_notification_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#email_notification_settings GooglePrivilegedAccessManagerSettings#email_notification_settings}
	EmailNotificationSettings *GooglePrivilegedAccessManagerSettingsEmailNotificationSettings `field:"optional" json:"emailNotificationSettings" yaml:"emailNotificationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#id GooglePrivilegedAccessManagerSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// service_account_approver_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#service_account_approver_settings GooglePrivilegedAccessManagerSettings#service_account_approver_settings}
	ServiceAccountApproverSettings *GooglePrivilegedAccessManagerSettingsServiceAccountApproverSettings `field:"optional" json:"serviceAccountApproverSettings" yaml:"serviceAccountApproverSettings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#timeouts GooglePrivilegedAccessManagerSettings#timeouts}
	Timeouts *GooglePrivilegedAccessManagerSettingsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

