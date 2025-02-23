from flask import Flask
import sys

app = Flask(__name__)

server_name = sys.argv[1] 
port = int(sys.argv[2]) 

@app.route('/')
def home():
    return f"Hello from {server_name} running on port {port}!"

if __name__ == '__main__':
    app.run(host="127.0.0.1", port=port)

