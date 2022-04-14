package handler

import (
    "entry_task/model"
)

const NameLenMax = 20
const NameLenMin = 2

const PassWordLenMax = 20
const PasswordLenMin = 2

const EmailLenMax = 20
const EmailLenMin = 2

func isUserNameValid(name string) bool {
    if len(name) < NameLenMin || len(name) > NameLenMax {
        return false
    }
    return  true
}

func isPasswordValid(p string) bool {
    if len(p) < PasswordLenMin || len(p) > PassWordLenMax{
        return false
    }
    return  true
}

func isUserTypeValid(userType model.UserType) bool{
    if !(userType == model.UserTypeSeller || userType == model.UserTypeCustomer) {
        return false
    }
    return  true
}

func isEmailValid(email string) bool{
    if len(email) < PasswordLenMin || len(p) > PassWordLenMax{
        return false
    }
    return  true
    return  true
}

func isPicUriValid(uri string) bool{
    return  true
}

func isShopNameValid(name string) bool{
    return  true
}
func isIntroductionValid(introduction string) bool{
    return  true
}

func isShopIDValid(introduction string) bool{
    return  true
}

func isProductIDValid(introduction string) bool{
    return  true
}

func isAttrValueValid(value string) bool{
    return  true
}
func isProductTitleValid(title string) bool{
    return  true
}
func isProductPriceValid(price int) bool{
    return  true
}

func isProductStockValid(price int) bool{
    return  true
}
