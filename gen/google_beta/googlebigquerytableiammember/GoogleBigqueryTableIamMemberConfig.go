package googlebigquerytableiammember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBigqueryTableIamMemberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table_iam_member#dataset_id GoogleBigqueryTableIamMember#dataset_id}.
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table_iam_member#member GoogleBigqueryTableIamMember#member}.
	Member *string `field:"required" json:"member" yaml:"member"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table_iam_member#role GoogleBigqueryTableIamMember#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table_iam_member#table_id GoogleBigqueryTableIamMember#table_id}.
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table_iam_member#condition GoogleBigqueryTableIamMember#condition}
	Condition *GoogleBigqueryTableIamMemberCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table_iam_member#id GoogleBigqueryTableIamMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table_iam_member#project GoogleBigqueryTableIamMember#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

