package queue

import (
	"hpcfusion/service/bash"
	"strings"
)

//IsInstall check PBSPro installation
func IsInstall() bool {
	if _, err := bash.Run("type qstat"); err != nil {
		return false
	}
	return true
}

//AllowHistory check history function
func AllowHistory() bool {
	b, _ := bash.Run("qmgr -c 'l s'")
	s := b.String()
	i := strings.Index(s, "job_history_enable")
	if i == -1 {
		return false
	}
	r := strings.Fields(s[i:])
	if b := r[2]; b == "True" {
		return true
	}
	return false
}
