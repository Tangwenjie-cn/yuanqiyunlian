
CREATE TABLE IF NOT EXISTS `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL COMMENT '账号',
  `password` varchar(64) DEFAULT NULL COMMENT '密码',
  `login_ip` varchar(20) DEFAULT NULL COMMENT '最后登录ip',
  `ltime` datetime DEFAULT NULL COMMENT '最后一次登录时间',
  `group_id` tinyint(4) NOT NULL COMMENT '权限分组id',
  `status` tinyint(4) DEFAULT '1' COMMENT '状态1:正常-1:禁用',
  `pid` int(11) DEFAULT '0' COMMENT '上级id',
  `ctime` datetime NOT NULL COMMENT '创建时间',
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username.Unique` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='管理员表';

INSERT INTO `admin` (`id`, `username`, `password`, `login_ip`, `ltime`, `group_id`, `status`, `pid`, `ctime`, `utime`) VALUES
	(1, 'admin', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', '127.0.0.1', '2025-12-25 18:54:44', 1, 1, 0, '2025-04-30 12:00:00', '2025-12-25 18:54:44');

CREATE TABLE IF NOT EXISTS `admin_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(150) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '账号',
  `body` text COLLATE utf8mb4_bin,
  `uri` varchar(500) COLLATE utf8mb4_bin DEFAULT NULL,
  `ip` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL,
  `ctime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `account` (`account`),
  KEY `ctime` (`ctime`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='管理员操作日志';

CREATE TABLE IF NOT EXISTS `apply_super` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '-1拒绝0待审核1已通过',
  `remark` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='申请成为超级推荐人';

CREATE TABLE IF NOT EXISTS `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(150) DEFAULT NULL COMMENT '标题',
  `thumb` varchar(255) DEFAULT NULL COMMENT '图片',
  `sid` int(11) DEFAULT '0' COMMENT '分类ID',
  `content` text COMMENT '内容',
  `status` tinyint(4) DEFAULT '1' COMMENT '状态1:正常-1:禁用',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1001 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='文章';

INSERT INTO `article` (`id`, `title`, `thumb`, `sid`, `content`, `status`, `ctime`, `utime`) VALUES
	(1, '联系我们', '', 1, '<p><br></p>', 1, '2025-11-01 15:57:56', '2025-12-24 14:46:23'),
	(2, '关于钱包', '', 1, '<ol><li>余额可用于购买消费也可提现</li><li>余额主要来源于推广用户消费奖励</li></ol>', 1, '2025-12-24 15:09:04', '2025-12-24 15:09:04'),
	(3, '提现说明', '', 1, '<ol><li>单次可提现10-200元</li><li>一周可提现一次</li><li>请务必同意订阅消息，可第一时间知道提现结果</li><li>提现手续费为提现金额的1%</li></ol>', 1, '2025-12-24 15:17:05', '2025-12-24 15:18:22'),
	(4, '推广说明', '', 1, '<ol><li>通过本页生成的海报图或者商品详情页生成的海报图分享给好友，好友通过图片二维码进入程序注册登录即可绑定推广关系</li><li>推广的用户在本程序内消费可获得消费金额的1%的推广奖励</li></ol>', 1, '2025-12-24 15:27:48', '2025-12-24 15:27:48'),
	(5, '超级推广人说明', '', 1, '<p><span style="color: rgb(225, 60, 57);"><strong>要求：</strong></span></p><ol><li>推广人数达50人</li><li>你至少有过1笔消费</li><li>提交前请确认是否达标，如管理员判断有恶意提交行为将封禁账号</li></ol><p><span style="color: rgb(225, 60, 57);"><strong>奖励：</strong></span></p><p>推广的用户消费，除了原有的推广奖励，还有额外的2%的奖励金</p>', 1, '2025-12-24 15:36:31', '2025-12-24 15:39:10');

CREATE TABLE IF NOT EXISTS `article_sort` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL COMMENT '名称',
  `thumb` varchar(150) DEFAULT NULL COMMENT '缩略图',
  `status` tinyint(4) DEFAULT '1' COMMENT '状态1:正常-1:禁用',
  `sort` int(11) DEFAULT NULL COMMENT '排序',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='文章类别';

INSERT INTO `article_sort` (`id`, `name`, `thumb`, `status`, `sort`, `ctime`, `utime`) VALUES
	(1, '帮助', '', 1, 0, '2025-04-30 23:50:00', '2025-05-21 11:13:53'),
	(2, '公告', '', 1, 1, '2025-04-30 23:50:00', NULL);

CREATE TABLE IF NOT EXISTS `auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上级id',
  `name` varchar(150) NOT NULL COMMENT '权限名称',
  `url` varchar(150) DEFAULT NULL COMMENT 'url地址',
  `path` varchar(150) DEFAULT NULL COMMENT '前端路由',
  `icon` varchar(100) DEFAULT NULL COMMENT '图标icon',
  `is_menu` tinyint(4) NOT NULL DEFAULT '1' COMMENT '菜单1:是,0:否',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=185 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='权限表';

INSERT INTO `auth` (`id`, `pid`, `name`, `url`, `path`, `icon`, `is_menu`, `sort`, `ctime`, `utime`) VALUES
	(5, 0, '权限管理', '', '', 'Lock', 1, 3, '2025-04-30 23:50:00', '2025-10-17 11:34:24'),
	(8, 5, '用户组', '/admin/auth/grouplist', '/auth/group', '', 1, 2, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(13, 8, '添加用户组', '/admin/auth/addGroup', NULL, '', 0, 1, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(15, 8, '删除用户组', '/admin/auth/delGroup', NULL, '', 0, 2, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(16, 8, '编辑用户组', '/admin/auth/editGroup', NULL, '', 0, 3, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(17, 8, '设置用户组权限', '/admin/auth/setGroupAuth', '', '', 0, 4, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(18, 5, '管理员', '/admin/auth/adminList', '/auth/admin', '', 1, 3, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(19, 18, '添加管理员', '/admin/auth/addAdmin', NULL, '', 0, 1, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(20, 18, '编辑管理员', '/admin/auth/editAdmin', NULL, '', 0, 2, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(21, 18, '删除管理员', '/admin/auth/delAdmin', NULL, '', 0, 3, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(33, 8, '获取用户组权限信息', '/admin/auth/getGroupAuth', NULL, '', 0, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(34, 0, '系统设置', '', '', 'Setting', 1, 2, '2025-04-30 23:50:00', '2025-10-17 11:33:38'),
	(36, 0, '用户管理', '', '', 'User', 1, 6, '2025-04-30 23:50:00', '2025-10-17 11:34:42'),
	(52, 0, '文章管理', '', '', 'Document', 1, 4, '2025-04-30 23:50:00', '2025-10-17 11:34:33'),
	(54, 52, '文章列表', '/admin/article/list', '/article/article', '', 1, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(77, 0, '图片管理', '', '', 'Picture', 1, 5, '2025-04-30 23:50:00', '2025-12-25 20:08:04'),
	(78, 52, '文章分类', '/admin/article/sort', '/article/sort', '', 1, 1, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(83, 34, '参数设置', '/admin/system/getSetAllInfo', '/system/saveset', '', 1, 2, '2025-04-30 23:50:00', '2025-10-17 11:33:52'),
	(84, 77, '图片列表', '/admin/gallery/index', '/gallery/index', '', 1, 1, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(85, 77, '图片分类', '/admin/gallery/sort', '/gallery/sort', '', 1, 2, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(86, 36, '用户列表', '/admin/user/list', '/user/index', '', 1, 1, '2025-04-30 23:50:00', '2025-11-22 16:01:03'),
	(87, 0, '首页', '', '/home', 'House', 1, 1, '2025-04-30 23:50:00', '2025-11-05 16:43:37'),
	(97, 54, '添加', '/admin/article/save', '', '', 0, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(98, 54, '编辑', '/admin/article/update', '', '', 0, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(99, 54, '删除', '/admin/article/del', '', '', 0, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(100, 78, '添加', '/admin/article/saveSort', '', '', 0, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(101, 78, '编辑', '/admin/article/updateSort', '', '', 0, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(102, 78, '删除', '/admin/article/delSort', '', '', 0, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(103, 84, '添加', '/admin/gallery/save', '', '', 0, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(104, 84, '编辑', '/admin/gallery/update', '', '', 0, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(105, 84, '删除', '/admin/gallery/del', '', '', 0, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(106, 85, '添加', '/admin/gallery/saveSort', '', '', 0, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(107, 85, '编辑', '/admin/gallery/updateSort', '', '', 0, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(108, 85, '删除', '/admin/gallery/delSort', '', '', 0, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(109, 86, '获取用户信息', '/admin/user/info', '', '', 0, 1, '2025-04-30 23:50:00', '2025-12-23 14:31:24'),
	(110, 83, '保存设置', '/admin/system/saveSet', '', '', 0, 1, '2025-04-30 23:50:00', '2025-10-17 11:34:16'),
	(111, 0, '商品管理', '', '', 'Goods', 1, 7, '2025-04-30 23:50:00', '2025-10-17 11:34:46'),
	(112, 111, '商品列表', '/admin/goods/index', '/goods/index', '', 1, 1, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(113, 111, '商品分类', '/admin/goods/sort', '/goods/sort', '', 1, 2, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(114, 87, '获取首页数据', '/admin/index/home', '', '', 0, 1, '2025-05-16 16:54:50', '2025-05-16 16:54:50'),
	(115, 0, '订单管理', '', '', 'Memo', 1, 8, '2025-05-26 10:41:10', '2025-10-17 11:34:52'),
	(116, 0, '装修布局', '', '', 'Brush', 1, 9, '2025-07-10 10:24:21', '2025-10-20 15:14:54'),
	(117, 116, '首页装修', '/admin//theme/index', '/theme/index', '', 1, 1, '2025-07-10 10:27:51', '2025-12-23 14:46:33'),
	(118, 116, '个人中心', '', '/theme/user', '', 1, 4, '2025-09-17 13:54:44', '2025-10-29 15:30:32'),
	(119, 116, '分类装修', '', '/theme/sort', '', 1, 2, '2025-09-17 13:59:34', '2025-09-28 15:45:34'),
	(121, 116, '页面链接', '/admin/theme/pages/list', '/theme/pages', '', 1, 6, '2025-09-17 14:01:45', '2025-12-23 14:52:09'),
	(123, 116, '底部导航', '', '/theme/tabBar', '', 1, 5, '2025-09-28 15:44:42', '2025-10-29 13:11:59'),
	(124, 116, '会员中心', '', '/theme/vip', '', 1, 3, '2025-10-17 11:32:24', '2025-10-17 11:39:10'),
	(125, 0, '会员设置', '', '', 'Trophy', 1, 10, '2025-10-17 13:18:57', '2025-10-17 13:18:57'),
	(126, 125, '会员列表', '/admin/vip/list', '/vip/index', '', 1, 1, '2025-10-17 13:20:57', '2025-12-23 14:55:03'),
	(127, 125, '会员特权', '/admin/vip/privilegeList', '/vip/privilege', '', 1, 2, '2025-10-17 13:21:34', '2025-12-23 14:56:50'),
	(128, 0, '卡密管理', '', '', 'Key', 1, 11, '2025-10-20 15:15:51', '2025-10-20 15:15:51'),
	(129, 128, '卡密列表', '/admin/kami/list', '/kami/index', '', 1, 1, '2025-10-20 15:18:05', '2025-12-23 14:59:22'),
	(130, 111, '商品栏目', '', '/goods/column', '', 1, 3, '2025-10-24 16:09:54', '2025-10-24 16:10:20'),
	(131, 116, '推广装修', '', '/theme/promotion', '', 1, 7, '2025-11-12 14:07:09', '2025-11-12 14:14:22'),
	(132, 36, '提现列表', '/admin/cash/list', '/user/cash', '', 1, 2, '2025-11-18 16:38:08', '2025-11-22 17:24:12'),
	(133, 36, '高级推广申请', '/admin/user/applySuper', '/user/tg', '', 1, 3, '2025-11-18 16:39:25', '2025-11-22 17:23:23'),
	(134, 115, '订单列表', '/admin/order/list', '/order/index', '', 1, 1, '2025-11-18 16:43:19', '2025-12-23 14:45:39'),
	(137, 86, '设置用户', '/admin/user/setUser', '', '', 0, 2, '2025-12-23 14:32:39', '2025-12-25 19:28:10'),
	(138, 86, '获取vip', '/admin/user/getUserVip', '', '', 0, 3, '2025-12-23 14:33:17', '2025-12-25 19:28:03'),
	(139, 86, '设置vip', '/admin/user/setUserVip', '', '', 0, 4, '2025-12-23 14:33:42', '2025-12-25 19:27:55'),
	(140, 133, '高级推广审核', '/admin/user/setApplySuper', '', '', 0, 1, '2025-12-23 14:34:47', '2025-12-25 19:28:49'),
	(141, 132, '提现审核', '/admin/cash/check', '', '', 0, 1, '2025-12-23 14:36:31', '2025-12-23 14:36:31'),
	(142, 112, '添加商品', '/admin/goods/save', '', '', 0, 1, '2025-12-23 14:37:35', '2025-12-23 14:37:35'),
	(143, 112, '获取商品信息', '/admin/goods/getGoods', '', '', 0, 2, '2025-12-23 14:38:13', '2025-12-23 14:38:13'),
	(144, 112, '更新商品', '/admin/goods/update', '', '', 0, 3, '2025-12-23 14:38:37', '2025-12-23 14:38:37'),
	(145, 112, '删除商品', '/admin/goods/del', '', '', 0, 4, '2025-12-23 14:39:04', '2025-12-23 14:39:04'),
	(146, 112, '自动导入文档', '/admin/goods/doc/upload', '', '', 0, 5, '2025-12-23 14:40:14', '2025-12-23 14:40:14'),
	(147, 113, '添加分类', '/admin/goods/addSort', '', '', 0, 1, '2025-12-23 14:40:49', '2025-12-23 14:40:49'),
	(148, 113, '更新分类', '/admin/goods/upSort', '', '', 0, 2, '2025-12-23 14:41:27', '2025-12-23 14:42:27'),
	(149, 113, '删除分类', '/admin/goods/delSort', '', '', 0, 3, '2025-12-23 14:41:56', '2025-12-23 14:41:56'),
	(150, 130, '添加栏目', '/admin/goods/column/save', '', '', 0, 1, '2025-12-23 14:42:57', '2025-12-23 14:44:36'),
	(151, 130, '更新栏目', '/admin/goods/column/edit', '', '', 0, 2, '2025-12-23 14:43:23', '2025-12-23 14:44:25'),
	(152, 130, '删除栏目', '/admin/goods/column/del', '', '', 0, 3, '2025-12-23 14:44:09', '2025-12-23 14:44:09'),
	(153, 117, '添加', '/admin/theme/index/save', '', '', 0, 1, '2025-12-23 14:47:50', '2025-12-23 14:47:50'),
	(154, 117, '装修更新', '/admin/theme/index/update', '', '', 0, 2, '2025-12-23 14:48:39', '2025-12-23 14:48:39'),
	(155, 117, '删除装修', '/admin/theme/index/del', '', '', 0, 3, '2025-12-23 14:49:10', '2025-12-23 14:49:27'),
	(156, 117, '设置装修', '/admin/theme/index/set', '', '', 0, 4, '2025-12-23 14:50:19', '2025-12-23 14:50:19'),
	(157, 117, '获取装修', '/admin/theme/index/get', '', '', 0, 5, '2025-12-23 14:50:58', '2025-12-23 14:50:58'),
	(158, 121, '添加', '/admin/theme/pages/save', '', '', 0, 1, '2025-12-23 14:53:03', '2025-12-23 14:54:09'),
	(159, 121, '更新', '/admin/theme/pages/update', '', '', 0, 2, '2025-12-23 14:53:29', '2025-12-23 14:53:29'),
	(160, 121, '删除', '/admin/theme/pages/del', '', '', 0, 3, '2025-12-23 14:53:48', '2025-12-23 14:53:48'),
	(161, 126, '添加', '/admin/vip/save', '', '', 0, 1, '2025-12-23 14:55:31', '2025-12-23 14:55:31'),
	(162, 126, '更新', '/admin/vip/edit', '', '', 0, 2, '2025-12-23 14:55:55', '2025-12-23 15:01:22'),
	(163, 126, '删除', '/admin/vip/del', '', '', 0, 3, '2025-12-23 14:56:18', '2025-12-23 14:56:18'),
	(164, 127, '添加', '/admin/vip/addPrivilege', '', '', 0, 1, '2025-12-23 14:57:24', '2025-12-23 14:57:24'),
	(165, 127, '编辑', '/admin/vip/editPrivilege', '', '', 0, 2, '2025-12-23 14:57:47', '2025-12-23 14:57:47'),
	(166, 127, '删除', '/admin/vip/delPrivilege', '', '', 0, 3, '2025-12-23 14:58:22', '2025-12-23 14:58:22'),
	(167, 129, '添加', '/admin/kami/save', '', '', 0, 1, '2025-12-23 14:59:53', '2025-12-23 14:59:53'),
	(168, 129, '更新', '/admin/kami/edit', '', '', 0, 2, '2025-12-23 15:00:25', '2025-12-23 15:01:11'),
	(169, 129, '删除', '/admin/kami/del', '', '', 0, 3, '2025-12-23 15:00:45', '2025-12-23 15:00:45'),
	(170, 86, '设置高级推广', '/admin/user/setUserSuper', '', '', 0, 5, '2025-12-25 19:27:39', '2025-12-25 19:27:39'),
	(171, 36, '余额记录', '/admin/user/balanceList', '/user/balance', '', 1, 4, '2025-12-26 14:15:46', '2025-12-26 14:15:46'),
	(172, 36, '商品获取记录', '/admin/user/goodsList', '/user/goodsList', '', 1, 5, '2025-12-26 15:50:24', '2025-12-26 16:00:06'),
	(173, 111, '任务记录', '/admin/goods/task', '/goods/task', '', 1, 4, '2025-12-26 16:06:57', '2025-12-26 16:06:57'),
	(174, 34, '微信小程序', '', '', '', 1, 5, '2025-12-26 19:28:18', '2025-12-29 14:01:42'),
	(175, 174, '上传小程序', '', '/system/wxMiniUpload', '', 1, 1, '2025-12-26 19:29:01', '2025-12-26 19:39:34'),
	(176, 34, '操作日志', '/admin/system/adminLogList', '/system/adminLog', '', 1, 3, '2025-12-29 14:00:49', '2025-12-29 14:00:49'),
  (177, 87, '刷新数据', '/admin/index/refresh', '', '', 0, 2, '2026-01-05 16:36:07', '2026-01-05 16:36:07'),
	(178, 87, '修改密码', '/admin/index/editPwd', '', '', 0, 3, '2026-01-05 16:38:10', '2026-01-05 16:38:10'),
	(179, 0, '积分营销', '', '', 'Coin', 1, 12, '2026-01-22 15:15:04', '2026-01-28 14:36:52'),
	(180, 179, '签到奖励', '/admin/sign/list', '/points/sign', '', 1, 1, '2026-01-22 15:16:54', '2026-01-28 13:38:42'),
	(181, 180, '新增', '/admin/sign/save', '', '', 0, 1, '2026-01-28 13:39:21', '2026-01-28 13:39:21'),
	(182, 180, '编辑', '/admin/sign/edit', '', '', 0, 2, '2026-01-28 13:39:54', '2026-01-28 13:39:54'),
	(183, 180, '删除', '/admin/sign/del', '', '', 0, 3, '2026-01-28 13:40:19', '2026-01-28 13:40:19'),
	(184, 179, '积分记录', '/admin/points/log', '/points/log', '', 1, 2, '2026-01-28 13:53:03', '2026-01-28 14:08:11');

CREATE TABLE IF NOT EXISTS `auth_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '用户组名称',
  `auth` text COMMENT '选中节点',
  `half_auth` varchar(1000) DEFAULT NULL COMMENT '半选中节点',
  `admin_id` int(11) DEFAULT '0' COMMENT '创建人id',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='权限组';

INSERT INTO `auth_group` (`id`, `name`, `auth`, `half_auth`, `admin_id`, `ctime`, `utime`) VALUES
	(1, '超级管理', NULL, NULL, 0, '2025-04-30 23:50:00', '2025-04-30 23:50:00'),
	(2, '普通管理员', '34,83,110,35,91,93,92,14,18,21,20,19,8,17,16,13,33,15,52,78,100,101,102,54,97,98,99', '5,6', 1, '2025-04-30 23:50:00', '2025-05-21 17:34:56');

CREATE TABLE IF NOT EXISTS `captcha` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1手机2邮箱',
  `account` varchar(150) COLLATE utf8mb4_bin DEFAULT NULL,
  `code` varchar(10) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '验证码',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1已使用',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `account` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='验证码';

CREATE TABLE IF NOT EXISTS `cash` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `third_order` varchar(150) DEFAULT NULL COMMENT '三方订单',
  `Type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1微信2支付宝3银行',
  `money` decimal(11,2) DEFAULT '0.00' COMMENT '提现金额',
  `fee` decimal(11,2) DEFAULT '0.00' COMMENT '手续费',
  `actual` decimal(11,2) DEFAULT '0.00' COMMENT '实付金额',
  `account` varchar(150) DEFAULT NULL COMMENT '账号',
  `realname` varchar(150) DEFAULT NULL COMMENT '姓名',
  `bank` varchar(150) DEFAULT NULL COMMENT '银行',
  `qrcode` varchar(255) DEFAULT NULL COMMENT '收款码',
  `status` tinyint(4) DEFAULT '0' COMMENT '-1拒绝0待处理1已完成2待确认',
  `uid` int(11) DEFAULT '0',
  `remark` varchar(255) DEFAULT NULL COMMENT '说明',
  `platform` varchar(50) DEFAULT NULL COMMENT '来源平台',
  `package_info` varchar(255) DEFAULT NULL COMMENT '微信转账',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='提现表';

CREATE TABLE IF NOT EXISTS `config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(50) DEFAULT NULL COMMENT '标题',
  `key` varchar(50) DEFAULT NULL COMMENT '参数名',
  `group` tinyint(2) DEFAULT NULL COMMENT '所属分组',
  `type` tinyint(2) DEFAULT '1' COMMENT '1:input;2:textarea;3:select;4:开关;5:多项选择;6:图片上传,7文件上传',
  `param` text,
  `value` text COMMENT '值',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `remark` varchar(255) DEFAULT NULL COMMENT '说明',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `key` (`key`)
) ENGINE=InnoDB AUTO_INCREMENT=121 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='系统参数';

INSERT INTO `config` (`id`, `title`, `key`, `group`, `type`, `param`, `value`, `sort`, `remark`, `ctime`, `utime`) VALUES
	(1, '参数分组', 'groups', 1, 5, '基础=1\n存储=2\n支付宝=4\n微信小程序=5\n业务=6\n文档转换=7\n阿里云=8\n腾讯云=9\n微信支付=10\n提现=11\n积分=12', '', 0, '', '2025-04-30 23:50:00', '2025-11-26 13:34:00'),
	(17, '站点logo', 'site_logo', 1, 6, '', '', 0, '', '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(29, 'AccessKeyId', 'aliyun_oss_accessKeyId', 8, 1, '', '', 1, '', '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(30, 'AccessKeySecret', 'aliyun_oss_accessKeySecret', 8, 1, '', '', 2, '', '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(38, '储存节点', 'aliyun_oss_region', 8, 1, NULL, '', 4, '如：杭州为cn-hangzhou', '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(39, '存储桶', 'aliyun_oss_bucket', 8, 1, '', '', 5, '', '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(46, '站点名称', 'site_name', 1, 1, NULL, '', 1, '', '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(54, 'APPID', 'alipay_appid', 4, 1, NULL, '', 2, '', '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(55, '支付宝公钥证书', 'alipay_alipayPublicCert', 4, 2, '', '', 6, 'alipayCertPublicKey_RSA2.crt', '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(56, '支付宝根证书', 'alipay_alipayRootCert', 4, 2, '', '', 5, 'alipayRootCert.crt', '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(57, '应用公钥证书', 'alipay_appPublicCert', 4, 2, '', '', 4, 'appCertPublicKey_xxxx.crt', '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(58, '应用私钥', 'alipay_privateKey', 4, 2, '', '', 3, '', '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(62, '网址', 'site_url', 1, 1, NULL, '', 2, NULL, '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(63, '支付宝启用', 'alipay_on', 4, 4, NULL, '1', 1, NULL, '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(64, '上传类型', 'upload_type', 2, 8, '本地存储=local\n阿里云=aliyun\n腾讯云=tencent', 'aliyun', 0, '设置文件存储位置', '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(65, '联系电话', 'telephone', 1, 1, NULL, '', 4, NULL, '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(66, '联系地址', 'address', 1, 1, '', '', 3, '', '2025-04-30 23:50:00', '2025-12-26 20:18:19'),
	(67, '起始页码', 'doc_start_page', 7, 1, '', '1', 0, '从此页码开始转换', '2025-08-20 14:46:49', '2025-12-26 20:18:19'),
	(68, '结束页码', 'doc_end_page', 7, 1, '', '3', 0, '结束页码（-1表示最后一页）', '2025-08-20 14:47:51', '2025-12-26 20:18:19'),
	(69, '分辨率', 'doc_img_dpi', 7, 1, '', '72', 0, '分辨率（网络推荐72-96，最高不超过150）', '2025-08-20 14:51:33', '2025-12-26 20:18:19'),
	(74, '上传大小限制', 'upload_size', 2, 1, '', '50', 1, '单位MB', '2025-08-26 14:10:55', '2025-12-26 20:18:19'),
	(75, '文件限制', 'upload_file_type', 2, 1, '', '.jpg|.png|.gif|.jpeg|.webp|.doc|.docx|.xls|.xlsx|.ppt|.pptx|.pdf', 2, '允许上传文件的后缀,以|隔开，例如: .jpg|.png|.gif|.jpeg', '2025-08-26 14:18:10', '2025-12-26 20:18:19'),
	(76, 'APPID', 'wx_mini_appid', 5, 1, '', '', 1, '', '2025-10-05 13:34:10', '2025-12-26 20:18:19'),
	(77, 'AppSecret', 'wx_mini_secret', 5, 1, '', '', 2, '', '2025-10-05 13:37:04', '2025-12-26 20:18:19'),
	(78, '开启分销', 'retail_on', 6, 4, '', '1', 1, '开启分销开关', '2025-10-21 17:14:22', '2025-12-26 20:18:19'),
	(79, '佣金类型', 'retail_type', 6, 8, '固定金额=1\n百分比=2', '2', 2, '分销佣金发放方式', '2025-10-21 17:37:59', '2025-12-26 20:18:19'),
	(80, '分销佣金', 'retail_yj', 6, 1, '', '2', 3, '固定值佣金直接填写金额；百分比填写百分比值，如5%填5', '2025-10-21 17:42:42', '2025-12-26 20:18:19'),
	(82, '超级推广', 'super_on', 6, 4, '', '1', 4, '是否开启超级推广人功能', '2025-10-23 16:54:42', '2025-12-26 20:18:19'),
	(83, '身份名称', 'super_name', 6, 1, '', '超级推广人', 5, '超级推广人的身份名称', '2025-10-23 16:57:03', '2025-12-26 20:18:19'),
	(84, '奖励方式', 'super_type', 6, 8, '固定金额=1\n百分比=2', '2', 6, '超级推广人的奖励方式', '2025-10-23 16:59:51', '2025-12-26 20:18:19'),
	(85, '推广佣金', 'super_yj', 6, 1, '', '3', 7, '固定值佣金直接填写金额；百分比填写百分比值，如5%填5', '2025-10-23 17:01:34', '2025-12-26 20:18:19'),
	(86, '提现功能', 'cash_on', 11, 4, '', '1', 1, '', '2025-11-07 16:24:22', '2025-12-26 20:18:19'),
	(87, '微信提现', 'cash_wxpay_type', 11, 8, '关闭=0\n收款码=1\n商家转账=2', '2', 2, '', '2025-11-07 16:27:52', '2025-12-26 20:18:19'),
	(88, '支付宝提现', 'cash_alipay_type', 11, 8, '关闭=0\n收款码=1\n商家转账=2', '1', 3, '', '2025-11-07 17:10:25', '2025-12-26 20:18:19'),
	(89, '银行卡提现', 'cash_bank_type', 11, 4, '关闭=0\n开启=1', '1', 4, '', '2025-11-07 17:11:41', '2025-12-26 20:18:19'),
	(90, '最低提现金额', 'cash_min_money', 11, 1, '', '10', 5, '', '2025-11-07 17:46:42', '2025-12-26 20:18:19'),
	(91, '最大提现金额', 'cash_max_money', 11, 1, '', '200', 6, '', '2025-11-07 17:48:09', '2025-12-26 20:18:19'),
	(92, '限制提现周期', 'cash_cycle', 11, 8, '无限制=0\n每日=1\n每周=2\n每月=3', '1', 7, '', '2025-11-07 17:50:51', '2025-12-26 20:18:19'),
	(93, '限制提现次数', 'cash_nums', 11, 1, '', '1', 8, '', '2025-11-18 14:13:26', '2025-12-26 20:18:19'),
	(94, '提现手续费', 'cash_fee', 11, 1, '', '1', 9, '单位百分比%，无手续费填0', '2025-11-18 14:36:39', '2025-12-26 20:18:19'),
	(95, '访问域名', 'aliyun_oss_url', 8, 1, '', '', 5, '', '2025-11-25 15:53:07', '2025-12-26 20:18:19'),
	(96, 'SecretId', 'tencent_oss_secretId', 9, 1, '', '', 1, '', '2025-11-25 16:23:58', '2025-12-26 20:18:19'),
	(97, 'SecretKey', 'tencent_oss_secretKey', 9, 1, '', '', 2, '', '2025-11-25 16:25:00', '2025-12-26 20:18:19'),
	(98, 'Bucket域名', 'tencent_oss_backetUrl', 9, 1, '', '', 3, '', '2025-11-25 16:33:06', '2025-12-26 20:18:19'),
	(99, '访问域名', 'tencent_oss_url', 9, 1, '', '', 4, '如有绑定自定义域名则填写', '2025-11-26 13:31:46', '2025-12-26 20:18:19'),
	(100, '支付开关', 'wxpay_on', 10, 4, '', '1', 1, '', '2025-11-26 14:09:23', '2025-12-26 20:18:19'),
	(101, '商户号', 'wxpay_mchId', 10, 1, '', '', 2, '', '2025-11-28 14:41:47', '2025-12-26 20:18:19'),
	(102, 'apiV3Key', 'wxpay_apiV3Key', 10, 1, '', '', 3, '', '2025-11-28 14:42:47', '2025-12-26 20:18:19'),
	(103, '私钥', 'wxpay_privateKey', 10, 2, '', '', 4, '证书apiclient_key.pem内容', '2025-11-28 14:43:35', '2025-12-26 20:18:19'),
	(104, '证书序列号', 'wxpay_serialNo', 10, 1, '', '', 5, '', '2025-11-28 14:44:20', '2025-12-26 20:18:19'),
	(105, '商品免费获取规则', 'goods_tg_type', 6, 8, '新注册用户=1\n进入推广商品页=2', '1', 8, '', '2025-11-28 16:42:43', '2025-12-26 20:18:19'),
	(106, '任务消息模板ID', 'wx_mini_task_templateId', 5, 1, '', '', 9, '', '2025-11-28 18:37:13', '2025-12-26 20:18:19'),
	(107, '提现消息模板ID', 'wx_mini_cash_templateId', 5, 1, '', '', 10, '', '2025-11-28 18:38:12', '2025-12-26 20:18:19'),
	(108, '支付公钥ID', 'wxpay_publicKeyID', 10, 1, '', '', 6, '', '2025-12-19 14:17:22', '2025-12-26 20:18:19'),
	(109, '公钥', 'wxpay_publicKey', 10, 2, '', '', 7, '公钥文件pub_key.pem内容', '2025-12-19 14:31:54', '2025-12-26 20:18:19'),
	(110, '上传密钥', 'wx_mini_upload_private_key', 5, 2, '', '', 3, '', '2025-12-26 20:17:44', '2025-12-26 20:18:19'),
  (111, '手机登录', 'phone_login', 6, 4, '', '1', 9, '使用手机登录，需设置阿里云或腾讯云短信API', '2025-12-30 18:22:36', '2025-12-30 20:31:16'),
	(112, '短信平台', 'sms_type', 6, 8, '关闭=0\n阿里云=1\n腾讯云=2', '1', 10, '发送短信平台阿里云或腾讯云', '2025-12-30 18:25:31', '2025-12-30 20:31:16'),
	(113, '短信签名', 'aliyun_sms_signName', 8, 1, '', '', 7, '', '2025-12-30 18:40:59', '2025-12-30 20:31:16'),
	(114, '短信模板CODE', 'aliyun_sms_templateCode', 8, 1, '', '', 8, '', '2025-12-30 18:42:26', '2025-12-30 20:31:16'),
	(115, '短信SDKAppID', 'tencent_sms_appid', 9, 1, '', '', 5, '', '2025-12-30 20:26:56', '2025-12-30 20:31:16'),
	(116, '短信签名', 'tencent_sms_signName', 9, 1, '', '', 6, '', '2025-12-30 20:28:09', '2025-12-30 20:31:16'),
	(117, '短信模板ID', 'tencent_sms_templateId', 9, 1, '', '', 7, '', '2025-12-30 20:29:27', '2025-12-30 20:31:16'),
  (118, '签到开关', 'sign_on', 12, 4, '', '1', 1, '', '2026-01-22 14:42:18', '2026-01-26 20:07:19'),
	(119, '签到模式', 'sign_mode', 12, 8, '无限制=1\n周循环=2\n月循环=3', '3', 2, '无限制，累积和连续签到不会清零，周循环，每周一会清理累积和连续的记录为0，重新开启计算，月循环，每月一号会清理累积和连续的记录为0，重新开启计算', '2026-01-22 14:44:40', '2026-01-26 20:07:19'),
	(120, '签到赠送积分', 'sign_give_points', 12, 1, '', '1', 3, '', '2026-01-22 14:48:08', '2026-01-26 20:07:19'),
  (121, '支付积分奖励', 'points_money_ratio', 12, 1, '', '0', 4, '订单交易金额发放积分的比例；0不发积分。如填10则代表满10元发放1积分，不足则不发放。', '2026-01-22 14:48:08', '2026-01-26 20:07:19');

CREATE TABLE IF NOT EXISTS `gl_correlation` (
  `id` int(11) AUTO_INCREMENT NOT NULL,
  `gid` int(11) DEFAULT '0',
  `lid` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `gid` (`gid`),
  KEY `lid` (`lid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='商品标签关联表';

CREATE TABLE IF NOT EXISTS `goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '标题',
  `thumb` varchar(255) NOT NULL DEFAULT '' COMMENT '封面',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型1文档2文章3视频4音频5网盘',
  `link` varchar(500) NOT NULL DEFAULT '' COMMENT '链接地址',
  `introduction` varchar(500) NOT NULL DEFAULT '' COMMENT '简介',
  `cost_price` decimal(20,2) NOT NULL DEFAULT '0.00' COMMENT '成本价',
  `original_price` decimal(20,2) NOT NULL DEFAULT '0.00' COMMENT '原价，划线价',
  `price` decimal(20,2) NOT NULL DEFAULT '0.00' COMMENT '价格',
  `points` int(11) NOT NULL DEFAULT '0' COMMENT '积分',
  `code` varchar(50) NOT NULL DEFAULT '' COMMENT '提取码',
  `pages` int(11) NOT NULL DEFAULT '1' COMMENT '总页数',
  `is_top` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '置顶',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '浏览数',
  `retail_on` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '分销-1关闭1开启',
  `retail_set` tinyint(4) NOT NULL DEFAULT '1' COMMENT '佣金设置1公用2单独设置',
  `retail_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '佣金方式1固定2百分比',
  `retail_yj` float NOT NULL DEFAULT '0' COMMENT '一级佣金',
  `is_adv` tinyint(4) NOT NULL DEFAULT '1' COMMENT '广告获取1开启-1关闭',
  `adv_nums` int(11) NOT NULL DEFAULT '0' COMMENT '广告次数',
  `is_share` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '分享获取1开启-1关闭',
  `share_nums` int(11) NOT NULL DEFAULT '0' COMMENT '任务获取分享数量',
  `status` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '-1下架1上架',
  `cid` int(11) NOT NULL DEFAULT '0' COMMENT '栏目ID',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `utime` (`utime`),
  KEY `price` (`price`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='商品';

CREATE TABLE IF NOT EXISTS `goods_column` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `ctime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='商品栏目';

CREATE TABLE IF NOT EXISTS `goods_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gid` int(11) NOT NULL DEFAULT '0',
  `content` text NOT NULL,
  `fee_content` text COMMENT '收费内容',
  PRIMARY KEY (`id`),
  KEY `gid` (`gid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='商品详情';


CREATE TABLE IF NOT EXISTS `goods_label` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1正常-1禁用',
  `ctime` datetime DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='商品标签';


CREATE TABLE IF NOT EXISTS `goods_sort` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '上级',
  `thumb` varchar(255) DEFAULT NULL,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `sort` int(11) NOT NULL DEFAULT '0',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1正常-1禁用',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='商品分类';

CREATE TABLE IF NOT EXISTS `goods_task` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0',
  `gid` int(11) NOT NULL DEFAULT '0',
  `adv` int(11) NOT NULL DEFAULT '0' COMMENT '任务要求广告',
  `adv_nums` int(11) NOT NULL DEFAULT '0' COMMENT '已完成广告',
  `share` int(11) NOT NULL DEFAULT '0' COMMENT '任务要求分享数量',
  `share_nums` int(11) NOT NULL DEFAULT '0' COMMENT '已完成分享数量',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `gid` (`gid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='任务获取商品';

CREATE TABLE IF NOT EXISTS `gs_correlation` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gid` int(11) DEFAULT '0',
  `sid` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `key` (`gid`,`sid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='商品分类关联表';

CREATE TABLE IF NOT EXISTS `img` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `thumb` varchar(255) DEFAULT NULL COMMENT '图片地址',
  `sid` smallint(6) DEFAULT NULL,
  `status` tinyint(4) DEFAULT '1' COMMENT '状态1:正常-1:禁用',
  `sort` smallint(6) DEFAULT NULL COMMENT '排序',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='图片表';

CREATE TABLE IF NOT EXISTS `img_sort` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '名称',
  `status` tinyint(4) DEFAULT '1' COMMENT '状态1:正常-1:禁用',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='图片分类';

CREATE TABLE IF NOT EXISTS `kami` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `key` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  `price` decimal(11,2) NOT NULL DEFAULT '0.00',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1已使用',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `key` (`key`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='卡密';

CREATE TABLE IF NOT EXISTS `order` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_no` varchar(150) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '订单号',
  `third_order` varchar(150) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '第三方订单/卡密',
  `type` tinyint(4) DEFAULT '0' COMMENT '1商品订单2vip订单',
  `uid` int(11) DEFAULT '0',
  `gid` int(11) DEFAULT '0',
  `vid` int(11) DEFAULT '0',
  `pay_type` tinyint(4) DEFAULT '0' COMMENT '1微信2支付宝3卡密4余额',
  `price` decimal(10,2) DEFAULT '0.00' COMMENT '价格',
  `discount_price` decimal(10,2) DEFAULT '0.00' COMMENT '抵扣/优惠',
  `discount_type` tinyint(4) DEFAULT '0' COMMENT '抵扣类型1优惠券2会员优惠',
  `pay_price` decimal(10,2) DEFAULT '0.00' COMMENT '支付金额',
  `retail_price` decimal(10,2) DEFAULT '0.00' COMMENT '推广分润',
  `super_price` decimal(10,2) DEFAULT '0.00' COMMENT '超级推广分润',
  `status` tinyint(4) DEFAULT '0' COMMENT '0待支付1已付款',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_no` (`order_no`),
  KEY `uid` (`uid`),
  KEY `gid` (`gid`),
  KEY `utime` (`utime`),
  KEY `vid` (`vid`),
  KEY `type` (`type`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='订单表';


CREATE TABLE IF NOT EXISTS `pages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `path` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `title` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `remark` varchar(300) COLLATE utf8mb4_bin DEFAULT NULL,
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='前端页面地址';

INSERT INTO `pages` (`id`, `path`, `title`, `remark`, `ctime`, `utime`) VALUES
	(1, '/pages/index/index', '首页', '', '2025-12-22 17:08:50', '2025-12-22 17:08:50'),
	(2, '/pages/sort/index', '分类', '', '2025-12-22 17:11:07', '2025-12-22 17:11:07'),
	(3, '/pages/vip/index', '会员中心', '', '2025-12-22 17:11:43', '2025-12-22 17:11:43'),
	(4, '/pages/user/index', '个人中心', '', '2025-12-22 17:12:11', '2025-12-22 17:12:11'),
	(5, '/pages/article/index', '文章页', '参数:文章id', '2025-12-22 17:13:59', '2025-12-22 17:15:13'),
	(6, '/pages/goods/info', '商品详情', '参数:商品id', '2025-12-22 17:15:08', '2025-12-22 17:15:08'),
	(7, '/pages/login/index', '登录', '参数:跳转地址back', '2025-12-22 17:16:52', '2025-12-22 17:16:52'),
	(8, '/pages/order/buy', '商品购买', '参数:商品id', '2025-12-22 17:18:17', '2025-12-22 17:18:17'),
	(9, '/pages/order/list', '我的订单', '', '2025-12-22 17:18:52', '2025-12-22 17:18:52'),
	(10, '/pages/order/buy-vip', '购买会员', '参数:会员id', '2025-12-22 17:20:08', '2025-12-22 17:20:08'),
	(11, '/pages/goods/index', '商品列表', '参数:type(1热门2指定3价格排序4最新)&sort_id(商品分类ID)', '2025-12-22 17:24:28', '2025-12-22 17:24:28'),
	(12, '/pages/user/set', '用户设置', '', '2025-12-22 17:24:58', '2025-12-22 17:24:58'),
	(13, '/pages/user/goods', '已获取商品', '', '2025-12-22 17:25:27', '2025-12-22 17:25:36'),
	(14, '/pages/user/wallet', '钱包', '', '2025-12-22 17:26:02', '2025-12-22 17:26:02'),
	(15, '/pages/user/bill', '余额记录', '', '2025-12-22 17:26:23', '2025-12-22 17:26:23'),
	(16, '/pages/cash/index', '提现', '', '2025-12-22 17:26:55', '2025-12-22 17:26:55'),
	(17, '/pages/user/promotion', '推广中心', '', '2025-12-22 17:27:26', '2025-12-22 17:27:26'),
	(18, '/pages/user/myPromotion', '推广列表', '', '2025-12-22 17:27:45', '2025-12-22 17:27:45'),
	(19, '/pages/user/subSuper', '申请高级推广', '', '2025-12-22 17:28:55', '2025-12-22 17:28:55'),
	(20, '/pages/cash/list', '提现列表', '', '2025-12-22 17:29:20', '2025-12-22 17:29:20'),
	(21, '/pages/task/goodsTask', '任务列表', '', '2025-12-22 17:29:54', '2025-12-22 17:34:13'),
	(23, '/pages/article/list', '文章列表', '参数:文章分类sid', '2025-12-22 18:39:58', '2025-12-22 18:39:58'),
  (24, '/pages/user/collect', '我的收藏', '', '2026-01-13 16:19:21', '2026-01-13 16:19:21'),
  (25, '/pages/points/mall', '积分商城', '', '2026-01-29 15:32:34', '2026-01-29 15:32:34'),
	(26, '/pages/points/sign', '签到', '', '2026-01-29 15:33:04', '2026-01-29 15:33:04'),
	(27, '/pages/points/log', '积分记录', '', '2026-01-29 15:34:07', '2026-01-29 15:34:07');

CREATE TABLE IF NOT EXISTS `privilege` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(150) COLLATE utf8mb4_bin DEFAULT NULL,
  `type` tinyint(4) DEFAULT '0' COMMENT '1商品类型2商品分类',
  `content` json DEFAULT NULL,
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='特权';


CREATE TABLE IF NOT EXISTS `svip` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(150) COLLATE utf8mb4_bin DEFAULT NULL,
  `desc` text COLLATE utf8mb4_bin COMMENT '描述说明',
  `days` int(11) DEFAULT '0' COMMENT '时长（天）',
  `price` decimal(11,2) DEFAULT '0.00' COMMENT '价格',
  `rebate` float DEFAULT '0' COMMENT '折扣',
  `points` int(11) DEFAULT '0' COMMENT '积分兑换',
  `status` tinyint(4) DEFAULT '0' COMMENT '-1禁用1启用',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `retail_on` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '分销-1关闭1开启',
  `retail_set` tinyint(4) NOT NULL DEFAULT '1' COMMENT '佣金设置1公用2单独设置',
  `retail_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '佣金方式1固定2百分比',
  `retail_yj` float NOT NULL DEFAULT '0' COMMENT '分销佣金',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;


CREATE TABLE IF NOT EXISTS `svip_privilege` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `vid` int(11) NOT NULL DEFAULT '0',
  `pid` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='svip特权关联表';

CREATE TABLE IF NOT EXISTS `theme` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1首页2分类页3会员4个人中心5导航6推广图',
  `title` varchar(50) NOT NULL DEFAULT '0' COMMENT '名称',
  `content` json DEFAULT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '-1未启用1启用',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='装修布局';

INSERT INTO `theme` (`id`, `type`, `title`, `content`, `status`, `ctime`, `utime`) VALUES
	(1, 1, '首页', '[{"h": 5, "i": "HeEc3N", "w": 1, "x": 0, "y": 0, "data": {"color": "#ffffff", "title": "首页", "backgroundColor": "#f56b59"}, "type": "status", "moved": false}, {"h": 15, "i": "SPQMpP", "w": 1, "x": 0, "y": 5, "data": {"list": [], "gradientEndColor": "#ffffff", "gradientStartColor": "#f56b59"}, "type": "search-carousel", "moved": false}, {"h": 3, "i": "De5JP6", "w": 1, "x": 0, "y": 20, "data": {"color": "#e21515", "fontSize": "14px", "backgroundColor": "#ffffff"}, "type": "notice", "moved": false}, {"h": 17, "i": "iNT6yr", "w": 1, "x": 0, "y": 23, "data": {"list": [], "color": "#000000", "arrange": 4, "fontSize": "14px", "backgroundColor": "#ffffff"}, "type": "classify", "moved": false}, {"h": 6, "i": "bckA7R", "w": 1, "x": 0, "y": 40, "data": {"link": "", "height": "120px", "imgUrl": ""}, "type": "image", "moved": false}, {"h": 21, "i": "dCrHs3", "w": 1, "x": 0, "y": 46, "data": {"list": [], "type": 1, "color": "#ffffff", "limit": "6", "title": "热门商品", "sortId": "0", "arrange": 2, "fontSize": "14px", "backgroundColor": "#ffffff", "gradientEndColor": "#c68686", "gradientStartColor": "#f56b59"}, "type": "goods", "moved": false}]', 1, '2025-09-29 09:45:29', '2026-01-06 17:46:37'),
	(2, 2, '分类', '[{"type": 2}]', 1, '2025-10-17 12:51:38', '2025-10-17 12:51:39'),
	(3, 3, '会员中心', '[{"color": "#ffffff", "title": "会员中心", "iconColor": "#efb814", "titleColor": "#ffffff", "backgroundColor": "#f56b59"}]', 1, '2025-10-17 12:52:07', '2025-10-29 19:00:10'),
	(4, 4, '个人中心', '[{"h": 5, "i": "QbmYnm", "w": 1, "x": 0, "y": 0, "data": {"color": "#ffffff", "title": "个人中心", "backgroundColor": "#f56b59"}, "type": "status", "moved": false}, {"h": 12, "i": "esYSTQ", "w": 1, "x": 0, "y": 5, "data": {"color": "#ffffff", "iconColor": "#efb814", "backgroundColor": "#f56b59", "gradientEndColor": "#5e5e40", "gradientStartColor": "#3e3e39"}, "type": "user", "moved": false}, {"h": 17, "i": "d5wJT6", "w": 1, "x": 0, "y": 17, "data": {"list": [{"link": "/pages/order/list", "name": "订单", "imgUrl": ""}, {"link": "/pages/user/goods", "name": "已获取", "imgUrl": ""}, {"link": "/pages/user/wallet", "name": "钱包", "imgUrl": ""}], "color": "#000000", "arrange": 3, "fontSize": "14px", "backgroundColor": "#ffffff"}, "type": "operate", "moved": false}, {"h": 17, "i": "AQD25P", "w": 1, "x": 0, "y": 34, "data": {"list": [{"link": "/pages/cash/index", "name": "提现", "imgUrl": ""}, {"link": "/pages/user/set", "name": "设置", "imgUrl": ""}, {"link": "/pages/task/goodsTask", "name": "任务", "imgUrl": ""}, {"link": "/pages/user/promotion", "name": "推广", "imgUrl": ""}], "color": "#000000", "arrange": 1, "fontSize": "14px", "backgroundColor": "#ffffff"}, "type": "operate", "moved": false}]', 1, '2025-10-17 12:52:29', '2026-01-06 17:48:53'),
	(5, 5, '导航', '[{"list": [{"link": "/pages/index/index", "name": "首页", "imgUrl": "", "selectImgUrl": ""}, {"link": "/pages/sort/index", "name": "分类", "imgUrl": "", "selectImgUrl": ""}, {"link": "/pages/vip/index", "name": "会员", "imgUrl": "", "selectImgUrl": ""}, {"link": "/pages/user/index", "name": "我的", "imgUrl": "", "selectImgUrl": ""}], "color": "#3d3d3d", "selectColor": "#ff0000", "backgroundColor": "#ffffff"}]', 1, '2025-10-29 13:31:44', '2026-01-06 17:51:28'),
	(6, 6, '推广图', '[{"color": "#f5e211", "title": "邀您使用源启云联", "imgUrl": "", "fontSize": "24px"}]', 1, '2025-11-12 14:47:57', '2026-01-06 17:55:12');
  
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '会员id',
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '上级用户',
  `pass` varchar(100) DEFAULT NULL COMMENT '密码',
  `token` varchar(150) DEFAULT NULL,
  `nickname` varchar(100) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(150) DEFAULT NULL COMMENT '头像',
  `balance` decimal(20,2) DEFAULT '0.00' COMMENT '余额',
  `points` int(11) DEFAULT '0' COMMENT '积分',
  `sign_num` int(11) DEFAULT '0' COMMENT '累计签到',
  `promotion` int(11) DEFAULT '0' COMMENT '推广人数',
  `status` tinyint(4) DEFAULT '1' COMMENT '1正常-1禁用',
  `is_super` tinyint(4) DEFAULT '0' COMMENT '超级身份',
  `super_set` tinyint(4) DEFAULT '0' COMMENT '1公用2单独设置',
  `super_type` tinyint(4) DEFAULT '0' COMMENT '佣金方式1固定2百分比',
  `super_yj` float DEFAULT '0',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1001 DEFAULT CHARSET=utf8 COMMENT='用户表';


CREATE TABLE IF NOT EXISTS `user_bill` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) DEFAULT '0' COMMENT '1推广2超级推广3提现4余额购买',
  `mode` tinyint(4) DEFAULT '0' COMMENT '1增加2减少',
  `before` decimal(20,2) DEFAULT '0.00' COMMENT '变化前',
  `after` decimal(20,2) DEFAULT '0.00' COMMENT '变化后',
  `money` decimal(20,2) DEFAULT '0.00' COMMENT '金额',
  `uid` int(11) DEFAULT '0',
  `remark` varchar(255) DEFAULT NULL COMMENT '注解',
  `tid` int(11) DEFAULT '0' COMMENT '关联表ID',
  `ctime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='余额记录';

CREATE TABLE IF NOT EXISTS `user_bind` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0',
  `openid` varchar(150) COLLATE utf8mb4_bin DEFAULT NULL,
  `phone` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL,
  `email` varchar(150) COLLATE utf8mb4_bin DEFAULT NULL,
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1手机2微信小程序3微信公众号',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `openid` (`openid`),
  KEY `phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户绑定信息';

CREATE TABLE IF NOT EXISTS `user_collect` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0',
  `gid` int(11) NOT NULL DEFAULT '0',
  `ctime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `gid` (`gid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户收藏';

CREATE TABLE IF NOT EXISTS `user_goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0',
  `gid` int(11) NOT NULL DEFAULT '0',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '获取方式1购买2任务',
  `ctime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `gid` (`gid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户拥有商品';

CREATE TABLE IF NOT EXISTS `user_svip` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0',
  `vid` int(11) NOT NULL DEFAULT '0',
  `expire_time` datetime DEFAULT NULL,
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `expire_time` (`expire_time`),
  KEY `utime` (`utime`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户svip记录';

CREATE TABLE IF NOT EXISTS `sign_reward` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `days` int(11) NOT NULL DEFAULT '0' COMMENT '签到天数',
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '1连续签到2累积签到',
  `points` int(11) NOT NULL DEFAULT '0' COMMENT '积分',
  `ctime` datetime DEFAULT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='签到奖励表';

CREATE TABLE IF NOT EXISTS `user_points` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0',
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '1签到2兑换vip3兑换商品',
  `points` int(11) NOT NULL DEFAULT '0',
  `mode` int(11) NOT NULL DEFAULT '1' COMMENT '1增加2减少',
  `before` int(11) NOT NULL DEFAULT '0' COMMENT '变更前',
  `after` int(11) NOT NULL DEFAULT '0' COMMENT '变更后',
  `remark` varchar(150) COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  `ctime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户积分记录';