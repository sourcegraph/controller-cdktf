package googleappengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleappengineflexibleappversion/internal"
)

type GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference interface {
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
	InternalValue() *GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization
	SetInternalValue(val *GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization)
	TargetReceivedBytesPerSecond() *float64
	SetTargetReceivedBytesPerSecond(val *float64)
	TargetReceivedBytesPerSecondInput() *float64
	TargetReceivedPacketsPerSecond() *float64
	SetTargetReceivedPacketsPerSecond(val *float64)
	TargetReceivedPacketsPerSecondInput() *float64
	TargetSentBytesPerSecond() *float64
	SetTargetSentBytesPerSecond(val *float64)
	TargetSentBytesPerSecondInput() *float64
	TargetSentPacketsPerSecond() *float64
	SetTargetSentPacketsPerSecond(val *float64)
	TargetSentPacketsPerSecondInput() *float64
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
	ResetTargetReceivedBytesPerSecond()
	ResetTargetReceivedPacketsPerSecond()
	ResetTargetSentBytesPerSecond()
	ResetTargetSentPacketsPerSecond()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference
type jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) InternalValue() *GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization {
	var returns *GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetReceivedBytesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReceivedBytesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetReceivedBytesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReceivedBytesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetReceivedPacketsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReceivedPacketsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetReceivedPacketsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetReceivedPacketsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetSentBytesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSentBytesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetSentBytesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSentBytesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetSentPacketsPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSentPacketsPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TargetSentPacketsPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSentPacketsPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference{}

	_jsii_.Create(
		"google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference_Override(g GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleAppEngineFlexibleAppVersion.GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetInternalValue(val *GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetTargetReceivedBytesPerSecond(val *float64) {
	if err := j.validateSetTargetReceivedBytesPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetReceivedBytesPerSecond",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetTargetReceivedPacketsPerSecond(val *float64) {
	if err := j.validateSetTargetReceivedPacketsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetReceivedPacketsPerSecond",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetTargetSentBytesPerSecond(val *float64) {
	if err := j.validateSetTargetSentBytesPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSentBytesPerSecond",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetTargetSentPacketsPerSecond(val *float64) {
	if err := j.validateSetTargetSentPacketsPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSentPacketsPerSecond",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ResetTargetReceivedBytesPerSecond() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetReceivedBytesPerSecond",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ResetTargetReceivedPacketsPerSecond() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetReceivedPacketsPerSecond",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ResetTargetSentBytesPerSecond() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetSentBytesPerSecond",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ResetTargetSentPacketsPerSecond() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetSentPacketsPerSecond",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

