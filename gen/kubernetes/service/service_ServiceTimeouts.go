package service


type ServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/service#create Service#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

