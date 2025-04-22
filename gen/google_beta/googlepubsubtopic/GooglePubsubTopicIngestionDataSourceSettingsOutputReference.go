package googlepubsubtopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlepubsubtopic/internal"
)

type GooglePubsubTopicIngestionDataSourceSettingsOutputReference interface {
	cdktf.ComplexObject
	AwsKinesis() GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference
	AwsKinesisInput() *GooglePubsubTopicIngestionDataSourceSettingsAwsKinesis
	AwsMsk() GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference
	AwsMskInput() *GooglePubsubTopicIngestionDataSourceSettingsAwsMsk
	AzureEventHubs() GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference
	AzureEventHubsInput() *GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubs
	CloudStorage() GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference
	CloudStorageInput() *GooglePubsubTopicIngestionDataSourceSettingsCloudStorage
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
	ConfluentCloud() GooglePubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference
	ConfluentCloudInput() *GooglePubsubTopicIngestionDataSourceSettingsConfluentCloud
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GooglePubsubTopicIngestionDataSourceSettings
	SetInternalValue(val *GooglePubsubTopicIngestionDataSourceSettings)
	PlatformLogsSettings() GooglePubsubTopicIngestionDataSourceSettingsPlatformLogsSettingsOutputReference
	PlatformLogsSettingsInput() *GooglePubsubTopicIngestionDataSourceSettingsPlatformLogsSettings
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
	PutAwsKinesis(value *GooglePubsubTopicIngestionDataSourceSettingsAwsKinesis)
	PutAwsMsk(value *GooglePubsubTopicIngestionDataSourceSettingsAwsMsk)
	PutAzureEventHubs(value *GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubs)
	PutCloudStorage(value *GooglePubsubTopicIngestionDataSourceSettingsCloudStorage)
	PutConfluentCloud(value *GooglePubsubTopicIngestionDataSourceSettingsConfluentCloud)
	PutPlatformLogsSettings(value *GooglePubsubTopicIngestionDataSourceSettingsPlatformLogsSettings)
	ResetAwsKinesis()
	ResetAwsMsk()
	ResetAzureEventHubs()
	ResetCloudStorage()
	ResetConfluentCloud()
	ResetPlatformLogsSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePubsubTopicIngestionDataSourceSettingsOutputReference
type jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) AwsKinesis() GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference {
	var returns GooglePubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference
	_jsii_.Get(
		j,
		"awsKinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) AwsKinesisInput() *GooglePubsubTopicIngestionDataSourceSettingsAwsKinesis {
	var returns *GooglePubsubTopicIngestionDataSourceSettingsAwsKinesis
	_jsii_.Get(
		j,
		"awsKinesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) AwsMsk() GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference {
	var returns GooglePubsubTopicIngestionDataSourceSettingsAwsMskOutputReference
	_jsii_.Get(
		j,
		"awsMsk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) AwsMskInput() *GooglePubsubTopicIngestionDataSourceSettingsAwsMsk {
	var returns *GooglePubsubTopicIngestionDataSourceSettingsAwsMsk
	_jsii_.Get(
		j,
		"awsMskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) AzureEventHubs() GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference {
	var returns GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference
	_jsii_.Get(
		j,
		"azureEventHubs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) AzureEventHubsInput() *GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubs {
	var returns *GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubs
	_jsii_.Get(
		j,
		"azureEventHubsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) CloudStorage() GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference {
	var returns GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference
	_jsii_.Get(
		j,
		"cloudStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) CloudStorageInput() *GooglePubsubTopicIngestionDataSourceSettingsCloudStorage {
	var returns *GooglePubsubTopicIngestionDataSourceSettingsCloudStorage
	_jsii_.Get(
		j,
		"cloudStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) ConfluentCloud() GooglePubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference {
	var returns GooglePubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference
	_jsii_.Get(
		j,
		"confluentCloud",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) ConfluentCloudInput() *GooglePubsubTopicIngestionDataSourceSettingsConfluentCloud {
	var returns *GooglePubsubTopicIngestionDataSourceSettingsConfluentCloud
	_jsii_.Get(
		j,
		"confluentCloudInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) InternalValue() *GooglePubsubTopicIngestionDataSourceSettings {
	var returns *GooglePubsubTopicIngestionDataSourceSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) PlatformLogsSettings() GooglePubsubTopicIngestionDataSourceSettingsPlatformLogsSettingsOutputReference {
	var returns GooglePubsubTopicIngestionDataSourceSettingsPlatformLogsSettingsOutputReference
	_jsii_.Get(
		j,
		"platformLogsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) PlatformLogsSettingsInput() *GooglePubsubTopicIngestionDataSourceSettingsPlatformLogsSettings {
	var returns *GooglePubsubTopicIngestionDataSourceSettingsPlatformLogsSettings
	_jsii_.Get(
		j,
		"platformLogsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGooglePubsubTopicIngestionDataSourceSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePubsubTopicIngestionDataSourceSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePubsubTopicIngestionDataSourceSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePubsubTopic.GooglePubsubTopicIngestionDataSourceSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePubsubTopicIngestionDataSourceSettingsOutputReference_Override(g GooglePubsubTopicIngestionDataSourceSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePubsubTopic.GooglePubsubTopicIngestionDataSourceSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference)SetInternalValue(val *GooglePubsubTopicIngestionDataSourceSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) PutAwsKinesis(value *GooglePubsubTopicIngestionDataSourceSettingsAwsKinesis) {
	if err := g.validatePutAwsKinesisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAwsKinesis",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) PutAwsMsk(value *GooglePubsubTopicIngestionDataSourceSettingsAwsMsk) {
	if err := g.validatePutAwsMskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAwsMsk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) PutAzureEventHubs(value *GooglePubsubTopicIngestionDataSourceSettingsAzureEventHubs) {
	if err := g.validatePutAzureEventHubsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAzureEventHubs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) PutCloudStorage(value *GooglePubsubTopicIngestionDataSourceSettingsCloudStorage) {
	if err := g.validatePutCloudStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudStorage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) PutConfluentCloud(value *GooglePubsubTopicIngestionDataSourceSettingsConfluentCloud) {
	if err := g.validatePutConfluentCloudParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfluentCloud",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) PutPlatformLogsSettings(value *GooglePubsubTopicIngestionDataSourceSettingsPlatformLogsSettings) {
	if err := g.validatePutPlatformLogsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPlatformLogsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) ResetAwsKinesis() {
	_jsii_.InvokeVoid(
		g,
		"resetAwsKinesis",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) ResetAwsMsk() {
	_jsii_.InvokeVoid(
		g,
		"resetAwsMsk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) ResetAzureEventHubs() {
	_jsii_.InvokeVoid(
		g,
		"resetAzureEventHubs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) ResetCloudStorage() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudStorage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) ResetConfluentCloud() {
	_jsii_.InvokeVoid(
		g,
		"resetConfluentCloud",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) ResetPlatformLogsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetPlatformLogsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

