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
					name:  "RAPID_GO_PORT"
					value: #port
				}, {
					name:  "RAPID_GO_POSTGRES_CONNECTION_STRING"
					value: "host=postgresql port=5432 user=postgres password=password dbname=postgres sslmode=disable"
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

#migration_labels: {
	app: "rapid-go-migration"
}

#current_directory: string @tag(dir,var=cwd)

migration: {
	apiVersion: "batch/v1"
	kind:       "Job"
	metadata: {
		name: "migration"
		labels: #migration_labels
	}
	spec: {
		template: {
			metadata: labels: #migration_labels
			spec: {
				containers: [{
					name:  "migrations"
					image: "migrate/migrate"
					imagePullPolicy: "IfNotPresent"
					args: ["-path=/migrations/", "-database", "postgres://postgres:password@postgresql:5432/postgres?sslmode=disable", "up"]
					volumeMounts: [{
						name: "migrations"
						mountPath: "/migrations"
					}]
				}]
				restartPolicy: "Never"
				volumes: [{
					name: "migrations"
					hostPath: {
						path: #current_directory + "/migrations"
						type: "Directory"
					}
				}]
			}
		}
	}
}
