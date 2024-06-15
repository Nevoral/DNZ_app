((doc,wind) => {
    const ExecFunc = (async (filePath, func) => {
        const [funcName, arg] = func.split('(');
        const args = arg.split(',');
        try {
            const module = await import(filePath);
            if (typeof module[funcName] === 'function') {
                module[funcName](...args);
            } else {
                console.warn(`Handler function "${funcName}" is not defined in the module.`);
            }
        } catch (error) {
            console.error(`Failed to load module or execute function "${funcName}":`, error);
        }
    });

    const handleCustomEvent = async (event, element) => {
        const attrs = Array.from(element.attributes);
        for (const attr of attrs) {
            if (attr.name.startsWith(`hx-`)) {
                const typeOfEvent = attr.name.slice(3);
                const requestInfo = await requestBuilder(element);
                if (!requestInfo) return;

                const responseText = requestInfo;

                const targetSelector = element.getAttribute('hx-target');
                const targetElements = parseTargetValues(targetSelector, element);

                targetElements.forEach(target => {
                    performSwap(element, target, responseText);
                });
                if (typeOfEvent === eventType) {
                    const [filepath, func] = attr.value.split('#');
                    if (filepath && func) {
                        ExecFunc(filepath, func)
                    } else {
                        console.warn(`Invalid action or handler name for attribute "${attr.name}".`);
                    }
                }
            }
        }
    };

    const docEventListeners = (eventTypes) => {
        eventTypes.forEach(eventType => {
            doc.addEventListener(eventType, async (event) => {
                let clickedElement = event.target;
                let eventFound = false;
                while (clickedElement) {
                    if (clickedElement.hasAttributes() && clickedElement.hasAttribute(`hx-trigger`)) {
                        const triggers = parseTriggerAttribute(clickedElement);
                        triggers.forEach(({ eventTypeEl, config }) => {
                            if (eventTypeEl === eventType) {
                                eventFound = true;
                            }
                        });
                        handleEvent(clickedElement, Array.from(clickedElement.attributes), eventType);
                        if (clickedElement.hasAttribute(`hx-stoppropagation`)) {
                            if (clickedElement.getAttribute('hx-stoppropagation') === eventType) {
                                event.stopPropagation();
                                break;
                            }
                        }
                    }
                    if (eventFound) {
                        break;
                    }
                    clickedElement = clickedElement.parentElement;
                }
            });
        });
    };

    doc.addEventListener('DOMContentLoaded', () => {
        // handleElem("-document", "click");
        docEventListeners(["click"]);
    });
})(document, window);

// ((e, t)=>{
//     const n = "__q_context__";
//     const r = window;
//     const o = new Set;
//     const p = "qvisible";
//     const b = "_qwikjson_";
//     const u = t => e.querySelectorAll(t);
//     const y = e => e && "function" == typeof e.then;
//     const d = (e, t, n = t.type) => {
//         u("[on" + e + "\\:" + n + "]").forEach((r => q(r, e, t, n)))
//     };
//     const m = t => {
//         if (void 0 === t[b]) {
//             let n = (t === e.documentElement ? e.body : t).lastElementChild;
//             for (; n;) {
//                 if ("SCRIPT" === n.tagName && "qwik/json" === n.getAttribute("type")) {
//                     t[b] = JSON.parse(n.textContent.replace(/\\x3C(\/?script)/gi, "<$1"));
//                     break
//                 }
//                 n = n.previousElementSibling
//             }
//         }
//     };
//     const q = async (t, r, o, i = o.type) => {
//         const c = "on" + r + ":" + i;
//         t.hasAttribute("preventdefault:" + i) && o.preventDefault();
//         const p = t._qc_, b = p && p.li.filter((e => e[0] === c));
//         if (b && b.length > 0) {
//             for (const e of b) {
//                 const n = e[1].getFn([t, o], (() => t.isConnected))(o, t), r = o.cancelBubble;
//                 y(n) && await n, r && o.stopPropagation()
//             }
//             return
//         }
//         const u = t.getAttribute(c);
//         if (u) {
//             const r = t.closest("[q\\:container]"), i = r.getAttribute("q:base"), c = r.getAttribute("q:version") || "unknown",
//                 p = r.getAttribute("q:manifest-hash") || "dev", b = new URL(i, e.baseURI);
//             for (const l of u.split("\n")) {
//                 const u = new URL(l, b), d = u.href, h = u.hash.replace(/^#?([^?[|]*).*$/, "$1") || "default",
//                     q = performance.now();
//                 let v, g, E;
//                 const _ = l.startsWith("#"),
//                     k = {qBase: i, qManifest: p, qVersion: c, href: d, symbol: h, element: t, reqTime: q};
//                 if (_) v = (r.qFuncs || [])[Number.parseInt(h)], v || (g = "sync", E = Error("sync handler error for symbol: " + h)); else {
//                     const e = u.href.split("#")[0];
//                     try {
//                         const t = import(e);
//                         m(r), v = (await t)[h]
//                     } catch (e) {
//                         g = "async", E = e
//                     }
//                 }
//                 if (!v) {
//                     w("qerror", s({importError: g, error: E}, k));
//                     break
//                 }
//                 const C = e[n];
//                 if (t.isConnected) try {
//                     e[n] = [t, o, u], _ || w("qsymbol", s({}, k));
//                     const r = v(o, t);
//                     y(r) && await r
//                 } catch (e) {
//                     w("qerror", s({error: e}, k))
//                 } finally {
//                     e[n] = C
//                 }
//             }
//         }
//     };
//     const w = (t, n) => {
//         e.dispatchEvent(new CustomEvent(t, {detail: n}))
//     };
//     const g = async e => {
//         let t = e.type.replace(/([A-Z])/g, (e => "-" + e.toLowerCase())), n = e.target;
//         for (d("-document", e, t); n && n.getAttribute;) {
//             const r = q(n, "", e, t);
//             let o = e.cancelBubble;
//             y(r) && await r, o = o || e.cancelBubble || n.hasAttribute("stoppropagation:" + e.type), n = e.bubbles && !0 !== o ? n.parentElement : null
//         }
//     };
//     const E = e => {
//         d("-window", e, e.type.replace(/([A-Z])/g, (e => "-" + e.toLowerCase())))
//     };
//     const _ = () => {
//         var n;
//         const s = e.readyState;
//         if (!t && ("interactive" == s || "complete" == s) && (t = 1, w("qinit"), (null != (n = r.requestIdleCallback) ? n : r.setTimeout).bind(r)((() => w("qidle"))), o.has("qvisible"))) {
//             const e = u("[on\\:qvisible]"), t = new IntersectionObserver((e => {
//                 for (const n of e) n.isIntersecting && (t.unobserve(n.target), q(n.target, "", new CustomEvent("qvisible", {detail: n})))
//             }));
//             e.forEach((e => t.observe(e)))
//         }
//     };
//     const C = t => {
//         for (const n of t) o.has(n) || (e.addEventListener(n, g, {capture: !0, passive: !1}), r.addEventListener(n, E, {capture: !0, passive: !1}), o.add(n))
//     };
//     if(!(n in e)){e[n]=0;const t=r.qwikevents;Array.isArray(t)&&C(t),r.qwikevents={push:(...e)=>C(e)},e.addEventListener("readystatechange", _, {capture: !1, passive: !1}),_()}})(document)})()