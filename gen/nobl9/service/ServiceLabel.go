package service


type ServiceLabel struct {
	// A key for the label, unique within the associated resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/service#key Service#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A set of values for a single key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/service#values Service#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

