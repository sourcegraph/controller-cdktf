package googlenetappvolume


type GoogleNetappVolumeCacheParametersCacheConfig struct {
	// Optional. Flag indicating whether a CIFS change notification is enabled for the FlexCache volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_volume#cifs_change_notify_enabled GoogleNetappVolume#cifs_change_notify_enabled}
	CifsChangeNotifyEnabled interface{} `field:"optional" json:"cifsChangeNotifyEnabled" yaml:"cifsChangeNotifyEnabled"`
}

