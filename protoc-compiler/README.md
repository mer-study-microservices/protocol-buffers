# Setting up Protoc Compiler

## Using `protoc` to generate code in any language 

```bash
protoc
```

```bash
mkdir python
mkdir java
```

```bash
# python
protoc -I=proto --python_out=python proto/*.proto
# java
protoc -I=proto --java_out=java proto/*.proto
```
