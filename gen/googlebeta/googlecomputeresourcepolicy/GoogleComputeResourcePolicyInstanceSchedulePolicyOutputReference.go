package googlecomputeresourcepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecomputeresourcepolicy/internal"
)

type GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExpirationTime() *string
	SetExpirationTime(val *string)
	ExpirationTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeResourcePolicyInstanceSchedulePolicy
	SetInternalValue(val *GoogleComputeResourcePolicyInstanceSchedulePolicy)
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	VmStartSchedule() GoogleComputeResourcePolicyInstanceSchedulePolicyVmStartScheduleOutputReference
	VmStartScheduleInput() *GoogleComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule
	VmStopSchedule() GoogleComputeResourcePolicyInstanceSchedulePolicyVmStopScheduleOutputReference
	VmStopScheduleInput() *GoogleComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule
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
	PutVmStartSchedule(value *GoogleComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule)
	PutVmStopSchedule(value *GoogleComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule)
	ResetExpirationTime()
	ResetStartTime()
	ResetVmStartSchedule()
	ResetVmStopSchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference
type jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) ExpirationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) ExpirationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) InternalValue() *GoogleComputeResourcePolicyInstanceSchedulePolicy {
	var returns *GoogleComputeResourcePolicyInstanceSchedulePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) VmStartSchedule() GoogleComputeResourcePolicyInstanceSchedulePolicyVmStartScheduleOutputReference {
	var returns GoogleComputeResourcePolicyInstanceSchedulePolicyVmStartScheduleOutputReference
	_jsii_.Get(
		j,
		"vmStartSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) VmStartScheduleInput() *GoogleComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule {
	var returns *GoogleComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule
	_jsii_.Get(
		j,
		"vmStartScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) VmStopSchedule() GoogleComputeResourcePolicyInstanceSchedulePolicyVmStopScheduleOutputReference {
	var returns GoogleComputeResourcePolicyInstanceSchedulePolicyVmStopScheduleOutputReference
	_jsii_.Get(
		j,
		"vmStopSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) VmStopScheduleInput() *GoogleComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule {
	var returns *GoogleComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule
	_jsii_.Get(
		j,
		"vmStopScheduleInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeResourcePolicyInstanceSchedulePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeResourcePolicy.GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference_Override(g GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeResourcePolicy.GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetExpirationTime(val *string) {
	if err := j.validateSetExpirationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationTime",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetInternalValue(val *GoogleComputeResourcePolicyInstanceSchedulePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) PutVmStartSchedule(value *GoogleComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule) {
	if err := g.validatePutVmStartScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVmStartSchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) PutVmStopSchedule(value *GoogleComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule) {
	if err := g.validatePutVmStopScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVmStopSchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) ResetExpirationTime() {
	_jsii_.InvokeVoid(
		g,
		"resetExpirationTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		g,
		"resetStartTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) ResetVmStartSchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetVmStartSchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) ResetVmStopSchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetVmStopSchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeResourcePolicyInstanceSchedulePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

