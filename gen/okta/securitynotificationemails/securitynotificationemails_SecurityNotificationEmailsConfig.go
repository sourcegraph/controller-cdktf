package securitynotificationemails

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityNotificationEmailsConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/security_notification_emails#id SecurityNotificationEmails#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Notifies end users about suspicious or unrecognized activity from their account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/security_notification_emails#report_suspicious_activity_enabled SecurityNotificationEmails#report_suspicious_activity_enabled}
	ReportSuspiciousActivityEnabled interface{} `field:"optional" json:"reportSuspiciousActivityEnabled" yaml:"reportSuspiciousActivityEnabled"`
	// Notifies end users of any activity on their account related to MFA factor enrollment.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/security_notification_emails#send_email_for_factor_enrollment_enabled SecurityNotificationEmails#send_email_for_factor_enrollment_enabled}
	SendEmailForFactorEnrollmentEnabled interface{} `field:"optional" json:"sendEmailForFactorEnrollmentEnabled" yaml:"sendEmailForFactorEnrollmentEnabled"`
	// Notifies end users that one or more factors have been reset for their account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/security_notification_emails#send_email_for_factor_reset_enabled SecurityNotificationEmails#send_email_for_factor_reset_enabled}
	SendEmailForFactorResetEnabled interface{} `field:"optional" json:"sendEmailForFactorResetEnabled" yaml:"sendEmailForFactorResetEnabled"`
	// Notifies end users about new sign-on activity.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/security_notification_emails#send_email_for_new_device_enabled SecurityNotificationEmails#send_email_for_new_device_enabled}
	SendEmailForNewDeviceEnabled interface{} `field:"optional" json:"sendEmailForNewDeviceEnabled" yaml:"sendEmailForNewDeviceEnabled"`
	// Notifies end users that the password for their account has changed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/security_notification_emails#send_email_for_password_changed_enabled SecurityNotificationEmails#send_email_for_password_changed_enabled}
	SendEmailForPasswordChangedEnabled interface{} `field:"optional" json:"sendEmailForPasswordChangedEnabled" yaml:"sendEmailForPasswordChangedEnabled"`
}

