package report

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ReportConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// dashboard block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#dashboard Report#dashboard}
	Dashboard *ReportDashboard `field:"required" json:"dashboard" yaml:"dashboard"`
	// A list of e-mail addresses that will receive the report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#email_recipients Report#email_recipients}
	EmailRecipients *[]*string `field:"required" json:"emailRecipients" yaml:"emailRecipients"`
	// The subject of the e-mail that will be sent every time the report runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#email_subject Report#email_subject}
	EmailSubject *string `field:"required" json:"emailSubject" yaml:"emailSubject"`
	// Whether the report is enabled or not. A disabled report will not run on the defined schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#enabled Report#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#label Report#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#schedule Report#schedule}
	Schedule *ReportSchedule `field:"required" json:"schedule" yaml:"schedule"`
	// A list of e-mail bcc addresses that will receive the report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#email_bcc_recipients Report#email_bcc_recipients}
	EmailBccRecipients *[]*string `field:"optional" json:"emailBccRecipients" yaml:"emailBccRecipients"`
	// The body of the e-mail that will be sent every time the report runs.
	//
	// This is used to add some context to the email.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#email_body Report#email_body}
	EmailBody *string `field:"optional" json:"emailBody" yaml:"emailBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#id Report#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

