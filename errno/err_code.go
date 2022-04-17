package errno

import "entry_task/model"

var CODE_SUCCESS = 0
var (
	ERR_INTERNAL              = model.Payload{Code: 10001, Msg: "ERR_INTERNAL"}
	ERR_PAGESIZE              = model.Payload{Code: 10002, Msg: "ERR_PAGESIZE"}
	ERR_PAGENUM               = model.Payload{Code: 10003, Msg: "ERR_PAGENUM"}
	ERR_INVALID_PARAM         = model.Payload{Code: 10004, Msg: "ERR_INVALID_PARAM"}
	ERR_USER_TYPE_PERMISSION  = model.Payload{Code: 10005, Msg: "ERR_USER_TYPE_PERMISSION"}
	ERR_MUST_USER_TYPE_SELLER = model.Payload{Code: 10006, Msg: "ERR_MUST_USER_TYPE_SELLER"}
	ERR_PARAM_PRODUCT_ID      = model.Payload{Code: 10007, Msg: "ERR_INVALID_PRODUCT_ID"}
	ERR_PARAM_SHOP_ID         = model.Payload{Code: 10008, Msg: "ERR_INVALID_SHOP_ID"}

	ERR_NO_LOGIN        = model.Payload{Code: 10101, Msg: "ERR_NO_LOGIN"}
	ERR_USER_NOT_EXIST  = model.Payload{Code: 10102, Msg: "ERR_USER_NOT_EXIST"}
	ERR_PASSWORD        = model.Payload{Code: 10103, Msg: "ERR_PASSWORD"}
	ERR_PARAM_USER_NAME = model.Payload{Code: 10104, Msg: "ERR_PARAM_USER_NAME"}
	ERR_PARAM_PASSWORD  = model.Payload{Code: 10105, Msg: "ERR_PARAM_PASSWORD"}
	ERR_PARAM_USER_TYPE = model.Payload{Code: 10106, Msg: "ERR_USER_TYPE"}
	//ERR_NO_PERMISSION   = model.Payload{Code: 10107, Msg: "ERR_NO_PERMISSION"}
	ERR_USERNAME_EXIST  = model.Payload{Code: 10107, Msg: "ERR_USERNAME_EXIST"}

	ERR_SHOP_NOT_EXIST  = model.Payload{Code: 10201, Msg: "ERR_SHOP_NOT_EXIST"}
	ERR_SHOP_ID_ILLEGAL = model.Payload{Code: 10202, Msg: "ERR_SHOP_ID_ILLEGAL"}
	ERR_SHOP_NAME_EXIST = model.Payload{Code: 10203, Msg: "ERR_SHOP_NAME_EXIST"}
//	ERR_SHOP_ID_LEN     = model.Payload{Code: 10204, Msg: "ERR_SHOP_ID_LEN"}
	ERR_PARAM_SHOP_NAME = model.Payload{Code: 10204, Msg: "ERR_PARAM_SHOP_NAME"}
	ERR_PARAM_SHOP_DESC = model.Payload{Code: 10206, Msg: "ERR_PARAM_SHOP_DESC"}
	ERR_NOT_SHOP_OWNER  = model.Payload{Code: 10207, Msg: "ERR_NOT_SHOP_OWNER"}

	ERR_PRODUCT_NO_EXIST     = model.Payload{Code: 10501, Msg: "ERR_PRODUCT_NO_EXIST"}
	ERR_PARAM_PRODUCT_TITLE  = model.Payload{Code: 10502, Msg: "ERR_PARAM_PRODUCT_TITLE"}
	ERR_BRAND_ID_NO_EXIST    = model.Payload{Code: 10503, Msg: "ERR_BRAND_ID_NO_EXIST"}
	ERR_PRICE_OUT_RANGE      = model.Payload{Code: 10504, Msg: "ERR_PRICE_OUT_RANGE"}
	ERR_STOCK_OUT_RANGE      = model.Payload{Code: 10505, Msg: "ERR_STOCK_OUT_RANGE"}
	ERR_CATEGORY_ID_NO_EXIST = model.Payload{Code: 10506, Msg: "ERR_CATEGORY_ID_NO_EXIST"}
	//ERR_CATEGORY_NO_EXIST    = model.Payload{Code: 10506, Msg: "ERR_CATEGORY_NO_EXIST"}
	ERR_COVERUIR_LEN         = model.Payload{Code: 10507, Msg: "ERR_COVERUIR_LEN"}
	ERR_SHOW_URI_NUM         = model.Payload{Code: 10508, Msg: "ERR_SHOW_URI_NUM"}
	ERR_SHOW_URI_LEN         = model.Payload{Code: 10509, Msg: "ERR_SHOW_URI_LEN"}
	ERR_SEARCH_KEY_LEN       = model.Payload{Code: 10510, Msg: "ERR_SEARCH_KEY_LEN"}
	ERR_NOT_PRODUCT_OWNER    = model.Payload{Code: 10511, Msg: "ERR_NOT_PRODUCT_OWNER"}
	ERR_ATTR_NAME            = model.Payload{Code: 10512, Msg: "ERR_ATTR_NAME"}
	ERR_PRODUCT_STATUS       = model.Payload{Code: 10513, Msg: "ERR_PRODUCT_STATUS"}
)

func OK(data interface{}) model.Payload {
	return model.Payload{
		Code: 0,
		Msg:  "ok",
		Data: data,
	}
}
