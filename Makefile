k8s-customize-controller:
	go build
run:
	./k8s_customize_controller -kubeconfig=/var/run/kubernetes/admin.kubeconfig
clan:
	rm -f k8s-customize-controller
