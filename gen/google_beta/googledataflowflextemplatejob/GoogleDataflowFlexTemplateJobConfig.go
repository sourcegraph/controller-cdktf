package googledataflowflextemplatejob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataflowFlexTemplateJobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataflow_flex_template_job#container_spec_gcs_path GoogleDataflowFlexTemplateJob#container_spec_gcs_path}.
	ContainerSpecGcsPath *string `field:"required" json:"containerSpecGcsPath" yaml:"containerSpecGcsPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataflow_flex_template_job#name GoogleDataflowFlexTemplateJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataflow_flex_template_job#id GoogleDataflowFlexTemplateJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataflow_flex_template_job#labels GoogleDataflowFlexTemplateJob#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataflow_flex_template_job#on_delete GoogleDataflowFlexTemplateJob#on_delete}.
	OnDelete *string `field:"optional" json:"onDelete" yaml:"onDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataflow_flex_template_job#parameters GoogleDataflowFlexTemplateJob#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataflow_flex_template_job#project GoogleDataflowFlexTemplateJob#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region in which the created job should run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataflow_flex_template_job#region GoogleDataflowFlexTemplateJob#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// If true, treat DRAINING and CANCELLING as terminal job states and do not wait for further changes before removing from terraform state and moving on.
	//
	// WARNING: this will lead to job name conflicts if you do not ensure that the job names are different, e.g. by embedding a release ID or by using a random_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_dataflow_flex_template_job#skip_wait_on_job_termination GoogleDataflowFlexTemplateJob#skip_wait_on_job_termination}
	SkipWaitOnJobTermination interface{} `field:"optional" json:"skipWaitOnJobTermination" yaml:"skipWaitOnJobTermination"`
}

