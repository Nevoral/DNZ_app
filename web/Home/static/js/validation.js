function addValidationListeners(element, validateFunc) {
    if (!element.hasAttribute('data-listeners-set')) {
        element.addEventListener('input', () => validateFunc(element.value));
        element.addEventListener('blur', () => validateFunc(element.value, true));
        element.setAttribute('data-listeners-set', 'true');
    }
}

function updateValidationUI(msgBoxId, shadowId, message, isValid) {
    let btnId = 'registerBtn';
    const inputTag = ["username", "email", "password"];
    if (!document.getElementById(btnId)) {
        btnId = 'loginBtn';
        inputTag.shift();
    }
    const button = document.getElementById(btnId);
    const msgBox = document.getElementById(msgBoxId);
    const underline = document.getElementById(shadowId);
    msgBox.textContent = message;
    msgBox.classList.remove('bg-red-500', 'bg-opacity-0', 'bg-opacity-75');
    underline.classList.remove('shadow-green-500', 'shadow-red-500');

    if (isValid) {
        msgBox.classList.add('bg-opacity-0');
        underline.classList.add('shadow-green-500');
        if (inputTag.filter(tag => document.getElementById(`${tag}MsgBox`).textContent !== '').length === 0 && document.getElementById(`terms`).checked) {
            button.removeAttribute('disabled');
            button.classList.add('hover:scale-105');
        }
    } else {
        msgBox.classList.add('bg-red-500', 'bg-opacity-75');
        underline.classList.add('shadow-red-500');
        if (!button.hasAttribute('disabled')) {
            button.setAttribute('disabled', 'disabled');
            button.classList.remove('hover:scale-105');
        }
    }
}