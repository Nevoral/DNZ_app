function validateUsername(username, finalValidation = false) {
    if (!username && finalValidation) {
        updateValidationUI('usernameMsgBox', 'usernameShadow', 'Uživatelské jméno je povinné.', false);
        return;
    }
    if (username.length < 3) {
        if (finalValidation) updateValidationUI('usernameMsgBox', 'usernameShadow', 'Uživatelské jméno je krátké min. 3 znaky.', false);
    } else {
        updateValidationUI('usernameMsgBox', 'usernameShadow', '', true);
    }
}

