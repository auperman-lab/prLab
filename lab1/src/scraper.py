import json
import ssl

import requests
import socket

from lab1.src.consts import linella_base_ip, linella_base_port, linella_base_url


def http_scraper(url):
    try:
        response = requests.get(url)
        response.raise_for_status()
        html_content = response.text
        return html_content

    except requests.RequestException as e:
        print(f"HTTP request failed: {e}")


def tcp_scraper(path):
    with socket.create_connection((linella_base_ip, linella_base_port)) as sock:
        context = ssl.create_default_context()
        with context.wrap_socket(sock, server_hostname=linella_base_url) as s:
            https_request = (
                f"GET {path} HTTP/1.1\r\n"
                f"Host: {linella_base_url}\r\n"
                "User-Agent: Python TCP Scraper\r\n"
                "Accept: text/html\r\n"
                "Connection: close\r\n\r\n"
            )

            s.sendall(https_request.encode('utf-8'))

            response = b""
            while True:
                chunk = s.recv(4096)
                if not chunk:
                    break
                response += chunk

            response_text = response.decode('utf-8', errors='ignore')
            headers, _, body = response_text.partition("\r\n\r\n")


            return body



def send_custom_post_request(path, data, username, password):
    with socket.create_connection(("127.0.0.1", 8080)) as sock:
        context = ssl.create_default_context()
        with context.wrap_socket(sock, server_hostname=linella_base_url) as s:
            json_data = json.dumps(data)

            # Construct the POST request
            https_request = (
                f"POST {path} HTTP/1.1\r\n"
                f"Host: {linella_base_url}\r\n"
                "User-Agent: Python TCP Scraper\r\n"
                "Accept: application/json\r\n"
                "Content-Type: application/json\r\n"
                f"Content-Length: {len(json_data)}\r\n"
                "Connection: close\r\n"
                f"Authorization: Basic {username}:{password}\r\n\r\n"  # Basic auth header
                f"{json_data}\r\n"
            )

            # Send the request
            s.sendall(https_request.encode('utf-8'))

            # Receive the response
            response = b""
            while True:
                chunk = s.recv(4096)
                if not chunk:
                    break
                response += chunk

            # Decode and return the response text
            response_text = response.decode('utf-8', errors='ignore')
            headers, _, body = response_text.partition("\r\n\r\n")
            return headers, body