package securitylabel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityLabelConfig struct {
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
	// The label to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/security_label#label SecurityLabel#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// The provider to apply the security label for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/security_label#label_provider SecurityLabel#label_provider}
	LabelProvider *string `field:"required" json:"labelProvider" yaml:"labelProvider"`
	// The name of the existing object to apply the security label to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/security_label#object_name SecurityLabel#object_name}
	ObjectName *string `field:"required" json:"objectName" yaml:"objectName"`
	// The type of the existing object to apply the security label to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/security_label#object_type SecurityLabel#object_type}
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/security_label#id SecurityLabel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

