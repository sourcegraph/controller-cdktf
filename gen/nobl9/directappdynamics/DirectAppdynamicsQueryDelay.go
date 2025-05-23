package directappdynamics


type DirectAppdynamicsQueryDelay struct {
	// Must be one of Minute or Second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_appdynamics#unit DirectAppdynamics#unit}
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Must be an integer greater than or equal to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/direct_appdynamics#value DirectAppdynamics#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

