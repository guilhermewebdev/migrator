FROM migrator:build as build
FROM debian:latest
COPY --from=build /migrator/bin/migrate /bin/migrate
