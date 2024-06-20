function togglePasswordVisibility() {
    const passwordInput = document.getElementById('password');
    const eyeSlash = document.getElementById('eyeSlash');
    if (passwordInput.type === 'password') {
        eyeSlash.style.display = "none";
        passwordInput.type = 'text';
    } else {
        eyeSlash.style.display = "block";
        passwordInput.type = 'password';
    }
}

function passwordValidate(password, finalValidation = false) {
    const criteria = [
        { regex: /(?=.*\d)/, message: "0-9" },
        { regex: /(?=.*[A-Z])/, message: "A-Z" },
        { regex: /(?=.*[a-z])/, message: "a-z" },
        { regex: /(?=.*[#?!@$%^&*-])/, message: "@#$%" },
        { regex: /.{8,20}/, message: "min délku 8" }
    ];

    let failedCriteria = criteria.filter(crit => !crit.regex.test(password));
    if (!password && finalValidation) {
        updateValidationUI('passwordMsgBox', 'passwordShadow', 'Heslo je povinné.', false);
        return;
    }
    if (failedCriteria.length > 0) {
        if (finalValidation) updateValidationUI('passwordMsgBox', 'passwordShadow', 'Heslo nesplňuje kritéria bezpečnosti: ' + formatList(failedCriteria.map(crit => crit.message)) + '.', false);
    } else {
        updateValidationUI('passwordMsgBox', 'passwordShadow', '', true);
    }
}

function formatList(items) {
    return items.length > 1 ?
        items.slice(0, -1).join(", ") + " a " + items[items.length - 1] :
        items[0];
}