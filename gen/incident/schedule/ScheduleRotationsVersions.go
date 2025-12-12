package schedule


type ScheduleRotationsVersions struct {
	// Defines the handover intervals for this rota, in order they should apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/schedule#handovers Schedule#handovers}
	Handovers interface{} `field:"required" json:"handovers" yaml:"handovers"`
	// Defines the next moment we'll trigger a handover.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/schedule#handover_start_at Schedule#handover_start_at}
	HandoverStartAt *string `field:"required" json:"handoverStartAt" yaml:"handoverStartAt"`
	// Controls how many people are on-call concurrently.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/schedule#layers Schedule#layers}
	Layers interface{} `field:"required" json:"layers" yaml:"layers"`
	// The incident.io ID of a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/schedule#users Schedule#users}
	Users *[]*string `field:"required" json:"users" yaml:"users"`
	// When this rotation config will be effective from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/schedule#effective_from Schedule#effective_from}
	EffectiveFrom *string `field:"optional" json:"effectiveFrom" yaml:"effectiveFrom"`
	// Optional restrictions that define when to schedule people for this rota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/schedule#working_intervals Schedule#working_intervals}
	WorkingIntervals interface{} `field:"optional" json:"workingIntervals" yaml:"workingIntervals"`
}

