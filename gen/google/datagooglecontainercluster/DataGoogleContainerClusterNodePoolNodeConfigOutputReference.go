package datagooglecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/datagooglecontainercluster/internal"
)

type DataGoogleContainerClusterNodePoolNodeConfigOutputReference interface {
	cdktn.ComplexObject
	AdvancedMachineFeatures() DataGoogleContainerClusterNodePoolNodeConfigAdvancedMachineFeaturesList
	BootDisk() DataGoogleContainerClusterNodePoolNodeConfigBootDiskList
	BootDiskKmsKey() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConfidentialNodes() DataGoogleContainerClusterNodePoolNodeConfigConfidentialNodesList
	ContainerdConfig() DataGoogleContainerClusterNodePoolNodeConfigContainerdConfigList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DiskSizeGb() *float64
	DiskType() *string
	EffectiveTaints() DataGoogleContainerClusterNodePoolNodeConfigEffectiveTaintsList
	EnableConfidentialStorage() cdktn.IResolvable
	EphemeralStorageLocalSsdConfig() DataGoogleContainerClusterNodePoolNodeConfigEphemeralStorageLocalSsdConfigList
	FastSocket() DataGoogleContainerClusterNodePoolNodeConfigFastSocketList
	FlexStart() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	GcfsConfig() DataGoogleContainerClusterNodePoolNodeConfigGcfsConfigList
	GuestAccelerator() DataGoogleContainerClusterNodePoolNodeConfigGuestAcceleratorList
	Gvnic() DataGoogleContainerClusterNodePoolNodeConfigGvnicList
	HostMaintenancePolicy() DataGoogleContainerClusterNodePoolNodeConfigHostMaintenancePolicyList
	ImageType() *string
	InternalValue() *DataGoogleContainerClusterNodePoolNodeConfig
	SetInternalValue(val *DataGoogleContainerClusterNodePoolNodeConfig)
	KubeletConfig() DataGoogleContainerClusterNodePoolNodeConfigKubeletConfigList
	Labels() cdktn.StringMap
	LinuxNodeConfig() DataGoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigList
	LocalNvmeSsdBlockConfig() DataGoogleContainerClusterNodePoolNodeConfigLocalNvmeSsdBlockConfigList
	LocalSsdCount() *float64
	LocalSsdEncryptionMode() *string
	LoggingVariant() *string
	MachineType() *string
	MaxRunDuration() *string
	Metadata() cdktn.StringMap
	MinCpuPlatform() *string
	NodeGroup() *string
	OauthScopes() *[]*string
	Preemptible() cdktn.IResolvable
	ReservationAffinity() DataGoogleContainerClusterNodePoolNodeConfigReservationAffinityList
	ResourceLabels() cdktn.StringMap
	ResourceManagerTags() cdktn.StringMap
	SandboxConfig() DataGoogleContainerClusterNodePoolNodeConfigSandboxConfigList
	SecondaryBootDisks() DataGoogleContainerClusterNodePoolNodeConfigSecondaryBootDisksList
	ServiceAccount() *string
	ShieldedInstanceConfig() DataGoogleContainerClusterNodePoolNodeConfigShieldedInstanceConfigList
	SoleTenantConfig() DataGoogleContainerClusterNodePoolNodeConfigSoleTenantConfigList
	Spot() cdktn.IResolvable
	StoragePools() *[]*string
	Tags() *[]*string
	Taint() DataGoogleContainerClusterNodePoolNodeConfigTaintList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WindowsNodeConfig() DataGoogleContainerClusterNodePoolNodeConfigWindowsNodeConfigList
	WorkloadMetadataConfig() DataGoogleContainerClusterNodePoolNodeConfigWorkloadMetadataConfigList
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleContainerClusterNodePoolNodeConfigOutputReference
type jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) AdvancedMachineFeatures() DataGoogleContainerClusterNodePoolNodeConfigAdvancedMachineFeaturesList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigAdvancedMachineFeaturesList
	_jsii_.Get(
		j,
		"advancedMachineFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) BootDisk() DataGoogleContainerClusterNodePoolNodeConfigBootDiskList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigBootDiskList
	_jsii_.Get(
		j,
		"bootDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) BootDiskKmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskKmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) ConfidentialNodes() DataGoogleContainerClusterNodePoolNodeConfigConfidentialNodesList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigConfidentialNodesList
	_jsii_.Get(
		j,
		"confidentialNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) ContainerdConfig() DataGoogleContainerClusterNodePoolNodeConfigContainerdConfigList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigContainerdConfigList
	_jsii_.Get(
		j,
		"containerdConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) EffectiveTaints() DataGoogleContainerClusterNodePoolNodeConfigEffectiveTaintsList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigEffectiveTaintsList
	_jsii_.Get(
		j,
		"effectiveTaints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) EnableConfidentialStorage() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enableConfidentialStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) EphemeralStorageLocalSsdConfig() DataGoogleContainerClusterNodePoolNodeConfigEphemeralStorageLocalSsdConfigList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigEphemeralStorageLocalSsdConfigList
	_jsii_.Get(
		j,
		"ephemeralStorageLocalSsdConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) FastSocket() DataGoogleContainerClusterNodePoolNodeConfigFastSocketList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigFastSocketList
	_jsii_.Get(
		j,
		"fastSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) FlexStart() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"flexStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) GcfsConfig() DataGoogleContainerClusterNodePoolNodeConfigGcfsConfigList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigGcfsConfigList
	_jsii_.Get(
		j,
		"gcfsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) GuestAccelerator() DataGoogleContainerClusterNodePoolNodeConfigGuestAcceleratorList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigGuestAcceleratorList
	_jsii_.Get(
		j,
		"guestAccelerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) Gvnic() DataGoogleContainerClusterNodePoolNodeConfigGvnicList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigGvnicList
	_jsii_.Get(
		j,
		"gvnic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) HostMaintenancePolicy() DataGoogleContainerClusterNodePoolNodeConfigHostMaintenancePolicyList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigHostMaintenancePolicyList
	_jsii_.Get(
		j,
		"hostMaintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) InternalValue() *DataGoogleContainerClusterNodePoolNodeConfig {
	var returns *DataGoogleContainerClusterNodePoolNodeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) KubeletConfig() DataGoogleContainerClusterNodePoolNodeConfigKubeletConfigList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigKubeletConfigList
	_jsii_.Get(
		j,
		"kubeletConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) Labels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) LinuxNodeConfig() DataGoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigList
	_jsii_.Get(
		j,
		"linuxNodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) LocalNvmeSsdBlockConfig() DataGoogleContainerClusterNodePoolNodeConfigLocalNvmeSsdBlockConfigList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigLocalNvmeSsdBlockConfigList
	_jsii_.Get(
		j,
		"localNvmeSsdBlockConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) LocalSsdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) LocalSsdEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSsdEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) LoggingVariant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingVariant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) MaxRunDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRunDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) Metadata() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) NodeGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) Preemptible() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"preemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) ReservationAffinity() DataGoogleContainerClusterNodePoolNodeConfigReservationAffinityList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigReservationAffinityList
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) ResourceLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"resourceLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) ResourceManagerTags() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"resourceManagerTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) SandboxConfig() DataGoogleContainerClusterNodePoolNodeConfigSandboxConfigList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigSandboxConfigList
	_jsii_.Get(
		j,
		"sandboxConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) SecondaryBootDisks() DataGoogleContainerClusterNodePoolNodeConfigSecondaryBootDisksList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigSecondaryBootDisksList
	_jsii_.Get(
		j,
		"secondaryBootDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) ShieldedInstanceConfig() DataGoogleContainerClusterNodePoolNodeConfigShieldedInstanceConfigList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigShieldedInstanceConfigList
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) SoleTenantConfig() DataGoogleContainerClusterNodePoolNodeConfigSoleTenantConfigList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigSoleTenantConfigList
	_jsii_.Get(
		j,
		"soleTenantConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) Spot() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"spot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) StoragePools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storagePools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) Taint() DataGoogleContainerClusterNodePoolNodeConfigTaintList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigTaintList
	_jsii_.Get(
		j,
		"taint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) WindowsNodeConfig() DataGoogleContainerClusterNodePoolNodeConfigWindowsNodeConfigList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigWindowsNodeConfigList
	_jsii_.Get(
		j,
		"windowsNodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) WorkloadMetadataConfig() DataGoogleContainerClusterNodePoolNodeConfigWorkloadMetadataConfigList {
	var returns DataGoogleContainerClusterNodePoolNodeConfigWorkloadMetadataConfigList
	_jsii_.Get(
		j,
		"workloadMetadataConfig",
		&returns,
	)
	return returns
}


func NewDataGoogleContainerClusterNodePoolNodeConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleContainerClusterNodePoolNodeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleContainerClusterNodePoolNodeConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataGoogleContainerCluster.DataGoogleContainerClusterNodePoolNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleContainerClusterNodePoolNodeConfigOutputReference_Override(d DataGoogleContainerClusterNodePoolNodeConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataGoogleContainerCluster.DataGoogleContainerClusterNodePoolNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference)SetInternalValue(val *DataGoogleContainerClusterNodePoolNodeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodePoolNodeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

