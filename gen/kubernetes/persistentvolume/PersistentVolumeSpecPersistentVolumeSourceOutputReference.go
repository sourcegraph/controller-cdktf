package persistentvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/persistentvolume/internal"
)

type PersistentVolumeSpecPersistentVolumeSourceOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() PersistentVolumeSpecPersistentVolumeSourceAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *PersistentVolumeSpecPersistentVolumeSourceAwsElasticBlockStore
	AzureDisk() PersistentVolumeSpecPersistentVolumeSourceAzureDiskOutputReference
	AzureDiskInput() *PersistentVolumeSpecPersistentVolumeSourceAzureDisk
	AzureFile() PersistentVolumeSpecPersistentVolumeSourceAzureFileOutputReference
	AzureFileInput() *PersistentVolumeSpecPersistentVolumeSourceAzureFile
	CephFs() PersistentVolumeSpecPersistentVolumeSourceCephFsOutputReference
	CephFsInput() *PersistentVolumeSpecPersistentVolumeSourceCephFs
	Cinder() PersistentVolumeSpecPersistentVolumeSourceCinderOutputReference
	CinderInput() *PersistentVolumeSpecPersistentVolumeSourceCinder
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference
	CsiInput() *PersistentVolumeSpecPersistentVolumeSourceCsi
	Fc() PersistentVolumeSpecPersistentVolumeSourceFcOutputReference
	FcInput() *PersistentVolumeSpecPersistentVolumeSourceFc
	FlexVolume() PersistentVolumeSpecPersistentVolumeSourceFlexVolumeOutputReference
	FlexVolumeInput() *PersistentVolumeSpecPersistentVolumeSourceFlexVolume
	Flocker() PersistentVolumeSpecPersistentVolumeSourceFlockerOutputReference
	FlockerInput() *PersistentVolumeSpecPersistentVolumeSourceFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() PersistentVolumeSpecPersistentVolumeSourceGcePersistentDiskOutputReference
	GcePersistentDiskInput() *PersistentVolumeSpecPersistentVolumeSourceGcePersistentDisk
	Glusterfs() PersistentVolumeSpecPersistentVolumeSourceGlusterfsOutputReference
	GlusterfsInput() *PersistentVolumeSpecPersistentVolumeSourceGlusterfs
	HostPath() PersistentVolumeSpecPersistentVolumeSourceHostPathOutputReference
	HostPathInput() *PersistentVolumeSpecPersistentVolumeSourceHostPath
	InternalValue() *PersistentVolumeSpecPersistentVolumeSource
	SetInternalValue(val *PersistentVolumeSpecPersistentVolumeSource)
	Iscsi() PersistentVolumeSpecPersistentVolumeSourceIscsiOutputReference
	IscsiInput() *PersistentVolumeSpecPersistentVolumeSourceIscsi
	Local() PersistentVolumeSpecPersistentVolumeSourceLocalOutputReference
	LocalInput() *PersistentVolumeSpecPersistentVolumeSourceLocal
	Nfs() PersistentVolumeSpecPersistentVolumeSourceNfsOutputReference
	NfsInput() *PersistentVolumeSpecPersistentVolumeSourceNfs
	PhotonPersistentDisk() PersistentVolumeSpecPersistentVolumeSourcePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *PersistentVolumeSpecPersistentVolumeSourcePhotonPersistentDisk
	Quobyte() PersistentVolumeSpecPersistentVolumeSourceQuobyteOutputReference
	QuobyteInput() *PersistentVolumeSpecPersistentVolumeSourceQuobyte
	Rbd() PersistentVolumeSpecPersistentVolumeSourceRbdOutputReference
	RbdInput() *PersistentVolumeSpecPersistentVolumeSourceRbd
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() PersistentVolumeSpecPersistentVolumeSourceVsphereVolumeOutputReference
	VsphereVolumeInput() *PersistentVolumeSpecPersistentVolumeSourceVsphereVolume
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAwsElasticBlockStore(value *PersistentVolumeSpecPersistentVolumeSourceAwsElasticBlockStore)
	PutAzureDisk(value *PersistentVolumeSpecPersistentVolumeSourceAzureDisk)
	PutAzureFile(value *PersistentVolumeSpecPersistentVolumeSourceAzureFile)
	PutCephFs(value *PersistentVolumeSpecPersistentVolumeSourceCephFs)
	PutCinder(value *PersistentVolumeSpecPersistentVolumeSourceCinder)
	PutCsi(value *PersistentVolumeSpecPersistentVolumeSourceCsi)
	PutFc(value *PersistentVolumeSpecPersistentVolumeSourceFc)
	PutFlexVolume(value *PersistentVolumeSpecPersistentVolumeSourceFlexVolume)
	PutFlocker(value *PersistentVolumeSpecPersistentVolumeSourceFlocker)
	PutGcePersistentDisk(value *PersistentVolumeSpecPersistentVolumeSourceGcePersistentDisk)
	PutGlusterfs(value *PersistentVolumeSpecPersistentVolumeSourceGlusterfs)
	PutHostPath(value *PersistentVolumeSpecPersistentVolumeSourceHostPath)
	PutIscsi(value *PersistentVolumeSpecPersistentVolumeSourceIscsi)
	PutLocal(value *PersistentVolumeSpecPersistentVolumeSourceLocal)
	PutNfs(value *PersistentVolumeSpecPersistentVolumeSourceNfs)
	PutPhotonPersistentDisk(value *PersistentVolumeSpecPersistentVolumeSourcePhotonPersistentDisk)
	PutQuobyte(value *PersistentVolumeSpecPersistentVolumeSourceQuobyte)
	PutRbd(value *PersistentVolumeSpecPersistentVolumeSourceRbd)
	PutVsphereVolume(value *PersistentVolumeSpecPersistentVolumeSourceVsphereVolume)
	ResetAwsElasticBlockStore()
	ResetAzureDisk()
	ResetAzureFile()
	ResetCephFs()
	ResetCinder()
	ResetCsi()
	ResetFc()
	ResetFlexVolume()
	ResetFlocker()
	ResetGcePersistentDisk()
	ResetGlusterfs()
	ResetHostPath()
	ResetIscsi()
	ResetLocal()
	ResetNfs()
	ResetPhotonPersistentDisk()
	ResetQuobyte()
	ResetRbd()
	ResetVsphereVolume()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PersistentVolumeSpecPersistentVolumeSourceOutputReference
type jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) AwsElasticBlockStore() PersistentVolumeSpecPersistentVolumeSourceAwsElasticBlockStoreOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) AwsElasticBlockStoreInput() *PersistentVolumeSpecPersistentVolumeSourceAwsElasticBlockStore {
	var returns *PersistentVolumeSpecPersistentVolumeSourceAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) AzureDisk() PersistentVolumeSpecPersistentVolumeSourceAzureDiskOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) AzureDiskInput() *PersistentVolumeSpecPersistentVolumeSourceAzureDisk {
	var returns *PersistentVolumeSpecPersistentVolumeSourceAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) AzureFile() PersistentVolumeSpecPersistentVolumeSourceAzureFileOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) AzureFileInput() *PersistentVolumeSpecPersistentVolumeSourceAzureFile {
	var returns *PersistentVolumeSpecPersistentVolumeSourceAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) CephFs() PersistentVolumeSpecPersistentVolumeSourceCephFsOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) CephFsInput() *PersistentVolumeSpecPersistentVolumeSourceCephFs {
	var returns *PersistentVolumeSpecPersistentVolumeSourceCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) Cinder() PersistentVolumeSpecPersistentVolumeSourceCinderOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) CinderInput() *PersistentVolumeSpecPersistentVolumeSourceCinder {
	var returns *PersistentVolumeSpecPersistentVolumeSourceCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) Csi() PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) CsiInput() *PersistentVolumeSpecPersistentVolumeSourceCsi {
	var returns *PersistentVolumeSpecPersistentVolumeSourceCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) Fc() PersistentVolumeSpecPersistentVolumeSourceFcOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) FcInput() *PersistentVolumeSpecPersistentVolumeSourceFc {
	var returns *PersistentVolumeSpecPersistentVolumeSourceFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) FlexVolume() PersistentVolumeSpecPersistentVolumeSourceFlexVolumeOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) FlexVolumeInput() *PersistentVolumeSpecPersistentVolumeSourceFlexVolume {
	var returns *PersistentVolumeSpecPersistentVolumeSourceFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) Flocker() PersistentVolumeSpecPersistentVolumeSourceFlockerOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) FlockerInput() *PersistentVolumeSpecPersistentVolumeSourceFlocker {
	var returns *PersistentVolumeSpecPersistentVolumeSourceFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) GcePersistentDisk() PersistentVolumeSpecPersistentVolumeSourceGcePersistentDiskOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) GcePersistentDiskInput() *PersistentVolumeSpecPersistentVolumeSourceGcePersistentDisk {
	var returns *PersistentVolumeSpecPersistentVolumeSourceGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) Glusterfs() PersistentVolumeSpecPersistentVolumeSourceGlusterfsOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) GlusterfsInput() *PersistentVolumeSpecPersistentVolumeSourceGlusterfs {
	var returns *PersistentVolumeSpecPersistentVolumeSourceGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) HostPath() PersistentVolumeSpecPersistentVolumeSourceHostPathOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) HostPathInput() *PersistentVolumeSpecPersistentVolumeSourceHostPath {
	var returns *PersistentVolumeSpecPersistentVolumeSourceHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) InternalValue() *PersistentVolumeSpecPersistentVolumeSource {
	var returns *PersistentVolumeSpecPersistentVolumeSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) Iscsi() PersistentVolumeSpecPersistentVolumeSourceIscsiOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) IscsiInput() *PersistentVolumeSpecPersistentVolumeSourceIscsi {
	var returns *PersistentVolumeSpecPersistentVolumeSourceIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) Local() PersistentVolumeSpecPersistentVolumeSourceLocalOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) LocalInput() *PersistentVolumeSpecPersistentVolumeSourceLocal {
	var returns *PersistentVolumeSpecPersistentVolumeSourceLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) Nfs() PersistentVolumeSpecPersistentVolumeSourceNfsOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) NfsInput() *PersistentVolumeSpecPersistentVolumeSourceNfs {
	var returns *PersistentVolumeSpecPersistentVolumeSourceNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PhotonPersistentDisk() PersistentVolumeSpecPersistentVolumeSourcePhotonPersistentDiskOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourcePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PhotonPersistentDiskInput() *PersistentVolumeSpecPersistentVolumeSourcePhotonPersistentDisk {
	var returns *PersistentVolumeSpecPersistentVolumeSourcePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) Quobyte() PersistentVolumeSpecPersistentVolumeSourceQuobyteOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) QuobyteInput() *PersistentVolumeSpecPersistentVolumeSourceQuobyte {
	var returns *PersistentVolumeSpecPersistentVolumeSourceQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) Rbd() PersistentVolumeSpecPersistentVolumeSourceRbdOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) RbdInput() *PersistentVolumeSpecPersistentVolumeSourceRbd {
	var returns *PersistentVolumeSpecPersistentVolumeSourceRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) VsphereVolume() PersistentVolumeSpecPersistentVolumeSourceVsphereVolumeOutputReference {
	var returns PersistentVolumeSpecPersistentVolumeSourceVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) VsphereVolumeInput() *PersistentVolumeSpecPersistentVolumeSourceVsphereVolume {
	var returns *PersistentVolumeSpecPersistentVolumeSourceVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}

