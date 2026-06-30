package ecsservice

type EcsServiceAlarms struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ecs_service#alarm_names EcsService#alarm_names}.
	AlarmNames *[]*string `field:"required" json:"alarmNames" yaml:"alarmNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ecs_service#enable EcsService#enable}.
	Enable any `field:"required" json:"enable" yaml:"enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ecs_service#rollback EcsService#rollback}.
	Rollback any `field:"required" json:"rollback" yaml:"rollback"`
}
