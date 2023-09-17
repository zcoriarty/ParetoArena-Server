## Run locally:

1. `docker build -t paretoarena .`
2. `docker run -p 8080:8080 paretoarena`


## Azure deployment:

### login

1. `az login --use-device-code`
2. `az acr login --name paretodev`

### Steps to push to Dev

1. `docker build -t paretodev .`
2. `docker tag paretodev paretodev.azurecr.io/paretodev:dev`
3. `docker push paretodev.azurecr.io/paretodev:dev`

### Steps to push to Prod

1. `docker build -t pareto .`
2. `docker tag pareto paretodev.azurecr.io/paretodev:latest`
3. `docker push paretodev.azurecr.io/paretodev:latest`
