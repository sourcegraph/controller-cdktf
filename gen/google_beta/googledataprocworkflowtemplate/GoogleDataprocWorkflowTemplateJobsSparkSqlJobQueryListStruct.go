package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryListStruct struct {
	// Required.
	//
	// The queries to execute. You do not need to end a query expression with a semicolon. Multiple queries can be specified in one string by separating each with a semicolon. Here is an example of a Dataproc API snippet that uses a QueryList to specify a HiveJob: "hiveJob": { "queryList": { "queries": [ "query1", "query2", "query3;query4", ] } }
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#queries GoogleDataprocWorkflowTemplate#queries}
	Queries *[]*string `field:"required" json:"queries" yaml:"queries"`
}

