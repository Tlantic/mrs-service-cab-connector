# mrs-service-cab-connector
Super Koch Client MRS Connector 

### Get Started

To run this module you have to clone the MRS Connector API project and point it to this specific 
connector:


- [mrs-service-connector-api](https://github.com/Tlantic/mrs-service-connector-api)


### Configurations

Set ENV Variables

- Catalog Binary

```
K8S_NAMESPACE=dev
K8S_SIDECAR=localhost:50051
KUBECONFIG=path/to/your/kubeconfig
CONFIG_PROVIDER=k8s
CONFIG_FORMAT=yaml
CONFIG_NAME=mrs-service-cab-connector-catalog
SERVICE_NAME=mrs-service-cab-connector-catalog
API_SERVICE=localhost:8080
```

### gRPC Ports

```
catalog: 8098
auth: 8095
``` 

