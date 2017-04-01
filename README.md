[![GoDoc](https://godoc.org/github.com/paybyphone/phpipam-sdk-go?status.svg)](https://godoc.org/github.com/paybyphone/phpipam-sdk-go)

# phpipam-sdk-go - Partial SDK for PHPIPAM

`phpipam-sdk-go` is a partial SDK for the [PHPIPAM][1] API.

[1]: https://phpipam.net/api/api_documentation/

This is a WIP and this README along with the rest of the code will develop until
it reaches an acceptable level of maturity that it can be used with some CLI
tools that we are developing to work with PHPIPAM, and possibly a Terraform
provider to help insert data gathered from AWS and beyond.

## Reference

See the [GoDoc][2] for the SDK usage details.

[2]: https://godoc.org/github.com/paybyphone/phpipam-sdk-go

## A Note on Custom Fields

PHPIPAM takes the (unfortunately common) approach to using `ATLER TABLE` to add
custom fields to tables where they can be added (addresses, subnets, VLANs,
VRFs, users, and devices in 1.2). More importantly, they are embedded in the
main resource object when making API requests. This makes custom fields hard to
implement, because since we do not know the names of these fields ahead of time,
we cannot model them into a strongly typed representation of each resource.

The custom field functions in each controller package here resolve this by
querying each controller ahead of time for the custom field schema before
returning or accepting a `map[string]interface{}` with each configured custom
field set (ensuring you only get the custom fields back, and also validating
incoming custom fields to ensure that non-custom field data is not being
updated).

Unfortunately, this approach - and the approach of using a statically typed
model of the resource data in general - means that you cannot use this SDK with
a system that has required custom fields set, as the API request will break on
the `NOT NULL` database constraint. If you plan on using this SDK, keep this in
mind and ensure that your custom fields are not required (try to choose sane
defaults in lieu).

## License

```
Copyright 2017 PayByPhone Technologies, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
