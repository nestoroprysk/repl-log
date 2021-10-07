docker_compose("./docker-compose.yml")

docker_build('node', '.',
  live_update = [
    sync('.', '/app'),
    run('go build'),
    restart_container()
  ])

docker_build('test', '.', dockerfile='Dockerfile.test',
  live_update = [
    sync('.', '/app'),
    restart_container()
  ])
