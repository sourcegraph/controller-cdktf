package cloudsecuritycomplianceframework

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudSecurityComplianceFrameworkConfig struct {
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
	// ID of the framework.
	//
	// This is not the full name of the framework.
	// This is the last part of the full name of the framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#framework_id CloudSecurityComplianceFramework#framework_id}
	FrameworkId *string `field:"required" json:"frameworkId" yaml:"frameworkId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#location CloudSecurityComplianceFramework#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#organization CloudSecurityComplianceFramework#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// cloud_control_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#cloud_control_details CloudSecurityComplianceFramework#cloud_control_details}
	CloudControlDetails interface{} `field:"optional" json:"cloudControlDetails" yaml:"cloudControlDetails"`
	// The description of the framework. The maximum length is 2000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#description CloudSecurityComplianceFramework#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Display name of the framework. The maximum length is 200 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#display_name CloudSecurityComplianceFramework#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#id CloudSecurityComplianceFramework#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#timeouts CloudSecurityComplianceFramework#timeouts}
	Timeouts *CloudSecurityComplianceFrameworkTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

