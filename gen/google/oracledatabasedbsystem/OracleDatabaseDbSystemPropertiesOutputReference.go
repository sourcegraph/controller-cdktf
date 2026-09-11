package oracledatabasedbsystem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/oracledatabasedbsystem/internal"
)

type OracleDatabaseDbSystemPropertiesOutputReference interface {
	cdktn.ComplexObject
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
	ComputeCount() *float64
	SetComputeCount(val *float64)
	ComputeCountInput() *float64
	ComputeModel() *string
	SetComputeModel(val *string)
	ComputeModelInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabaseEdition() *string
	SetDatabaseEdition(val *string)
	DatabaseEditionInput() *string
	DataCollectionOptions() OracleDatabaseDbSystemPropertiesDataCollectionOptionsOutputReference
	DataCollectionOptionsInput() *OracleDatabaseDbSystemPropertiesDataCollectionOptions
	DataStorageSizeGb() *float64
	SetDataStorageSizeGb(val *float64)
	DataStorageSizeGbInput() *float64
	DbHome() OracleDatabaseDbSystemPropertiesDbHomeOutputReference
	DbHomeInput() *OracleDatabaseDbSystemPropertiesDbHome
	DbSystemOptions() OracleDatabaseDbSystemPropertiesDbSystemOptionsOutputReference
	DbSystemOptionsInput() *OracleDatabaseDbSystemPropertiesDbSystemOptions
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	// Experimental.
	Fqn() *string
	Hostname() *string
	HostnamePrefix() *string
	SetHostnamePrefix(val *string)
	HostnamePrefixInput() *string
	InitialDataStorageSizeGb() *float64
	SetInitialDataStorageSizeGb(val *float64)
	InitialDataStorageSizeGbInput() *float64
	InternalValue() *OracleDatabaseDbSystemProperties
	SetInternalValue(val *OracleDatabaseDbSystemProperties)
	LicenseModel() *string
	SetLicenseModel(val *string)
	LicenseModelInput() *string
	LifecycleState() *string
	MemorySizeGb() *float64
	SetMemorySizeGb(val *float64)
	MemorySizeGbInput() *float64
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	Ocid() *string
	PrivateIp() *string
	SetPrivateIp(val *string)
	PrivateIpInput() *string
	RecoStorageSizeGb() *float64
	SetRecoStorageSizeGb(val *float64)
	RecoStorageSizeGbInput() *float64
	Shape() *string
	SetShape(val *string)
	ShapeInput() *string
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
	TimeZone() OracleDatabaseDbSystemPropertiesTimeZoneOutputReference
	TimeZoneInput() *OracleDatabaseDbSystemPropertiesTimeZone
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
	PutDataCollectionOptions(value *OracleDatabaseDbSystemPropertiesDataCollectionOptions)
	PutDbHome(value *OracleDatabaseDbSystemPropertiesDbHome)
	PutDbSystemOptions(value *OracleDatabaseDbSystemPropertiesDbSystemOptions)
	PutTimeZone(value *OracleDatabaseDbSystemPropertiesTimeZone)
	ResetComputeModel()
	ResetDataCollectionOptions()
	ResetDataStorageSizeGb()
	ResetDbHome()
	ResetDbSystemOptions()
	ResetDomain()
	ResetHostnamePrefix()
	ResetMemorySizeGb()
	ResetNodeCount()
	ResetPrivateIp()
	ResetRecoStorageSizeGb()
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

// The jsii proxy struct for OracleDatabaseDbSystemPropertiesOutputReference
type jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ComputeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ComputeModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) DatabaseEdition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) DatabaseEditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseEditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) DataCollectionOptions() OracleDatabaseDbSystemPropertiesDataCollectionOptionsOutputReference {
	var returns OracleDatabaseDbSystemPropertiesDataCollectionOptionsOutputReference
	_jsii_.Get(
		j,
		"dataCollectionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) DataCollectionOptionsInput() *OracleDatabaseDbSystemPropertiesDataCollectionOptions {
	var returns *OracleDatabaseDbSystemPropertiesDataCollectionOptions
	_jsii_.Get(
		j,
		"dataCollectionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) DataStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) DataStorageSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) DbHome() OracleDatabaseDbSystemPropertiesDbHomeOutputReference {
	var returns OracleDatabaseDbSystemPropertiesDbHomeOutputReference
	_jsii_.Get(
		j,
		"dbHome",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) DbHomeInput() *OracleDatabaseDbSystemPropertiesDbHome {
	var returns *OracleDatabaseDbSystemPropertiesDbHome
	_jsii_.Get(
		j,
		"dbHomeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) DbSystemOptions() OracleDatabaseDbSystemPropertiesDbSystemOptionsOutputReference {
	var returns OracleDatabaseDbSystemPropertiesDbSystemOptionsOutputReference
	_jsii_.Get(
		j,
		"dbSystemOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) DbSystemOptionsInput() *OracleDatabaseDbSystemPropertiesDbSystemOptions {
	var returns *OracleDatabaseDbSystemPropertiesDbSystemOptions
	_jsii_.Get(
		j,
		"dbSystemOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) HostnamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) HostnamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) InitialDataStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDataStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) InitialDataStorageSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDataStorageSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) InternalValue() *OracleDatabaseDbSystemProperties {
	var returns *OracleDatabaseDbSystemProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) MemorySizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) MemorySizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) PrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) RecoStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recoStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) RecoStorageSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recoStorageSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) TimeZone() OracleDatabaseDbSystemPropertiesTimeZoneOutputReference {
	var returns OracleDatabaseDbSystemPropertiesTimeZoneOutputReference
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) TimeZoneInput() *OracleDatabaseDbSystemPropertiesTimeZone {
	var returns *OracleDatabaseDbSystemPropertiesTimeZone
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewOracleDatabaseDbSystemPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OracleDatabaseDbSystemPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseDbSystemPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseDbSystem.OracleDatabaseDbSystemPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseDbSystemPropertiesOutputReference_Override(o OracleDatabaseDbSystemPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseDbSystem.OracleDatabaseDbSystemPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetComputeCount(val *float64) {
	if err := j.validateSetComputeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeCount",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetComputeModel(val *string) {
	if err := j.validateSetComputeModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeModel",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetDatabaseEdition(val *string) {
	if err := j.validateSetDatabaseEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseEdition",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetDataStorageSizeGb(val *float64) {
	if err := j.validateSetDataStorageSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStorageSizeGb",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetHostnamePrefix(val *string) {
	if err := j.validateSetHostnamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnamePrefix",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetInitialDataStorageSizeGb(val *float64) {
	if err := j.validateSetInitialDataStorageSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialDataStorageSizeGb",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetInternalValue(val *OracleDatabaseDbSystemProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetMemorySizeGb(val *float64) {
	if err := j.validateSetMemorySizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySizeGb",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetPrivateIp(val *string) {
	if err := j.validateSetPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIp",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetRecoStorageSizeGb(val *float64) {
	if err := j.validateSetRecoStorageSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoStorageSizeGb",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetShape(val *string) {
	if err := j.validateSetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) PutDataCollectionOptions(value *OracleDatabaseDbSystemPropertiesDataCollectionOptions) {
	if err := o.validatePutDataCollectionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDataCollectionOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) PutDbHome(value *OracleDatabaseDbSystemPropertiesDbHome) {
	if err := o.validatePutDbHomeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDbHome",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) PutDbSystemOptions(value *OracleDatabaseDbSystemPropertiesDbSystemOptions) {
	if err := o.validatePutDbSystemOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDbSystemOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) PutTimeZone(value *OracleDatabaseDbSystemPropertiesTimeZone) {
	if err := o.validatePutTimeZoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeZone",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ResetComputeModel() {
	_jsii_.InvokeVoid(
		o,
		"resetComputeModel",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ResetDataCollectionOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetDataCollectionOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ResetDataStorageSizeGb() {
	_jsii_.InvokeVoid(
		o,
		"resetDataStorageSizeGb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ResetDbHome() {
	_jsii_.InvokeVoid(
		o,
		"resetDbHome",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ResetDbSystemOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetDbSystemOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ResetDomain() {
	_jsii_.InvokeVoid(
		o,
		"resetDomain",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ResetHostnamePrefix() {
	_jsii_.InvokeVoid(
		o,
		"resetHostnamePrefix",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ResetMemorySizeGb() {
	_jsii_.InvokeVoid(
		o,
		"resetMemorySizeGb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ResetNodeCount() {
	_jsii_.InvokeVoid(
		o,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ResetPrivateIp() {
	_jsii_.InvokeVoid(
		o,
		"resetPrivateIp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ResetRecoStorageSizeGb() {
	_jsii_.InvokeVoid(
		o,
		"resetRecoStorageSizeGb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseDbSystemPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

