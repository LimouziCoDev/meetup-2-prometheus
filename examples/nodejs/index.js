var express = require("express");
var app = express();

const Prometheus = require("prom-client");
const httpRequestDurationMicroseconds = new Prometheus.Histogram({
  name: "http_request_duration_microseconds",
  help: "Duration of HTTP requests in ms",
  labelNames: ["handler", "code"],
  // buckets for response time from 0.1ms to 500ms
  buckets: [0.1, 5, 15, 50, 100, 200, 300, 400, 500]
});

// Runs before each requests
app.use((req, res, next) => {
  res.locals.startEpoch = Date.now();
  next();
});

app.get("/demo/", (req, res, next) => {
  res.send("Welcome to my demo");
  next();
});

// Metrics endpoint
app.get("/metrics", (req, res) => {
  res.set("Content-Type", Prometheus.register.contentType);
  res.end(Prometheus.register.metrics());
});

// Runs after each requests
app.use((req, res, next) => {
  const responseTimeInMs = Date.now() - res.locals.startEpoch;

  httpRequestDurationMicroseconds
    .labels(req.route.path, res.statusCode)
    .observe(responseTimeInMs);

  next();
});

app.listen(8888, () => {
  console.log("Server running on port 8888");
});

// Graceful shutdown
process.on("SIGTERM", () => {
  clearInterval(metricsInterval);

  server.close(err => {
    if (err) {
      console.error(err);
      process.exit(1);
    }

    process.exit(0);
  });
});
