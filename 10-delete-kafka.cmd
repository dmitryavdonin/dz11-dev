kubectl delete all --all -n kafka

kubectl -n kafka delete -f "https://strimzi.io/install/latest?namespace=kafka"

kubectl delete namespace kafka
