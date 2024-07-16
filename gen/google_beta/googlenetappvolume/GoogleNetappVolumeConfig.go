package googlenetappvolume

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetappVolumeConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Capacity of the volume (in GiB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#capacity_gib GoogleNetappVolume#capacity_gib}
	CapacityGib *string `field:"required" json:"capacityGib" yaml:"capacityGib"`
	// Name of the pool location.
	//
	// Usually a region name, expect for some STANDARD service level pools which require a zone name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#location GoogleNetappVolume#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the volume. Needs to be unique per location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#name GoogleNetappVolume#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The protocol of the volume.
	//
	// Allowed combinations are '['NFSV3']', '['NFSV4']', '['SMB']', '['NFSV3', 'NFSV4']', '['SMB', 'NFSV3']' and '['SMB', 'NFSV4']'. Possible values: ["NFSV3", "NFSV4", "SMB"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#protocols GoogleNetappVolume#protocols}
	Protocols *[]*string `field:"required" json:"protocols" yaml:"protocols"`
	// Share name (SMB) or export path (NFS) of the volume. Needs to be unique per location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#share_name GoogleNetappVolume#share_name}
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// Name of the storage pool to create the volume in. Pool needs enough spare capacity to accomodate the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#storage_pool GoogleNetappVolume#storage_pool}
	StoragePool *string `field:"required" json:"storagePool" yaml:"storagePool"`
	// backup_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#backup_config GoogleNetappVolume#backup_config}
	BackupConfig *GoogleNetappVolumeBackupConfig `field:"optional" json:"backupConfig" yaml:"backupConfig"`
	// Policy to determine if the volume should be deleted forcefully.
	//
	// Volumes may have nested snapshot resources. Deleting such a volume will fail.
	// Setting this parameter to FORCE will delete volumes including nested snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#deletion_policy GoogleNetappVolume#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#description GoogleNetappVolume#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// export_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#export_policy GoogleNetappVolume#export_policy}
	ExportPolicy *GoogleNetappVolumeExportPolicy `field:"optional" json:"exportPolicy" yaml:"exportPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#id GoogleNetappVolume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Flag indicating if the volume is a kerberos volume or not, export policy rules control kerberos security modes (krb5, krb5i, krb5p).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#kerberos_enabled GoogleNetappVolume#kerberos_enabled}
	KerberosEnabled interface{} `field:"optional" json:"kerberosEnabled" yaml:"kerberosEnabled"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#labels GoogleNetappVolume#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#project GoogleNetappVolume#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// restore_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#restore_parameters GoogleNetappVolume#restore_parameters}
	RestoreParameters *GoogleNetappVolumeRestoreParameters `field:"optional" json:"restoreParameters" yaml:"restoreParameters"`
	// List of actions that are restricted on this volume. Possible values: ["DELETE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#restricted_actions GoogleNetappVolume#restricted_actions}
	RestrictedActions *[]*string `field:"optional" json:"restrictedActions" yaml:"restrictedActions"`
	// Security Style of the Volume.
	//
	// Use UNIX to use UNIX or NFSV4 ACLs for file permissions.
	// Use NTFS to use NTFS ACLs for file permissions. Can only be set for volumes which use SMB together with NFS as protocol. Possible values: ["NTFS", "UNIX"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#security_style GoogleNetappVolume#security_style}
	SecurityStyle *string `field:"optional" json:"securityStyle" yaml:"securityStyle"`
	// Settings for volumes with SMB access. Possible values: ["ENCRYPT_DATA", "BROWSABLE", "CHANGE_NOTIFY", "NON_BROWSABLE", "OPLOCKS", "SHOW_SNAPSHOT", "SHOW_PREVIOUS_VERSIONS", "ACCESS_BASED_ENUMERATION", "CONTINUOUSLY_AVAILABLE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#smb_settings GoogleNetappVolume#smb_settings}
	SmbSettings *[]*string `field:"optional" json:"smbSettings" yaml:"smbSettings"`
	// If enabled, a NFS volume will contain a read-only .snapshot directory which provides access to each of the volume's snapshots. Will enable "Previous Versions" support for SMB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#snapshot_directory GoogleNetappVolume#snapshot_directory}
	SnapshotDirectory interface{} `field:"optional" json:"snapshotDirectory" yaml:"snapshotDirectory"`
	// snapshot_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#snapshot_policy GoogleNetappVolume#snapshot_policy}
	SnapshotPolicy *GoogleNetappVolumeSnapshotPolicy `field:"optional" json:"snapshotPolicy" yaml:"snapshotPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#timeouts GoogleNetappVolume#timeouts}
	Timeouts *GoogleNetappVolumeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Unix permission the mount point will be created with. Default is 0770. Applicable for UNIX security style volumes only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#unix_permissions GoogleNetappVolume#unix_permissions}
	UnixPermissions *string `field:"optional" json:"unixPermissions" yaml:"unixPermissions"`
}

