package zonesettingsoverride

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/zonesettingsoverride/internal"
)

type ZoneSettingsOverrideInitialSettingsList interface {
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
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ZoneSettingsOverrideInitialSettingsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZoneSettingsOverrideInitialSettingsList
type jsiiProxy_ZoneSettingsOverrideInitialSettingsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewZoneSettingsOverrideInitialSettingsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ZoneSettingsOverrideInitialSettingsList {
	_init_.Initialize()

	if err := validateNewZoneSettingsOverrideInitialSettingsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZoneSettingsOverrideInitialSettingsList{}

	_jsii_.Create(
		"cloudflare.zoneSettingsOverride.ZoneSettingsOverrideInitialSettingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewZoneSettingsOverrideInitialSettingsList_Override(z ZoneSettingsOverrideInitialSettingsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.zoneSettingsOverride.ZoneSettingsOverrideInitialSettingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		z,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsList) Get(index *float64) ZoneSettingsOverrideInitialSettingsOutputReference {
	if err := z.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ZoneSettingsOverrideInitialSettingsOutputReference

	_jsii_.Invoke(
		z,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := z.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		z,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

