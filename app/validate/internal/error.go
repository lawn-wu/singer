package common

import (
	"singer/global"
	"strings"
)

// GetError
// @Description: 获取验证失败信息
// @param msg
// @param err
// @return returnMsg
func GetError(msg map[string]string, err string) (returnMsg string) {
	returnMsg = "Parameter error, non-compliance with rules"
	if global.Config.System.Debug {
		returnMsg = err
	}
	// 这里的err字符串变量信息可能为："Key: 'SendSMS.Mobile' Error:Field validation for 'Mobile' failed on the 'checkMobile' tag\nKey: 'SendSMS.Type' Error:Field validation for 'Type' failed on the 'required' tag"
	// 选择第一个验证未通过的字段信息
	firstNotAllowMsg := strings.Split(err, "\n")[0]
	// 字段值在字符串的开始标识和结束标识
	fieldStartTag, fieldEndTag := "Error:Field validation for '", "' failed on the"
	// 字段值开始位置
	fieldStart := strings.Index(firstNotAllowMsg, fieldStartTag)
	// 字段值结束位置
	fieldEnd := strings.Index(firstNotAllowMsg, fieldEndTag)
	if fieldStart < 0 || fieldEnd < 0 {
		return
	}
	// 获取验证未通过的字段值
	field := firstNotAllowMsg[fieldStart+len(fieldStartTag) : fieldEnd]
	// 规则值在字符串中的开始标识和结束标识
	ruleStartTag, ruleEndTag := "failed on the '", "' tag"
	// 规则值开始位置
	ruleStart := strings.Index(firstNotAllowMsg, ruleStartTag)
	// 规则值结束位置
	ruleEnd := strings.Index(firstNotAllowMsg, ruleEndTag)
	if ruleStart < 0 || ruleEnd < 0 {
		return
	}
	// 获取验证未通过的规则
	rule := firstNotAllowMsg[ruleStart+len(ruleStartTag) : ruleEnd]
	if _, ok := msg[field+"."+rule]; ok {
		returnMsg = msg[field+"."+rule]
	}
	return
}
