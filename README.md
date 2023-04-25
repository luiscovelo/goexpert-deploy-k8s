kind create cluster --name=goexpert

kubectl cluster-info --context kind-goexpert

kubectl get nodes

kubectl apply -f k8s/deployment.yaml

kubectl get pods

kubectl delete pod server-8687bc9d45-m7dc2

kubectl apply -f k8s/services.yaml

kubectl get services

kubectl port-forward service/server-service 8080:8080

probes => verificação de saude de forma geral.

startupProbe => ao subir o pod ele faz a validação
readinessProbe => durante a execução ele faz a validação, se o serviço está ok para receber request
livenessProbe => durante a execução ele faz a validação, se serviço não estiver OK ele restarta