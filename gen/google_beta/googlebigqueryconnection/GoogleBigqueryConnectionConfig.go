package googlebigqueryconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBigqueryConnectionConfig struct {
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
	// aws block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#aws GoogleBigqueryConnection#aws}
	Aws *GoogleBigqueryConnectionAws `field:"optional" json:"aws" yaml:"aws"`
	// azure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#azure GoogleBigqueryConnection#azure}
	Azure *GoogleBigqueryConnectionAzure `field:"optional" json:"azure" yaml:"azure"`
	// cloud_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#cloud_resource GoogleBigqueryConnection#cloud_resource}
	CloudResource *GoogleBigqueryConnectionCloudResource `field:"optional" json:"cloudResource" yaml:"cloudResource"`
	// cloud_spanner block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#cloud_spanner GoogleBigqueryConnection#cloud_spanner}
	CloudSpanner *GoogleBigqueryConnectionCloudSpanner `field:"optional" json:"cloudSpanner" yaml:"cloudSpanner"`
	// cloud_sql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#cloud_sql GoogleBigqueryConnection#cloud_sql}
	CloudSql *GoogleBigqueryConnectionCloudSql `field:"optional" json:"cloudSql" yaml:"cloudSql"`
	// Optional connection id that should be assigned to the created connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#connection_id GoogleBigqueryConnection#connection_id}
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// A descriptive description for the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#description GoogleBigqueryConnection#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A descriptive name for the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#friendly_name GoogleBigqueryConnection#friendly_name}
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#id GoogleBigqueryConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. The Cloud KMS key that is used for encryption.
	//
	// Example: projects/[kms_project_id]/locations/[region]/keyRings/[key_region]/cryptoKeys/[key]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#kms_key_name GoogleBigqueryConnection#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// The geographic location where the connection should reside.
	//
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1.
	// Spanner Connections same as spanner region
	// AWS allowed regions are aws-us-east-1
	// Azure allowed regions are azure-eastus2
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#location GoogleBigqueryConnection#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#project GoogleBigqueryConnection#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// spark block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#spark GoogleBigqueryConnection#spark}
	Spark *GoogleBigqueryConnectionSpark `field:"optional" json:"spark" yaml:"spark"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_connection#timeouts GoogleBigqueryConnection#timeouts}
	Timeouts *GoogleBigqueryConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

