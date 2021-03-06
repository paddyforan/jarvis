# Resource Files

Resource files are machine-readable files that define the resources made available by the APIs being generated. They are YAML documents that describe the resources, along with their interactions and properties.

## Sample

```yml
id: message
name: Message
description: A string of data that is meant to be processed.
parent: mq/queue
url_slug: id
url_prefix: messages
properties:
- id: id
  type: string
  description: A unique, API-generated identifier for this resource.
  permissions:
  - r
- id: body
  type: string
  description: The data that is meant to be processed.
  permissions:
  - r
  - w
- id: timeout
  type: duration
  description: The amount of time a single reservation lasts for this message by default.
    This can be overwritten on a per-request basis.
  default: 60
  maximum: 86400
  minimum: 30
  permissions:
  - r
  - w
- id: delay
  type: duration
  description: The number of seconds to delay putting a message on the queue. The
    message will not be available until this time has elapsed.
  default: 0
  maximum: 604800
  permissions:
  - r
  - w
- id: expires_in
  type: duration
  description: The number of seconds to keep a message on the queue before it is automatically
    deleted.
  default: 604800
  maximum: 2592000
  permissions:
  - r
  - w
interactions:
- id: delete
  verb: destroy
  description: Remove a message from the queue.
  name: Delete a Message
- id: peek
  verb: list
  description: Retrieve messages from the queue without reserving them.
  name: Peek at Messages
  params:
  - id: "n"
    type: int
    description: The maximum number of messages to return.
    default: 1
    maximum: 100
- id: push
  verb: create
  description: Add messages to the end of the queue.
  name: Push Messages
```

## Properties

The properties of the resource object determine the resource's behaviour and properties.

<table>
<tr><th>Field</th><th>Required</th><th>Description</th></tr>
<tr><td>id</td><td>Yes</td><td>An API-unique ID for the resource. Will be used as the key for this resource in request and response objects.</td></tr>
<tr><td>name</td><td>Yes</td><td>A human-friendly name for the resource.</td></tr>
<tr><td>description</td><td>Yes</td><td>A human-friendly description of the resource.</td></tr>
<tr><td>parent</td><td>No</td><td>The ID of the resource this resource is a child of, if this resource has a parent. The ID must be in the form &quot;{API ID}/{RESOURCE ID}&quot;.</td></tr>
<tr><td>parent_is_collection</td><td>No</td><td>When set to &quot;true&quot;, the parent's slug will not be used when constructing a URL. Instead, the parent's prefix will immediately precede this resource's prefix.</td></tr>
<tr><td>url_slug</td><td>Yes</td><td>The property whose value will be used as a slug when constructing URLs for this resource.</td></tr>
<tr><td>url_prefix</td><td>Yes</td><td>The URL prefix that will precede the slug. This should be a short slug that describes the collection of resources.</td></tr>
<tr><td>plural_id</td><td>No</td><td>The plural form of the id for this resource, to be used as the key for this resource in request and response objects containing more than one of the resource. If not set, defaults to url_prefix.</td></tr>
<tr><td>properties</td><td>Yes</td><td>Property objects describing the properties of the resource.</td></tr>
<tr><td>interactions</td><td>No</td><td>Interaction objects describing the possible actions that can be performed against the resource.</td></tr>
</table>

Property objects have their own properties, describing the constraints of the property:

<table>
<tr><th>Field</th><th>Required</th><th>Description</th></tr>
<tr><td>id</td><td>Yes</td><td>A resource-unique ID for the property.</td></tr>
<tr><td>type</td><td>Yes</td><td>The type of value expected by the property. Should be one of the following: string, bytes, duration, datetime, int, float, boolean, array, object, pointer</td></tr>
<tr><td>description</td><td>Yes</td><td>A human-friendly description of the property.</td></tr>
<tr><td>format</td><td>No</td><td>A regular expression that the value of the property must match. Requests with properties not matching this format will be considered invalid unless specifically overridden in the interaction.</td></tr>
<tr><td>maximum</td><td>No</td><td>A maximum value, as an int, for the value of the property. For strings, bytes, and arrays, the length is compared to the maximum value. For durations, datetimes, ints, and floats, the value is compared to the maximum value. Objects and pointers cannot have maximum values.</td></tr>
<tr><td>minimum</td><td>No</td><td>A minimum value, as an int, for the value of the property. For strings, bytes, and arrays, the length is compared to the minimum value. For durations, datetimes, ints, and floats, the value is compared to the minimum value. Objects and pointers cannot have minimum values.</td></tr>
<tr><td>default</td><td>No</td><td>A default value that will be used if the property is omitted. Properties without a default value are considered required and will cause a request to be considered invalid if they are not specified. The word &ldquo;nil&rdquo; can be used to signify that, by default, a property is not set.</td></tr>
<tr><td>value_type</td><td>No</td><td>For pointers, the type of the value the pointer is pointing to. Requests pointing to other types will be considered invalid.</td></tr>
<tr><td>permissions</td><td>No</td><td>An array of permissions (&quot;r&quot; for read, &quot;w&quot; for write) that clients have for this property.</td></tr>
<tr><td>repeated</td><td>No</td><td><strong>Used only in URL parameters.</strong> If set to true, the param is expected to be repeated (e.g., ?param=a&param=b&param=c).</td></tr>
</table>

Interaction objects have their own properties, describing the constraints and requirements of the interaction:

<table>
<tr><th>Field</th><th>Required</th><th>Description</th></tr>
<tr><td>id</td><td>Yes</td><td>A resource-unique ID for the interaction.</td></tr>
<tr><td>name</td><td>Yes</td><td>A human-friendly identified for the interaction.</td></tr>
<tr><td>verb</td><td>Yes</td><td>A description of what the interaction does to the resource. Accepted values are: create, get, list, update, destroy</td></tr>
<tr><td>accept_many</td><td>No</td><td>If set to &quot;true&quot;, the request will expect an array of objects in the request, not just one.</td></tr>
<tr><td>description</td><td>Yes</td><td>A human-friendly description of the interaction.</td></tr>
<tr><td>params</td><td>No</td><td>An array of property objects describing URL parameters that are accepted or required for this request.</td></tr>
</table>
