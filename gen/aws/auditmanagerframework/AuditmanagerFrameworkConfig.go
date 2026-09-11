package auditmanagerframework

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AuditmanagerFrameworkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/auditmanager_framework#name AuditmanagerFramework#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/auditmanager_framework#compliance_type AuditmanagerFramework#compliance_type}.
	ComplianceType *string `field:"optional" json:"complianceType" yaml:"complianceType"`
	// control_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/auditmanager_framework#control_sets AuditmanagerFramework#control_sets}
	ControlSets interface{} `field:"optional" json:"controlSets" yaml:"controlSets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/auditmanager_framework#description AuditmanagerFramework#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/auditmanager_framework#tags AuditmanagerFramework#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

