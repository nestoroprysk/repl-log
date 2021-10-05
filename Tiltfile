docker_compose("./docker-compose.yml")

docker_build('master', '.', dockerfile='Dockerfile.master',
  live_update = [
    sync('.', '/app'),
    run('go build -o /app/master cmd/master/main.go'),
    restart_container()
  ])

docker_build('secondary', '.', dockerfile='Dockerfile.secondary',
  live_update = [
    sync('.', '/app'),
    run('go build -o /app/secondary cmd/secondary/main.go'),
    restart_container()
  ])

docker_build('test', '.', dockerfile='Dockerfile.test',
  live_update = [
    sync('.', '/app'),
    restart_container()
  ])
