api_gateway="./api-gateway/api-gateway-k8s.yaml"
auth_svc="./auth-svc/auth-svc-k8s.yaml"
auth_svc_db="./auth-svc/auth-postgres-k8s.yaml"
mail_svc="./mail-svc/mail-svc-k8s.yaml"
vault_svc="./vault-svc/vault-svc-k8s.yaml"
vault_svc_db="./vault-svc/vault-mongo-k8s.yaml"

kubectl apply --filename "$api_gateway"
kubectl apply --filename "$auth_svc"
kubectl apply --filename "$auth_svc_db"
kubectl apply --filename "$mail_svc"
kubectl apply --filename "$vault_svc"
kubectl apply --filename "$vault_svc_db"

