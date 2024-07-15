package googleloggingprojectsink

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleLoggingProjectSinkConfig struct {
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
	// The destination of the sink (or, in other words, where logs are written to).
	//
	// Can be a Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples: "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" The writer associated with the sink must have access to write to the above resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_project_sink#destination GoogleLoggingProjectSink#destination}
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The name of the logging sink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_project_sink#name GoogleLoggingProjectSink#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// bigquery_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_project_sink#bigquery_options GoogleLoggingProjectSink#bigquery_options}
	BigqueryOptions *GoogleLoggingProjectSinkBigqueryOptions `field:"optional" json:"bigqueryOptions" yaml:"bigqueryOptions"`
	// A service account provided by the caller that will be used to write the log entries.
	//
	// The format must be serviceAccount:some@email. This field can only be specified if you are routing logs to a destination outside this sink's project. If not specified, a Logging service account will automatically be generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_project_sink#custom_writer_identity GoogleLoggingProjectSink#custom_writer_identity}
	CustomWriterIdentity *string `field:"optional" json:"customWriterIdentity" yaml:"customWriterIdentity"`
	// A description of this sink. The maximum length of the description is 8000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_project_sink#description GoogleLoggingProjectSink#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set to True, then this sink is disabled and it does not export any log entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_project_sink#disabled GoogleLoggingProjectSink#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// exclusions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_project_sink#exclusions GoogleLoggingProjectSink#exclusions}
	Exclusions interface{} `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_project_sink#filter GoogleLoggingProjectSink#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_project_sink#id GoogleLoggingProjectSink#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the project to create the sink in.
	//
	// If omitted, the project associated with the provider is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_project_sink#project GoogleLoggingProjectSink#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Whether or not to create a unique identity associated with this sink.
	//
	// If false (the legacy behavior), then the writer_identity used is serviceAccount:cloud-logs@system.gserviceaccount.com. If true (default), then a unique service account is created and used for this sink. If you wish to publish logs across projects, you must set unique_writer_identity to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_project_sink#unique_writer_identity GoogleLoggingProjectSink#unique_writer_identity}
	UniqueWriterIdentity interface{} `field:"optional" json:"uniqueWriterIdentity" yaml:"uniqueWriterIdentity"`
}

