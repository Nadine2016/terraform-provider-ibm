// Code generated by go-swagger; DO NOT EDIT.

package l_baas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams creates a new GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams object
// with the default values initialized.
func NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams() *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams {
	var ()
	return &GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParamsWithTimeout creates a new GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParamsWithTimeout(timeout time.Duration) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams {
	var ()
	return &GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams{

		timeout: timeout,
	}
}

// NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParamsWithContext creates a new GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParamsWithContext(ctx context.Context) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams {
	var ()
	return &GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams{

		Context: ctx,
	}
}

// NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParamsWithHTTPClient creates a new GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParamsWithHTTPClient(client *http.Client) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams {
	var ()
	return &GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams{
		HTTPClient: client,
	}
}

/*GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams contains all the parameters to send to the API endpoint
for the get load balancers ID listeners listener ID policies policy ID rules rule ID operation typically these are written to a http.Request
*/
type GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The load balancer identifier

	*/
	ID string
	/*ListenerID
	  The listener identifier

	*/
	ListenerID string
	/*PolicyID
	  The policy identifier

	*/
	PolicyID string
	/*RuleID
	  The rule identifier

	*/
	RuleID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) WithTimeout(timeout time.Duration) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) WithContext(ctx context.Context) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) WithHTTPClient(client *http.Client) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) WithGeneration(generation int64) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) WithID(id string) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) SetID(id string) {
	o.ID = id
}

// WithListenerID adds the listenerID to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) WithListenerID(listenerID string) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams {
	o.SetListenerID(listenerID)
	return o
}

// SetListenerID adds the listenerId to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) SetListenerID(listenerID string) {
	o.ListenerID = listenerID
}

// WithPolicyID adds the policyID to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) WithPolicyID(policyID string) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) SetPolicyID(policyID string) {
	o.PolicyID = policyID
}

// WithRuleID adds the ruleID to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) WithRuleID(ruleID string) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WithVersion adds the version to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) WithVersion(version string) *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get load balancers ID listeners listener ID policies policy ID rules rule ID params
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param listener_id
	if err := r.SetPathParam("listener_id", o.ListenerID); err != nil {
		return err
	}

	// path param policy_id
	if err := r.SetPathParam("policy_id", o.PolicyID); err != nil {
		return err
	}

	// path param rule_id
	if err := r.SetPathParam("rule_id", o.RuleID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
