#!/usr/bin/env python
# Reflects the requests from HTTP methods GET, POST, PUT, and DELETE
# Author: Liam Huang (Yuuki) <liamhuang0205@gmail.com>

try:
    from http.server import HTTPServer, BaseHTTPRequestHandler
except:
    from BaseHTTPServer import HTTPServer, BaseHTTPRequestHandler
from optparse import OptionParser
import json

class RequestHandler(BaseHTTPRequestHandler):

    def do_GET(self):
        request_path = self.path

        self.send_response(200)
        self.send_header('Content-Type', 'application/json')
        self.end_headers()
        json_string = json.dumps(dict(self.headers))
        self.wfile.write(json_string)

        print('%sBegin of Headers%s' % ('-' * 5, '-' * 5))
        for k, v in self.headers.items():
            print('%s: %s' % (k, v))
        print('%sEnd of Headers%s' % ('-' * 5, '-' * 5))

        return None

    do_POST   = do_GET
    do_PUT    = do_GET
    do_DELETE = do_GET

def main():
    port = 8080
    print('Listening on all interfaces:%s' % port)
    server = HTTPServer(('', port), RequestHandler)
    server.serve_forever()

if __name__ == "__main__":
    parser = OptionParser()
    parser.usage = ("Creates an HTTP-header-echo-server.")
    (options, args) = parser.parse_args()
    main()
