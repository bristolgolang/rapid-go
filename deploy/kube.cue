deployment: {
	apiVersion: "apps/v1"
	kind:       "Deployment"
	metadata: {
		name: "api"
		labels: app: "rapid-go"
	}
	spec: {
		replicas: 1
		selector: matchLabels: app: "rapid-go"
		template: {
			metadata: labels: app: "rapid-go"
			spec: containers: [{
				name:  "rapid-go"
				image: "rapid-go"
				imagePullPolicy: "Never"
				ports: [{
					containerPort: 32400
				}]
				readinessProbe: httpGet: {
					port: 32400
					path: "/ready"
				}
			}]
		}
	}
}
service: {
	apiVersion: "v1"
	kind:       "Service"
	metadata: name: "api"
	spec: {
		selector: app: "rapid-go"
		ports: [{
			protocol:   "TCP"
			port:       32400
			targetPort: 32400
		}]
	}
}