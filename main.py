import socket
import time
import json

def main():
    server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    server.bind(('127.0.0.1', 50051))
    server.listen(5)
    print("🚀 Madi Engine Core (Python) is running on 127.0.0.1:50051")

    while True:
        client, addr = server.accept()
        try:
            while True:
                payload = json.dumps({
                    "status": "ONLINE",
                    "system_load": 45.2,
                    "temperature": 38.4,
                    "vortex_rpm": 1200.0,
                    "active_threads": 8
                }) + "|^"
                client.sendall(payload.encode('utf-8'))
                time.sleep(1)
        except Exception:
            client.close()

if __name__ == '__main__':
    main()
