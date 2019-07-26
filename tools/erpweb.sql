/*
Navicat MySQL Data Transfer

Source Server         : minic630
Source Server Version : 50626
Source Host           : minic630.linkdood.cn:11306
Source Database       : erpweb

Target Server Type    : MYSQL
Target Server Version : 50626
File Encoding         : 65001

Date: 2019-07-10 17:28:05
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '账号信息表',
  `cardid` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '工号',
  `status` int(4) NOT NULL COMMENT '账号状态 1：正常；2:冻结; 3:停用；4：锁定',
  `password` varchar(500) CHARACTER SET utf8mb4 NOT NULL COMMENT '密码',
  PRIMARY KEY (`id`),
  UNIQUE KEY `cardid` (`cardid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='账号信息表';

-- ----------------------------
-- Table structure for arrivalbill
-- ----------------------------
DROP TABLE IF EXISTS `arrivalbill`;
CREATE TABLE `arrivalbill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `arrivalbillcode` varchar(50) CHARACTER SET utf8mb4 NOT NULL COMMENT '到货单编码',
  `contractcode` varchar(50) CHARACTER SET utf8mb4 NOT NULL COMMENT '采购合同编号',
  `handler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '制单人',
  `status` tinyint(5) NOT NULL COMMENT '状态(1:制单;2:已入库)',
  `warehouseid` bigint(20) NOT NULL COMMENT '仓库id(仓库表主键)',
  `createat` datetime NOT NULL COMMENT '制单时间',
  `indate` datetime DEFAULT NULL COMMENT '入库日期',
  `storehandler` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '入库操作员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='到货订单表';

-- ----------------------------
-- Table structure for arrivaldetail
-- ----------------------------
DROP TABLE IF EXISTS `arrivaldetail`;
CREATE TABLE `arrivaldetail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `arrivalbillcode` varchar(50) CHARACTER SET utf8mb4 NOT NULL COMMENT '到货单编码',
  `mattercode` varchar(50) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料编码',
  `name` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料名称',
  `unit` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '单位',
  `arrivalnum` bigint(20) NOT NULL COMMENT '到货数量',
  `putinnum` bigint(20) NOT NULL DEFAULT '0' COMMENT '入库数量',
  `price` bigint(20) NOT NULL COMMENT '单价',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='到货详情表';

-- ----------------------------
-- Table structure for customer
-- ----------------------------
DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `custcode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '客户编码',
  `name` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '客户名称',
  `zipcode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '邮编',
  `postaddress` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '通讯地址',
  `taxnum` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '税号',
  `depositbank` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '开户银行',
  `bankaccount` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '银行账号',
  `railwayadmin` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '铁路局',
  `maintainsection` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '机务段',
  `remark` varchar(1000) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='客户信息表';

-- ----------------------------
-- Table structure for department
-- ----------------------------
DROP TABLE IF EXISTS `department`;
CREATE TABLE `department` (
  `id` bigint(20) NOT NULL,
  `compID` bigint(20) NOT NULL COMMENT '公司ID',
  `name` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '部门名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='部门信息表';

-- ----------------------------
-- Table structure for employee
-- ----------------------------
DROP TABLE IF EXISTS `employee`;
CREATE TABLE `employee` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '人员id',
  `cardid` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '工号',
  `name` varchar(50) CHARACTER SET utf8mb4 NOT NULL COMMENT '姓名',
  `sex` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别 0男 1女',
  `compID` tinyint(4) NOT NULL DEFAULT '0' COMMENT '所属分公司ID  0总部1北京2杭州',
  `deptID` int(4) NOT NULL COMMENT '部门ID',
  `dutyID` int(4) DEFAULT NULL COMMENT '岗位ID',
  `health` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '身体状况',
  `height` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '身高',
  `nativeplace` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '籍贯',
  `nation` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '民族',
  `maritalstatus` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '婚姻状况',
  `education` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '学历',
  `university` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '毕业院校',
  `major` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '专业',
  `qualification` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '专业资格',
  `trialsalary` bigint(20) DEFAULT NULL COMMENT '试用期工资',
  `salary` bigint(20) DEFAULT NULL COMMENT '转正后工资',
  `idnumber` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '身份证号',
  `address1` varchar(500) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '户口地址',
  `postcode1` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '邮编',
  `address2` varchar(500) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '现住地址',
  `postcode2` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '邮编',
  `contactnumber` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '联系电话',
  `phonenumber` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '手机',
  `email` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT 'email',
  `emergencycontact` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '紧急情况联系人',
  `contactnumber1` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '联系电话',
  `address3` varchar(500) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '现在何处',
  `trialexpired` date NOT NULL COMMENT '试用到期',
  `entrydate` date NOT NULL COMMENT '入职日期',
  `birthday` date DEFAULT NULL COMMENT '出生日期',
  `contractbegindate` date NOT NULL COMMENT '合同开始日期',
  `contractenddate` date NOT NULL COMMENT '合同结束日期',
  PRIMARY KEY (`id`),
  UNIQUE KEY `cardid` (`cardid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='人员信息表';

-- ----------------------------
-- Table structure for financeflow
-- ----------------------------
DROP TABLE IF EXISTS `financeflow`;
CREATE TABLE `financeflow` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `type` tinyint(5) NOT NULL COMMENT '类型(1：维修合同；2：销售合同；3：采购合同)',
  `direction` tinyint(5) NOT NULL COMMENT '资金流向(1：流入；-1：流出)',
  `account` bigint(20) NOT NULL COMMENT '金额',
  `paymethod` tinyint(5) NOT NULL COMMENT '支付方式(1：现金；2：银行转账；3：支票)',
  `billcode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '发票编号',
  `attachment` varchar(1000) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '附件',
  `remark` varchar(1000) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '备注',
  `handler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '操作人',
  `operdate` datetime NOT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='财务流水表';

-- ----------------------------
-- Table structure for inquiry
-- ----------------------------
DROP TABLE IF EXISTS `inquiry`;
CREATE TABLE `inquiry` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `inquirycode` bigint(20) DEFAULT NULL COMMENT '询价单号',
  `type` tinyint(5) NOT NULL COMMENT '类型(1：维修；2：销售)',
  `custcode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '客户编码',
  `createat` datetime NOT NULL COMMENT '制单日期',
  `deadline` date DEFAULT NULL COMMENT '截止日期',
  `status` tinyint(5) NOT NULL COMMENT '状态(1：未回复；2：已回复)',
  `handler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '制单人',
  `replyhandler` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '回复人',
  `replydate` date DEFAULT NULL COMMENT '回复日期',
  `remark` text CHARACTER SET utf8mb4 COMMENT '备注',
  `attachment` text CHARACTER SET utf8mb4 COMMENT '附件',
  `validity` int(5) DEFAULT NULL COMMENT '价格有效期(单位：月)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `inquirycode` (`inquirycode`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='询价表';

-- ----------------------------
-- Table structure for inquirydetail
-- ----------------------------
DROP TABLE IF EXISTS `inquirydetail`;
CREATE TABLE `inquirydetail` (
  `inquirycode` bigint(20) NOT NULL COMMENT '询价单号',
  `mattercode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料编号',
  `num` bigint(20) NOT NULL COMMENT '数量',
  `price` bigint(20) DEFAULT NULL COMMENT '价格',
  PRIMARY KEY (`inquirycode`,`mattercode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='询价详情表';

-- ----------------------------
-- Table structure for leave
-- ----------------------------
DROP TABLE IF EXISTS `leaves`;
CREATE TABLE `leaves` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `employeeid` bigint(20) NOT NULL,
  `leaveat` date DEFAULT NULL COMMENT '离职日期',
  `reason` text CHARACTER SET utf8mb4 COMMENT '离职原因',
  PRIMARY KEY (`id`),
  UNIQUE KEY `employeeid` (`employeeid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='离职登记表';

-- ----------------------------
-- Table structure for marketcontract
-- ----------------------------
DROP TABLE IF EXISTS `marketcontract`;
CREATE TABLE `marketcontract` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `contractcode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '合同编号',
  `type` tinyint(5) NOT NULL COMMENT '类型 1:维修合同;2:销售合同',
  `custcode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '客户编号',
  `handler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '制单人',
  `execstatus` tinyint(5) DEFAULT NULL COMMENT '执行状态(1:制单；2:审核；3:执行中；6:结束)',
  `signdate` date DEFAULT NULL COMMENT '签订日期',
  `deadline` date DEFAULT NULL COMMENT '结束日期',
  `createat` datetime DEFAULT NULL COMMENT '制单日期',
  `amount` bigint(20) DEFAULT NULL COMMENT '金额',
  `settlestatus` tinyint(5) DEFAULT NULL COMMENT '结算状态(1:未结算；2:已结算；3:部分结算)',
  `settleamount` bigint(20) DEFAULT NULL COMMENT '结算金额',
  `attachment` varchar(200) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '附件',
  `remark` text CHARACTER SET utf8mb4 COMMENT '备注',
  `currentreviewer` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '当前审核人',
  `currentreviewindex` tinyint(5) DEFAULT '-1' COMMENT '当前审核序号',
  `relatedcode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '关联合同编号',
  `vehicles` text CHARACTER SET utf8mb4 NOT NULL COMMENT '车辆列表(车辆编号的json数组)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `contractcode` (`contractcode`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='市场合同表';

-- ----------------------------
-- Table structure for matter
-- ----------------------------
DROP TABLE IF EXISTS `matter`;
CREATE TABLE `matter` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `mattercode` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '物料编码',
  `name` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料名称',
  `brand` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '品牌',
  `unit` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '单位',
  `class` bigint(20) NOT NULL COMMENT '物料分类id',
  `param` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '规格参数',
  `grossweight` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '毛重',
  `netweight` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '净重',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mattercode` (`mattercode`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='物料信息表';

-- ----------------------------
-- Table structure for matterplan
-- ----------------------------
DROP TABLE IF EXISTS `matterplan`;
CREATE TABLE `matterplan` (
  `itemid` bigint(20) NOT NULL,
  `mattercode` varchar(20) COLLATE utf8_bin NOT NULL,
  `plannum` bigint(20) NOT NULL,
  PRIMARY KEY (`itemid`,`mattercode`,`plannum`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='备料计划表';

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` bigint(20) NOT NULL COMMENT '菜单id',
  `title` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '标题',
  `icon` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '图标',
  `parentID` bigint(20) NOT NULL COMMENT '父ID',
  `path` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '路径',
  `component` varchar(500) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '组件名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单表（本表所有数据是预置的，不需要增改删，只需要查询）';

-- ----------------------------
-- Table structure for operlog
-- ----------------------------
DROP TABLE IF EXISTS `operlog`;
CREATE TABLE `operlog` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `operator` bigint(20) NOT NULL COMMENT '操作人ID(employee表的id字段)',
  `detail` text CHARACTER SET utf8mb4 NOT NULL COMMENT '操作详情',
  `doat` datetime NOT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='操作日志表';

-- ----------------------------
-- Table structure for outofdetail
-- ----------------------------
DROP TABLE IF EXISTS `outofdetail`;
CREATE TABLE `outofdetail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `outcode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '出库单编号',
  `mattercode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '物料编码',
  `num` bigint(20) NOT NULL COMMENT '出库数量',
  `price` bigint(20) NOT NULL COMMENT '单价',
  `value` bigint(20) NOT NULL COMMENT '总价',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='出库详情表';

-- ----------------------------
-- Table structure for outofstore
-- ----------------------------
DROP TABLE IF EXISTS `outofstore`;
CREATE TABLE `outofstore` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `outcode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '出库单编号',
  `warehouseid` bigint(20) NOT NULL COMMENT '仓库id',
  `type` tinyint(5) NOT NULL COMMENT '类型(1：维修领料；2：销售出库；3：调拨出库)',
  `relatedcode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '关联单据号(当类型为销售出库时，此为销售合同编号，决定出库列表)',
  `outdate` datetime NOT NULL COMMENT '出库时间',
  `storehandler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '出库操作人',
  `pickhandler` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '领料人',
  `contractcode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '维修合同编号',
  `vehiclecode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '车辆编号',
  `itemname` varchar(1000) CHARACTER SET utf8mb4 NOT NULL COMMENT '维修项目名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='物料出库表';

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `userID` bigint(20) NOT NULL COMMENT 'employee的id字段',
  `read` varchar(1000) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '只读菜单列表(为逗号分隔的数字,返回时按数组)',
  `write` varchar(1000) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '可写菜单列表(为逗号分隔的数字,返回时按数组)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `userID` (`userID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限信息表';

-- ----------------------------
-- Table structure for purchasecontract
-- ----------------------------
DROP TABLE IF EXISTS `purchasecontract`;
CREATE TABLE `purchasecontract` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `contractcode` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '采购合同编号',
  `handler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '采购员',
  `currentreviewer` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '当前审核人',
  `currentreviewindex` tinyint(5) DEFAULT NULL COMMENT '当前审核人序号',
  `status` tinyint(5) NOT NULL COMMENT '状态(1:制单;2:审核；3:执行中；4:执行完；5:已结算)',
  `suppcode` varchar(50) CHARACTER SET utf8mb4 NOT NULL COMMENT '供应商编号',
  `type` tinyint(5) NOT NULL COMMENT '源类型(1：配件销售合同；2：维修合同；3：消耗品)',
  `relatedcode` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '关联合同号(消耗品时可以为空)',
  `receiveaddress` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '收货地址',
  `receiver` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '收货人',
  `receiverphone` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '收货人电话',
  `taxsign` tinyint(5) NOT NULL COMMENT '含税标志(1：含税价；2：不含税价)',
  `taxrate` tinyint(5) NOT NULL COMMENT '税率',
  `ralatedinquirycode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '关联询价单号',
  `settlestatus` tinyint(5) NOT NULL COMMENT '结算状态(1:未结算；2:已结算；3:部分结算)',
  `settleamount` bigint(20) DEFAULT '0' COMMENT '已结算金额',
  `amount` bigint(20) NOT NULL COMMENT '金额',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='采购合同表';

-- ----------------------------
-- Table structure for purchasedetail
-- ----------------------------
DROP TABLE IF EXISTS `purchasedetail`;
CREATE TABLE `purchasedetail` (
  `contractcode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '采购合同编号',
  `mattercode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料编码',
  `num` bigint(20) DEFAULT NULL COMMENT '采购数量',
  `price` bigint(20) DEFAULT NULL COMMENT '单价',
  `value` bigint(20) DEFAULT NULL COMMENT '总价',
  PRIMARY KEY (`contractcode`,`mattercode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='采购合同详情表';

-- ----------------------------
-- Table structure for putindetail
-- ----------------------------
DROP TABLE IF EXISTS `putindetail`;
CREATE TABLE `putindetail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `incode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '入库单编号(同putinstore的incode字段)',
  `mattercode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料编码',
  `realnum` bigint(20) NOT NULL COMMENT '入库数量',
  `num` bigint(20) NOT NULL COMMENT '到货单数量',
  `price` bigint(20) NOT NULL COMMENT '单价',
  `value` bigint(20) NOT NULL COMMENT '总价',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='入库详情表';

-- ----------------------------
-- Table structure for putinstore
-- ----------------------------
DROP TABLE IF EXISTS `putinstore`;
CREATE TABLE `putinstore` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `incode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '入库单编号',
  `warehouseid` bigint(20) NOT NULL COMMENT '仓库id',
  `source` int(5) NOT NULL COMMENT '来源(1：采购到货单；2：调拨入库；3：盘亏盘盈)',
  `relatedcode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '关联单据号',
  `indate` datetime NOT NULL COMMENT '入库时间',
  `storehandler` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '入库操作人(cardid)',
  `purchasehandler` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '采购处理人(cardid)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='物料入库表';

-- ----------------------------
-- Table structure for repaircost
-- ----------------------------
DROP TABLE IF EXISTS `repaircost`;
CREATE TABLE `repaircost` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `type` tinyint(5) NOT NULL COMMENT '费用类型(1：物料;2:人工；3：其他；4：外协)',
  `extend` text CHARACTER SET utf8mb4 COMMENT '扩展信息(json;如人工的出发和返回时间)',
  `unit` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '单位',
  `num` bigint(20) NOT NULL COMMENT '数量',
  `price` bigint(20) NOT NULL COMMENT '单价',
  `remark` text CHARACTER SET utf8mb4 COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='维修项费用表';

-- ----------------------------
-- Table structure for repairitem
-- ----------------------------
DROP TABLE IF EXISTS `repairitem`;
CREATE TABLE `repairitem` (
  `contractcode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '合同编号',
  `itemname` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '项目名称',
  `status` tinyint(5) NOT NULL COMMENT '状态(3:进行中；4：完工)',
  `fault` text CHARACTER SET utf8mb4 COMMENT '故障现象',
  `causeanalysis` text CHARACTER SET utf8mb4 COMMENT '原因分析',
  `measures` text CHARACTER SET utf8mb4 COMMENT '修复措施',
  `vehiclecode` varchar(20) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`contractcode`,`itemname`,`vehiclecode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='车辆维修表';

-- ----------------------------
-- Table structure for review
-- ----------------------------
DROP TABLE IF EXISTS `review`;
CREATE TABLE `review` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `type` tinyint(5) DEFAULT NULL COMMENT '类型(1:采购合同审核)',
  `detail` varchar(1000) CHARACTER SET utf8mb4 NOT NULL COMMENT '详情([{end:0, cardids:['''']}])',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='审核流程表';

-- ----------------------------
-- Table structure for reviewresult
-- ----------------------------
DROP TABLE IF EXISTS `reviewresult`;
CREATE TABLE `reviewresult` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '类型(1:采购合同审核)',
  `type` tinyint(5) DEFAULT NULL,
  `relatedcode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '关联单号',
  `reviewer` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '审核人工号',
  `opinion` text CHARACTER SET utf8mb4 COMMENT '审核意见',
  `result` tinyint(5) NOT NULL COMMENT '1:审核通过;2:驳回',
  `reviewtime` datetime DEFAULT NULL COMMENT '审核时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `type` (`type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='审核结果表';

-- ----------------------------
-- Table structure for saledetail
-- ----------------------------
DROP TABLE IF EXISTS `saledetail`;
CREATE TABLE `saledetail` (
  `contractid` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '合同编号',
  `mattercode` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '物料编码',
  `num` bigint(20) NOT NULL,
  `price` bigint(20) NOT NULL COMMENT '单价',
  PRIMARY KEY (`contractid`,`mattercode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='销售详情表';

-- ----------------------------
-- Table structure for stock
-- ----------------------------
DROP TABLE IF EXISTS `stock`;
CREATE TABLE `stock` (
  `warehouseid` bigint(20) NOT NULL COMMENT '仓库id',
  `mattercode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '物料编码',
  `num` bigint(20) NOT NULL COMMENT '数量',
  `averageprice` bigint(20) NOT NULL COMMENT '均价',
  PRIMARY KEY (`warehouseid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='库存表';

-- ----------------------------
-- Table structure for supplier
-- ----------------------------
DROP TABLE IF EXISTS `supplier`;
CREATE TABLE `supplier` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `suppcode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '供应商编号',
  `name` varchar(200) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '供应商名称',
  `address` varchar(500) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '地址',
  `zipcode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '邮编',
  `fax` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '传真',
  `website` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '官网',
  `depositbank` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '开户行',
  `bankaccount` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '银行账号',
  `taxrate` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '税率',
  `paymethod` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '付款方式',
  `taxnum` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '税号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `suppcode` (`suppcode`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='供应商信息表';

-- ----------------------------
-- Table structure for supplyrelation
-- ----------------------------
DROP TABLE IF EXISTS `supplyrelation`;
CREATE TABLE `supplyrelation` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `supplierid` bigint(20) NOT NULL COMMENT '供应商主键',
  `matterid` bigint(20) NOT NULL COMMENT '物料主键',
  PRIMARY KEY (`id`,`supplierid`,`matterid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='供货关系表';

-- ----------------------------
-- Table structure for vehicle
-- ----------------------------
DROP TABLE IF EXISTS `vehicle`;
CREATE TABLE `vehicle` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `custcode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '客户编码',
  `vehiclecode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '车辆编码',
  `name` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '车辆名称',
  `type` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '车型',
  `line` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '线路',
  `productdate` datetime DEFAULT NULL COMMENT '出厂日期',
  `manufacturer` varchar(200) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '生产厂家',
  `remark` varchar(1000) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='车辆信息表';

-- ----------------------------
-- Table structure for warehouse
-- ----------------------------
DROP TABLE IF EXISTS `warehouse`;
CREATE TABLE `warehouse` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '仓库名称',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='仓库信息表';
