package googleoracledatabasecloudexadatainfrastructure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleoracledatabasecloudexadatainfrastructure/internal"
)

type GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference interface {
	cdktf.ComplexObject
	ActivatedStorageCount() *float64
	AdditionalStorageCount() *float64
	AvailableStorageSizeGb() *float64
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
	CpuCount() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomerContacts() GoogleOracleDatabaseCloudExadataInfrastructurePropertiesCustomerContactsList
	CustomerContactsInput() interface{}
	DataStorageSizeTb() *float64
	DbNodeStorageSizeGb() *float64
	DbServerVersion() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleOracleDatabaseCloudExadataInfrastructureProperties
	SetInternalValue(val *GoogleOracleDatabaseCloudExadataInfrastructureProperties)
	MaintenanceWindow() GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference
	MaintenanceWindowInput() *GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow
	MaxCpuCount() *float64
	MaxDataStorageTb() *float64
	MaxDbNodeStorageSizeGb() *float64
	MaxMemoryGb() *float64
	MemorySizeGb() *float64
	MonthlyDbServerVersion() *string
	MonthlyStorageServerVersion() *string
	NextMaintenanceRunId() *string
	NextMaintenanceRunTime() *string
	NextSecurityMaintenanceRunTime() *string
	Ocid() *string
	OciUrl() *string
	Shape() *string
	SetShape(val *string)
	ShapeInput() *string
	State() *string
	StorageCount() *float64
	SetStorageCount(val *float64)
	StorageCountInput() *float64
	StorageServerVersion() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalStorageSizeGb() *float64
	SetTotalStorageSizeGb(val *float64)
	TotalStorageSizeGbInput() *float64
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
	PutCustomerContacts(value interface{})
	PutMaintenanceWindow(value *GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow)
	ResetComputeCount()
	ResetCustomerContacts()
	ResetMaintenanceWindow()
	ResetStorageCount()
	ResetTotalStorageSizeGb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference
type jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ActivatedStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activatedStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) AdditionalStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) AvailableStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ComputeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) CpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) CustomerContacts() GoogleOracleDatabaseCloudExadataInfrastructurePropertiesCustomerContactsList {
	var returns GoogleOracleDatabaseCloudExadataInfrastructurePropertiesCustomerContactsList
	_jsii_.Get(
		j,
		"customerContacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) CustomerContactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerContactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) DataStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) DbNodeStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) DbServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) InternalValue() *GoogleOracleDatabaseCloudExadataInfrastructureProperties {
	var returns *GoogleOracleDatabaseCloudExadataInfrastructureProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MaintenanceWindow() GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference {
	var returns GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MaintenanceWindowInput() *GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow {
	var returns *GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MaxCpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MaxDataStorageTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataStorageTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MaxDbNodeStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDbNodeStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MaxMemoryGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MemorySizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MonthlyDbServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyDbServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MonthlyStorageServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyStorageServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) NextMaintenanceRunId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceRunId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) NextMaintenanceRunTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceRunTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) NextSecurityMaintenanceRunTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextSecurityMaintenanceRunTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) StorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) StorageCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) StorageServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) TotalStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) TotalStorageSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalStorageSizeGbInput",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOracleDatabaseCloudExadataInfrastructure.GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference_Override(g GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOracleDatabaseCloudExadataInfrastructure.GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetComputeCount(val *float64) {
	if err := j.validateSetComputeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetInternalValue(val *GoogleOracleDatabaseCloudExadataInfrastructureProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetShape(val *string) {
	if err := j.validateSetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetStorageCount(val *float64) {
	if err := j.validateSetStorageCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCount",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetTotalStorageSizeGb(val *float64) {
	if err := j.validateSetTotalStorageSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalStorageSizeGb",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) PutCustomerContacts(value interface{}) {
	if err := g.validatePutCustomerContactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomerContacts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) PutMaintenanceWindow(value *GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow) {
	if err := g.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ResetComputeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetComputeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ResetCustomerContacts() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomerContacts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ResetStorageCount() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ResetTotalStorageSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetTotalStorageSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

