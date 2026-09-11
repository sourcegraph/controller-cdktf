package zonesettingsoverride

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/zonesettingsoverride/internal"
)

type ZoneSettingsOverrideInitialSettingsSecurityHeaderList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ZoneSettingsOverrideInitialSettingsSecurityHeaderOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZoneSettingsOverrideInitialSettingsSecurityHeaderList
type jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewZoneSettingsOverrideInitialSettingsSecurityHeaderList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ZoneSettingsOverrideInitialSettingsSecurityHeaderList {
	_init_.Initialize()

	if err := validateNewZoneSettingsOverrideInitialSettingsSecurityHeaderListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zoneSettingsOverride.ZoneSettingsOverrideInitialSettingsSecurityHeaderList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewZoneSettingsOverrideInitialSettingsSecurityHeaderList_Override(z ZoneSettingsOverrideInitialSettingsSecurityHeaderList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.zoneSettingsOverride.ZoneSettingsOverrideInitialSettingsSecurityHeaderList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		z,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := z.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		z,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList) Get(index *float64) ZoneSettingsOverrideInitialSettingsSecurityHeaderOutputReference {
	if err := z.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ZoneSettingsOverrideInitialSettingsSecurityHeaderOutputReference

	_jsii_.Invoke(
		z,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := z.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		z,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsSecurityHeaderList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

