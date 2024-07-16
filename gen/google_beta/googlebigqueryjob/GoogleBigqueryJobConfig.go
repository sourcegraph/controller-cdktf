package googlebigqueryjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBigqueryJobConfig struct {
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
	// The ID of the job.
	//
	// The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#job_id GoogleBigqueryJob#job_id}
	JobId *string `field:"required" json:"jobId" yaml:"jobId"`
	// copy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#copy GoogleBigqueryJob#copy}
	Copy *GoogleBigqueryJobCopy `field:"optional" json:"copy" yaml:"copy"`
	// extract block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#extract GoogleBigqueryJob#extract}
	Extract *GoogleBigqueryJobExtract `field:"optional" json:"extract" yaml:"extract"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#id GoogleBigqueryJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#job_timeout_ms GoogleBigqueryJob#job_timeout_ms}
	JobTimeoutMs *string `field:"optional" json:"jobTimeoutMs" yaml:"jobTimeoutMs"`
	// The labels associated with this job. You can use these to organize and group your jobs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#labels GoogleBigqueryJob#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// load block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#load GoogleBigqueryJob#load}
	Load *GoogleBigqueryJobLoad `field:"optional" json:"load" yaml:"load"`
	// The geographic location of the job. The default value is US.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#location GoogleBigqueryJob#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#project GoogleBigqueryJob#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#query GoogleBigqueryJob#query}
	Query *GoogleBigqueryJobQuery `field:"optional" json:"query" yaml:"query"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#timeouts GoogleBigqueryJob#timeouts}
	Timeouts *GoogleBigqueryJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

