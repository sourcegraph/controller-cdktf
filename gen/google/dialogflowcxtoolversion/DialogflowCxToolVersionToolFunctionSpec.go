package dialogflowcxtoolversion


type DialogflowCxToolVersionToolFunctionSpec struct {
	// Optional.
	//
	// The JSON schema is encapsulated in a [google.protobuf.Struct](https://protobuf.dev/reference/protobuf/google.protobuf/#struct) to describe the input of the function.
	// This input is a JSON object that contains the function's parameters as properties of the object
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#input_schema DialogflowCxToolVersion#input_schema}
	InputSchema *string `field:"optional" json:"inputSchema" yaml:"inputSchema"`
	// Optional.
	//
	// The JSON schema is encapsulated in a [google.protobuf.Struct](https://protobuf.dev/reference/protobuf/google.protobuf/#struct) to describe the output of the function.
	// This output is a JSON object that contains the function's parameters as properties of the object
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#output_schema DialogflowCxToolVersion#output_schema}
	OutputSchema *string `field:"optional" json:"outputSchema" yaml:"outputSchema"`
}

