<a name="readme-top"></a>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#prerequisites">Prerequisites</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Welcome to the Terraform Schema Converter, a tool designed to simplify the development of a Terraform Provider by streamlining the conversion proccess of data representation between a Terraform SDKv2 resource schema object and a Swagger API generated models.

It is often that Terraform provider developers find themselves writing repeatative and complex code to convert the data representation of the Terraform resource schema object to match the target API structure.
With this tool, the developer simply configures a mapping object to be consumed by a converter object that will handle the two-way conversion between the Terraform resource object and the Swagger API, elimnating most of the typing and renaming complexities and allowing the developer to focus on the actual logic required for the resource CRUD implementation.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- Prerequisites -->
## Prerequisites

- Make sure your Terraform provider project is using Terraform SDKv2
- Basic familiarity with Terraform Plugin Provider SDKv2

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Usage

To use the converter, it is required to instantiate the generic converter struct object, often referred to as the converter object which receives a mapping object as an input and a type of a model as a generic type.
With the converter object instantiated, you can now call the construct model or fill terraform schema functions according to your needs.

The following examples have been taken from the project tests:

Converter Mapping & Converter Object  
[converter-mapping-object]

Construct  
[converter-construct-model]

Fill Terraform Schema  
[converter-fill-terraform-schema]


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
[converter-mapping-object]: https://github.com/TeraSky-OSS/terraform-schema-converter/blob/99b0fa845b5bed6d77f572a9a1e0bac38e73e78b/test/data/converter_mapping.go#L6-L115
[converter-construct-model]: https://github.com/TeraSky-OSS/terraform-schema-converter/blob/99b0fa845b5bed6d77f572a9a1e0bac38e73e78b/test/construct_model_test.go#L6-L67
[converter-fill-terraform-schema]: https://github.com/TeraSky-OSS/terraform-schema-converter/blob/99b0fa845b5bed6d77f572a9a1e0bac38e73e78b/test/fill_tf_schema_test.go#L6-L53