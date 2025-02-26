// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func comparePubsubSubscriptionExpirationPolicy(_, old, new string, _ *schema.ResourceData) bool {
	trimmedNew := strings.TrimLeft(new, "0")
	trimmedOld := strings.TrimLeft(old, "0")
	if strings.Contains(trimmedNew, ".") {
		trimmedNew = strings.TrimRight(strings.TrimSuffix(trimmedNew, "s"), "0") + "s"
	}
	if strings.Contains(trimmedOld, ".") {
		trimmedOld = strings.TrimRight(strings.TrimSuffix(trimmedOld, "s"), "0") + "s"
	}
	return trimmedNew == trimmedOld
}

func GetPubsubSubscriptionCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//pubsub.googleapis.com/projects/{{project}}/subscriptions/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetPubsubSubscriptionApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "pubsub.googleapis.com/Subscription",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/pubsub/v1/rest",
				DiscoveryName:        "Subscription",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetPubsubSubscriptionApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandPubsubSubscriptionName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	topicProp, err := expandPubsubSubscriptionTopic(d.Get("topic"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("topic"); !isEmptyValue(reflect.ValueOf(topicProp)) && (ok || !reflect.DeepEqual(v, topicProp)) {
		obj["topic"] = topicProp
	}
	labelsProp, err := expandPubsubSubscriptionLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	pushConfigProp, err := expandPubsubSubscriptionPushConfig(d.Get("push_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("push_config"); !isEmptyValue(reflect.ValueOf(pushConfigProp)) && (ok || !reflect.DeepEqual(v, pushConfigProp)) {
		obj["pushConfig"] = pushConfigProp
	}
	ackDeadlineSecondsProp, err := expandPubsubSubscriptionAckDeadlineSeconds(d.Get("ack_deadline_seconds"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ack_deadline_seconds"); !isEmptyValue(reflect.ValueOf(ackDeadlineSecondsProp)) && (ok || !reflect.DeepEqual(v, ackDeadlineSecondsProp)) {
		obj["ackDeadlineSeconds"] = ackDeadlineSecondsProp
	}
	messageRetentionDurationProp, err := expandPubsubSubscriptionMessageRetentionDuration(d.Get("message_retention_duration"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("message_retention_duration"); !isEmptyValue(reflect.ValueOf(messageRetentionDurationProp)) && (ok || !reflect.DeepEqual(v, messageRetentionDurationProp)) {
		obj["messageRetentionDuration"] = messageRetentionDurationProp
	}
	retainAckedMessagesProp, err := expandPubsubSubscriptionRetainAckedMessages(d.Get("retain_acked_messages"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("retain_acked_messages"); !isEmptyValue(reflect.ValueOf(retainAckedMessagesProp)) && (ok || !reflect.DeepEqual(v, retainAckedMessagesProp)) {
		obj["retainAckedMessages"] = retainAckedMessagesProp
	}
	expirationPolicyProp, err := expandPubsubSubscriptionExpirationPolicy(d.Get("expiration_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("expiration_policy"); ok || !reflect.DeepEqual(v, expirationPolicyProp) {
		obj["expirationPolicy"] = expirationPolicyProp
	}

	return resourcePubsubSubscriptionEncoder(d, config, obj)
}

func resourcePubsubSubscriptionEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "name")
	return obj, nil
}

func expandPubsubSubscriptionName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/subscriptions/{{name}}")
}

func expandPubsubSubscriptionTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	project, err := getProject(d, config)
	if err != nil {
		return "", err
	}

	topic := d.Get("topic").(string)

	re := regexp.MustCompile("projects\\/(.*)\\/topics\\/(.*)")
	match := re.FindStringSubmatch(topic)
	if len(match) == 3 {
		return topic, nil
	} else {
		// If no full topic given, we expand it to a full topic on the same project
		fullTopic := fmt.Sprintf("projects/%s/topics/%s", project, topic)
		d.Set("topic", fullTopic)
		return fullTopic, nil
	}
}

func expandPubsubSubscriptionLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandPubsubSubscriptionPushConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedOidcToken, err := expandPubsubSubscriptionPushConfigOidcToken(original["oidc_token"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOidcToken); val.IsValid() && !isEmptyValue(val) {
		transformed["oidcToken"] = transformedOidcToken
	}

	transformedPushEndpoint, err := expandPubsubSubscriptionPushConfigPushEndpoint(original["push_endpoint"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPushEndpoint); val.IsValid() && !isEmptyValue(val) {
		transformed["pushEndpoint"] = transformedPushEndpoint
	}

	transformedAttributes, err := expandPubsubSubscriptionPushConfigAttributes(original["attributes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAttributes); val.IsValid() && !isEmptyValue(val) {
		transformed["attributes"] = transformedAttributes
	}

	return transformed, nil
}

func expandPubsubSubscriptionPushConfigOidcToken(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServiceAccountEmail, err := expandPubsubSubscriptionPushConfigOidcTokenServiceAccountEmail(original["service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !isEmptyValue(val) {
		transformed["serviceAccountEmail"] = transformedServiceAccountEmail
	}

	transformedAudience, err := expandPubsubSubscriptionPushConfigOidcTokenAudience(original["audience"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAudience); val.IsValid() && !isEmptyValue(val) {
		transformed["audience"] = transformedAudience
	}

	return transformed, nil
}

func expandPubsubSubscriptionPushConfigOidcTokenServiceAccountEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionPushConfigOidcTokenAudience(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionPushConfigPushEndpoint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionPushConfigAttributes(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandPubsubSubscriptionAckDeadlineSeconds(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionMessageRetentionDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionRetainAckedMessages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSubscriptionExpirationPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTtl, err := expandPubsubSubscriptionExpirationPolicyTtl(original["ttl"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTtl); val.IsValid() && !isEmptyValue(val) {
		transformed["ttl"] = transformedTtl
	}

	return transformed, nil
}

func expandPubsubSubscriptionExpirationPolicyTtl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
