package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplateJobs struct {
	// Required.
	//
	// The step id. The id must be unique among all jobs within the template. The step id is used as prefix for job id, as job `goog-dataproc-workflow-step-id` label, and in prerequisiteStepIds field from other steps. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#step_id GoogleDataprocWorkflowTemplate#step_id}
	StepId *string `field:"required" json:"stepId" yaml:"stepId"`
	// hadoop_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#hadoop_job GoogleDataprocWorkflowTemplate#hadoop_job}
	HadoopJob *GoogleDataprocWorkflowTemplateJobsHadoopJob `field:"optional" json:"hadoopJob" yaml:"hadoopJob"`
	// hive_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#hive_job GoogleDataprocWorkflowTemplate#hive_job}
	HiveJob *GoogleDataprocWorkflowTemplateJobsHiveJob `field:"optional" json:"hiveJob" yaml:"hiveJob"`
	// Optional.
	//
	// The labels to associate with this job. Label keys must be between 1 and 63 characters long, and must conform to the following regular expression: p{Ll}p{Lo}{0,62} Label values must be between 1 and 63 characters long, and must conform to the following regular expression: [p{Ll}p{Lo}p{N}_-]{0,63} No more than 32 labels can be associated with a given job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#labels GoogleDataprocWorkflowTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// pig_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#pig_job GoogleDataprocWorkflowTemplate#pig_job}
	PigJob *GoogleDataprocWorkflowTemplateJobsPigJob `field:"optional" json:"pigJob" yaml:"pigJob"`
	// Optional.
	//
	// The optional list of prerequisite job step_ids. If not specified, the job will start at the beginning of workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#prerequisite_step_ids GoogleDataprocWorkflowTemplate#prerequisite_step_ids}
	PrerequisiteStepIds *[]*string `field:"optional" json:"prerequisiteStepIds" yaml:"prerequisiteStepIds"`
	// presto_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#presto_job GoogleDataprocWorkflowTemplate#presto_job}
	PrestoJob *GoogleDataprocWorkflowTemplateJobsPrestoJob `field:"optional" json:"prestoJob" yaml:"prestoJob"`
	// pyspark_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#pyspark_job GoogleDataprocWorkflowTemplate#pyspark_job}
	PysparkJob *GoogleDataprocWorkflowTemplateJobsPysparkJob `field:"optional" json:"pysparkJob" yaml:"pysparkJob"`
	// scheduling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#scheduling GoogleDataprocWorkflowTemplate#scheduling}
	Scheduling *GoogleDataprocWorkflowTemplateJobsScheduling `field:"optional" json:"scheduling" yaml:"scheduling"`
	// spark_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#spark_job GoogleDataprocWorkflowTemplate#spark_job}
	SparkJob *GoogleDataprocWorkflowTemplateJobsSparkJob `field:"optional" json:"sparkJob" yaml:"sparkJob"`
	// spark_r_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#spark_r_job GoogleDataprocWorkflowTemplate#spark_r_job}
	SparkRJob *GoogleDataprocWorkflowTemplateJobsSparkRJob `field:"optional" json:"sparkRJob" yaml:"sparkRJob"`
	// spark_sql_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#spark_sql_job GoogleDataprocWorkflowTemplate#spark_sql_job}
	SparkSqlJob *GoogleDataprocWorkflowTemplateJobsSparkSqlJob `field:"optional" json:"sparkSqlJob" yaml:"sparkSqlJob"`
}

