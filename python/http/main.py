from http.server import BaseHTTPRequestHandler, HTTPServer

class HTTPRequestHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == '/':
            self.send_response(200)
            self.send_header("Content-type", "text/html")
            self.end_headers()
            self.wfile.write(b'ping')

def run(server_class=HTTPServer, handler_class=HTTPRequestHandler):
    server_address = ('', 8080)
    httpd =server_class(server_address, handler_class)
    print("starting server on port 8080")
    httpd.serve_forever()


if __name__=='__main__':
    run()