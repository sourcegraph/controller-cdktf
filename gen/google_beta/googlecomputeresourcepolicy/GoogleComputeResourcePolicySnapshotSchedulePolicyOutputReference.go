package googlecomputeresourcepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeresourcepolicy/internal"
)

type GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeResourcePolicySnapshotSchedulePolicy
	SetInternalValue(val *GoogleComputeResourcePolicySnapshotSchedulePolicy)
	RetentionPolicy() GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference
	RetentionPolicyInput() *GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy
	Schedule() GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference
	ScheduleInput() *GoogleComputeResourcePolicySnapshotSchedulePolicySchedule
	SnapshotProperties() GoogleComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference
	SnapshotPropertiesInput() *GoogleComputeResourcePolicySnapshotSchedulePolicySnapshotProperties
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutRetentionPolicy(value *GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy)
	PutSchedule(value *GoogleComputeResourcePolicySnapshotSchedulePolicySchedule)
	PutSnapshotProperties(value *GoogleComputeResourcePolicySnapshotSchedulePolicySnapshotProperties)
	ResetRetentionPolicy()
	ResetSnapshotProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference
type jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) InternalValue() *GoogleComputeResourcePolicySnapshotSchedulePolicy {
	var returns *GoogleComputeResourcePolicySnapshotSchedulePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) RetentionPolicy() GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference {
	var returns GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference
	_jsii_.Get(
		j,
		"retentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) RetentionPolicyInput() *GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy {
	var returns *GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy
	_jsii_.Get(
		j,
		"retentionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) Schedule() GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference {
	var returns GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) ScheduleInput() *GoogleComputeResourcePolicySnapshotSchedulePolicySchedule {
	var returns *GoogleComputeResourcePolicySnapshotSchedulePolicySchedule
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) SnapshotProperties() GoogleComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference {
	var returns GoogleComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference
	_jsii_.Get(
		j,
		"snapshotProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) SnapshotPropertiesInput() *GoogleComputeResourcePolicySnapshotSchedulePolicySnapshotProperties {
	var returns *GoogleComputeResourcePolicySnapshotSchedulePolicySnapshotProperties
	_jsii_.Get(
		j,
		"snapshotPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeResourcePolicySnapshotSchedulePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeResourcePolicy.GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference_Override(g GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeResourcePolicy.GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference)SetInternalValue(val *GoogleComputeResourcePolicySnapshotSchedulePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) PutRetentionPolicy(value *GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy) {
	if err := g.validatePutRetentionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRetentionPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) PutSchedule(value *GoogleComputeResourcePolicySnapshotSchedulePolicySchedule) {
	if err := g.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) PutSnapshotProperties(value *GoogleComputeResourcePolicySnapshotSchedulePolicySnapshotProperties) {
	if err := g.validatePutSnapshotPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSnapshotProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) ResetRetentionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRetentionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) ResetSnapshotProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetSnapshotProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

