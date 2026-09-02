package googlecomputefuturereservation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputefuturereservation/internal"
)

type GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GuestAccelerators() GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesGuestAcceleratorsList
	GuestAcceleratorsInput() interface{}
	InternalValue() *GoogleComputeFutureReservationSpecificSkuPropertiesInstanceProperties
	SetInternalValue(val *GoogleComputeFutureReservationSpecificSkuPropertiesInstanceProperties)
	LocalSsds() GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesLocalSsdsList
	LocalSsdsInput() interface{}
	LocationHint() *string
	SetLocationHint(val *string)
	LocationHintInput() *string
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	MaintenanceFreezeDurationHours() *float64
	SetMaintenanceFreezeDurationHours(val *float64)
	MaintenanceFreezeDurationHoursInput() *float64
	MaintenanceInterval() *string
	SetMaintenanceInterval(val *string)
	MaintenanceIntervalInput() *string
	MinCpuPlatform() *string
	SetMinCpuPlatform(val *string)
	MinCpuPlatformInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutGuestAccelerators(value interface{})
	PutLocalSsds(value interface{})
	ResetGuestAccelerators()
	ResetLocalSsds()
	ResetLocationHint()
	ResetMachineType()
	ResetMaintenanceFreezeDurationHours()
	ResetMaintenanceInterval()
	ResetMinCpuPlatform()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference
type jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) GuestAccelerators() GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesGuestAcceleratorsList {
	var returns GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesGuestAcceleratorsList
	_jsii_.Get(
		j,
		"guestAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) GuestAcceleratorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestAcceleratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) InternalValue() *GoogleComputeFutureReservationSpecificSkuPropertiesInstanceProperties {
	var returns *GoogleComputeFutureReservationSpecificSkuPropertiesInstanceProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) LocalSsds() GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesLocalSsdsList {
	var returns GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesLocalSsdsList
	_jsii_.Get(
		j,
		"localSsds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) LocalSsdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSsdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) LocationHint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationHint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) LocationHintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationHintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) MaintenanceFreezeDurationHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceFreezeDurationHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) MaintenanceFreezeDurationHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maintenanceFreezeDurationHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) MaintenanceInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) MaintenanceIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeFutureReservation.GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference_Override(g GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeFutureReservation.GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference)SetInternalValue(val *GoogleComputeFutureReservationSpecificSkuPropertiesInstanceProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference)SetLocationHint(val *string) {
	if err := j.validateSetLocationHintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationHint",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference)SetMaintenanceFreezeDurationHours(val *float64) {
	if err := j.validateSetMaintenanceFreezeDurationHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceFreezeDurationHours",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference)SetMaintenanceInterval(val *string) {
	if err := j.validateSetMaintenanceIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) PutGuestAccelerators(value interface{}) {
	if err := g.validatePutGuestAcceleratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGuestAccelerators",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) PutLocalSsds(value interface{}) {
	if err := g.validatePutLocalSsdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLocalSsds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) ResetGuestAccelerators() {
	_jsii_.InvokeVoid(
		g,
		"resetGuestAccelerators",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) ResetLocalSsds() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalSsds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) ResetLocationHint() {
	_jsii_.InvokeVoid(
		g,
		"resetLocationHint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) ResetMaintenanceFreezeDurationHours() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceFreezeDurationHours",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) ResetMaintenanceInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		g,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

