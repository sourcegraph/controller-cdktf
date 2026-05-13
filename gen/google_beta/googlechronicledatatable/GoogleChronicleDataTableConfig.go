package googlechronicledatatable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleChronicleDataTableConfig struct {
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
	// The ID to use for the data table.
	//
	// This is also the display name for
	// the data table. It must satisfy the following requirements:
	// - Starts with letter.
	// - Contains only letters, numbers and underscore.
	// - Must be unique and has length < 256.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#data_table_id GoogleChronicleDataTable#data_table_id}
	DataTableId *string `field:"required" json:"dataTableId" yaml:"dataTableId"`
	// A user-provided description of the data table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#description GoogleChronicleDataTable#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#instance GoogleChronicleDataTable#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#location GoogleChronicleDataTable#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// column_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#column_info GoogleChronicleDataTable#column_info}
	ColumnInfo interface{} `field:"optional" json:"columnInfo" yaml:"columnInfo"`
	// The policy governing the deletion of the data table.
	//
	// If set to 'FORCE', allows the deletion of the data table even if it contains rows.
	// If set to 'DEFAULT',or if the field is omitted, the data table must be empty before it can be deleted.
	// Possible values: DEFAULT, FORCE
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#deletion_policy GoogleChronicleDataTable#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#id GoogleChronicleDataTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#project GoogleChronicleDataTable#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// User-provided TTL of the data table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#row_time_to_live GoogleChronicleDataTable#row_time_to_live}
	RowTimeToLive *string `field:"optional" json:"rowTimeToLive" yaml:"rowTimeToLive"`
	// scope_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#scope_info GoogleChronicleDataTable#scope_info}
	ScopeInfo *GoogleChronicleDataTableScopeInfo `field:"optional" json:"scopeInfo" yaml:"scopeInfo"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#timeouts GoogleChronicleDataTable#timeouts}
	Timeouts *GoogleChronicleDataTableTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

