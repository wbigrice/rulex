// Copyright (C) 2023 wwhai
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package apis

import (
	"strings"

	"github.com/gin-gonic/gin"
	archsupport "github.com/hootrhino/rulex/bspsupport"
	common "github.com/hootrhino/rulex/plugin/http_server/common"
	"github.com/hootrhino/rulex/typex"
)

/*
*
* 信号强度
*
 */
func Get4GCSQ(c *gin.Context, ruleEngine typex.RuleX) {
	c.JSON(common.HTTP_OK, common.OkWithData(archsupport.RhinoPiGet4GCSQ()))
}

// (1,"CHINA MOBILE","CMCC","46000",0),
// (3,"CHN-UNICOM","UNICOM","46001",7),
// +COPS: 0,0,\"CHINA MOBILE\",7
func Get4GCOPS(c *gin.Context, ruleEngine typex.RuleX) {
	result, err := archsupport.RhinoPiGetCOPS()
	if err != nil {
		c.JSON(common.HTTP_OK, common.Error400(err))
	} else {
		cm := "UNKNOWN"
		if strings.Contains(result, "CMCC") {
			cm = "中国移动"
		}
		if strings.Contains(result, "UNICOM") {
			cm = "中国联通"
		}
		c.JSON(common.HTTP_OK, common.OkWithData(cm))
	}
}

/*
*
* 获取电话卡ICCID
*
 */
func Get4GICCID(c *gin.Context, ruleEngine typex.RuleX) {
	result, err := archsupport.RhinoPiGetCOPS()
	if err != nil {
		c.JSON(common.HTTP_OK, common.Error400(err))
	} else {
		c.JSON(common.HTTP_OK, common.OkWithData(result))
	}
}
