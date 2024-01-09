package alertpolicy


type AlertPolicyAlertMethod struct {
	// The name of the previously defined alert method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/alert_policy#name AlertPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Project name the Alert Method is in, must conform to the naming convention from [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names). If not defined, Nobl9 returns a default value for this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/alert_policy#project AlertPolicy#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

