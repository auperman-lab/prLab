import requests


def html_scraper(url):
    try:
        response = requests.get(url)
        response.raise_for_status()
        html_content = response.text
        return html_content

    except requests.RequestException as e:
        print(f"HTTP request failed: {e}")