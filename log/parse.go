package log

import (
	"bytes"
	"strings"
	"time"

	"gopkg.in/alog.v1/utils"
)

// GetTimeData 获取时间数据
func GetTimeData(data time.Time) map[string]interface{} {
	t := data.Format("2006-01-02 15:04:05.000")
	ts := strings.Split(t, " ")
	tDates := strings.Split(ts[0], "-")
	tTime := strings.Split(ts[1], ".")
	tTimes := strings.Split(tTime[0], ":")
	tData := map[string]interface{}{
		"Year":        tDates[0],
		"Month":       tDates[1],
		"Day":         tDates[2],
		"Hour":        tTimes[0],
		"Minute":      tTimes[1],
		"Second":      tTimes[2],
		"MilliSecond": tTime[1],
	}
	return tData
}

// ParseTime 解析时间模板
func ParseTime(tmpl interface{}, data time.Time) string {
	tData := GetTimeData(data)
	buf, err := utils.NewParseTmpl(tmpl).Parse(tData)
	if err != nil {
		return ""
	}
	return buf.String()
}

// ParseFileName 解析日志文件名模板
func ParseFileName(tmpl interface{}, data *LogItem) string {
	fData := GetTimeData((*data).Time)
	fData["Level"] = (*data).Level.ToString()
	fData["Tag"] = (*data).Tag
	buf, err := utils.NewParseTmpl(tmpl).Parse(fData)
	if err != nil {
		return ""
	}
	return buf.String()
}

// ParseLogItemWithBuffer 将日志项模板解析到缓冲区
func ParseLogItemToBuffer(tmpl interface{}, timeTmpl interface{}, data *LogItem) *bytes.Buffer {
	tData := GetTimeData((*data).Time)
	timeBuf, err := utils.NewParseTmpl(timeTmpl).Parse(tData)
	if err != nil {
		return bytes.NewBufferString("")
	}
	lData := map[string]interface{}{
		"ID":            (*data).ID,
		"Time":          timeBuf.String(),
		"Level":         (*data).Level.ToString(),
		"Tag":           (*data).Tag,
		"Message":       (*data).Message,
		"FileName":      (*data).File.Name,
		"ShortFileName": (*data).File.ShortName,
		"FileFuncName":  (*data).File.FuncName,
		"FileLine":      (*data).File.Line,
	}
	buf, err := utils.NewParseTmpl(tmpl).Parse(lData)
	if err != nil {
		return bytes.NewBufferString("")
	}
	return buf
}

// ParseLogItem 解析日志项模板
func ParseLogItem(tmpl interface{}, timeTmpl interface{}, data *LogItem) string {
	buf := ParseLogItemToBuffer(tmpl, timeTmpl, data)
	buf.WriteByte('\n')
	return buf.String()
}
