package link

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LinkConfig struct {
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
	// A collection of field mappings on which to link source and target datasets.
	//
	// Each element of the array can be written as a colon separated string, e.g.
	// `source_column:target_column`. If the source and target fields have the
	// same name, the target field name can be omitted, i.e. `col:col` can be
	// written as `col`.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/link#fields Link#fields}
	Fields *[]*string `field:"required" json:"fields" yaml:"fields"`
	// OID for the source dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/link#source Link#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// OID for the target dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/link#target Link#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// OID of the workspace this object is contained in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/link#workspace Link#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/link#id Link#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A human-readable label for the link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/link#label Link#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
}

