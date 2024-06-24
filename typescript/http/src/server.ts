import http, { IncomingMessage, ServerResponse } from "http";
const hostName = "127.0.0.1";
const port = 8080;

const server = http.createServer((req: IncomingMessage, res: ServerResponse) => {
  res.statusCode = 200;
  res.setHeader("Content-type", "application/json");
  const url = req.url;
  const method = req.method;

  if (method === "GET") {
    if (url === "/") {
      res.end(JSON.stringify({ message: "ping" }));
    }
  } else {
    res.statusCode = 405;
    res.end(JSON.stringify({ error: "Method not Allowed" }));
  }
});

server.listen(port, hostName, () => {
  console.log(`Server running at http://${hostName}:${port}/`);
});
