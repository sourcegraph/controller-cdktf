package googlebigquerydatasetaccess

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBigqueryDatasetAccessAConfig struct {
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
	// A unique ID for this dataset, without the project name.
	//
	// The ID
	// must contain only letters (a-z, A-Z), numbers (0-9), or
	// underscores (_). The maximum length is 1,024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#dataset_id GoogleBigqueryDatasetAccessA#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// dataset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#dataset GoogleBigqueryDatasetAccessA#dataset}
	Dataset *GoogleBigqueryDatasetAccessDatasetA `field:"optional" json:"dataset" yaml:"dataset"`
	// A domain to grant access to. Any users signed in with the domain specified will be granted the specified access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#domain GoogleBigqueryDatasetAccessA#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// An email address of a Google Group to grant access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#group_by_email GoogleBigqueryDatasetAccessA#group_by_email}
	GroupByEmail *string `field:"optional" json:"groupByEmail" yaml:"groupByEmail"`
	// Some other type of member that appears in the IAM Policy but isn't a user, group, domain, or special group.
	//
	// For example: 'allUsers'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#iam_member GoogleBigqueryDatasetAccessA#iam_member}
	IamMember *string `field:"optional" json:"iamMember" yaml:"iamMember"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#id GoogleBigqueryDatasetAccessA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#project GoogleBigqueryDatasetAccessA#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Describes the rights granted to the user specified by the other member of the access object.
	//
	// Basic, predefined, and custom roles are
	// supported. Predefined roles that have equivalent basic roles are
	// swapped by the API to their basic counterparts, and will show a diff
	// post-create. See
	// [official docs](https://cloud.google.com/bigquery/docs/access-control).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#role GoogleBigqueryDatasetAccessA#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// routine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#routine GoogleBigqueryDatasetAccessA#routine}
	Routine *GoogleBigqueryDatasetAccessRoutineA `field:"optional" json:"routine" yaml:"routine"`
	// A special group to grant access to.
	//
	// Possible values include:
	// * 'projectOwners': Owners of the enclosing project.
	// * 'projectReaders': Readers of the enclosing project.
	// * 'projectWriters': Writers of the enclosing project.
	// * 'allAuthenticatedUsers': All authenticated BigQuery users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#special_group GoogleBigqueryDatasetAccessA#special_group}
	SpecialGroup *string `field:"optional" json:"specialGroup" yaml:"specialGroup"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#timeouts GoogleBigqueryDatasetAccessA#timeouts}
	Timeouts *GoogleBigqueryDatasetAccessTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// An email address of a user to grant access to. For example: fred@example.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#user_by_email GoogleBigqueryDatasetAccessA#user_by_email}
	UserByEmail *string `field:"optional" json:"userByEmail" yaml:"userByEmail"`
	// view block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#view GoogleBigqueryDatasetAccessA#view}
	View *GoogleBigqueryDatasetAccessViewA `field:"optional" json:"view" yaml:"view"`
}

