package googleprivilegedaccessmanagerentitlement

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GooglePrivilegedAccessManagerEntitlementConfig struct {
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
	// eligible_users block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#eligible_users GooglePrivilegedAccessManagerEntitlement#eligible_users}
	EligibleUsers interface{} `field:"required" json:"eligibleUsers" yaml:"eligibleUsers"`
	// The ID to use for this Entitlement.
	//
	// This will become the last part of the resource name.
	// This value should be 4-63 characters, and valid characters are "[a-z]", "[0-9]", and "-". The first character should be from [a-z].
	// This value should be unique among all other Entitlements under the specified 'parent'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#entitlement_id GooglePrivilegedAccessManagerEntitlement#entitlement_id}
	EntitlementId *string `field:"required" json:"entitlementId" yaml:"entitlementId"`
	// The region of the Entitlement resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#location GooglePrivilegedAccessManagerEntitlement#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The maximum amount of time for which access would be granted for a request.
	//
	// A requester can choose to ask for access for less than this duration but never more.
	// Format: calculate the time in seconds and concatenate it with 's' i.e. 2 hours = "7200s", 45 minutes = "2700s"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#max_request_duration GooglePrivilegedAccessManagerEntitlement#max_request_duration}
	MaxRequestDuration *string `field:"required" json:"maxRequestDuration" yaml:"maxRequestDuration"`
	// Format: projects/{project-id|project-number} or organizations/{organization-number} or folders/{folder-number}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#parent GooglePrivilegedAccessManagerEntitlement#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// privileged_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#privileged_access GooglePrivilegedAccessManagerEntitlement#privileged_access}
	PrivilegedAccess *GooglePrivilegedAccessManagerEntitlementPrivilegedAccess `field:"required" json:"privilegedAccess" yaml:"privilegedAccess"`
	// requester_justification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#requester_justification_config GooglePrivilegedAccessManagerEntitlement#requester_justification_config}
	RequesterJustificationConfig *GooglePrivilegedAccessManagerEntitlementRequesterJustificationConfig `field:"required" json:"requesterJustificationConfig" yaml:"requesterJustificationConfig"`
	// additional_notification_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#additional_notification_targets GooglePrivilegedAccessManagerEntitlement#additional_notification_targets}
	AdditionalNotificationTargets *GooglePrivilegedAccessManagerEntitlementAdditionalNotificationTargets `field:"optional" json:"additionalNotificationTargets" yaml:"additionalNotificationTargets"`
	// approval_workflow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#approval_workflow GooglePrivilegedAccessManagerEntitlement#approval_workflow}
	ApprovalWorkflow *GooglePrivilegedAccessManagerEntitlementApprovalWorkflow `field:"optional" json:"approvalWorkflow" yaml:"approvalWorkflow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#id GooglePrivilegedAccessManagerEntitlement#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#timeouts GooglePrivilegedAccessManagerEntitlement#timeouts}
	Timeouts *GooglePrivilegedAccessManagerEntitlementTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

