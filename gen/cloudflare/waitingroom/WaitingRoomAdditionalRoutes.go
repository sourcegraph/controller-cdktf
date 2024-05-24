package waitingroom


type WaitingRoomAdditionalRoutes struct {
	// The additional host name for which the waiting room to be applied on (no wildcards).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/waiting_room#host WaitingRoom#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The path within the additional host to enable the waiting room on. Defaults to `/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/waiting_room#path WaitingRoom#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

