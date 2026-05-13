package googlebigqueryroutine


type GoogleBigqueryRoutineExternalRuntimeOptions struct {
	// Amount of CPU provisioned for a Python UDF container instance. For more information, see [Configure container limits for Python UDFs](https://cloud.google.com/bigquery/docs/user-defined-functions-python#configure-container-limits).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_routine#container_cpu GoogleBigqueryRoutine#container_cpu}
	ContainerCpu *float64 `field:"optional" json:"containerCpu" yaml:"containerCpu"`
	// Amount of memory provisioned for a Python UDF container instance.
	//
	// Format:
	// {number}{unit} where unit is one of "M", "G", "Mi" and "Gi" (e.g. 1G,
	// 512Mi). If not specified, the default value is 512Mi. For more information,
	// see [Configure container limits for Python
	// UDFs](https://cloud.google.com/bigquery/docs/user-defined-functions-python#configure-container-limits)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_routine#container_memory GoogleBigqueryRoutine#container_memory}
	ContainerMemory *string `field:"optional" json:"containerMemory" yaml:"containerMemory"`
	// Maximum number of rows in each batch sent to the external runtime.
	//
	// If
	// absent or if 0, BigQuery dynamically decides the number of rows in a batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_routine#max_batching_rows GoogleBigqueryRoutine#max_batching_rows}
	MaxBatchingRows *string `field:"optional" json:"maxBatchingRows" yaml:"maxBatchingRows"`
	// Fully qualified name of the connection whose service account will be used to execute the code in the container. Format: '"projects/{project_id}/locations/{location_id}/connections/{connection_id}"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_routine#runtime_connection GoogleBigqueryRoutine#runtime_connection}
	RuntimeConnection *string `field:"optional" json:"runtimeConnection" yaml:"runtimeConnection"`
	// Language runtime version. Example: 'python-3.11'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_routine#runtime_version GoogleBigqueryRoutine#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
}

