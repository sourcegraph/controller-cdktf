package googlecomputesecuritypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputesecuritypolicy/internal"
)

type GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference interface {
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
	ConfidenceThreshold() *float64
	SetConfidenceThreshold(val *float64)
	ConfidenceThresholdInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExpirationSec() *float64
	SetExpirationSec(val *float64)
	ExpirationSecInput() *float64
	// Experimental.
	Fqn() *string
	ImpactedBaselineThreshold() *float64
	SetImpactedBaselineThreshold(val *float64)
	ImpactedBaselineThresholdInput() *float64
	InternalValue() *GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfig
	SetInternalValue(val *GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfig)
	LoadThreshold() *float64
	SetLoadThreshold(val *float64)
	LoadThresholdInput() *float64
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
	ResetConfidenceThreshold()
	ResetExpirationSec()
	ResetImpactedBaselineThreshold()
	ResetLoadThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference
type jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) ConfidenceThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confidenceThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) ConfidenceThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confidenceThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) ExpirationSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) ExpirationSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) ImpactedBaselineThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"impactedBaselineThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) ImpactedBaselineThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"impactedBaselineThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) InternalValue() *GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfig {
	var returns *GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) LoadThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) LoadThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loadThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeSecurityPolicy.GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference_Override(g GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeSecurityPolicy.GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference)SetConfidenceThreshold(val *float64) {
	if err := j.validateSetConfidenceThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidenceThreshold",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference)SetExpirationSec(val *float64) {
	if err := j.validateSetExpirationSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference)SetImpactedBaselineThreshold(val *float64) {
	if err := j.validateSetImpactedBaselineThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"impactedBaselineThreshold",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference)SetInternalValue(val *GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference)SetLoadThreshold(val *float64) {
	if err := j.validateSetLoadThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadThreshold",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) ResetConfidenceThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetConfidenceThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) ResetExpirationSec() {
	_jsii_.InvokeVoid(
		g,
		"resetExpirationSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) ResetImpactedBaselineThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetImpactedBaselineThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) ResetLoadThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetLoadThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

