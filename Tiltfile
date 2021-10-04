docker_compose("./docker-compose.yml")

docker_build('master', '.', dockerfile='Dockerfile.master',
  live_update = [
    sync('.', '/app'),
    run('go build -o /app/master cmd/master/main.go'),
    restart_container()
  ])

docker_build('secondary-1', '.', dockerfile='Dockerfile.secondary-1',
  live_update = [
    sync('.', '/app'),
    run('go build -o /app/secondary cmd/secondary/main.go'),
    restart_container()
  ])

docker_build('secondary-2', '.', dockerfile='Dockerfile.secondary-2',
  live_update = [
    sync('.', '/app'),
    run('go build -o /app/secondary cmd/secondary/main.go'),
    restart_container()
  ])
