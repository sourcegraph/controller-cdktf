package googleappengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleappengineflexibleappversion/internal"
)

type GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference interface {
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
	InternalValue() *GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilization
	SetInternalValue(val *GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilization)
	TargetReadBytesPerSecond() *float64
	SetTargetReadBytesPerSecond(val *float64)
	TargetReadBytesPerSecondInput() *float64
	TargetReadOpsPerSecond() *float64
	SetTargetReadOpsPerSecond(val *float64)
	TargetReadOpsPerSecondInput() *float64
	TargetWriteBytesPerSecond() *float64
	SetTargetWriteBytesPerSecond(val *float64)
	TargetWriteBytesPerSecondInput() *float64
	TargetWriteOpsPerSecond() *float64
	SetTargetWriteOpsPerSecond(val *float64)
	TargetWriteOpsPerSecondInput() *float64
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
	ResetTargetReadBytesPerSecond()
	ResetTargetReadOpsPerSecond()
	ResetTargetWriteBytesPerSecond()
	ResetTargetWriteOpsPerSecond()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference
type jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) InternalValue() *GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilization {
	var returns *GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilization
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetReadBytesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReadBytesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetReadBytesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReadBytesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetReadOpsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReadOpsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetReadOpsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReadOpsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetWriteBytesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetWriteBytesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetWriteBytesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetWriteBytesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetWriteOpsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetWriteOpsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TargetWriteOpsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetWriteOpsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference{}

	_jsii_.Create(
		"google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference_Override(g GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetInternalValue(val *GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilization) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetTargetReadBytesPerSecond(val *float64) {
	if err := j.validateSetTargetReadBytesPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetReadBytesPerSecond",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetTargetReadOpsPerSecond(val *float64) {
	if err := j.validateSetTargetReadOpsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetReadOpsPerSecond",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetTargetWriteBytesPerSecond(val *float64) {
	if err := j.validateSetTargetWriteBytesPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetWriteBytesPerSecond",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetTargetWriteOpsPerSecond(val *float64) {
	if err := j.validateSetTargetWriteOpsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetWriteOpsPerSecond",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ResetTargetReadBytesPerSecond() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetReadBytesPerSecond",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ResetTargetReadOpsPerSecond() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetReadOpsPerSecond",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ResetTargetWriteBytesPerSecond() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetWriteBytesPerSecond",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ResetTargetWriteOpsPerSecond() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetWriteOpsPerSecond",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

