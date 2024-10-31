from bs4 import BeautifulSoup
from lab1.src.model.product import Product
from lab1.src.scraper import http_scraper, tcp_scraper


def parse_products(html_content):
    soup = BeautifulSoup(html_content, 'html.parser')  # Parse the HTML content
    products = []

    for item in soup.find_all('div', {"class": ["products-catalog-content__item", "products-catalog-content__item_marked"]}):

        name_div = item.find("div", {"class": "products-catalog-content__body"})
        p_name = name_div.find("a", {"class": "products-catalog-content__name"}).string  # task2

        p_link = name_div.find("a", {"class": "products-catalog-content__name"})
        p_link = p_link["href"]

        # product_html = http_scraper("https://" + base_url + p_link)
        product_html = tcp_scraper(p_link)
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

def parse_all_categories(html_content):
    soup = BeautifulSoup(html_content, 'html.parser')
    category_div = soup.find("div", {"class": "catalog__lft"})
    category_list = {}
    for li in  category_div.find("ul").select("li"):
        if not li.find_parent("ul", {"class":"level2"}):
            span = li.find("span")
            if span is not None:
                category_list[span.get_text(strip=True)] = list()
            else:
                category_list[li.find("a").get_text(strip=True)] = list()
        else:
            category_list[list(category_list.keys())[-1]].append(li.find("a").get_text(strip=True))

    for category in category_list.keys():
        print(category, category_list[category])
    return

