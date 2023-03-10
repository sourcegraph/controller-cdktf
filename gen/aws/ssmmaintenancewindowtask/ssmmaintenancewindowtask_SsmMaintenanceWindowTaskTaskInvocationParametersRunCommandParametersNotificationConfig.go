package ssmmaintenancewindowtask


type SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#notification_arn SsmMaintenanceWindowTask#notification_arn}.
	NotificationArn *string `field:"optional" json:"notificationArn" yaml:"notificationArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#notification_events SsmMaintenanceWindowTask#notification_events}.
	NotificationEvents *[]*string `field:"optional" json:"notificationEvents" yaml:"notificationEvents"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssm_maintenance_window_task#notification_type SsmMaintenanceWindowTask#notification_type}.
	NotificationType *string `field:"optional" json:"notificationType" yaml:"notificationType"`
}

