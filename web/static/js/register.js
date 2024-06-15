document.addEventListener('DOMContentLoaded', () => {
    const ord = new Order();
    const products = [];
    document.addEventListener('htmx:afterSwap', function(event) {
        if (event.detail.target.id === 'popupContainer') {
            let popupWindow = window.open('', 'Form Popup', 'width=600,height=400');
            let content = document.getElementById('popupContainer').innerHTML;
            popupWindow.document.write(content);
            document.getElementById('popupContainer').style.display = 'none';
        }
        const xhr = event.detail.xhr;
        if (xhr && xhr.responseURL.includes('/food') || xhr.responseURL.includes('/drink')) {
            document.querySelectorAll('[id^="product-"]').forEach(product => {
                const id = parseInt(product.getAttribute('data-id'), 10);
                const title = product.getAttribute('data-title');
                const price = parseInt(product.getAttribute('data-price'), 10);
                const pr = new Product(id, title, price, ord, product);
                if (consist(products, id)) {
                    products.push(pr)
                }
            });
        }
    });
});

function consist(prod, id) {
    for (let i = 0; i < prod.length; i++) {
        if(prod[i].id === id) {
            return false;
        }
    }
    return true;
}

class Order {
    constructor() {
        this.items = new Map(); // Using a Map to store items and quantities
        this.totalPrice = 0;
        this.orderListElement = document.getElementById('summary');
        this.totalPriceElement = document.getElementById('totalPrice');
        document.getElementById('sendOrder').addEventListener('click', () => this.cashOut());
        document.getElementById('clearOrder').addEventListener('click', () => this.clearOrder());
    }

    addItem(product) {
        if (this.items.has(product.title)) {
            this.items.get(product.title).quantity++;
        } else {
            this.items.set(product.title, {id: product.id, quantity: 1, price: product.price});
        }
        this.updateOrderDisplay();
    }

    removeItem(product) {
        if (this.items.has(product.title)) {
            const item = this.items.get(product.title);
            if (item.quantity > 1) {
                item.quantity--;
            } else {
                this.items.delete(product.title);
            }
        } else {
            return false;
        }
        this.updateOrderDisplay();
        return true;
    }

    updateOrderDisplay() {
        this.orderListElement.innerHTML = '';
        this.totalPrice = 0;
        this.items.forEach((item, title) => {
            const li = document.createElement('li');
            li.className = 'flex flex-row justify-between';

            // Create the container for the quantity and product name
            const divProduct = document.createElement('div');
            divProduct.className = 'flex flex-grow';

            // Create and append the quantity span
            const spanQuantity = document.createElement('span');
            spanQuantity.textContent = `${item.quantity}x`;
            divProduct.appendChild(spanQuantity);

            // Create and append the product name span
            const spanProduct = document.createElement('span');
            spanProduct.className = 'mx-2';
            spanProduct.textContent = title;
            divProduct.appendChild(spanProduct);

            // Create and append the container for the price
            const divPrice = document.createElement('div');
            divPrice.textContent = `${item.price * item.quantity} Kč`;

            // Append product and price divs to the list item
            li.appendChild(divProduct);
            li.appendChild(divPrice);
            this.orderListElement.appendChild(li);
            this.totalPrice += item.quantity * item.price;
        });
        this.totalPriceElement.textContent = `Celkem: ${this.totalPrice} Kč`;
    }

    cashOut() {
        const orderDetails = Array.from(this.items.entries()).map(([title, item]) =>
            `${item.quantity},${title},${item.price}`
        ).join(";");
        fetch('/order', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({order: orderDetails, total: this.totalPrice})
            // body: JSON.stringify({order: `${orderDetails}`, total: this.totalPrice})
        }).then(response => response.text())
            .then(data => console.log('Order logged:', data))
            .catch(error => console.error('Error logging order:', error));

        this.items.clear();
        this.updateOrderDisplay();
    }

    clearOrder() {
        const orderDetails = Array.from(this.items.entries()).map(([title, item]) =>
            `${item.id}:${item.quantity}`
        ).join(';');
        fetch('/clear-order?cat=food', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(orderDetails)
        }).then(response => response.text())
            .then(data => {console.log('Order logged:', data)
                swapHtmlContent("MenuArea", data)
            })
            .catch(error => console.error('Error logging order:', error));

        this.items.clear();
        this.updateOrderDisplay();
    }

}

class Product {
    constructor(id, title, price, order, element) {
        this.id = id;
        this.title = title;
        this.price = price;
        this.Ord = order;
        this.productElement = element;


        // Event Listener for the product div
        this.productElement.addEventListener('click', (event) => this.incrementCount(event));
        // Find and add event listener to the button
        const decrementButton = this.productElement.querySelector(`#decrementButton-${this.id}`);
        decrementButton.addEventListener('click', (event) => this.decrementCount(event));
    }

    incrementCount(event) {
        if (event.target.tagName !== 'BUTTON') {
            this.Ord.addItem(this);
            this.changeNumber(1);
        }
    }

    decrementCount(event) {
        event.stopPropagation();
        if (this.Ord.removeItem(this)) {
            this.changeNumber(-1);
        }
    }
    changeNumber(number) {
        const countElement = document.getElementById('count-'+this.id);
        if (countElement) {
            const currentCount = parseInt(countElement.textContent, 10);
            if (!isNaN(currentCount)) {
                countElement.textContent = (currentCount + number).toString();
            } else {
                console.error('The content of the element is not a valid number');
            }
        } else {
            console.error('Element with id "count-1" not found');
        }
    }
}

/**
 * Swaps the inner HTML of an element with new content.
 *
 * @param {string} elementId - The ID of the DOM element to update.
 * @param {string} newHtmlContent - The new HTML content to set.
 */
function swapHtmlContent(elementId, newHtmlContent) {
    const element = document.getElementById(elementId);
    if (element) {
        element.innerHTML = newHtmlContent;
        console.log(newHtmlContent);
    } else {
        console.error('Element not found: ', elementId);
    }
}