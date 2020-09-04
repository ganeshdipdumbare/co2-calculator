# co2 emission calculation CLI tool

The simple CLI tool written in Golang using Cobra library for checking CO2 emission during the journey.

## Description  

The CLI tool is written using clean architecture and it is divided into following modules -  
- app - business logic layer independent of external libs or APIs. As same business layer can be used for writing REST API wihtout having modifications in the other layers.
- cmd - the layer is responsible for communicating with app layer to get co2 emission value for the journey. It also validates if the required paramaters are passed to CLI tool.
- model - the layer is responsible for holding structure details for the core business logic(app layer)

## Usage
### prerequisite
- go v1.14 
### commands
- test - to run unit test cases (app should always have 100% test coverage)  
```
make test
```
- to build   
```
make build
```  
- to install   
```
make install
``` 

After build, the binary can be used as follows -  
- help   
```
./co2-calculator --help
```
- run  
```
./co2-calculator --distance 1800.5 --transportation-method large-petrol-car
```

Note - 
- ``transportation-method`` and ``distance`` flags are required, other flags are optional.  
- Default value for ```output``` flag is "kg" and ```unit-of-distance``` is "km".

## Improvements
- add 100% test coverage for cmd