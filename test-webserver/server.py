from http.server import BaseHTTPRequestHandler, HTTPServer

routes = {
    "/admin": 200,
    "/login": 403,
    "/status": 200,
    "/debug": 500,
    "/health": 200,
}

class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        status = routes.get(self.path, 404)
        self.send_response(status)
        self.send_header('Content-type', 'text/plain')
        self.end_headers()
        self.wfile.write(f"Path: {self.path} => Status: {status}".encode())

    def log_message(self, format, *args):
        return  # disable logging to stdout

if __name__ == "__main__":
    server = HTTPServer(("", 8000), Handler)
    print("Server läuft auf http://localhost:8000")
    server.serve_forever()