package persistentvolume


type PersistentVolumeSpec struct {
	// Contains all ways the volume can be mounted. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#access-modes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume#access_modes PersistentVolume#access_modes}
	AccessModes *[]*string `field:"required" json:"accessModes" yaml:"accessModes"`
	// A description of the persistent volume's resources and capacity. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#capacity.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume#capacity PersistentVolume#capacity}
	Capacity *map[string]*string `field:"required" json:"capacity" yaml:"capacity"`
	// persistent_volume_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume#persistent_volume_source PersistentVolume#persistent_volume_source}
	PersistentVolumeSource *PersistentVolumeSpecPersistentVolumeSource `field:"required" json:"persistentVolumeSource" yaml:"persistentVolumeSource"`
	// claim_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume#claim_ref PersistentVolume#claim_ref}
	ClaimRef *PersistentVolumeSpecClaimRef `field:"optional" json:"claimRef" yaml:"claimRef"`
	// A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume#mount_options PersistentVolume#mount_options}
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// node_affinity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume#node_affinity PersistentVolume#node_affinity}
	NodeAffinity *PersistentVolumeSpecNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	// What happens to a persistent volume when released from its claim.
	//
	// Valid options are Retain (default) and Recycle. Recycling must be supported by the volume plugin underlying this persistent volume. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#recycling-policy
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume#persistent_volume_reclaim_policy PersistentVolume#persistent_volume_reclaim_policy}
	PersistentVolumeReclaimPolicy *string `field:"optional" json:"persistentVolumeReclaimPolicy" yaml:"persistentVolumeReclaimPolicy"`
	// A description of the persistent volume's class. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#class.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume#storage_class_name PersistentVolume#storage_class_name}
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// Defines if a volume is intended to be used with a formatted filesystem.
	//
	// or to remain in raw block state.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/persistent_volume#volume_mode PersistentVolume#volume_mode}
	VolumeMode *string `field:"optional" json:"volumeMode" yaml:"volumeMode"`
}
