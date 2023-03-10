package googlegameservicesgameserverconfig


type GoogleGameServicesGameServerConfigFleetConfigs struct {
	// The fleet spec, which is sent to Agones to configure fleet.
	//
	// The spec can be passed as inline json but it is recommended to use a file reference
	// instead. File references can contain the json or yaml format of the fleet spec. Eg:
	//
	// fleet_spec = jsonencode(yamldecode(file("fleet_configs.yaml")))
	// fleet_spec = file("fleet_configs.json")
	//
	// The format of the spec can be found :
	// 'https://agones.dev/site/docs/reference/fleet/'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_game_services_game_server_config#fleet_spec GoogleGameServicesGameServerConfig#fleet_spec}
	FleetSpec *string `field:"required" json:"fleetSpec" yaml:"fleetSpec"`
	// The name of the FleetConfig.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_game_services_game_server_config#name GoogleGameServicesGameServerConfig#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

