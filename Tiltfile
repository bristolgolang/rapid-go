deployment = local("cue export deploy/kube.cue -e deployment --out yaml")
service = local("cue export deploy/kube.cue -e service --out yaml")

docker_build(
    'rapid-go',
    context='.',
    dockerfile='./Dockerfile',
)

k8s_yaml([deployment, service])