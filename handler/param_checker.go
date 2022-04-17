package handler

import (
	"entry_task/errno"
	"entry_task/model"
	"entry_task/pkg/data_manager"
)

const NameLenMin = 2
const NameLenMax = 20

const PasswordLenMin = 2
const PassWordLenMax = 20

const EmailLenMin = 2
const EmailLenMax = 20

const StrorUriMax = 20

const DefaultPageSize = 20

var productStatus = map[model.PStatus]bool{
	model.ProudctStatusNormal:   true,
	model.ProudctStatusAuditing: true,
	model.ProudctStatusDeleted:  true,
	model.ProudctStatusEdit:     true,
	model.ProudctStatusForbiden: true,
	model.ProudctStatusOff:      true,
}

func isUserNameValid(name string) bool {
	if len(name) < NameLenMin || len(name) > NameLenMax {
		return false
	}
	return true
}

func isPasswordValid(p string) bool {
	if len(p) < PasswordLenMin || len(p) > PassWordLenMax {
		return false
	}
	return true
}

func isUserTypeValid(userType model.UserType) bool {
	if !(userType == model.UserTypeSeller || userType == model.UserTypeCustomer) {
		return false
	}
	return true
}

func isEmailValid(email string) bool {
	if len(email) < PasswordLenMin || len(email) > PassWordLenMax {
		return false
	}
	return true
}

func isPicUriValid(uri string) bool {
	if len(uri) > StrorUriMax {
		return false
	}
	return true
}

func isShopNameValid(name string) bool {
	return true
}
func isIntroductionValid(introduction string) bool {
	return true
}

func isShopIDValid(introduction string) bool {
	if len(introduction) == 0 || len(introduction) > 36 {
		return false
	}
	return true
}

func isProductIDValid(introduction string) bool {
	return true
}

func isAttrValueValid(value string) bool {
	return true
}
func isProductTitleValid(title string) bool {
	return true
}
func isProductPriceValid(price uint32) bool {
	return true
}

func isProductStockValid(price uint32) bool {
	return true
}
func isProductStatusValid(status model.PStatus) bool {
	if _, ok := productStatus[status]; ok {
		return true
	}
	return false
}

func isAttrValid(atts []model.AttrInfo) bool {
	attrNames := make(map[string]bool)
	for _, attr := range atts {
		if _, ok := attrNames[attr.Name]; ok {
			return false
		}
		_, err := data_manager.GetAttr(attr.Name)
		if err != nil {
			return false
		}
		attrNames[attr.Name] = true
	}
	return true
}

func checkProduct(product *model.Product) model.Payload {
	if !isProductTitleValid(product.Title) {
		return errno.ERR_PARAM_PRODUCT_TITLE
	}
	if !isProductPriceValid(product.Price) {
		return errno.ERR_PARAM_PRODUCT_TITLE
	}
	if !isProductStockValid(product.Stock) {
		return errno.ERR_PARAM_PRODUCT_TITLE
	}
	if product.BrandID != 0 {
		_, err := data_manager.GetBrand(product.BrandID)
		if err != nil {
			return errno.ERR_BRAND_ID_NO_EXIST
		}
	}
	if product.CategoryID != 0 {
		_, err := data_manager.GetCategory(product.CategoryID)
		if err != nil {
			return errno.ERR_CATEGORY_ID_NO_EXIST
		}
	}
	return errno.OK(nil)
}
