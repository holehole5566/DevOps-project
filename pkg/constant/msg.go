package constant

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	SERVER_ERROR:          "fail",
	INVALID_PARAMS:        "error parameters",
	ERROR_GET_ARTICLE_FAIL:   "fail to get the tour",
	// 100xx tour
	ERROR_GET_ARTICLE_ID_NOT_NUM:         "param id is not a number",
	ERROR_GET_ARTICLE_NO_RECORD:          "tour not found",
	ERROR_ADD_ARTICLE_FORMAT_INCORRECT:   "collects id null",
	ERROR_ADD_ARTICLE_NO_COLLECTS_RECORD: "collects record not found",


}

var SuccessMsg = "Success"

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[SERVER_ERROR]
}