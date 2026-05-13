package googleapigeesecurityfeedback

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApigeeSecurityFeedbackConfig struct {
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
	// feedback_contexts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_feedback#feedback_contexts GoogleApigeeSecurityFeedback#feedback_contexts}
	FeedbackContexts interface{} `field:"required" json:"feedbackContexts" yaml:"feedbackContexts"`
	// Resource ID of the security feedback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_feedback#feedback_id GoogleApigeeSecurityFeedback#feedback_id}
	FeedbackId *string `field:"required" json:"feedbackId" yaml:"feedbackId"`
	// The type of feedback being submitted. Possible values: ["EXCLUDED_DETECTION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_feedback#feedback_type GoogleApigeeSecurityFeedback#feedback_type}
	FeedbackType *string `field:"required" json:"feedbackType" yaml:"feedbackType"`
	// The Apigee Organization associated with the Apigee Security Feedback, in the format 'organizations/{{org_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_feedback#org_id GoogleApigeeSecurityFeedback#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Optional text the user can provide for additional, unstructured context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_feedback#comment GoogleApigeeSecurityFeedback#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The display name of the feedback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_feedback#display_name GoogleApigeeSecurityFeedback#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_feedback#id GoogleApigeeSecurityFeedback#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The reason for the feedback. Possible values: ["INTERNAL_SYSTEM", "NON_RISK_CLIENT", "NAT", "PENETRATION_TEST", "OTHER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_feedback#reason GoogleApigeeSecurityFeedback#reason}
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_security_feedback#timeouts GoogleApigeeSecurityFeedback#timeouts}
	Timeouts *GoogleApigeeSecurityFeedbackTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

