package service


type ServiceStatus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/service#slo_count Service#slo_count}.
	SloCount *float64 `field:"optional" json:"sloCount" yaml:"sloCount"`
}

