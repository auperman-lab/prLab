from bs4 import BeautifulSoup
from lab1.src.consts import base_url
from lab1.src.model.product import Product
from lab1.src.scraper import html_scraper


def parse_products(html_content):
    soup = BeautifulSoup(html_content, 'html.parser')  # Parse the HTML content
    products = []

    for item in soup.find_all('div', {"class": ["products-catalog-content__item", "products-catalog-content__item_marked"]}):

        name_div = item.find("div", {"class": "products-catalog-content__body"})
        p_name = name_div.find("a", {"class": "products-catalog-content__name"}).string  # task2

        p_link = name_div.find("a", {"class": "products-catalog-content__name"})
        p_link = p_link["href"]

        product_html = html_scraper(base_url + p_link)
        p_category, p_subcategory = parse_categories(product_html)


        if name_div.find("span", {"class": ["price-products-catalog-content__static"]}):
            price_now = name_div.find("span", {"class": ["price-products-catalog-content__static"]}).get_text(strip=True)
            price_old = ""
            discount = ""
        else:
            price_old = name_div.find("span", {"class": "price-products-catalog-content__old"}).get_text(strip=True)
            price_now = name_div.find("span", {"class": "price-products-catalog-content__new"}).get_text(strip=True)
            discount = name_div.find("div", {"class": "price-products-catalog-content__discount"}).get_text(strip=True)

        product = Product(
            name=p_name,
            link=p_link,
            price_old= price_old,
            price_now=price_now,
            discount=discount,
            category=p_category,
            sub_category=p_subcategory
        )
        try:
            product.validate()
            products.append(product)
        except Exception as e:
            print(e)


    return products


def parse_categories(html_content):
    soup = BeautifulSoup(html_content, 'html.parser')
    product_breadcrumbs = [i.get_text(strip=True) for i in soup.find("ul", {"class": "breadcrumbs"}).find_all("li")]

    p_category = product_breadcrumbs[2]
    p_subcategory = product_breadcrumbs[3]
    return p_category, p_subcategory