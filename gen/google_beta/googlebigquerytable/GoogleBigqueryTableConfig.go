package googlebigquerytable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBigqueryTableConfig struct {
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
	// The dataset ID to create the table in. Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#dataset_id GoogleBigqueryTable#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// A unique ID for the resource. Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#table_id GoogleBigqueryTable#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
	// Specifies column names to use for data clustering.
	//
	// Up to four top-level columns are allowed, and should be specified in descending priority order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#clustering GoogleBigqueryTable#clustering}
	Clustering *[]*string `field:"optional" json:"clustering" yaml:"clustering"`
	// Whether or not to allow Terraform to destroy the instance.
	//
	// Unless this field is set to false in Terraform state, a terraform destroy or terraform apply that would delete the instance will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#deletion_protection GoogleBigqueryTable#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The field description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#description GoogleBigqueryTable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#encryption_configuration GoogleBigqueryTable#encryption_configuration}
	EncryptionConfiguration *GoogleBigqueryTableEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The time when this table expires, in milliseconds since the epoch.
	//
	// If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#expiration_time GoogleBigqueryTable#expiration_time}
	ExpirationTime *float64 `field:"optional" json:"expirationTime" yaml:"expirationTime"`
	// external_data_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#external_data_configuration GoogleBigqueryTable#external_data_configuration}
	ExternalDataConfiguration *GoogleBigqueryTableExternalDataConfiguration `field:"optional" json:"externalDataConfiguration" yaml:"externalDataConfiguration"`
	// A descriptive name for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#friendly_name GoogleBigqueryTable#friendly_name}
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#id GoogleBigqueryTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A mapping of labels to assign to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#labels GoogleBigqueryTable#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// materialized_view block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#materialized_view GoogleBigqueryTable#materialized_view}
	MaterializedView *GoogleBigqueryTableMaterializedView `field:"optional" json:"materializedView" yaml:"materializedView"`
	// The ID of the project in which the resource belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#project GoogleBigqueryTable#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// range_partitioning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#range_partitioning GoogleBigqueryTable#range_partitioning}
	RangePartitioning *GoogleBigqueryTableRangePartitioning `field:"optional" json:"rangePartitioning" yaml:"rangePartitioning"`
	// A JSON schema for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#schema GoogleBigqueryTable#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// time_partitioning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#time_partitioning GoogleBigqueryTable#time_partitioning}
	TimePartitioning *GoogleBigqueryTableTimePartitioning `field:"optional" json:"timePartitioning" yaml:"timePartitioning"`
	// view block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_bigquery_table#view GoogleBigqueryTable#view}
	View *GoogleBigqueryTableView `field:"optional" json:"view" yaml:"view"`
}

