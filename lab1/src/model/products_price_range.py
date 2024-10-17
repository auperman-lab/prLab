from functools import reduce
from datetime import datetime

class ProductsPriceRange:
    def __init__(self, products, lower, upper):
        self.products = products
        self.filtered_products = self._filter_products(lower,upper)
        self.total_price = self._sum_prices()
        self.updated_at = datetime.utcnow()

    def _filter_products(self, min_price, max_price):
        filtered_products = list(filter(
            lambda p: p.price_now and min_price <= float(p.price_now) <= max_price,
            self.products
        ))
        return filtered_products

    def _sum_prices(self):
        total_price = reduce(
        lambda acc, p: acc + p.price_now,
            self.filtered_products,
            0.0
        )
        return total_price

    def __str__(self):
        return (
            f"Filtered Products: {[p.name for p in self.filtered_products]} \n\n"
            f"Total Price: {self.total_price:.2f} MDL \n"
            f"Update_at: {self.updated_at}"
        )

