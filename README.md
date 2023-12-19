<a name="readme-top"></a>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#prerequisites">Prerequisites</a></li>
    <li>
     <a href="#how-it-works?">How It Works?</a>
     <ul>
     <li><a href="#usage">Usage</a></li>
     <li><a href="#mapping-types">Mapping Types</a></li>
     <li><a href="#converter-mapping-object">Converter Mapping Object</a></li>
     </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Welcome to the Terraform Schema Converter, a tool designed to simplify the development of a Terraform Provider by streamlining the conversion proccess of data representation between a Terraform SDKv2 resource schema object and a Swagger API generated models.

It is often that Terraform provider developers find themselves writing repeatative and complex code to convert the data representation of the Terraform resource schema object to match the target API structure.  
With this tool, the developer simply configures a mapping object to be consumed by a converter object that will handle the two-way conversion between the Terraform resource object and the Swagger API models, elimnating most of the typing and renaming complexities and allowing the developer to focus on the actual logic required for the resource CRUD implementation.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- Prerequisites -->
## Prerequisites

- Make sure your Terraform provider project is using [Terraform SDKv2][tf-sdk-v2].
- Basic familiarity with Terraform Plugin Provider SDKv2

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- How It Works? -->
## How It Works?

<!-- Usage -->
### Usage

To use the converter, it is required to instantiate the generic converter struct object, often referred to as the converter object which receives a mapping object as an input and a type of a model as a generic type.
With the converter object instantiated, you can now call the construct model or fill terraform schema functions according to your needs.

The following examples have been taken from the project tests:

Converter Mapping & Converter Object  
https://github.com/TeraSky-OSS/terraform-schema-converter/blob/99b0fa845b5bed6d77f572a9a1e0bac38e73e78b/test/data/converter_mapping.go#L6-L115

Construct  
https://github.com/TeraSky-OSS/terraform-schema-converter/blob/99b0fa845b5bed6d77f572a9a1e0bac38e73e78b/test/construct_model_test.go#L6-L67

Fill Terraform Schema  
https://github.com/TeraSky-OSS/terraform-schema-converter/blob/99b0fa845b5bed6d77f572a9a1e0bac38e73e78b/test/fill_tf_schema_test.go#L6-L53


<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- Mapping Types -->
### Mapping Types

The [mapping types][mapping-types] are the building blocks of a converter mapping object which instructs the converter what should be the operation for each field matching between the Terraform object and the model object.

#### Built-in String

This is not a custom type defined in the library but rather a built-in type that instructs the converter that the field mapped to it is a leaf field that should be copied with the value of either the Terraform field or the model field depending on the operation at work.
Leaf fields in the converter mapping will always include the full model path concatenated by a separator.
Example: “subscriptions.individualCustomer.fullName”

The converter library has a predefined default model path separator and a helper function to build the full model path.
Model fields that store an array of structs should be marked with an array marker which is also predefined in the converter library along with a helper function to build that convention.

In the case of Map types, sometimes the developer would want to copy all of the keys between the Terraform object and the model object without specifying them explicitly in the mapping.
The converter library has a predefined marker to instruct the converter to just copy all of the keys as is without renaming them or manipulating them in any way.

#### BlockToStruct

This type of mapping instructs the converter that the Terraform field mapped to it is essentially an unrepeatable block configured as TypeList with MaxItem 1 where the single item is a ‘map[string]interface{}’ that should be mapped to a model field of type ‘struct’.

This custom type is operating as a ‘map[string]interface{}’ under the hood and it expects that each key defined inside it will be inner fields of the Terraform block that will map to the inner fields in the model object.
This type is also useful for model fields of type ‘interface{}’ if the expected value is a ‘map[string]interface’.

#### BlockToStructSlice

This type of mapping instructs the converter that the Terraform field mapped to it is essentially an unrepeatable block configured as TypeList with MaxItem 1 where the single item is a ‘map[string]interface{}’ that should be mapped to a list of ‘structs’ in the model object.
Using this type requires delegating the listing of the items to the inner fields in the Terraform block as explained above.

This custom type is operating as a ‘[]map[string]interface{}’ under the hood and it expects that each item defined inside it will be a ‘BlockToStruct’ that will be mapped to an item of the model object structs array.
The inner leaf fields should mark with an array marker the corresponding part in the full model field path.

#### BlockSliceToStructSlice

