package common

const (
	HandlerSuccess		= 300000
	HandlerReadBodyErr	= 300010
	HandlerReadPathErr	= 300011

	HandlerDBInsertErr	= 300022
	HandlerDBSelectErr	= 300023
	HandlerDBUpdateErr	= 300024
	HandlerDBDeleteErr	= 300025

	HandlerPasswordErr	= 300036

	HandlerReadBodyErrMsg = "read body params err, please check params"
	HandlerReadPathErrMsg = "read path params err, please check params"

)
