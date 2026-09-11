package dataobservereport

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataObserveReportConfig struct {
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
	// Resource ID for this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/report#id DataObserveReport#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// created_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/report#created_by DataObserveReport#created_by}
	CreatedBy interface{} `field:"optional" json:"createdBy" yaml:"createdBy"`
	// dashboard block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/report#dashboard DataObserveReport#dashboard}
	Dashboard interface{} `field:"optional" json:"dashboard" yaml:"dashboard"`
	// A list of e-mail bcc addresses that will receive the report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/report#email_bcc_recipients DataObserveReport#email_bcc_recipients}
	EmailBccRecipients *[]*string `field:"optional" json:"emailBccRecipients" yaml:"emailBccRecipients"`
	// A list of e-mail addresses that will receive the report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/report#email_recipients DataObserveReport#email_recipients}
	EmailRecipients *[]*string `field:"optional" json:"emailRecipients" yaml:"emailRecipients"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/report#schedule DataObserveReport#schedule}
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// updated_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/report#updated_by DataObserveReport#updated_by}
	UpdatedBy interface{} `field:"optional" json:"updatedBy" yaml:"updatedBy"`
}

