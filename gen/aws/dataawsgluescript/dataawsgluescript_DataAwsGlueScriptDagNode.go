package dataawsgluescript


type DataAwsGlueScriptDagNode struct {
	// args block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script#args DataAwsGlueScript#args}
	Args interface{} `field:"required" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script#id DataAwsGlueScript#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script#node_type DataAwsGlueScript#node_type}.
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/glue_script#line_number DataAwsGlueScript#line_number}.
	LineNumber *float64 `field:"optional" json:"lineNumber" yaml:"lineNumber"`
}

