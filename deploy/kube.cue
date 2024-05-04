import "strconv"

#name: "rapid-go"

#port: "32400"

#labels: {
	app: "rapid-go"
}

deployment: {
	apiVersion: "apps/v1"
	kind:       "Deployment"
	metadata: {
		name: "api"
		labels: #labels
	}
	spec: {
		replicas: 1
		selector: matchLabels: #labels
		template: {
			metadata: labels: #labels
			spec: containers: [{
				name:  #name
				image: "rapid-go"
				imagePullPolicy: "Never"
				env: [{
					name:  "PORT"
					value: #port
				}]
				ports: [{
					containerPort: strconv.Atoi(#port)
				}]
				readinessProbe: httpGet: {
					port: strconv.Atoi(#port)
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
		selector: #labels
		ports: [{
			protocol:   "TCP"
			port:       strconv.Atoi(#port)
			targetPort: strconv.Atoi(#port)
		}]
	}
}
