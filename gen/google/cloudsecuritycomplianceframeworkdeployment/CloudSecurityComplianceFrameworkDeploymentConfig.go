package cloudsecuritycomplianceframeworkdeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudSecurityComplianceFrameworkDeploymentConfig struct {
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
	// cloud_control_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#cloud_control_metadata CloudSecurityComplianceFrameworkDeployment#cloud_control_metadata}
	CloudControlMetadata interface{} `field:"required" json:"cloudControlMetadata" yaml:"cloudControlMetadata"`
	// framework block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#framework CloudSecurityComplianceFrameworkDeployment#framework}
	Framework *CloudSecurityComplianceFrameworkDeploymentFramework `field:"required" json:"framework" yaml:"framework"`
	// User provided identifier.
	//
	// It should be unique in scope of a parent.
	// This is optional and if not provided, a random UUID will be generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#framework_deployment_id CloudSecurityComplianceFrameworkDeployment#framework_deployment_id}
	FrameworkDeploymentId *string `field:"required" json:"frameworkDeploymentId" yaml:"frameworkDeploymentId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#location CloudSecurityComplianceFrameworkDeployment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#organization CloudSecurityComplianceFrameworkDeployment#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// target_resource_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#target_resource_config CloudSecurityComplianceFrameworkDeployment#target_resource_config}
	TargetResourceConfig *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfig `field:"required" json:"targetResourceConfig" yaml:"targetResourceConfig"`
	// User provided description of the Framework deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#description CloudSecurityComplianceFrameworkDeployment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#id CloudSecurityComplianceFrameworkDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#timeouts CloudSecurityComplianceFrameworkDeployment#timeouts}
	Timeouts *CloudSecurityComplianceFrameworkDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

