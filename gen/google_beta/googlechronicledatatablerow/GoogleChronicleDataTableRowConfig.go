package googlechronicledatatablerow

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleChronicleDataTableRowConfig struct {
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
	// The ID of the parent DataTable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table_row#data_table_id GoogleChronicleDataTableRow#data_table_id}
	DataTableId *string `field:"required" json:"dataTableId" yaml:"dataTableId"`
	// The Chronicle instance ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table_row#instance GoogleChronicleDataTableRow#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The GCP location of the Chronicle instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table_row#location GoogleChronicleDataTableRow#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// All column values for a single row.
	//
	// The values should be in the same order
	// as the columns of the data tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table_row#values GoogleChronicleDataTableRow#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table_row#id GoogleChronicleDataTableRow#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table_row#project GoogleChronicleDataTableRow#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// User-provided TTL of the data table row.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table_row#row_time_to_live GoogleChronicleDataTableRow#row_time_to_live}
	RowTimeToLive *string `field:"optional" json:"rowTimeToLive" yaml:"rowTimeToLive"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table_row#timeouts GoogleChronicleDataTableRow#timeouts}
	Timeouts *GoogleChronicleDataTableRowTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

