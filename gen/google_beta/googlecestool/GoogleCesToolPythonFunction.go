package googlecestool


type GoogleCesToolPythonFunction struct {
	// The name of the Python function to execute.
	//
	// Must match a Python function
	// name defined in the python code. Case sensitive. If the name is not
	// provided, the first function defined in the python code will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#name GoogleCesTool#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Python code to execute for the tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#python_code GoogleCesTool#python_code}
	PythonCode *string `field:"optional" json:"pythonCode" yaml:"pythonCode"`
}

