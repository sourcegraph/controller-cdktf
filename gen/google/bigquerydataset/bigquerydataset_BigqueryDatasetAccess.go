package bigquerydataset


type BigqueryDatasetAccess struct {
	// dataset block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset#dataset BigqueryDataset#dataset}
	Dataset *BigqueryDatasetAccessDataset `field:"optional" json:"dataset" yaml:"dataset"`
	// A domain to grant access to. Any users signed in with the domain specified will be granted the specified access.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset#domain BigqueryDataset#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// An email address of a Google Group to grant access to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset#group_by_email BigqueryDataset#group_by_email}
	GroupByEmail *string `field:"optional" json:"groupByEmail" yaml:"groupByEmail"`
	// Describes the rights granted to the user specified by the other member of the access object.
	//
	// Basic, predefined, and custom roles
	// are supported. Predefined roles that have equivalent basic roles
	// are swapped by the API to their basic counterparts. See
	// [official docs](https://cloud.google.com/bigquery/docs/access-control).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset#role BigqueryDataset#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// A special group to grant access to. Possible values include:.
	//
	// 'projectOwners': Owners of the enclosing project.
	//
	//
	// 'projectReaders': Readers of the enclosing project.
	//
	//
	// 'projectWriters': Writers of the enclosing project.
	//
	//
	// 'allAuthenticatedUsers': All authenticated BigQuery users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset#special_group BigqueryDataset#special_group}
	SpecialGroup *string `field:"optional" json:"specialGroup" yaml:"specialGroup"`
	// An email address of a user to grant access to. For example: fred@example.com.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset#user_by_email BigqueryDataset#user_by_email}
	UserByEmail *string `field:"optional" json:"userByEmail" yaml:"userByEmail"`
	// view block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset#view BigqueryDataset#view}
	View *BigqueryDatasetAccessView `field:"optional" json:"view" yaml:"view"`
}