func NewPersistentVolumeSpecPersistentVolumeSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PersistentVolumeSpecPersistentVolumeSourceOutputReference {
	_init_.Initialize()

	if err := validateNewPersistentVolumeSpecPersistentVolumeSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolume.PersistentVolumeSpecPersistentVolumeSourceOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPersistentVolumeSpecPersistentVolumeSourceOutputReference_Override(p PersistentVolumeSpecPersistentVolumeSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolume.PersistentVolumeSpecPersistentVolumeSourceOutputReference",
		[]any{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) SetInternalValue(val *PersistentVolumeSpecPersistentVolumeSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutAwsElasticBlockStore(value *PersistentVolumeSpecPersistentVolumeSourceAwsElasticBlockStore) {
	if err := p.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAwsElasticBlockStore",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutAzureDisk(value *PersistentVolumeSpecPersistentVolumeSourceAzureDisk) {
	if err := p.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAzureDisk",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutAzureFile(value *PersistentVolumeSpecPersistentVolumeSourceAzureFile) {
	if err := p.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAzureFile",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutCephFs(value *PersistentVolumeSpecPersistentVolumeSourceCephFs) {
	if err := p.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCephFs",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutCinder(value *PersistentVolumeSpecPersistentVolumeSourceCinder) {
	if err := p.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCinder",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutCsi(value *PersistentVolumeSpecPersistentVolumeSourceCsi) {
	if err := p.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCsi",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutFc(value *PersistentVolumeSpecPersistentVolumeSourceFc) {
	if err := p.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFc",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutFlexVolume(value *PersistentVolumeSpecPersistentVolumeSourceFlexVolume) {
	if err := p.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFlexVolume",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutFlocker(value *PersistentVolumeSpecPersistentVolumeSourceFlocker) {
	if err := p.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFlocker",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutGcePersistentDisk(value *PersistentVolumeSpecPersistentVolumeSourceGcePersistentDisk) {
	if err := p.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGcePersistentDisk",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutGlusterfs(value *PersistentVolumeSpecPersistentVolumeSourceGlusterfs) {
	if err := p.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGlusterfs",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutHostPath(value *PersistentVolumeSpecPersistentVolumeSourceHostPath) {
	if err := p.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHostPath",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutIscsi(value *PersistentVolumeSpecPersistentVolumeSourceIscsi) {
	if err := p.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIscsi",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutLocal(value *PersistentVolumeSpecPersistentVolumeSourceLocal) {
	if err := p.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLocal",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutNfs(value *PersistentVolumeSpecPersistentVolumeSourceNfs) {
	if err := p.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNfs",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutPhotonPersistentDisk(value *PersistentVolumeSpecPersistentVolumeSourcePhotonPersistentDisk) {
	if err := p.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPhotonPersistentDisk",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutQuobyte(value *PersistentVolumeSpecPersistentVolumeSourceQuobyte) {
	if err := p.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQuobyte",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutRbd(value *PersistentVolumeSpecPersistentVolumeSourceRbd) {
	if err := p.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRbd",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) PutVsphereVolume(value *PersistentVolumeSpecPersistentVolumeSourceVsphereVolume) {
	if err := p.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putVsphereVolume",
		[]any{value},
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		p,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		p,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		p,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		p,
		"resetCephFs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		p,
		"resetCinder",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		p,
		"resetCsi",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		p,
		"resetFc",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		p,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		p,
		"resetFlocker",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		p,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		p,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		p,
		"resetHostPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		p,
		"resetIscsi",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		p,
		"resetLocal",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		p,
		"resetNfs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		p,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		p,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		p,
		"resetRbd",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		p,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		p,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeSpecPersistentVolumeSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
