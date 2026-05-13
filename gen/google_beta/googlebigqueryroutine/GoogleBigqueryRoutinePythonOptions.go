package googlebigqueryroutine


type GoogleBigqueryRoutinePythonOptions struct {
	// The name of the function defined in Python code as the entry point when the Python UDF is invoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_routine#entry_point GoogleBigqueryRoutine#entry_point}
	EntryPoint *string `field:"required" json:"entryPoint" yaml:"entryPoint"`
	// A list of Python package names along with versions to be installed. Example: ["pandas>=2.1", "google-cloud-translate==3.11"]. For more information, see [Use third-party packages](https://cloud.google.com/bigquery/docs/user-defined-functions-python#third-party-packages).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_routine#packages GoogleBigqueryRoutine#packages}
	Packages *[]*string `field:"optional" json:"packages" yaml:"packages"`
}

