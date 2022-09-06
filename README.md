# Golang Bootstrap Project

The Golang bootstrap project will help you set up your golang project and start focusing on the implementation in a seconds.
the project has the followings:

* main file
* main function
* API service with Echo framework
* example API endpoint for retrieving user with user service and repo as well models
* environment variables with viper
* output logs with logrus
* mockery library for generating mock files

---

### Getting started

The policy enforcer is written in a way so that it can serve multiple purposes. to get it to run you need to the following
```bash
#  Run
go run golang_bootstrap_project/cmd/bootstrapProject

```
1. Endpoint will be exposed :4000

---

### Generate mocks
access the directory where your service is and execure 
```bash
#  Run

mockery --name=ServiceName

```

mock folder will be created for you with dedicated mock service
