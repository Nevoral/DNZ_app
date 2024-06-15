import parseTime from "./parseTime";

// Function to parse hx-swap attribute and return swap configuration
const parseSwapAttribute = (element) => {
    const swapConfig = {
        swapStyle: 'innerHTML',
        swapDelay: 0,
        settleDelay: 20,
        transition: false,
        scroll: null,
        show: null,
        focusScroll: null,
    };

    if (element.hasAttribute('hx-swap')) {
        const swapValue = element.getAttribute('hx-swap');
        const parts = swapValue.split(' ');
        swapConfig.swapStyle = parts[0];

        parts.slice(1).forEach(part => {
            const [key, value] = part.split(':');
            switch (key) {
                case 'swap':
                    swapConfig.swapDelay = parseTime(value);
                    break;
                case 'settle':
                    swapConfig.settleDelay = parseTime(value);
                    break;
                case 'transition':
                    swapConfig.transition = value === 'true';
                    break;
                case 'scroll':
                    swapConfig.scroll = value;
                    break;
                case 'show':
                    swapConfig.show = value;
                    break;
                case 'focus-scroll':
                    swapConfig.focusScroll = value === 'true';
                    break;
            }
        });
    }

    return swapConfig;
};

// Helper function to handle scroll modifiers
const handleScroll = (target, scrollValue) => {
    if (scrollValue.startsWith('window')) {
        const position = scrollValue.split(':')[1];
        window.scrollTo({
            top: position === 'top' ? 0 : document.body.scrollHeight,
            behavior: 'smooth'
        });
    } else {
        const [selector, position] = scrollValue.split(':');
        const element = selector === '' ? target : document.querySelector(selector);
        if (element) {
            element.scrollIntoView({ block: position, behavior: 'smooth' });
        }
    }
};

// Helper function to handle show modifiers
const handleShow = (target, showValue) => {
    const [selector, position] = showValue.split(':');
    const element = selector === '' ? target : document.querySelector(selector);
    if (element) {
        element.scrollIntoView({ block: position, behavior: 'smooth' });
    }
};

// Function to perform the swap based on the hx-swap configuration
const performSwap = (element, targetElement, content) => {
    const swapConfig = parseSwapAttribute(element);
    setTimeout(() => {
        switch (swapConfig.swapStyle) {
            case 'innerHTML':
                targetElement.innerHTML = content;
                break;
            case 'outerHTML':
                targetElement.outerHTML = content;
                break;
            case 'beforebegin':
                targetElement.insertAdjacentHTML('beforebegin', content);
                break;
            case 'afterbegin':
                targetElement.insertAdjacentHTML('afterbegin', content);
                break;
            case 'beforeend':
                targetElement.insertAdjacentHTML('beforeend', content);
                break;
            case 'afterend':
                targetElement.insertAdjacentHTML('afterend', content);
                break;
            case 'delete':
                targetElement.remove();
                break;
            case 'none':
                // Do nothing
                break;
            default:
                targetElement.innerHTML = content;
        }

        setTimeout(() => {
            // Handle scroll and show modifiers
            if (swapConfig.scroll) {
                handleScroll(targetElement, swapConfig.scroll);
            }
            if (swapConfig.show) {
                handleShow(targetElement, swapConfig.show);
            }
            if (swapConfig.focusScroll) {
                targetElement.focus({ preventScroll: !swapConfig.focusScroll });
            }
        }, swapConfig.settleDelay);
    }, swapConfig.swapDelay);
};

export default performSwap;