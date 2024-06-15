const parseTargetValues = (value, element) => {
    let targets = [];

    const addTarget = (selector) => {
        let target;
        if (selector === 'this') {
            target = element;
        }
        else if (selector.startsWith('closest ')) {
            const cssSelector = selector.slice(8).trim();
            target = element.closest(cssSelector);
        }
        else if (selector.startsWith('next ')) {
            const cssSelector = selector.slice(5).trim();
            target = element.nextElementSibling;
            while (target && !target.matches(cssSelector)) {
                target = target.nextElementSibling;
            }
        }
        else if (selector.startsWith('previous ')) {
            const cssSelector = selector.slice(9).trim();
            target = element.previousElementSibling;
            while (target && !target.matches(cssSelector)) {
                target = target.previousElementSibling;
            }
        }
        else if (selector.startsWith('find ')) {
            const cssSelector = selector.slice(5).trim();
            target = element.querySelector(cssSelector);
        }
        else {
            target = document.querySelector(selector);
        }
        if (target) {
            targets.push(target);
        }
    };

    // Split multiple selectors if comma-separated
    const selectors = value.match("/(\[[^\]]+\])|(\S+)/g");
    if (selectors) {
        selectors.forEach(selector => {
            // Remove square brackets and split by comma if necessary
            selector = selector.replace(/[\[\]]/g, '');
            const subSelectors = selector.split(',').map(s => s.trim());
            subSelectors.forEach(subSelector => addTarget(subSelector));
        });
    }

    return targets;
};
export default parseTargetValues;
