{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}ServiceDirectoryService{% endblock %}
{% block body %}


<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>Service Directory</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/service-directory/docs/">/service-directory/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>v1beta1.projects.locations.namespaces.services</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a href="/service-directory/docs/reference/rest/v1beta1/projects.locations.namespaces.services">/service-directory/docs/reference/rest/v1beta1/projects.locations.namespaces.services</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpservicedirectoryservice<br>gcpservicedirectoryservices<br>servicedirectoryservice</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>servicedirectory.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>servicedirectoryservices.servicedirectory.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>Yes</td>
</tr>

<tr>
    <td>Supports IAM Conditions</td>
    <td>No</td>
</tr>

<tr>
    <td>Supports IAM Audit Configs</td>
    <td>No</td>
</tr>
<tr>
    <td>IAM External Reference Format</td>
    <td>
        
        <p>{% verbatim %}{{namespace}}/services/{{service_id}}{% endverbatim %}</p>
        
    </td>
</tr>


<tr>
<td>{{product_name_short}} Default Average Reconcile Interval In Seconds</td>
<td>600</td>
</tr>
</tbody>
</table>

## Custom Resource Definition Properties


### Annotations
<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>cnrm.cloud.google.com/state-into-spec</code></td>
    </tr>
</tbody>
</table>


### Spec
#### Schema
```yaml
namespaceRef:
  external: string
  name: string
  namespace: string
resourceID: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>namespaceRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The ServiceDirectoryNamespace that this service belongs to.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>namespaceRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `name` field of a `ServiceDirectoryNamespace` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>namespaceRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>namespaceRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Optional. The serviceId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>



### Status
#### Schema
```yaml
conditions:
- lastTransitionTime: string
  message: string
  reason: string
  status: string
  type: string
name: string
observedGeneration: integer
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observation of the resource's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>name</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The resource name for the service in the
format 'projects/*/locations/*/namespaces/*/services/*'.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Typical Use Case
```yaml
# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

apiVersion: servicedirectory.cnrm.cloud.google.com/v1beta1
kind: ServiceDirectoryService
metadata:
  name: servicedirectoryservice-sample
  labels:
    label-one: value-one
spec:
  namespaceRef:
    name: servicedirectoryservice-dep
---
apiVersion: servicedirectory.cnrm.cloud.google.com/v1beta1
kind: ServiceDirectoryNamespace
metadata:
  name: servicedirectoryservice-dep
spec:
  location: us-central1
  projectRef:
    external: ${PROJECT_ID?}
```


{% endblock %}

