#### OpenAPI Object


| Field Name   | Type                                    | Description                                                  |
| ------------ | --------------------------------------- | ------------------------------------------------------------ |
| openapi      | string                                  | **REQUIRED**                                                 |
| info         | [Info Object](#Info Object)             | **REQUIRED**                                                 |
| servers      | [[Server Object](#Server Object)]       |                                                              |
| paths        | [Path Item Object](#Path Item Object)   | **REQUIRED**. /{path} The available paths and operations for the API. |
| components   | [Components Object](#Components Object) |                                                              |
| security     | object                                  |                                                              |
| tags         | object                                  |                                                              |
| externalDocs | object                                  |                                                              |



#### Info Object

| Field Name     | Type                                                         | Description                                                  |
| -------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| title          | `string`                                                     | **REQUIRED**. The title of the application.                  |
| description    | `string`                                                     | A short description of the application. [CommonMark syntax](http://spec.commonmark.org/) MAY be used for rich text representation. |
| termsOfService | `string`                                                     | A URL to the Terms of Service for the API. MUST be in the format of a URL. |
| contact        | [Contact Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#contactObject) | The contact information for the exposed API.                 |
| license        | [License Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#licenseObject) | The license information for the exposed API.                 |
| version        | `string`                                                     | **REQUIRED**. The version of the OpenAPI document (which is distinct from the [OpenAPI Specification version](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#oasVersion) or the API implementation version). |



#### Server Object

| Field Name  | Type                                                         | Description                                                  |
| ----------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| url         | `string`                                                     | **REQUIRED**. A URL to the target host. This URL supports Server Variables and MAY be relative, to indicate that the host location is relative to the location where the OpenAPI document is being served. Variable substitutions will be made when a variable is named in `{`brackets`}`. |
| description | `string`                                                     | An optional string describing the host designated by the URL. [CommonMark syntax](http://spec.commonmark.org/) MAY be used for rich text representation. |
| variables   | Map[`string`, [Server Variable Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#serverVariableObject)] | A map between a variable name and its value. The value is used for substitution in the server's URL template. |



#### Path Item Object

| Field Name  | Type                                                         | Description                                                  |
| ----------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| $ref        | `string`                                                     | Allows for an external definition of this path item. The referenced structure MUST be in the format of a [Path Item Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#pathItemObject). If there are conflicts between the referenced definition and this Path Item's definition, the behavior is *undefined*. |
| summary     | `string`                                                     | An optional, string summary, intended to apply to all operations in this path. |
| description | `string`                                                     | An optional, string description, intended to apply to all operations in this path. [CommonMark syntax](http://spec.commonmark.org/) MAY be used for rich text representation. |
| get         | [Operation Object](#Operation Object)                        | A definition of a GET operation on this path.                |
| put         | [Operation Object](#Operation Object)                        | A definition of a PUT operation on this path.                |
| post        | [Operation Object](#Operation Object)                        | A definition of a POST operation on this path.               |
| delete      | [Operation Object](#Operation Object)                        | A definition of a DELETE operation on this path.             |
| options     | [Operation Object](#Operation Object)                        | A definition of a OPTIONS operation on this path.            |
| head        | [Operation Object](#Operation Object)                        | A definition of a HEAD operation on this path.               |
| patch       | [Operation Object](#Operation Object)                        | A definition of a PATCH operation on this path.              |
| trace       | [Operation Object](#Operation Object)                        | A definition of a TRACE operation on this path.              |
| servers     | [[Server Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#serverObject)] | An alternative `server` array to service all operations in this path. |
| parameters  | [[Parameter Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject)] | A list of parameters that are applicable for all the operations described under this path. These parameters can be overridden at the operation level, but cannot be removed there. The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination of a [name](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterName) and [location](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterIn). The list can use the [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject) to link to parameters that are defined at the [OpenAPI Object's components/parameters](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#componentsParameters). |





#### Components Object

| Field Name      | Type                                                         | Description                                                  |
| --------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| schemas         | Map[`string`, [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#schemaObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject)] | An object to hold reusable [Schema Objects](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#schemaObject). |
| responses       | Map[`string`, [Response Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#responseObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject)] | An object to hold reusable [Response Objects](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#responseObject). |
| parameters      | Map[`string`, [Parameter Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject)] | An object to hold reusable [Parameter Objects](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterObject). |
| examples        | Map[`string`, [Example Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#exampleObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject)] | An object to hold reusable [Example Objects](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#exampleObject). |
| requestBodies   | Map[`string`, [Request Body Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#requestBodyObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject)] | An object to hold reusable [Request Body Objects](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#requestBodyObject). |
| headers         | Map[`string`, [Header Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#headerObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject)] | An object to hold reusable [Header Objects](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#headerObject). |
| securitySchemes | Map[`string`, [Security Scheme Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#securitySchemeObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject)] | An object to hold reusable [Security Scheme Objects](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#securitySchemeObject). |
| links           | Map[`string`, [Link Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#linkObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject)] | An object to hold reusable [Link Objects](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#linkObject). |
| callbacks       | Map[`string`, [Callback Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#callbackObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject)] | An object to hold reusable [Callback Objects](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#callbackObject). |



#### Operation Object

| Field Name   | Type                                                         | Description                                                  |
| ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| tags         | [`string`]                                                   | A list of tags for API documentation control. Tags can be used for logical grouping of operations by resources or any other qualifier. |
| summary      | `string`                                                     | A short summary of what the operation does.                  |
| description  | `string`                                                     | A verbose explanation of the operation behavior. [CommonMark syntax](http://spec.commonmark.org/) MAY be used for rich text representation. |
| externalDocs | [External Documentation Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#externalDocumentationObject) | Additional external documentation for this operation.        |
| operationId  | `string`                                                     | Unique string used to identify the operation. The id MUST be unique among all operations described in the API. Tools and libraries MAY use the operationId to uniquely identify an operation, therefore, it is RECOMMENDED to follow common programming naming conventions. |
| parameters   | [Parameter Object](#Parameter Object)                        | A list of parameters that are applicable for this operation. If a parameter is already defined at the [Path Item](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#pathItemParameters), the new definition will override it but can never remove it. The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination of a [name](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterName) and [location](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterIn). The list can use the [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject) to link to parameters that are defined at the [OpenAPI Object's components/parameters](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#componentsParameters). |
| requestBody  | [Request Body Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#requestBodyObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject) | The request body applicable for this operation. The `requestBody` is only supported in HTTP methods where the HTTP 1.1 specification [RFC7231](https://tools.ietf.org/html/rfc7231#section-4.3.1) has explicitly defined semantics for request bodies. In other cases where the HTTP spec is vague, `requestBody` SHALL be ignored by consumers. |
| responses    | [Responses Object](#Response Object)                         | **REQUIRED**. The list of possible responses as they are returned from executing this operation. |
| callbacks    | Map[`string`, [Callback Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#callbackObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject)] | A map of possible out-of band callbacks related to the parent operation. The key is a unique identifier for the Callback Object. Each value in the map is a [Callback Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#callbackObject) that describes a request that may be initiated by the API provider and the expected responses. The key value used to identify the callback object is an expression, evaluated at runtime, that identifies a URL to use for the callback operation. |
| deprecated   | `boolean`                                                    | Declares this operation to be deprecated. Consumers SHOULD refrain from usage of the declared operation. Default value is `false`. |
| security     | [[Security Requirement Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#securityRequirementObject)] | A declaration of which security mechanisms can be used for this operation. The list of values includes alternative security requirement objects that can be used. Only one of the security requirement objects need to be satisfied to authorize a request. This definition overrides any declared top-level [`security`](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#oasSecurity). To remove a top-level security declaration, an empty array can be used. |
| servers      | [[Server Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#serverObject)] | An alternative `server` array to service this operation. If an alternative `server` object is specified at the Path Item Object or Root level, it will be overridden by this value. |



#### Request Body Object

| Field Name  | Type                                                         | Description                                                  |
| ----------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| description | `string`                                                     | A brief description of the request body. This could contain examples of use. [CommonMark syntax](http://spec.commonmark.org/) MAY be used for rich text representation. |
| content     | Map[`string`, [Media Type Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#mediaTypeObject)] | **REQUIRED**. The content of the request body. The key is a media type or [media type range](https://tools.ietf.org/html/rfc7231#appendix-D) and the value describes it. For requests that match multiple keys, only the most specific key is applicable. e.g. text/plain overrides text/* |
| required    | `boolean`                                                    | Determines if the request body is required in the request. Defaults to `false`. |



#### Response Object

Describes a single response from an API Operation, including design-time, static `links` to operations based on the response.

| Field Name  | Type                                                         | Description                                                  |
| ----------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| description | `string`                                                     | **REQUIRED**. A short description of the response. [CommonMark syntax](http://spec.commonmark.org/) MAY be used for rich text representation. |
| headers     | Map[`string`, [Header Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#headerObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject)] | Maps a header name to its definition. [RFC7230](https://tools.ietf.org/html/rfc7230#page-22) states header names are case insensitive. If a response header is defined with the name `"Content-Type"`, it SHALL be ignored. |
| content     | Map[`string`, [Media Type Object](#Media Type Object)]       | A map containing descriptions of potential response payloads. The key is a media type or [media type range](https://tools.ietf.org/html/rfc7231#appendix-D) and the value describes it. For responses that match multiple keys, only the most specific key is applicable. e.g. text/plain overrides text/* |
| links       | Map[`string`, [Link Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#linkObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject)] | A map of operations links that can be followed from the response. The key of the map is a short name for the link, following the naming constraints of the names for [Component Objects](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#componentsObject). |



#### Media Type Object

Each Media Type Object provides schema and examples for the media type identified by its key.

| Field Name | Type                                                         | Description                                                  |
| ---------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| schema     | [Schema Object](#Schema Object) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject) | The schema defining the type used for the request body.      |
| example    | Any                                                          | Example of the media type. The example object SHOULD be in the correct format as specified by the media type. The `example` object is mutually exclusive of the `examples` object. Furthermore, if referencing a `schema` which contains an example, the `example` value SHALL *override* the example provided by the schema. |
| examples   | Map[ `string`, [Example Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#exampleObject) \| [Reference Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#referenceObject)] | Examples of the media type. Each example object SHOULD match the media type and specified schema if present. The `examples` object is mutually exclusive of the `example` object. Furthermore, if referencing a `schema` which contains an example, the `examples` value SHALL *override* the example provided by the schema. |
| encoding   | Map[`string`, [Encoding Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#encodingObject)] | A map between a property name and its encoding information. The key, being the property name, MUST exist in the schema as a property. The encoding object SHALL only apply to `requestBody` objects when the media type is `multipart` or `application/x-www-form-urlencoded`. |



#### Schema Object

The Schema Object allows the definition of input and output data types. These types can be objects, but also primitives and arrays. This object is an extended subset of the [JSON Schema Specification Wright Draft 00](http://json-schema.org/).

For more information about the properties, see [JSON Schema Core](https://tools.ietf.org/html/draft-wright-json-schema-00) and [JSON Schema Validation](https://tools.ietf.org/html/draft-wright-json-schema-validation-00). Unless stated otherwise, the property definitions follow the JSON Schema.

##### Properties

The following properties are taken directly from the JSON Schema definition and follow the same specifications:

- title
- multipleOf
- maximum
- exclusiveMaximum
- minimum
- exclusiveMinimum
- maxLength
- minLength
- pattern (This string SHOULD be a valid regular expression, according to the [ECMA 262 regular expression](https://www.ecma-international.org/ecma-262/5.1/#sec-7.8.5) dialect)
- maxItems
- minItems
- uniqueItems
- maxProperties
- minProperties
- required
- enum

The following properties are taken from the JSON Schema definition but their definitions were adjusted to the OpenAPI Specification.

- type - Value MUST be a string. Multiple types via an array are not supported.
- allOf - Inline or referenced schema MUST be of a [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#schemaObject) and not a standard JSON Schema.
- oneOf - Inline or referenced schema MUST be of a [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#schemaObject) and not a standard JSON Schema.
- anyOf - Inline or referenced schema MUST be of a [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#schemaObject) and not a standard JSON Schema.
- not - Inline or referenced schema MUST be of a [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#schemaObject) and not a standard JSON Schema.
- items - Value MUST be an object and not an array. Inline or referenced schema MUST be of a [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#schemaObject) and not a standard JSON Schema. `items` MUST be present if the `type` is `array`.
- properties - Property definitions MUST be a [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#schemaObject) and not a standard JSON Schema (inline or referenced).
- additionalProperties - Value can be boolean or object. Inline or referenced schema MUST be of a [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#schemaObject) and not a standard JSON Schema.
- description - [CommonMark syntax](http://spec.commonmark.org/) MAY be used for rich text representation.
- format - See [Data Type Formats](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#dataTypeFormat) for further details. While relying on JSON Schema's defined formats, the OAS offers a few additional predefined formats.
- default - The default value represents what would be assumed by the consumer of the input as the value of the schema if one is not provided. Unlike JSON Schema, the value MUST conform to the defined type for the Schema Object defined at the same level. For example, if `type` is `string`, then `default` can be `"foo"` but cannot be `1`.





#### Parameter Object

Describes a single operation parameter.

A unique parameter is defined by a combination of a [name](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterName) and [location](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterIn).

##### Parameter Locations

There are four possible parameter locations specified by the `in` field:

- path - Used together with [Path Templating](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#pathTemplating), where the parameter value is actually part of the operation's URL. This does not include the host or base path of the API. For example, in `/items/{itemId}`, the path parameter is `itemId`.
- query - Parameters that are appended to the URL. For example, in `/items?id=###`, the query parameter is `id`.
- header - Custom headers that are expected as part of the request. Note that [RFC7230](https://tools.ietf.org/html/rfc7230#page-22) states header names are case insensitive.
- cookie - Used to pass a specific cookie value to the API.

##### Fixed Fields

| Field Name      | Type      | Description                                                  |
| --------------- | --------- | ------------------------------------------------------------ |
| name            | `string`  | **REQUIRED**. The name of the parameter. Parameter names are *case sensitive*.If [`in`](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterIn) is `"path"`, the `name` field MUST correspond to the associated path segment from the [path](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#pathsPath) field in the [Paths Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#pathsObject). See [Path Templating](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#pathTemplating) for further information.If [`in`](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterIn) is `"header"` and the `name` field is `"Accept"`, `"Content-Type"` or `"Authorization"`, the parameter definition SHALL be ignored.For all other cases, the `name` corresponds to the parameter name used by the [`in`](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterIn) property. |
| in              | `string`  | **REQUIRED**. The location of the parameter. Possible values are "query", "header", "path" or "cookie". |
| description     | `string`  | A brief description of the parameter. This could contain examples of use. [CommonMark syntax](http://spec.commonmark.org/) MAY be used for rich text representation. |
| required        | `boolean` | Determines whether this parameter is mandatory. If the [parameter location](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterIn) is "path", this property is **REQUIRED** and its value MUST be `true`. Otherwise, the property MAY be included and its default value is `false`. |
| deprecated      | `boolean` | Specifies that a parameter is deprecated and SHOULD be transitioned out of usage. |
| allowEmptyValue | `boolean` | Sets the ability to pass empty-valued parameters. This is valid only for `query` parameters and allows sending a parameter with an empty value. Default value is `false`. If [`style`](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#parameterStyle) is used, and if behavior is `n/a` (cannot be serialized), the value of `allowEmptyValue` SHALL be ignored. |



####  Data Types

Primitive data types in the OAS are based on the types supported by the [JSON Schema Specification Wright Draft 00](https://tools.ietf.org/html/draft-wright-json-schema-00#section-4.2). Note that `integer` as a type is also supported and is defined as a JSON number without a fraction or exponent part. `null` is not supported as a type (see [`nullable`](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#schemaNullable) for an alternative solution). Models are defined using the [Schema Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#schemaObject), which is an extended subset of JSON Schema Specification Wright Draft 00.

Primitives have an optional modifier property: `format`. OAS uses several known formats to define in fine detail the data type being used. However, to support documentation needs, the `format` property is an open `string`-valued property, and can have any value. Formats such as `"email"`, `"uuid"`, and so on, MAY be used even though undefined by this specification. Types that are not accompanied by a `format` property follow the type definition in the JSON Schema. Tools that do not recognize a specific `format` MAY default back to the `type` alone, as if the `format` is not specified.

The formats defined by the OAS are:

| Common Name | [`type`](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#dataTypes) | [`format`](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md#dataTypeFormat) | Comments                                                     |
| ----------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| integer     | `integer`                                                    | `int32`                                                      | signed 32 bits                                               |
| long        | `integer`                                                    | `int64`                                                      | signed 64 bits                                               |
| float       | `number`                                                     | `float`                                                      |                                                              |
| double      | `number`                                                     | `double`                                                     |                                                              |
| string      | `string`                                                     |                                                              |                                                              |
| byte        | `string`                                                     | `byte`                                                       | base64 encoded characters                                    |
| binary      | `string`                                                     | `binary`                                                     | any sequence of octets                                       |
| boolean     | `boolean`                                                    |                                                              |                                                              |
| date        | `string`                                                     | `date`                                                       | As defined by `full-date` - [RFC3339](http://xml2rfc.ietf.org/public/rfc/html/rfc3339.html#anchor14) |
| dateTime    | `string`                                                     | `date-time`                                                  | As defined by `date-time` - [RFC3339](http://xml2rfc.ietf.org/public/rfc/html/rfc3339.html#anchor14) |
| password    | `string`                                                     | `password`                                                   | A hint to UIs to obscure input.                              |





```
type: object
required:
- name
properties:
  name:
    type: string
  address:
    $ref: '#/components/schemas/Address'
  age:
    type: integer
    format: int32
    minimum: 0
```