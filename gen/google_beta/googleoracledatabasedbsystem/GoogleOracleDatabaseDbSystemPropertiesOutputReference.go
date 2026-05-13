package googleoracledatabasedbsystem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleoracledatabasedbsystem/internal"
)

type GoogleOracleDatabaseDbSystemPropertiesOutputReference interface {
	cdktf.ComplexObject
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
	DataCollectionOptions() GoogleOracleDatabaseDbSystemPropertiesDataCollectionOptionsOutputReference
	DataCollectionOptionsInput() *GoogleOracleDatabaseDbSystemPropertiesDataCollectionOptions
	DataStorageSizeGb() *float64
	SetDataStorageSizeGb(val *float64)
	DataStorageSizeGbInput() *float64
	DbHome() GoogleOracleDatabaseDbSystemPropertiesDbHomeOutputReference
	DbHomeInput() *GoogleOracleDatabaseDbSystemPropertiesDbHome
	DbSystemOptions() GoogleOracleDatabaseDbSystemPropertiesDbSystemOptionsOutputReference
	DbSystemOptionsInput() *GoogleOracleDatabaseDbSystemPropertiesDbSystemOptions
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
	InternalValue() *GoogleOracleDatabaseDbSystemProperties
	SetInternalValue(val *GoogleOracleDatabaseDbSystemProperties)
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
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() GoogleOracleDatabaseDbSystemPropertiesTimeZoneOutputReference
	TimeZoneInput() *GoogleOracleDatabaseDbSystemPropertiesTimeZone
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
	PutDataCollectionOptions(value *GoogleOracleDatabaseDbSystemPropertiesDataCollectionOptions)
	PutDbHome(value *GoogleOracleDatabaseDbSystemPropertiesDbHome)
	PutDbSystemOptions(value *GoogleOracleDatabaseDbSystemPropertiesDbSystemOptions)
	PutTimeZone(value *GoogleOracleDatabaseDbSystemPropertiesTimeZone)
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
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseDbSystemPropertiesOutputReference
type jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ComputeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ComputeModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) DatabaseEdition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) DatabaseEditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseEditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) DataCollectionOptions() GoogleOracleDatabaseDbSystemPropertiesDataCollectionOptionsOutputReference {
	var returns GoogleOracleDatabaseDbSystemPropertiesDataCollectionOptionsOutputReference
	_jsii_.Get(
		j,
		"dataCollectionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) DataCollectionOptionsInput() *GoogleOracleDatabaseDbSystemPropertiesDataCollectionOptions {
	var returns *GoogleOracleDatabaseDbSystemPropertiesDataCollectionOptions
	_jsii_.Get(
		j,
		"dataCollectionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) DataStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) DataStorageSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) DbHome() GoogleOracleDatabaseDbSystemPropertiesDbHomeOutputReference {
	var returns GoogleOracleDatabaseDbSystemPropertiesDbHomeOutputReference
	_jsii_.Get(
		j,
		"dbHome",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) DbHomeInput() *GoogleOracleDatabaseDbSystemPropertiesDbHome {
	var returns *GoogleOracleDatabaseDbSystemPropertiesDbHome
	_jsii_.Get(
		j,
		"dbHomeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) DbSystemOptions() GoogleOracleDatabaseDbSystemPropertiesDbSystemOptionsOutputReference {
	var returns GoogleOracleDatabaseDbSystemPropertiesDbSystemOptionsOutputReference
	_jsii_.Get(
		j,
		"dbSystemOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) DbSystemOptionsInput() *GoogleOracleDatabaseDbSystemPropertiesDbSystemOptions {
	var returns *GoogleOracleDatabaseDbSystemPropertiesDbSystemOptions
	_jsii_.Get(
		j,
		"dbSystemOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) HostnamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) HostnamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) InitialDataStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDataStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) InitialDataStorageSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDataStorageSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) InternalValue() *GoogleOracleDatabaseDbSystemProperties {
	var returns *GoogleOracleDatabaseDbSystemProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) MemorySizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) MemorySizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) PrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) RecoStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recoStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) RecoStorageSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recoStorageSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) TimeZone() GoogleOracleDatabaseDbSystemPropertiesTimeZoneOutputReference {
	var returns GoogleOracleDatabaseDbSystemPropertiesTimeZoneOutputReference
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) TimeZoneInput() *GoogleOracleDatabaseDbSystemPropertiesTimeZone {
	var returns *GoogleOracleDatabaseDbSystemPropertiesTimeZone
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseDbSystemPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseDbSystemPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseDbSystemPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOracleDatabaseDbSystem.GoogleOracleDatabaseDbSystemPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseDbSystemPropertiesOutputReference_Override(g GoogleOracleDatabaseDbSystemPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOracleDatabaseDbSystem.GoogleOracleDatabaseDbSystemPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetComputeCount(val *float64) {
	if err := j.validateSetComputeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetComputeModel(val *string) {
	if err := j.validateSetComputeModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeModel",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetDatabaseEdition(val *string) {
	if err := j.validateSetDatabaseEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseEdition",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetDataStorageSizeGb(val *float64) {
	if err := j.validateSetDataStorageSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStorageSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetHostnamePrefix(val *string) {
	if err := j.validateSetHostnamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnamePrefix",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetInitialDataStorageSizeGb(val *float64) {
	if err := j.validateSetInitialDataStorageSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialDataStorageSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetInternalValue(val *GoogleOracleDatabaseDbSystemProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetMemorySizeGb(val *float64) {
	if err := j.validateSetMemorySizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetPrivateIp(val *string) {
	if err := j.validateSetPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIp",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetRecoStorageSizeGb(val *float64) {
	if err := j.validateSetRecoStorageSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoStorageSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetShape(val *string) {
	if err := j.validateSetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) PutDataCollectionOptions(value *GoogleOracleDatabaseDbSystemPropertiesDataCollectionOptions) {
	if err := g.validatePutDataCollectionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataCollectionOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) PutDbHome(value *GoogleOracleDatabaseDbSystemPropertiesDbHome) {
	if err := g.validatePutDbHomeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDbHome",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) PutDbSystemOptions(value *GoogleOracleDatabaseDbSystemPropertiesDbSystemOptions) {
	if err := g.validatePutDbSystemOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDbSystemOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) PutTimeZone(value *GoogleOracleDatabaseDbSystemPropertiesTimeZone) {
	if err := g.validatePutTimeZoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeZone",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ResetComputeModel() {
	_jsii_.InvokeVoid(
		g,
		"resetComputeModel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ResetDataCollectionOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetDataCollectionOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ResetDataStorageSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetDataStorageSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ResetDbHome() {
	_jsii_.InvokeVoid(
		g,
		"resetDbHome",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ResetDbSystemOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetDbSystemOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ResetDomain() {
	_jsii_.InvokeVoid(
		g,
		"resetDomain",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ResetHostnamePrefix() {
	_jsii_.InvokeVoid(
		g,
		"resetHostnamePrefix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ResetMemorySizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetMemorySizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ResetNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ResetPrivateIp() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ResetRecoStorageSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetRecoStorageSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseDbSystemPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

