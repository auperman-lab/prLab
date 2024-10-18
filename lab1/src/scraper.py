import ssl

import requests
import socket

from lab1.src.consts import base_ip, base_port, base_url


def http_scraper(url):
    try:
        response = requests.get(url)
        response.raise_for_status()
        html_content = response.text
        return html_content

    except requests.RequestException as e:
        print(f"HTTP request failed: {e}")


def tcp_scraper(path):
    with socket.create_connection((base_ip, base_port)) as sock:
        context = ssl.create_default_context()
        with context.wrap_socket(sock, server_hostname=base_url) as s:
            https_request = (
                f"GET {path} HTTP/1.1\r\n"
                f"Host: {base_url}\r\n"
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