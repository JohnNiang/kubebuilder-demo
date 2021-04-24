# Kubebuilder demo

Firstly, [kind](https://github.com/kubernetes-sigs/kind) a ligth weight tool is essential to run our controller in local development.
Secondly, [kustomize](https://kustomize.io/) and [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder) are needed as well.

Last but not least, let's get start:

```bash
# Build project
make 

# Install demo CRD into k8s cluster in kind
make install

# Run demo controller monitoring demo CRD
make run

# Create demo resource with demo kind 
kubectl apply -f config/samples

# Check controller output
```