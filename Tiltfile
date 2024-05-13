load('ext://helm_resource', 'helm_resource', 'helm_repo')

deployment = local("cue export deploy/kube.cue -e deployment --out yaml")
service = local("cue export deploy/kube.cue -e service --out yaml")
migration = local("cue export deploy/kube.cue -T -e migration --out yaml")

docker_build(
    'rapid-go',
    context='.',
    dockerfile='./Dockerfile',
)

k8s_yaml([deployment, service, migration])
k8s_resource(workload='api', port_forwards=32400)

helm_repo('bitnami', 'https://charts.bitnami.com/bitnami')
helm_resource('postgresql', 'bitnami/postgresql', resource_deps=['bitnami'], flags=['--set=global.postgresql.auth.password=password'])
k8s_resource(workload='postgresql', port_forwards=5432)