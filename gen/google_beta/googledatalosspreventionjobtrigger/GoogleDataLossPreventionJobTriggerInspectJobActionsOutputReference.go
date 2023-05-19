package googledatalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventionjobtrigger/internal"
)

type GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference interface {
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
	Deidentify() GoogleDataLossPreventionJobTriggerInspectJobActionsDeidentifyOutputReference
	DeidentifyInput() *GoogleDataLossPreventionJobTriggerInspectJobActionsDeidentify
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JobNotificationEmails() GoogleDataLossPreventionJobTriggerInspectJobActionsJobNotificationEmailsOutputReference
	JobNotificationEmailsInput() *GoogleDataLossPreventionJobTriggerInspectJobActionsJobNotificationEmails
	PublishFindingsToCloudDataCatalog() GoogleDataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogOutputReference
	PublishFindingsToCloudDataCatalogInput() *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog
	PublishSummaryToCscc() GoogleDataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCsccOutputReference
	PublishSummaryToCsccInput() *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCscc
	PublishToStackdriver() GoogleDataLossPreventionJobTriggerInspectJobActionsPublishToStackdriverOutputReference
	PublishToStackdriverInput() *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishToStackdriver
	PubSub() GoogleDataLossPreventionJobTriggerInspectJobActionsPubSubOutputReference
	PubSubInput() *GoogleDataLossPreventionJobTriggerInspectJobActionsPubSub
	SaveFindings() GoogleDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputReference
	SaveFindingsInput() *GoogleDataLossPreventionJobTriggerInspectJobActionsSaveFindings
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
	PutDeidentify(value *GoogleDataLossPreventionJobTriggerInspectJobActionsDeidentify)
	PutJobNotificationEmails(value *GoogleDataLossPreventionJobTriggerInspectJobActionsJobNotificationEmails)
	PutPublishFindingsToCloudDataCatalog(value *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog)
	PutPublishSummaryToCscc(value *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCscc)
	PutPublishToStackdriver(value *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishToStackdriver)
	PutPubSub(value *GoogleDataLossPreventionJobTriggerInspectJobActionsPubSub)
	PutSaveFindings(value *GoogleDataLossPreventionJobTriggerInspectJobActionsSaveFindings)
	ResetDeidentify()
	ResetJobNotificationEmails()
	ResetPublishFindingsToCloudDataCatalog()
	ResetPublishSummaryToCscc()
	ResetPublishToStackdriver()
	ResetPubSub()
	ResetSaveFindings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference
type jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) Deidentify() GoogleDataLossPreventionJobTriggerInspectJobActionsDeidentifyOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobActionsDeidentifyOutputReference
	_jsii_.Get(
		j,
		"deidentify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) DeidentifyInput() *GoogleDataLossPreventionJobTriggerInspectJobActionsDeidentify {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobActionsDeidentify
	_jsii_.Get(
		j,
		"deidentifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) JobNotificationEmails() GoogleDataLossPreventionJobTriggerInspectJobActionsJobNotificationEmailsOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobActionsJobNotificationEmailsOutputReference
	_jsii_.Get(
		j,
		"jobNotificationEmails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) JobNotificationEmailsInput() *GoogleDataLossPreventionJobTriggerInspectJobActionsJobNotificationEmails {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobActionsJobNotificationEmails
	_jsii_.Get(
		j,
		"jobNotificationEmailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PublishFindingsToCloudDataCatalog() GoogleDataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogOutputReference
	_jsii_.Get(
		j,
		"publishFindingsToCloudDataCatalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PublishFindingsToCloudDataCatalogInput() *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog
	_jsii_.Get(
		j,
		"publishFindingsToCloudDataCatalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PublishSummaryToCscc() GoogleDataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCsccOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCsccOutputReference
	_jsii_.Get(
		j,
		"publishSummaryToCscc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PublishSummaryToCsccInput() *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCscc {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCscc
	_jsii_.Get(
		j,
		"publishSummaryToCsccInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PublishToStackdriver() GoogleDataLossPreventionJobTriggerInspectJobActionsPublishToStackdriverOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobActionsPublishToStackdriverOutputReference
	_jsii_.Get(
		j,
		"publishToStackdriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PublishToStackdriverInput() *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishToStackdriver {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishToStackdriver
	_jsii_.Get(
		j,
		"publishToStackdriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PubSub() GoogleDataLossPreventionJobTriggerInspectJobActionsPubSubOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobActionsPubSubOutputReference
	_jsii_.Get(
		j,
		"pubSub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PubSubInput() *GoogleDataLossPreventionJobTriggerInspectJobActionsPubSub {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobActionsPubSub
	_jsii_.Get(
		j,
		"pubSubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) SaveFindings() GoogleDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputReference
	_jsii_.Get(
		j,
		"saveFindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) SaveFindingsInput() *GoogleDataLossPreventionJobTriggerInspectJobActionsSaveFindings {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobActionsSaveFindings
	_jsii_.Get(
		j,
		"saveFindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionJobTriggerInspectJobActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference_Override(g GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PutDeidentify(value *GoogleDataLossPreventionJobTriggerInspectJobActionsDeidentify) {
	if err := g.validatePutDeidentifyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeidentify",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PutJobNotificationEmails(value *GoogleDataLossPreventionJobTriggerInspectJobActionsJobNotificationEmails) {
	if err := g.validatePutJobNotificationEmailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJobNotificationEmails",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PutPublishFindingsToCloudDataCatalog(value *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) {
	if err := g.validatePutPublishFindingsToCloudDataCatalogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPublishFindingsToCloudDataCatalog",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PutPublishSummaryToCscc(value *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCscc) {
	if err := g.validatePutPublishSummaryToCsccParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPublishSummaryToCscc",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PutPublishToStackdriver(value *GoogleDataLossPreventionJobTriggerInspectJobActionsPublishToStackdriver) {
	if err := g.validatePutPublishToStackdriverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPublishToStackdriver",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PutPubSub(value *GoogleDataLossPreventionJobTriggerInspectJobActionsPubSub) {
	if err := g.validatePutPubSubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPubSub",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) PutSaveFindings(value *GoogleDataLossPreventionJobTriggerInspectJobActionsSaveFindings) {
	if err := g.validatePutSaveFindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSaveFindings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) ResetDeidentify() {
	_jsii_.InvokeVoid(
		g,
		"resetDeidentify",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) ResetJobNotificationEmails() {
	_jsii_.InvokeVoid(
		g,
		"resetJobNotificationEmails",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) ResetPublishFindingsToCloudDataCatalog() {
	_jsii_.InvokeVoid(
		g,
		"resetPublishFindingsToCloudDataCatalog",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) ResetPublishSummaryToCscc() {
	_jsii_.InvokeVoid(
		g,
		"resetPublishSummaryToCscc",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) ResetPublishToStackdriver() {
	_jsii_.InvokeVoid(
		g,
		"resetPublishToStackdriver",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) ResetPubSub() {
	_jsii_.InvokeVoid(
		g,
		"resetPubSub",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) ResetSaveFindings() {
	_jsii_.InvokeVoid(
		g,
		"resetSaveFindings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

