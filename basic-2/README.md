# Protocol Buffers Basics 2

* [Defining Multiple Messages in the Same File](#defining-multiple-messages-in-the-same-file)
* [Nesting Types](#nesting-types)
* [Importing Types](#importing-types)
* [Packages](#packages)

## Defining Multiple Messages in the Same File 

Create a message `Date` and add it to `Person` as a field for a birthday.: [same-level-message.proto](https://github.com/mer-study-microservices/protocol-buffers/blob/master/basic-2/same-level-message.proto)

* It is possible, in the same `.proto` file, to define multiple types 

## Nesting Types

Add a field `Address` and addd it to `Person` to have multiple addresses.: [nested-messages.proto](https://github.com/mer-study-microservices/protocol-buffers/blob/master/basic-2/nested-messages.proto)

* It is possible to define types within types 
* The reasons could be:
	* Avoiding naming conflicts
	* Enforcing some level of "locality" for that type 
* You can nest types as deeply as you want. 



## Importing Types

Move `Date` out of `Person` and import the date file.: [imports-types](https://github.com/mer-study-microservices/protocol-buffers/tree/master/basic-2/imports-types)

* It is possible to have different types in `.proto` files.
* It is useful to re-use code and import other `.proto` files. 

## Packages

[packages-example](https://github.com/mer-study-microservices/protocol-buffers/tree/master/basic-2/packages-example)

* It is very important to define the packages in which your protocol buffer messages live
	* when your code gets compiled, it will be placed at the package you indicated. 
	* It is also helps to prevent name conflicts between messages (my.package.Person)

