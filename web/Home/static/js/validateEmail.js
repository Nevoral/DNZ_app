function emailValidate(email, finalValidation = false) {
    if (!email && finalValidation) {
        updateValidationUI('emailMsgBox', 'emailShadow', 'E-mail je povinný.', false);
        return;
    }
    if (!/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(email.toLowerCase())) {
        if (finalValidation) updateValidationUI('emailMsgBox', 'emailShadow', 'E-mailová addresa není správná.', false);
    } else {
        updateValidationUI('emailMsgBox', 'emailShadow', '', true);
    }
}