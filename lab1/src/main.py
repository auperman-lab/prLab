from lab1.src.consts import base_url
from lab1.src.model.products_price_range import ProductsPriceRange
from lab1.src.parser import parse_products
from lab1.src.scraper import html_scraper
from lab1.src.value_converter import convert_price


def main():
    milk_url = f"{base_url}/ro/catalog/produse_lactate?page=1"

    html = html_scraper(milk_url)
    products = parse_products(html)

    for product in products:
        product.price_eur  = convert_price(product)

    cheap_products = ProductsPriceRange(products, 0, 20)
    print(cheap_products)



if __name__ == "__main__":
    main()