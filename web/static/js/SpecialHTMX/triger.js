import parseTime from "./parseTime";
import requestBuilder from "./ajaxBuilder";
import parseTargetValues from "./target";
import performSwap from "./swap";
// Function to parse hx-trigger attribute and attach appropriate event listeners
const parseTriggerAttribute = (element) => {
    if (element.hasAttribute('hx-trigger')) {
        const triggers = element.getAttribute('hx-trigger').split(',').map(t => t.trim());
        const triggerAttrs = [];
        triggers.forEach(trigger => {
            const [eventType, ...modifiers] = trigger.split(' ');
            const config = {
                once: false,
                changed: false,
                delay: 0,
                throttle: 0,
                from: null,
                target: null,
                consume: false,
                queue: 'last'
            };

            modifiers.forEach(mod => {
                const [key, value] = mod.split(':');
                switch (key) {
                    case 'once':
                        config.once = true;
                        break;
                    case 'changed':
                        config.changed = true;
                        break;
                    case 'delay':
                        config.delay = parseTime(value);
                        break;
                    case 'throttle':
                        config.throttle = parseTime(value);
                        break;
                    case 'from':
                        config.from = value;
                        break;
                    case 'target':
                        config.target = value;
                        break;
                    case 'consume':
                        config.consume = true;
                        break;
                    case 'queue':
                        config.queue = value;
                        break;
                }
            });
            triggerAttrs.push({eventType, config})
        });
        return triggerAttrs
    }
};

// Main handler function
const handleEvent = async (event) => {
    const element = event.currentTarget;

    // Parse and build the request
    const requestInfo = await requestBuilder(element);
    if (!requestInfo) return;

    const responseText = requestInfo;

    // Determine the target elements
    const targetSelector = element.getAttribute('hx-target');
    const targetElements = parseTargetValues(targetSelector, element);

    // Perform the swap on each target element
    targetElements.forEach(target => {
        performSwap(element, target, responseText);
    });
};

// Function to attach event listeners based on hx-trigger
const attachTriggers = (element) => {
    const triggers = parseTriggerAttribute(element);
    triggers.forEach(({ eventType, config }) => {
        if (eventType.startsWith('every ')) {
            const interval = parseTime(eventType.slice(6));
            setInterval(() => handleEvent(new Event('poll')), interval);
        } else {
            const eventHandler = (event) => {
                if (config.once) {
                    event.target.removeEventListener(eventType, eventHandler);
                }

                if (config.consume) {
                    event.stopPropagation();
                }

                if (config.delay > 0) {
                    clearTimeout(element._hxTimeout);
                    element._hxTimeout = setTimeout(() => handleEvent(event), config.delay);
                } else if (config.throttle > 0) {
                    if (!element._hxLastEvent || (Date.now() - element._hxLastEvent) > config.throttle) {
                        element._hxLastEvent = Date.now();
                        handleEvent(event);
                    }
                } else {
                    handleEvent(event);
                }
            };
        }
    });
};