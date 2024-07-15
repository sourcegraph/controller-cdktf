package googledataprocjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataprocJobConfig struct {
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
	// placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#placement GoogleDataprocJob#placement}
	Placement *GoogleDataprocJobPlacement `field:"required" json:"placement" yaml:"placement"`
	// By default, you can only delete inactive jobs within Dataproc.
	//
	// Setting this to true, and calling destroy, will ensure that the job is first cancelled before issuing the delete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#force_delete GoogleDataprocJob#force_delete}
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// hadoop_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#hadoop_config GoogleDataprocJob#hadoop_config}
	HadoopConfig *GoogleDataprocJobHadoopConfig `field:"optional" json:"hadoopConfig" yaml:"hadoopConfig"`
	// hive_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#hive_config GoogleDataprocJob#hive_config}
	HiveConfig *GoogleDataprocJobHiveConfig `field:"optional" json:"hiveConfig" yaml:"hiveConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#id GoogleDataprocJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. The labels to associate with this job.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// 				Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#labels GoogleDataprocJob#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// pig_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#pig_config GoogleDataprocJob#pig_config}
	PigConfig *GoogleDataprocJobPigConfig `field:"optional" json:"pigConfig" yaml:"pigConfig"`
	// presto_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#presto_config GoogleDataprocJob#presto_config}
	PrestoConfig *GoogleDataprocJobPrestoConfig `field:"optional" json:"prestoConfig" yaml:"prestoConfig"`
	// The project in which the cluster can be found and jobs subsequently run against.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#project GoogleDataprocJob#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// pyspark_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#pyspark_config GoogleDataprocJob#pyspark_config}
	PysparkConfig *GoogleDataprocJobPysparkConfig `field:"optional" json:"pysparkConfig" yaml:"pysparkConfig"`
	// reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#reference GoogleDataprocJob#reference}
	Reference *GoogleDataprocJobReference `field:"optional" json:"reference" yaml:"reference"`
	// The Cloud Dataproc region.
	//
	// This essentially determines which clusters are available for this job to be submitted to. If not specified, defaults to global.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#region GoogleDataprocJob#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// scheduling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#scheduling GoogleDataprocJob#scheduling}
	Scheduling *GoogleDataprocJobScheduling `field:"optional" json:"scheduling" yaml:"scheduling"`
	// spark_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#spark_config GoogleDataprocJob#spark_config}
	SparkConfig *GoogleDataprocJobSparkConfig `field:"optional" json:"sparkConfig" yaml:"sparkConfig"`
	// sparksql_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#sparksql_config GoogleDataprocJob#sparksql_config}
	SparksqlConfig *GoogleDataprocJobSparksqlConfig `field:"optional" json:"sparksqlConfig" yaml:"sparksqlConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#timeouts GoogleDataprocJob#timeouts}
	Timeouts *GoogleDataprocJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

