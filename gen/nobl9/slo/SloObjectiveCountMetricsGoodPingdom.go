package slo


type SloObjectiveCountMetricsGoodPingdom struct {
	// Pingdom uptime or transaction check's ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#check_id Slo#check_id}
	CheckId *string `field:"required" json:"checkId" yaml:"checkId"`
	// Pingdom check type - uptime or transaction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#check_type Slo#check_type}
	CheckType *string `field:"optional" json:"checkType" yaml:"checkType"`
	// Optional for the Uptime checks. Use it to filter the Pingdom check results by status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#status Slo#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

