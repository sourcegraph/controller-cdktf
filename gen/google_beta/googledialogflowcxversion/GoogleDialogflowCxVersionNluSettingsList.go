package googledialogflowcxversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxversion/internal"
)

type GoogleDialogflowCxVersionNluSettingsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) GoogleDialogflowCxVersionNluSettingsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxVersionNluSettingsList
type jsiiProxy_GoogleDialogflowCxVersionNluSettingsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_GoogleDialogflowCxVersionNluSettingsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxVersionNluSettingsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxVersionNluSettingsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxVersionNluSettingsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxVersionNluSettingsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxVersionNluSettingsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleDialogflowCxVersionNluSettingsList {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxVersionNluSettingsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxVersionNluSettingsList{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxVersion.GoogleDialogflowCxVersionNluSettingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxVersionNluSettingsList_Override(g GoogleDialogflowCxVersionNluSettingsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxVersion.GoogleDialogflowCxVersionNluSettingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxVersionNluSettingsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxVersionNluSettingsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxVersionNluSettingsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxVersionNluSettingsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := g.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		g,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxVersionNluSettingsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxVersionNluSettingsList) Get(index *float64) GoogleDialogflowCxVersionNluSettingsOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleDialogflowCxVersionNluSettingsOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxVersionNluSettingsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxVersionNluSettingsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

