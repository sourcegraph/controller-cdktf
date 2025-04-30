package googledataprocbatch

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataprocBatchConfig struct {
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
	// The ID to use for the batch, which will become the final component of the batch's resource name.
	//
	// This value must be 4-63 characters. Valid characters are /[a-z][0-9]-/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#batch_id GoogleDataprocBatch#batch_id}
	BatchId *string `field:"optional" json:"batchId" yaml:"batchId"`
	// environment_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#environment_config GoogleDataprocBatch#environment_config}
	EnvironmentConfig *GoogleDataprocBatchEnvironmentConfig `field:"optional" json:"environmentConfig" yaml:"environmentConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#id GoogleDataprocBatch#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels to associate with this batch.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#labels GoogleDataprocBatch#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location in which the batch will be created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#location GoogleDataprocBatch#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#project GoogleDataprocBatch#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// pyspark_batch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#pyspark_batch GoogleDataprocBatch#pyspark_batch}
	PysparkBatch *GoogleDataprocBatchPysparkBatch `field:"optional" json:"pysparkBatch" yaml:"pysparkBatch"`
	// runtime_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#runtime_config GoogleDataprocBatch#runtime_config}
	RuntimeConfig *GoogleDataprocBatchRuntimeConfig `field:"optional" json:"runtimeConfig" yaml:"runtimeConfig"`
	// spark_batch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#spark_batch GoogleDataprocBatch#spark_batch}
	SparkBatch *GoogleDataprocBatchSparkBatch `field:"optional" json:"sparkBatch" yaml:"sparkBatch"`
	// spark_r_batch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#spark_r_batch GoogleDataprocBatch#spark_r_batch}
	SparkRBatch *GoogleDataprocBatchSparkRBatch `field:"optional" json:"sparkRBatch" yaml:"sparkRBatch"`
	// spark_sql_batch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#spark_sql_batch GoogleDataprocBatch#spark_sql_batch}
	SparkSqlBatch *GoogleDataprocBatchSparkSqlBatch `field:"optional" json:"sparkSqlBatch" yaml:"sparkSqlBatch"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#timeouts GoogleDataprocBatch#timeouts}
	Timeouts *GoogleDataprocBatchTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

