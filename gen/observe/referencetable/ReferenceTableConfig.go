package referencetable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReferenceTableConfig struct {
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
	// MD5 checksum of the source file. Can be computed using `filemd5("<source_file>")`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/reference_table#checksum ReferenceTable#checksum}
	Checksum *string `field:"required" json:"checksum" yaml:"checksum"`
	// The name of the reference table name. Must be unique within workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/reference_table#label ReferenceTable#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// The path to a CSV file containing the reference table data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/reference_table#source_file ReferenceTable#source_file}
	SourceFile *string `field:"required" json:"sourceFile" yaml:"sourceFile"`
	// Description for the reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/reference_table#description ReferenceTable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/reference_table#id ReferenceTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The field that should be used for the OPAL label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/reference_table#label_field ReferenceTable#label_field}
	LabelField *string `field:"optional" json:"labelField" yaml:"labelField"`
	// The primary key of the reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/reference_table#primary_key ReferenceTable#primary_key}
	PrimaryKey *[]*string `field:"optional" json:"primaryKey" yaml:"primaryKey"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/reference_table#schema ReferenceTable#schema}
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}

