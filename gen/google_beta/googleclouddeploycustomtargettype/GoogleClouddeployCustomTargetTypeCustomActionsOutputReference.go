package googleclouddeploycustomtargettype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleclouddeploycustomtargettype/internal"
)

type GoogleClouddeployCustomTargetTypeCustomActionsOutputReference interface {
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
	DeployAction() *string
	SetDeployAction(val *string)
	DeployActionInput() *string
	// Experimental.
	Fqn() *string
	IncludeSkaffoldModules() GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesList
	IncludeSkaffoldModulesInput() interface{}
	InternalValue() *GoogleClouddeployCustomTargetTypeCustomActions
	SetInternalValue(val *GoogleClouddeployCustomTargetTypeCustomActions)
	RenderAction() *string
	SetRenderAction(val *string)
	RenderActionInput() *string
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
	PutIncludeSkaffoldModules(value interface{})
	ResetIncludeSkaffoldModules()
	ResetRenderAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleClouddeployCustomTargetTypeCustomActionsOutputReference
type jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) DeployAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) DeployActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) IncludeSkaffoldModules() GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesList {
	var returns GoogleClouddeployCustomTargetTypeCustomActionsIncludeSkaffoldModulesList
	_jsii_.Get(
		j,
		"includeSkaffoldModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) IncludeSkaffoldModulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSkaffoldModulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) InternalValue() *GoogleClouddeployCustomTargetTypeCustomActions {
	var returns *GoogleClouddeployCustomTargetTypeCustomActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) RenderAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renderAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) RenderActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renderActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleClouddeployCustomTargetTypeCustomActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleClouddeployCustomTargetTypeCustomActionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleClouddeployCustomTargetTypeCustomActionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployCustomTargetType.GoogleClouddeployCustomTargetTypeCustomActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleClouddeployCustomTargetTypeCustomActionsOutputReference_Override(g GoogleClouddeployCustomTargetTypeCustomActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployCustomTargetType.GoogleClouddeployCustomTargetTypeCustomActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference)SetDeployAction(val *string) {
	if err := j.validateSetDeployActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deployAction",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference)SetInternalValue(val *GoogleClouddeployCustomTargetTypeCustomActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference)SetRenderAction(val *string) {
	if err := j.validateSetRenderActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renderAction",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) PutIncludeSkaffoldModules(value interface{}) {
	if err := g.validatePutIncludeSkaffoldModulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIncludeSkaffoldModules",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) ResetIncludeSkaffoldModules() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeSkaffoldModules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) ResetRenderAction() {
	_jsii_.InvokeVoid(
		g,
		"resetRenderAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleClouddeployCustomTargetTypeCustomActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

