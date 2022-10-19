## Generate api spec
Note: run the generation in a different folder, not on root

### prerequisites
- see details https://github.com/OpenAPITools/openapi-generator#1---installation
```
wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.2.0/openapi-generator-cli-6.2.0.jar -O openapi-generator-cli.jar
```
or for windows 
```
Invoke-WebRequest -OutFile openapi-generator-cli.jar https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.2.0/openapi-generator-cli-6.2.0.jar
```
### install java8+ to run the generator
```
java -jar openapi-generator-cli.jar help
```

command to generate 
```
java -jar $HOME\bin\openapi-generator-cli.jar generate -g go-gin-server  -i .\spec.yaml --additional-properties=packageName=models
```