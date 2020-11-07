# docker-nghttp

A golang based docker image for providing angular content.

# Examples

    docker run -v /your/data/dir:/etc/nghttp/web matthiasschild/nghttp

Or here an example with docker-compose and traefik:

    version: "3"
    
    services:
      web:
        image: matthiasschild/nghttp
        restart: always
        volumes:
          - /your/data/dir:/etc/nghttp/web
        labels:
          - "traefik.enable=true"
          - "traefik.backend=finstat-web"
          - "traefik.frontend.rule=Host:example.org"
