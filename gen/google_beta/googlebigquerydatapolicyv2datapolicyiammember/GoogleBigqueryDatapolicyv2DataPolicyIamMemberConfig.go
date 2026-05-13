package googlebigquerydatapolicyv2datapolicyiammember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBigqueryDatapolicyv2DataPolicyIamMemberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy_iam_member#data_policy_id GoogleBigqueryDatapolicyv2DataPolicyIamMember#data_policy_id}.
	DataPolicyId *string `field:"required" json:"dataPolicyId" yaml:"dataPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy_iam_member#member GoogleBigqueryDatapolicyv2DataPolicyIamMember#member}.
	Member *string `field:"required" json:"member" yaml:"member"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy_iam_member#role GoogleBigqueryDatapolicyv2DataPolicyIamMember#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy_iam_member#condition GoogleBigqueryDatapolicyv2DataPolicyIamMember#condition}
	Condition *GoogleBigqueryDatapolicyv2DataPolicyIamMemberCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy_iam_member#id GoogleBigqueryDatapolicyv2DataPolicyIamMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy_iam_member#location GoogleBigqueryDatapolicyv2DataPolicyIamMember#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy_iam_member#project GoogleBigqueryDatapolicyv2DataPolicyIamMember#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

