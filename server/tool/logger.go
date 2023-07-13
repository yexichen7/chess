/**
* @Author: yexichen
* @Date:2023/7/11-19:23
* @Desc
**/

package tool

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger() {
	Logger, _ = zap.NewProduction()
}
