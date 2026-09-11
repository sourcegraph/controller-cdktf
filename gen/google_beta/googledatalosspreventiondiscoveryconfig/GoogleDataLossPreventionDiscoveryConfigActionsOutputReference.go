package googledatalosspreventiondiscoveryconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventiondiscoveryconfig/internal"
)

type GoogleDataLossPreventionDiscoveryConfigActionsOutputReference interface {
	cdktn.ComplexObject
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
	ExportData() GoogleDataLossPreventionDiscoveryConfigActionsExportDataOutputReference
	ExportDataInput() *GoogleDataLossPreventionDiscoveryConfigActionsExportData
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PublishToChronicle() GoogleDataLossPreventionDiscoveryConfigActionsPublishToChronicleOutputReference
	PublishToChronicleInput() *GoogleDataLossPreventionDiscoveryConfigActionsPublishToChronicle
	PublishToDataplexCatalog() GoogleDataLossPreventionDiscoveryConfigActionsPublishToDataplexCatalogOutputReference
	PublishToDataplexCatalogInput() *GoogleDataLossPreventionDiscoveryConfigActionsPublishToDataplexCatalog
	PublishToScc() GoogleDataLossPreventionDiscoveryConfigActionsPublishToSccOutputReference
	PublishToSccInput() *GoogleDataLossPreventionDiscoveryConfigActionsPublishToScc
	PubSubNotification() GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotificationOutputReference
	PubSubNotificationInput() *GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotification
	TagResources() GoogleDataLossPreventionDiscoveryConfigActionsTagResourcesOutputReference
	TagResourcesInput() *GoogleDataLossPreventionDiscoveryConfigActionsTagResources
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutExportData(value *GoogleDataLossPreventionDiscoveryConfigActionsExportData)
	PutPublishToChronicle(value *GoogleDataLossPreventionDiscoveryConfigActionsPublishToChronicle)
	PutPublishToDataplexCatalog(value *GoogleDataLossPreventionDiscoveryConfigActionsPublishToDataplexCatalog)
	PutPublishToScc(value *GoogleDataLossPreventionDiscoveryConfigActionsPublishToScc)
	PutPubSubNotification(value *GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotification)
	PutTagResources(value *GoogleDataLossPreventionDiscoveryConfigActionsTagResources)
	ResetExportData()
	ResetPublishToChronicle()
	ResetPublishToDataplexCatalog()
	ResetPublishToScc()
	ResetPubSubNotification()
	ResetTagResources()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionDiscoveryConfigActionsOutputReference
type jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) ExportData() GoogleDataLossPreventionDiscoveryConfigActionsExportDataOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigActionsExportDataOutputReference
	_jsii_.Get(
		j,
		"exportData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) ExportDataInput() *GoogleDataLossPreventionDiscoveryConfigActionsExportData {
	var returns *GoogleDataLossPreventionDiscoveryConfigActionsExportData
	_jsii_.Get(
		j,
		"exportDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) PublishToChronicle() GoogleDataLossPreventionDiscoveryConfigActionsPublishToChronicleOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigActionsPublishToChronicleOutputReference
	_jsii_.Get(
		j,
		"publishToChronicle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) PublishToChronicleInput() *GoogleDataLossPreventionDiscoveryConfigActionsPublishToChronicle {
	var returns *GoogleDataLossPreventionDiscoveryConfigActionsPublishToChronicle
	_jsii_.Get(
		j,
		"publishToChronicleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) PublishToDataplexCatalog() GoogleDataLossPreventionDiscoveryConfigActionsPublishToDataplexCatalogOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigActionsPublishToDataplexCatalogOutputReference
	_jsii_.Get(
		j,
		"publishToDataplexCatalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) PublishToDataplexCatalogInput() *GoogleDataLossPreventionDiscoveryConfigActionsPublishToDataplexCatalog {
	var returns *GoogleDataLossPreventionDiscoveryConfigActionsPublishToDataplexCatalog
	_jsii_.Get(
		j,
		"publishToDataplexCatalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) PublishToScc() GoogleDataLossPreventionDiscoveryConfigActionsPublishToSccOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigActionsPublishToSccOutputReference
	_jsii_.Get(
		j,
		"publishToScc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) PublishToSccInput() *GoogleDataLossPreventionDiscoveryConfigActionsPublishToScc {
	var returns *GoogleDataLossPreventionDiscoveryConfigActionsPublishToScc
	_jsii_.Get(
		j,
		"publishToSccInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) PubSubNotification() GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotificationOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotificationOutputReference
	_jsii_.Get(
		j,
		"pubSubNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) PubSubNotificationInput() *GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotification {
	var returns *GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotification
	_jsii_.Get(
		j,
		"pubSubNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) TagResources() GoogleDataLossPreventionDiscoveryConfigActionsTagResourcesOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigActionsTagResourcesOutputReference
	_jsii_.Get(
		j,
		"tagResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) TagResourcesInput() *GoogleDataLossPreventionDiscoveryConfigActionsTagResources {
	var returns *GoogleDataLossPreventionDiscoveryConfigActionsTagResources
	_jsii_.Get(
		j,
		"tagResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionDiscoveryConfigActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDataLossPreventionDiscoveryConfigActionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionDiscoveryConfigActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionDiscoveryConfig.GoogleDataLossPreventionDiscoveryConfigActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionDiscoveryConfigActionsOutputReference_Override(g GoogleDataLossPreventionDiscoveryConfigActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionDiscoveryConfig.GoogleDataLossPreventionDiscoveryConfigActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) PutExportData(value *GoogleDataLossPreventionDiscoveryConfigActionsExportData) {
	if err := g.validatePutExportDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExportData",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) PutPublishToChronicle(value *GoogleDataLossPreventionDiscoveryConfigActionsPublishToChronicle) {
	if err := g.validatePutPublishToChronicleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPublishToChronicle",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) PutPublishToDataplexCatalog(value *GoogleDataLossPreventionDiscoveryConfigActionsPublishToDataplexCatalog) {
	if err := g.validatePutPublishToDataplexCatalogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPublishToDataplexCatalog",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) PutPublishToScc(value *GoogleDataLossPreventionDiscoveryConfigActionsPublishToScc) {
	if err := g.validatePutPublishToSccParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPublishToScc",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) PutPubSubNotification(value *GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotification) {
	if err := g.validatePutPubSubNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPubSubNotification",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) PutTagResources(value *GoogleDataLossPreventionDiscoveryConfigActionsTagResources) {
	if err := g.validatePutTagResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTagResources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) ResetExportData() {
	_jsii_.InvokeVoid(
		g,
		"resetExportData",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) ResetPublishToChronicle() {
	_jsii_.InvokeVoid(
		g,
		"resetPublishToChronicle",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) ResetPublishToDataplexCatalog() {
	_jsii_.InvokeVoid(
		g,
		"resetPublishToDataplexCatalog",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) ResetPublishToScc() {
	_jsii_.InvokeVoid(
		g,
		"resetPublishToScc",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) ResetPubSubNotification() {
	_jsii_.InvokeVoid(
		g,
		"resetPubSubNotification",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) ResetTagResources() {
	_jsii_.InvokeVoid(
		g,
		"resetTagResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

