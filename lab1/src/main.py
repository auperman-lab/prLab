from lab1.src.consts import base_url
from lab1.src.parser import parse_products
from lab1.src.scraper import html_scraper



def main():
    milk_url = f"{base_url}/ro/catalog/produse_lactate?page=1"

    html = html_scraper(milk_url)
    products = parse_products(html)

    for p in products:
        print(p)



if __name__ == "__main__":
    main()