This type of mapping instructs the converter that the Terraform field mapped to it is essentially a repeatable block configured as TypeList where each item is a ‘map[string]interface{}’ that should be mapped to a list of ‘structs’ in the model object.

This custom type is operating as a ‘[]map[string]interface{}’ under the hood and it expects that each item defined inside it will be a ‘BlockToStruct’ that will be mapped to an item of the model object structs array.
The inner leaf fields should mark with an array marker the corresponding part in the full model field path.


#### Map

This type of mapping instructs the converter that the Terraform field mapped to it is essentially a map field configured as TypeMap that should be mapped to a map of Go primitive such as ‘map[string]string’ but can also be useful when the model field is of type ‘interface{}’.

When there is no need to explicitly define what keys in the map should be converted, the AllKeysMarker can be used to instruct the converter to copy all of the keys.

This custom type is operating as a ‘map[string]interface{}’ under the hood and it expects that each item defined inside it will be a key that will be mapped to an item in the model object map.
To use the AllKeysMarker, the developer should define a single key in the map at the form of the marker and its value will be the full model field path concatenated with the AllKeysMarker as well.


#### ListToStruct

This type of mapping instructs the converter that the Terraform field mapped to it is essentially a list of primitives configured as TypeList with where each item is one of the primitive types that should be mapped to a model field of type ‘struct’ which contains only one field that should be set and there is no reason to define it as a Terraform block in the Terraform resource schema.

This custom type is operating as a ‘[]string’ under the hood and it expects to have a single string value that maps to the full model field path.


#### Evaluated Field

This type of mapping instructs the converter to initiate a custom function for converting the data at the defined point in the hierarchy granting the developer full control over how a value should be converted.

This custom type is operating as a ‘struct’ under the hood and it expects 2 attributes:
- The full model field path.
- The conversion function, which receives the evaluation mode (construct or flatten) and the value at the defined hierarchy and it returns the converted value.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Mapping Types -->
### Converter Mapping Object

The mapping object is a ‘BlockToStruct’ object that maps the terraform fields to the struct fields by using the special mapping types defined in the converter library and it is being used as a two-way mapper for both the construct and flatten operation.

An example of a converter mapping object can be found here.

When defining a mapping object, the developer would have to follow the following rules:
- The mapping object will always be a ‘BlockToStruct’ variable.
- The keys of the mapping will always be of type ‘String’ and should be set with the name of the field in the Terraform object.
- The hierarchy of the mapping should follow the hierarchy of the Terraform object schema that is being converted.
- The values of each key in the mapping can be one of the mapping types or a string of the concatenated paths in the model object hierarchy that matches the Terraform field.

The mapping object also supports a functionality of unpacking the mapping object into a higher level of hierarchy which is useful when having a data source that stores an array of the root object.

Consider the following structure:
```go
type BusinessesListing struct {
   Businesses []*SomeBusiness `json:"businesses,omitempty"`
  
   TotalCount int `json:"totalCount,omitempty"`
}


type SomeBusiness struct {
   ...
}
```
Suppose you build a resource for creating a new business, the Terraform schema of the resource and the converter mapping object will be defined to support a single business construct.

Now, suppose there is a requirement to build a data source that will list businesses and return an array of businesses.
Such a data source would probably include fields that are dedicated to the listing operation against the API but it will also have to include an array of businesses as a computed field which will add another level of hierarchy in comparison to the resource converter mapping.

To solve this problem, the mapping object provides a functionality to create a new mapping object of a higher hierarchy by giving it the root model path that matches the top of the new hierarchy tree.

Therefore, in this case, the data source converter mapping will include the businesses mapping by unpacking the resource converter mapping with the model path level of “businesses”.
The unpack functionality will recursively build a new mapping object containing all of the model field paths defined in the existing mapping object.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->
## Roadmap

- [ ] Move the generic model notation from the converter object initialization to the construct model function
- [ ] Create a construct function to build a JSON object in addition to the model object construct
- [ ] Model field path should be changed to a dedicated type of array of paths instead of a concatenated string
- [ ] Evaluate function should also return an error to be propagated in the entire conversion process
- [ ] Consider creating functions for converting that will create and dispose a converter object, eliminating the need to hold the converter objects in memory for the lifespan of the provider process

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the Apache 2.0 License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[tf-sdk-v2]: https://developer.hashicorp.com/terraform/plugin/sdkv2
[mapping-types]: https://github.com/TeraSky-OSS/terraform-schema-converter/blob/main/pkg/maptypes/map_types.go