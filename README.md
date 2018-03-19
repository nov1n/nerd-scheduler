### PRE

* kubectl label nodes --all temperature-

### PREP

* watch -n 2 kubectl get po
* kubectl proxy

### DEMO

* go run ./annotator/main.go
* vim deployments/pauser.yaml
* kubectl create -f deployments/pauser.yaml
* kubectl get po
* -- Scheduled as expected
* vim deployments/pauser_custom.yaml
* -- See spec.schedulerName
* kubectl create -f deployments/pauser_custom.yaml
* -- Pending
* ./nerd-scheduler
* -- Scheduled
