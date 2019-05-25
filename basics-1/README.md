# Protocol Buffers Basics 1

* [Scalar Types](#scalar-types)
    * [Scalar Types Number](#scalar-types-number)
    * [Scalar Types Boolean](#scalar-types-boolean)
    * [Scalar Types String](#scalar-types-string)
    * [Scalar Types Bytes](#scalar-types-bytes)
* [Tag](#tag)
* [Repeated Fields](#repeated-fields)
* [Comments](#comments)
* [Default Values for Fields](#default-values-for-fields)
* [Enums](#enums)

## Scalar Types

[scalar-types.proto](https://github.com/mer-study-microservices/protocol-buffers/blob/master/basics-1/scalar-types.proto)

```
int32 (age)
string (first name)
string (last name)
bytes (small picture)
bool (profile verified)
float (height)
```


### Scalar Types Number

* Numbers can take various forms based on what values you expect them to have: `double`, `float`, `int32`, `int64`, `uint32`, `uint64`, `sint32`, `sint64`, `fixed32`, `fixed64`, `sfixed32`, `sfixed64`
* `Integer`: For now, let's use `int32`
* `Floating point numbers`:
	* `float` (32 bits)
	* `double` (64 bits) - for more precision (if you really need it)

### Scalar Types Boolean

* Boolean can hold the value `True` or `False`
* It is represented as `bool` in protobuf

### Scalar Types String

* String represents an arbitrary length of text 
* It is represented as `string` in Protobuf
* A string must always contain UTF-8 encoded or 7-bit ASCII text. 

### Scalar Types Bytes

* Bytes represents any sequence of byte array. 
* It is represented as `bytes` in Protobuf
* It will be up to you to interpret what these bytes mean 
* For example you could use these bytes to include a small image 

## Tag

* In Protocol Buffers, field names are not important! (but when programming the fields are important)
* For protobuf the important element is the `tag`
* Smallest tag: 1,
* Largest tag: 2<sup>29</sup> - 1, or 536,870,911
* You also cannot use the numbers 19000 through 19999
* Tags numbered from 1 to 15 use `1 byte` in space, so use them for frequently populated fields 
* Tags numbered from 16 to 2047 use `2 bytes`

## Repeated Fields

Add a list of phone numbers to our Person example!: [repeated-fields.proto](https://github.com/mer-study-microservices/protocol-buffers/blob/master/basics-1/repeated-fields.proto)

Key word `repeated` means a list. 

* To make a "list" or an "array", you can use the concept of repeated fields
* The list can take any number (0 or more) of elements you want 

## Comments

[comments.proto](https://github.com/mer-study-microservices/protocol-buffers/blob/master/basics-1/comments.proto)

* It is possible to embed comments in your `.proto` file 
* It is actually recommended to use comments as a form of documentation for your schemas.

## Default values for Fields

* All fields, if not specified or unknown, will take a default value 

* `bool`: false
* number (`int32`, etc...): 0
* `string`: empty string
* `bytes`: empty bytes
* `enum`: first value
* `repeated`: empty list

## Enums

Add an Enum to `Person` for the field `Eye Color`: [enums.proto](https://github.com/mer-study-microservices/protocol-buffers/blob/master/basics-1/enums.proto)

* If you know all the values a field can take in advance, you can leverage the `Enum` type
* **The first value of an Enum is the default value**
* `Enum` must start by the tag 0 (which is the default value)