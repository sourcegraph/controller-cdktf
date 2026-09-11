package googlechroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechroniclefeed/internal"
)

type GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	Get(index *float64) GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList
type jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList {
	_init_.Initialize()

	if err := validateNewGoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList_Override(g GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := g.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		g,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList) Get(index *float64) GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsProofpointOnDemandSettingsAuthenticationHeaderKeyValuesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

