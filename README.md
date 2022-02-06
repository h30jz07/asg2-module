# ETI Assignment 2 Package 3.7 - Modules
Module Microservice for ETI Assignment 1 - by Ho Jun Zhi (S10195174G)
## Design Consideration of Microservices and Architecture
The design of the package consists of 1 microservices, excluding the frontend - the module-backend microservice and module-vue microservice.
The backend acts as a data access layer, communicating with other packages within Edu-Fi. The frontend communicates only with the backend to get the data it needs which is then displayed. 
Originally, the backend was omitted because package 3.7 could act as a frontend as it only had view funtionality. Additionally, no database was required as data could be directly accessed from other packages in Edu-Fi. However, as a microservice, additional functionalities could come about that would create the need for a database. Thus, the backend microservice was created - to allow for elasticity of the Module Package's features, as well as act as a anti-corruption layer.

## Architecture Diagram
![Application Architecture Diagram](Asg2_Architecture_Diagram.png)

## Link to container image
[Link to frontend docker image](https://hub.docker.com/repository/docker/h30jz07/module-vue)
[Link to backend docker image](https://hub.docker.com/repository/docker/h30jz07/module-backend)

## Instructions for setting up and running microservices
