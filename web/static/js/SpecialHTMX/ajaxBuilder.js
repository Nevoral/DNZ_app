const parseAttributes = (element) => {
    const supportedAttributes = ['hx-get', 'hx-post', 'hx-put', 'hx-patch', 'hx-delete'];
    for (const attr of supportedAttributes) {
        if (element.hasAttribute(attr)) {
            return { method: attr.split('-')[1].toUpperCase(), url: element.getAttribute(attr) };
        }
    }
    return null;
};

// Function to collect parameters from specified elements
const collectParams = (element) => {
    const params = new FormData();
    const addParams = (el) => {
        if (el.tagName === 'FORM') {
            new FormData(el).forEach((value, key) => {
                params.append(key, value);
            });
        } else if (el.name && (el.value !== undefined)) {
            params.append(el.name, el.value);
        }
    };
    addParams(element);

    const form = element.closest('form');
    if (form) addParams(form);

    if (element.hasAttribute('hx-include')) {
        const includeSelector = element.getAttribute('hx-include');
        const includedElements = document.querySelectorAll(includeSelector);
        includedElements.forEach(addParams);
    }

    return params;
};

// Function to filter parameters based on hx-params attribute
const filterParams = (element, params) => {
    if (element.hasAttribute('hx-params')) {
        const hxParams = element.getAttribute('hx-params');
        const filteredParams = new URLSearchParams();
        if (hxParams === 'none') {
            return filteredParams; // Include no parameters
        } else if (hxParams.startsWith('not ')) {
            const excludeList = hxParams.slice(4).split(',').map(p => p.trim());
            params.forEach((value, key) => {
                if (!excludeList.includes(key)) {
                    filteredParams.append(key, value);
                }
            });
        } else {
            const includeList = hxParams.split(',').map(p => p.trim());
            includeList.forEach(key => {
                if (params.has(key)) {
                    filteredParams.append(key, params.get(key));
                }
            });
        }
        return filteredParams;
    }
    return new URLSearchParams(params); // Default to include all parameters
};

const requestBuilder = async (element) => {
    const requestInfo = parseAttributes(element);
    if (requestInfo) {
        const {method, url} = requestInfo;
        const config = {
            method: method,
            headers: {
                'Content-Type': 'application/json'
            },
        };
        if (method !== 'GET') {
            let params = collectParams(element);
            config.body = filterParams(element, params);
        }
        try {
            const response = await fetch(url, config);
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return await response.json();
        } catch (error) {
            console.error('Fetch error:', error);
            throw error;
        }
    }
};

export default requestBuilder;