# Docker CPU Usage Stats

This project builds a simple static golang binary and corresponding Docker image that queries the Docker socket to retrieve the list of running containers and, for each, outputs the total CPU usage and elapsed time since the container was started.

The Docker image can be run as follows:
```bash
docker run -v /var/run/docker.sock:/var/run/docker.sock --rm dcurrie/cpu-usage
