package apihubplugin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/apihubplugin/internal"
)

type ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList
type jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

func NewApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList {
	_init_.Initialize()

	if err := validateNewApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList{}

	_jsii_.Create(
		"@cdktf/provider-google.apihubPlugin.ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList_Override(a ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apihubPlugin.ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]any{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) Get(index *float64) ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]any{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) Resolve(_context cdktf.IResolveContext) any {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		a,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginConfigTemplateAdditionalConfigTemplateEnumOptionsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
