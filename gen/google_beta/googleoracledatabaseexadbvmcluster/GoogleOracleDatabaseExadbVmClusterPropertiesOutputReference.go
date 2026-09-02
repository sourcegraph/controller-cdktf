package googleoracledatabaseexadbvmcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleoracledatabaseexadbvmcluster/internal"
)

type GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference interface {
	cdktn.ComplexObject
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
	DataCollectionOptions() GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference
	DataCollectionOptionsInput() *GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptions
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
	InternalValue() *GoogleOracleDatabaseExadbVmClusterProperties
	SetInternalValue(val *GoogleOracleDatabaseExadbVmClusterProperties)
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
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeZone() GoogleOracleDatabaseExadbVmClusterPropertiesTimeZoneOutputReference
	TimeZoneInput() *GoogleOracleDatabaseExadbVmClusterPropertiesTimeZone
	VmFileSystemStorage() GoogleOracleDatabaseExadbVmClusterPropertiesVmFileSystemStorageOutputReference
	VmFileSystemStorageInput() *GoogleOracleDatabaseExadbVmClusterPropertiesVmFileSystemStorage
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
	PutDataCollectionOptions(value *GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptions)
	PutTimeZone(value *GoogleOracleDatabaseExadbVmClusterPropertiesTimeZone)
	PutVmFileSystemStorage(value *GoogleOracleDatabaseExadbVmClusterPropertiesVmFileSystemStorage)
	ResetAdditionalEcpuCountPerNode()
	ResetClusterName()
	ResetDataCollectionOptions()
	ResetLicenseModel()
	ResetScanListenerPortTcp()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference
type jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) AdditionalEcpuCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalEcpuCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) AdditionalEcpuCountPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalEcpuCountPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) DataCollectionOptions() GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference {
	var returns GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference
	_jsii_.Get(
		j,
		"dataCollectionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) DataCollectionOptionsInput() *GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptions {
	var returns *GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptions
	_jsii_.Get(
		j,
		"dataCollectionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) EnabledEcpuCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enabledEcpuCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) EnabledEcpuCountPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enabledEcpuCountPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ExascaleDbStorageVault() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exascaleDbStorageVault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ExascaleDbStorageVaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exascaleDbStorageVaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) GiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"giVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) GridImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gridImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) GridImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gridImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) HostnamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) HostnamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) InternalValue() *GoogleOracleDatabaseExadbVmClusterProperties {
	var returns *GoogleOracleDatabaseExadbVmClusterProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) MemorySizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) OciUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ScanListenerPortTcp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ScanListenerPortTcpInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ShapeAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ShapeAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) TimeZone() GoogleOracleDatabaseExadbVmClusterPropertiesTimeZoneOutputReference {
	var returns GoogleOracleDatabaseExadbVmClusterPropertiesTimeZoneOutputReference
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) TimeZoneInput() *GoogleOracleDatabaseExadbVmClusterPropertiesTimeZone {
	var returns *GoogleOracleDatabaseExadbVmClusterPropertiesTimeZone
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) VmFileSystemStorage() GoogleOracleDatabaseExadbVmClusterPropertiesVmFileSystemStorageOutputReference {
	var returns GoogleOracleDatabaseExadbVmClusterPropertiesVmFileSystemStorageOutputReference
	_jsii_.Get(
		j,
		"vmFileSystemStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) VmFileSystemStorageInput() *GoogleOracleDatabaseExadbVmClusterPropertiesVmFileSystemStorage {
	var returns *GoogleOracleDatabaseExadbVmClusterPropertiesVmFileSystemStorage
	_jsii_.Get(
		j,
		"vmFileSystemStorageInput",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseExadbVmClusterPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseExadbVmClusterPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseExadbVmCluster.GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseExadbVmClusterPropertiesOutputReference_Override(g GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseExadbVmCluster.GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetAdditionalEcpuCountPerNode(val *float64) {
	if err := j.validateSetAdditionalEcpuCountPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalEcpuCountPerNode",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetEnabledEcpuCountPerNode(val *float64) {
	if err := j.validateSetEnabledEcpuCountPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledEcpuCountPerNode",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetExascaleDbStorageVault(val *string) {
	if err := j.validateSetExascaleDbStorageVaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exascaleDbStorageVault",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetGridImageId(val *string) {
	if err := j.validateSetGridImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gridImageId",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetHostnamePrefix(val *string) {
	if err := j.validateSetHostnamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnamePrefix",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetInternalValue(val *GoogleOracleDatabaseExadbVmClusterProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetScanListenerPortTcp(val *float64) {
	if err := j.validateSetScanListenerPortTcpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanListenerPortTcp",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetShapeAttribute(val *string) {
	if err := j.validateSetShapeAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shapeAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) PutDataCollectionOptions(value *GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptions) {
	if err := g.validatePutDataCollectionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataCollectionOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) PutTimeZone(value *GoogleOracleDatabaseExadbVmClusterPropertiesTimeZone) {
	if err := g.validatePutTimeZoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeZone",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) PutVmFileSystemStorage(value *GoogleOracleDatabaseExadbVmClusterPropertiesVmFileSystemStorage) {
	if err := g.validatePutVmFileSystemStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVmFileSystemStorage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ResetAdditionalEcpuCountPerNode() {
	_jsii_.InvokeVoid(
		g,
		"resetAdditionalEcpuCountPerNode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ResetClusterName() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ResetDataCollectionOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetDataCollectionOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		g,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ResetScanListenerPortTcp() {
	_jsii_.InvokeVoid(
		g,
		"resetScanListenerPortTcp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

