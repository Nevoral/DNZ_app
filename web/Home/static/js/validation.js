function addValidationListeners(fieldId, validateFunc) {
    document.getElementById(fieldId).addEventListener('focus', function() {
        if (!this.hasAttribute('data-listeners-set')) {
            this.addEventListener('input', () => validateFunc(this.value));
            this.addEventListener('blur', () => validateFunc(this.value, true));
            this.setAttribute('data-listeners-set', 'true');
        }
    });
}

function updateValidationUI(msgBoxId, shadowId, message, isValid) {
    const button = document.getElementById('RegisterBTN');
    const inputTag = ["username", "email", "password"];
    const msgBox = document.getElementById(msgBoxId);
    const underline = document.getElementById(shadowId);
    msgBox.textContent = message;
    msgBox.classList.remove('bg-red-500', 'bg-opacity-0', 'bg-opacity-75');
    underline.classList.remove('shadow-green-500', 'shadow-red-500');

    if (isValid) {
        msgBox.classList.add('bg-opacity-0');
        underline.classList.add('shadow-green-500');
        if (inputTag.filter(tag => document.getElementById(`${tag}MsgBox`).textContent !== '').length === 0 ) {
            button.removeAttribute('disabled');
            button.classList.remove('disabled:hover:from-cyan-600/100', 'disabled:hover:to-green-600/100');
            button.classList.add('hover:from-cyan-600/100', 'hover:to-green-600/100');
        }
    } else {
        msgBox.classList.add('bg-red-500', 'bg-opacity-75');
        underline.classList.add('shadow-red-500');
        if (!button.hasAttribute('disabled')) {
            button.setAttribute('disabled', 'disabled');
            button.classList.add('disabled:hover:from-cyan-600/100', 'disabled:hover:to-green-600/100');
            button.classList.remove('hover:from-cyan-600/100', 'hover:to-green-600/100');
        }
    }
}

addValidationListeners('username', validateUsername);
addValidationListeners('email', validateEmail);
addValidationListeners('password', validatePassword);