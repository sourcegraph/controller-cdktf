package oracledatabaseexadbvmcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/oracledatabaseexadbvmcluster/internal"
)

type OracleDatabaseExadbVmClusterPropertiesOutputReference interface {
	cdktf.ComplexObject
	AdditionalEcpuCountPerNode() *float64
	SetAdditionalEcpuCountPerNode(val *float64)
	AdditionalEcpuCountPerNodeInput() *float64
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataCollectionOptions() OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference
	DataCollectionOptionsInput() *OracleDatabaseExadbVmClusterPropertiesDataCollectionOptions
	EnabledEcpuCountPerNode() *float64
	SetEnabledEcpuCountPerNode(val *float64)
	EnabledEcpuCountPerNodeInput() *float64
	ExascaleDbStorageVault() *string
	SetExascaleDbStorageVault(val *string)
	ExascaleDbStorageVaultInput() *string
	// Experimental.
	Fqn() *string
	GiVersion() *string
	GridImageId() *string
	SetGridImageId(val *string)
	GridImageIdInput() *string
	Hostname() *string
	HostnamePrefix() *string
	SetHostnamePrefix(val *string)
	HostnamePrefixInput() *string
	InternalValue() *OracleDatabaseExadbVmClusterProperties
	SetInternalValue(val *OracleDatabaseExadbVmClusterProperties)
	LicenseModel() *string
	SetLicenseModel(val *string)
	LicenseModelInput() *string
	LifecycleState() *string
	MemorySizeGb() *float64
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	OciUri() *string
	ScanListenerPortTcp() *float64
	SetScanListenerPortTcp(val *float64)
	ScanListenerPortTcpInput() *float64
	ShapeAttribute() *string
	SetShapeAttribute(val *string)
	ShapeAttributeInput() *string
	SshPublicKeys() *[]*string
	SetSshPublicKeys(val *[]*string)
	SshPublicKeysInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() OracleDatabaseExadbVmClusterPropertiesTimeZoneOutputReference
	TimeZoneInput() *OracleDatabaseExadbVmClusterPropertiesTimeZone
	VmFileSystemStorage() OracleDatabaseExadbVmClusterPropertiesVmFileSystemStorageOutputReference
	VmFileSystemStorageInput() *OracleDatabaseExadbVmClusterPropertiesVmFileSystemStorage
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	PutDataCollectionOptions(value *OracleDatabaseExadbVmClusterPropertiesDataCollectionOptions)
	PutTimeZone(value *OracleDatabaseExadbVmClusterPropertiesTimeZone)
	PutVmFileSystemStorage(value *OracleDatabaseExadbVmClusterPropertiesVmFileSystemStorage)
	ResetAdditionalEcpuCountPerNode()
	ResetClusterName()
	ResetDataCollectionOptions()
	ResetLicenseModel()
	ResetScanListenerPortTcp()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseExadbVmClusterPropertiesOutputReference
type jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) AdditionalEcpuCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalEcpuCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) AdditionalEcpuCountPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalEcpuCountPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) DataCollectionOptions() OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference {
	var returns OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference
	_jsii_.Get(
		j,
		"dataCollectionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) DataCollectionOptionsInput() *OracleDatabaseExadbVmClusterPropertiesDataCollectionOptions {
	var returns *OracleDatabaseExadbVmClusterPropertiesDataCollectionOptions
	_jsii_.Get(
		j,
		"dataCollectionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) EnabledEcpuCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enabledEcpuCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) EnabledEcpuCountPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enabledEcpuCountPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ExascaleDbStorageVault() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exascaleDbStorageVault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ExascaleDbStorageVaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exascaleDbStorageVaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) GiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"giVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) GridImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gridImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) GridImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gridImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) HostnamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) HostnamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) InternalValue() *OracleDatabaseExadbVmClusterProperties {
	var returns *OracleDatabaseExadbVmClusterProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) MemorySizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) OciUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ScanListenerPortTcp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ScanListenerPortTcpInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ShapeAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ShapeAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) TimeZone() OracleDatabaseExadbVmClusterPropertiesTimeZoneOutputReference {
	var returns OracleDatabaseExadbVmClusterPropertiesTimeZoneOutputReference
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) TimeZoneInput() *OracleDatabaseExadbVmClusterPropertiesTimeZone {
	var returns *OracleDatabaseExadbVmClusterPropertiesTimeZone
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) VmFileSystemStorage() OracleDatabaseExadbVmClusterPropertiesVmFileSystemStorageOutputReference {
	var returns OracleDatabaseExadbVmClusterPropertiesVmFileSystemStorageOutputReference
	_jsii_.Get(
		j,
		"vmFileSystemStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) VmFileSystemStorageInput() *OracleDatabaseExadbVmClusterPropertiesVmFileSystemStorage {
	var returns *OracleDatabaseExadbVmClusterPropertiesVmFileSystemStorage
	_jsii_.Get(
		j,
		"vmFileSystemStorageInput",
		&returns,
	)
	return returns
}


func NewOracleDatabaseExadbVmClusterPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OracleDatabaseExadbVmClusterPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseExadbVmClusterPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseExadbVmCluster.OracleDatabaseExadbVmClusterPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseExadbVmClusterPropertiesOutputReference_Override(o OracleDatabaseExadbVmClusterPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseExadbVmCluster.OracleDatabaseExadbVmClusterPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetAdditionalEcpuCountPerNode(val *float64) {
	if err := j.validateSetAdditionalEcpuCountPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalEcpuCountPerNode",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetEnabledEcpuCountPerNode(val *float64) {
	if err := j.validateSetEnabledEcpuCountPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledEcpuCountPerNode",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetExascaleDbStorageVault(val *string) {
	if err := j.validateSetExascaleDbStorageVaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exascaleDbStorageVault",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetGridImageId(val *string) {
	if err := j.validateSetGridImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gridImageId",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetHostnamePrefix(val *string) {
	if err := j.validateSetHostnamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnamePrefix",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetInternalValue(val *OracleDatabaseExadbVmClusterProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetScanListenerPortTcp(val *float64) {
	if err := j.validateSetScanListenerPortTcpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanListenerPortTcp",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetShapeAttribute(val *string) {
	if err := j.validateSetShapeAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shapeAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) PutDataCollectionOptions(value *OracleDatabaseExadbVmClusterPropertiesDataCollectionOptions) {
	if err := o.validatePutDataCollectionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDataCollectionOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) PutTimeZone(value *OracleDatabaseExadbVmClusterPropertiesTimeZone) {
	if err := o.validatePutTimeZoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeZone",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) PutVmFileSystemStorage(value *OracleDatabaseExadbVmClusterPropertiesVmFileSystemStorage) {
	if err := o.validatePutVmFileSystemStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putVmFileSystemStorage",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ResetAdditionalEcpuCountPerNode() {
	_jsii_.InvokeVoid(
		o,
		"resetAdditionalEcpuCountPerNode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ResetClusterName() {
	_jsii_.InvokeVoid(
		o,
		"resetClusterName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ResetDataCollectionOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetDataCollectionOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		o,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ResetScanListenerPortTcp() {
	_jsii_.InvokeVoid(
		o,
		"resetScanListenerPortTcp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

