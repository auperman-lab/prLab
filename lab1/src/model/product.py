from lab1.src.value_converter import convert_price


class Product:
    def __init__(self, name, price_now, price_old, discount, link, category, sub_category):
        self.name = name
        self.price_now = price_now
        self.price_old = price_old
        self.discount = discount
        self.link = link
        self.category = category
        self.sub_category = sub_category
        self.price_eur = 0

    def __str__(self):
        return (
            f"Product Name: {self.name}\n"
            f"Old Price: {self.price_old} MDL\n"
            f"Current Price: {self.price_now} MDL\n"
            f"Discount: {self.discount}\n"
            f"Link: {self.link}\n"
            f"Category: {self.category}\n"
            f"Sub-Category: {self.sub_category}\n"
            f"EUR Price: {self.price_eur} EUR\n"
        )

    def validate(self):
        # Validate name
        if not isinstance(self.name, str) or not self.name.strip():
            raise ValueError("Invalid product name")

        # Validate price_old
        if isinstance(self.price_now, str) and '/' in self.price_now:
            price_part, weight_part = self.price_now.split('/')
            self.name = self.name + " per " +  weight_part
            if not self._is_valid_price(price_part):
                raise ValueError("Invalid price_old format" + self.name)
            self.price_now = float(price_part)
        elif not self._is_valid_price(self.price_now):
            raise ValueError("Invalid price_old" + self.name)
        else:
            self.price_now = float(self.price_now)


        # Validate price_new (it can be empty)
        if self.price_old and not self._is_valid_price(self.price_old):
            raise ValueError("Invalid price_new" + self.name)
        elif self.price_old:
            self.price_old = float(self.price_old)
        else:
            self.price_old = 0

        # Validate discount (it can be empty but if present, should be a valid percentage)
        if self.discount and not self._is_valid_discount(self.discount):
            raise ValueError("Invalid discount" + self.name)

        # Validate category and sub-category
        if not isinstance(self.category, str) or not self.category:
            raise ValueError("Invalid category" + self.name)

        if not isinstance(self.sub_category, str) or not self.sub_category:
            raise ValueError("Invalid sub-category" + self.name)

    def _is_valid_price(self, price):
        try:
            return isinstance(float(price), float) and float(price) >= 0
        except ValueError:
            return False

    def _is_valid_discount(self, discount):
        return discount.startswith('-') and discount.endswith('%') and discount[1:-1].isdigit()
