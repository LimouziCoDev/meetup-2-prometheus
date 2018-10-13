FROM alpine

# Copy the binary file and set it as entrypoint
ADD app /
ENTRYPOINT ["/app"]

# The service listens on port 8888 by default.
EXPOSE 8888