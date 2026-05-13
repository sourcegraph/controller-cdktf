package googlechroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechroniclefeed/internal"
)

type GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference interface {
	cdktf.ComplexObject
	AzureSasToken() *string
	SetAzureSasToken(val *string)
	AzureSasTokenInput() *string
	AzureStorageConnectionString() *string
	SetAzureStorageConnectionString(val *string)
	AzureStorageConnectionStringInput() *string
	AzureStorageContainer() *string
	SetAzureStorageContainer(val *string)
	AzureStorageContainerInput() *string
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
	ConsumerGroup() *string
	SetConsumerGroup(val *string)
	ConsumerGroupInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EventHubConnectionString() *string
	SetEventHubConnectionString(val *string)
	EventHubConnectionStringInput() *string
	EventHubNamespace() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleChronicleFeedDetailsAzureEventHubSettings
	SetInternalValue(val *GoogleChronicleFeedDetailsAzureEventHubSettings)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	ResetAzureSasToken()
	ResetAzureStorageConnectionString()
	ResetAzureStorageContainer()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference
type jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) AzureSasToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureSasToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) AzureSasTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureSasTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) AzureStorageConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) AzureStorageConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageConnectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) AzureStorageContainer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) AzureStorageContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) ConsumerGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) ConsumerGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) EventHubConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) EventHubConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubConnectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) EventHubNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) InternalValue() *GoogleChronicleFeedDetailsAzureEventHubSettings {
	var returns *GoogleChronicleFeedDetailsAzureEventHubSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleFeedDetailsAzureEventHubSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference_Override(g GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetAzureSasToken(val *string) {
	if err := j.validateSetAzureSasTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureSasToken",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetAzureStorageConnectionString(val *string) {
	if err := j.validateSetAzureStorageConnectionStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureStorageConnectionString",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetAzureStorageContainer(val *string) {
	if err := j.validateSetAzureStorageContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureStorageContainer",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetConsumerGroup(val *string) {
	if err := j.validateSetConsumerGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroup",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetEventHubConnectionString(val *string) {
	if err := j.validateSetEventHubConnectionStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventHubConnectionString",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetInternalValue(val *GoogleChronicleFeedDetailsAzureEventHubSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) ResetAzureSasToken() {
	_jsii_.InvokeVoid(
		g,
		"resetAzureSasToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) ResetAzureStorageConnectionString() {
	_jsii_.InvokeVoid(
		g,
		"resetAzureStorageConnectionString",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) ResetAzureStorageContainer() {
	_jsii_.InvokeVoid(
		g,
		"resetAzureStorageContainer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

