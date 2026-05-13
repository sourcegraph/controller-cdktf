package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference interface {
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
	InternalValue() *GoogleContainerClusterAutopilotClusterPolicyConfig
	SetInternalValue(val *GoogleContainerClusterAutopilotClusterPolicyConfig)
	NoStandardNodePools() interface{}
	SetNoStandardNodePools(val interface{})
	NoStandardNodePoolsInput() interface{}
	NoSystemImpersonation() interface{}
	SetNoSystemImpersonation(val interface{})
	NoSystemImpersonationInput() interface{}
	NoSystemMutation() interface{}
	SetNoSystemMutation(val interface{})
	NoSystemMutationInput() interface{}
	NoUnsafeWebhooks() interface{}
	SetNoUnsafeWebhooks(val interface{})
	NoUnsafeWebhooksInput() interface{}
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
	ResetNoStandardNodePools()
	ResetNoSystemImpersonation()
	ResetNoSystemMutation()
	ResetNoUnsafeWebhooks()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference
type jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) InternalValue() *GoogleContainerClusterAutopilotClusterPolicyConfig {
	var returns *GoogleContainerClusterAutopilotClusterPolicyConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) NoStandardNodePools() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noStandardNodePools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) NoStandardNodePoolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noStandardNodePoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) NoSystemImpersonation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSystemImpersonation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) NoSystemImpersonationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSystemImpersonationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) NoSystemMutation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSystemMutation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) NoSystemMutationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSystemMutationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) NoUnsafeWebhooks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnsafeWebhooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) NoUnsafeWebhooksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noUnsafeWebhooksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterAutopilotClusterPolicyConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterAutopilotClusterPolicyConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterAutopilotClusterPolicyConfigOutputReference_Override(g GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference)SetInternalValue(val *GoogleContainerClusterAutopilotClusterPolicyConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference)SetNoStandardNodePools(val interface{}) {
	if err := j.validateSetNoStandardNodePoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noStandardNodePools",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference)SetNoSystemImpersonation(val interface{}) {
	if err := j.validateSetNoSystemImpersonationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSystemImpersonation",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference)SetNoSystemMutation(val interface{}) {
	if err := j.validateSetNoSystemMutationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSystemMutation",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference)SetNoUnsafeWebhooks(val interface{}) {
	if err := j.validateSetNoUnsafeWebhooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noUnsafeWebhooks",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) ResetNoStandardNodePools() {
	_jsii_.InvokeVoid(
		g,
		"resetNoStandardNodePools",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) ResetNoSystemImpersonation() {
	_jsii_.InvokeVoid(
		g,
		"resetNoSystemImpersonation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) ResetNoSystemMutation() {
	_jsii_.InvokeVoid(
		g,
		"resetNoSystemMutation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) ResetNoUnsafeWebhooks() {
	_jsii_.InvokeVoid(
		g,
		"resetNoUnsafeWebhooks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterAutopilotClusterPolicyConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

