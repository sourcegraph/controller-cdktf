package googlebigquerydatapolicyv2datapolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBigqueryDatapolicyv2DataPolicyConfig struct {
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
	// User-assigned (human readable) ID of the data policy that needs to be unique within a project.
	//
	// Used as {data_policy_id} in part of the resource
	// name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy#data_policy_id GoogleBigqueryDatapolicyv2DataPolicy#data_policy_id}
	DataPolicyId *string `field:"required" json:"dataPolicyId" yaml:"dataPolicyId"`
	// Type of data policy. Possible values: DATA_MASKING_POLICY RAW_DATA_ACCESS_POLICY COLUMN_LEVEL_SECURITY_POLICY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy#data_policy_type GoogleBigqueryDatapolicyv2DataPolicy#data_policy_type}
	DataPolicyType *string `field:"required" json:"dataPolicyType" yaml:"dataPolicyType"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy#location GoogleBigqueryDatapolicyv2DataPolicy#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// data_masking_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy#data_masking_policy GoogleBigqueryDatapolicyv2DataPolicy#data_masking_policy}
	DataMaskingPolicy *GoogleBigqueryDatapolicyv2DataPolicyDataMaskingPolicy `field:"optional" json:"dataMaskingPolicy" yaml:"dataMaskingPolicy"`
	// The etag for this Data Policy.
	//
	// This field is used for UpdateDataPolicy calls. If Data Policy exists, this
	// field is required and must match the server's etag. It will also be
	// populated in the response of GetDataPolicy, CreateDataPolicy, and
	// UpdateDataPolicy calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy#etag GoogleBigqueryDatapolicyv2DataPolicy#etag}
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// The list of IAM principals that have Fine Grained Access to the underlying data goverened by this data policy.
	//
	// Uses the [IAM V2 principal
	// syntax](https://cloud.google.com/iam/docs/principal-identifiers#v2) Only
	// supports principal types users, groups, serviceaccounts, cloudidentity.
	// This field is supported in V2 Data Policy only. In case of V1 data policies
	// (i.e. verion = 1 and policy_tag is set), this field is not populated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy#grantees GoogleBigqueryDatapolicyv2DataPolicy#grantees}
	Grantees *[]*string `field:"optional" json:"grantees" yaml:"grantees"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy#id GoogleBigqueryDatapolicyv2DataPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy#project GoogleBigqueryDatapolicyv2DataPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_datapolicyv2_data_policy#timeouts GoogleBigqueryDatapolicyv2DataPolicy#timeouts}
	Timeouts *GoogleBigqueryDatapolicyv2DataPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

