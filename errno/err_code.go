package errno

type ErrCode int

/*
const   (
    SUCCESS ErrCode = iota
    ERR_INTERNAL = 10001
    ERR_NO_LOGIN  = 10004
    ERR_USER_NOT_EXIST = 10101
    ERR_PASSWORD = 10102
    ERR_USER_NAME_LEN = 10103
    ERR_PASSWORD_LEN = 10104
    ERR_USER_TYPE = 10105


    ERR_PRODUCT_NO_EXIST = 10501
    ERR_PRODUCT_TITLE_LEN = 10502
    ERR_BRAND_NO_EXIST = 10503
    ERR_PRICE_OUT_RANGE   = 10504
    ERR_STOCK_OUT_RANGE = 10505
    ERR_CATEGORY_NO_EXIST = 10506
    ERR_COVERUIR_LEN = 10507
    ERR_SHOW_URI_NUM = 10508
    ERR_SHOW_URI_LEN = 10509
    ERR_SEARCH_KEY_LEN = 10510
)*/
var CODE_SUCCESS = 0
var (
	ERR_INTERNAL             = Payload{Code: 10001, Msg: "ERR_INTERNAL"}
	ERR_PAGESIZE             = Payload{Code: 10002, Msg: "ERR_PAGESIZE"}
	ERR_PAGENUM              = Payload{Code: 10003, Msg: "ERR_PAGENUM"}
	ERR_INVALID_PARAM        = Payload{Code: 10004, Msg: "ERR_INVALID_PARAM"}
	ERR_USER_TYPE_PERMISSION = Payload{Code: 10005, Msg: "ERR_USER_TYPE_PERMISSION"}
	ERR_MUST_USER_TYPE_SELLER = Payload{Code: 10005, Msg: "ERR_MUST_USER_TYPE_SELLER"}

	ERR_NO_LOGIN       = Payload{Code: 10101, Msg: "ERR_NO_LOGIN"}
	ERR_USER_NOT_EXIST = Payload{Code: 10102, Msg: "ERR_USER_NOT_EXIST"}
	ERR_PASSWORD       = Payload{Code: 10103, Msg: "ERR_PASSWORD"}
	ERR_USER_NAME_LEN  = Payload{Code: 10104, Msg: "ERR_USER_NAME_LEN"}
	ERR_PASSWORD_LEN   = Payload{Code: 10105, Msg: "ERR_PASSWORD_LEN"}
	ERR_USER_TYPE      = Payload{Code: 10106, Msg: "ERR_USER_TYPE"}
	ERR_NO_PERMISSION  = Payload{Code: 10107, Msg: "ERR_NO_PERMISSION"}
	ERR_USERNAME_EXIST = Payload{Code: 10108, Msg: "ERR_USERNAME_EXIST"}

	ERR_SHOP_NOT_EXIST  = Payload{Code: 10201, Msg: "ERR_SHOP_NOT_EXIST"}
	ERR_SHOP_ID_ILLEGAL = Payload{Code: 10202, Msg: "ERR_SHOP_ID_ILLEGAL"}
	ERR_SHOP_NAME_EXIST = Payload{Code: 10203, Msg: "ERR_SHOP_NAME_EXIST"}
	ERR_SHOP_ID_LEN     = Payload{Code: 10204, Msg: "ERR_SHOP_ID_LEN"}

	ERR_PRODUCT_NO_EXIST  = Payload{Code: 10501, Msg: "ERR_PRODUCT_NO_EXIST"}
	ERR_PRODUCT_TITLE_LEN = Payload{Code: 10502, Msg: "ERR_PRODUCT_TITLE_LEN"}
	ERR_BRAND_NO_EXIST    = Payload{Code: 10503, Msg: "ERR_BRAND_NO_EXIST"}
	ERR_PRICE_OUT_RANGE   = Payload{Code: 10504, Msg: "ERR_PRICE_OUT_RANGE"}
	ERR_STOCK_OUT_RANGE   = Payload{Code: 10505, Msg: "ERR_STOCK_OUT_RANGE"}
	ERR_CATEGORY_NO_EXIST = Payload{Code: 10506, Msg: "ERR_CATEGORY_NO_EXIST"}
	ERR_COVERUIR_LEN      = Payload{Code: 10507, Msg: "ERR_COVERUIR_LEN"}
	ERR_SHOW_URI_NUM      = Payload{Code: 10508, Msg: "ERR_SHOW_URI_NUM"}
	ERR_SHOW_URI_LEN      = Payload{Code: 10509, Msg: "ERR_SHOW_URI_LEN"}
	ERR_SEARCH_KEY_LEN    = Payload{Code: 10510, Msg: "ERR_SEARCH_KEY_LEN"}
)

type Payload struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func OK(data interface{}) Payload {
	return Payload{
		Code: 0,
		Msg:  "ok",
		Data: data,
	}
}
