package services

import (
	"github.com/mnhkahn/togo/dmltogo"
	"testing"
)

func TestDexToBin(t *testing.T) {
	res, _ := dmltogo.DmlToGo("CREATE TABLE `t_bulletin` (\n  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',\n  `bulletin_code` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '公告编码',\n  `title` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',\n  `bulletin_type` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '公告类型',\n  `detail` text COLLATE utf8mb4_general_ci COMMENT '公告详情',\n  `is_top` tinyint NOT NULL DEFAULT '0' COMMENT '0启用 1 删除',\n  `is_pop_ups` tinyint NOT NULL DEFAULT '0' COMMENT '0启用 1 删除',\n  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '0启用 1 删除',\n  `start_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '开始时间',\n  `end_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '结束时间',\n  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',\n  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',\n  `creator` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',\n  PRIMARY KEY (`id`),\n  KEY `idx_bulletin_code` (`bulletin_code`)\n) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='公告表'")
	t.Log(string(res))
}
