package functionresource


type FunctionResourceArg struct {
	// The argument type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#type FunctionResource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// An expression to be used as default value if the parameter is not specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#default FunctionResource#default}
	Default *string `field:"optional" json:"default" yaml:"default"`
	// The argument mode. One of: IN, OUT, INOUT, or VARIADIC.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#mode FunctionResource#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The argument name. The name may be required for some languages or depending on the argument mode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#name FunctionResource#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

