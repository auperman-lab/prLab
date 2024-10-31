from lab1.src.consts import base_url
from lab1.src.parser import parse_all_categories
from lab1.src.scraper import http_scraper


def main():
    milk_path = "/ro/catalog/produse_lactate?page=1"
    milk_url = "https://" + base_url + milk_path
    categories_path = "/ro/catalog"
    categories_url = "https://"+base_url+ categories_path

    html = http_scraper(categories_url)

    parse_all_categories(html)

    # html = tcp_scraper(milk_path)
    #
    #
    # products = parse_products(html)
    #




if __name__ == "__main__":
    main()