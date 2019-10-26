package models

import (
	"erpweb/util"

	"github.com/astaxie/beego"
)

type Menu struct {
	MenuID    int64  `json:"menuID" orm:"column(menuID);pk"`    //菜单id
	Title     string `json:"title" orm:"column(title)"`         //标题
	Icon      string `json:"icon" orm:"column(icon)"`           //图标
	ParentID  int64  `json:"parentID" orm:"column(parentID)"`   //父ID
	Path      string `json:"path" orm:"column(path)"`           //路径
	Component string `json:"component" orm:"column(component)"` //组件名
}

type WebMenu struct {
	MenuID    int64      `json:"menuID"`
	Title     string     `json:"title"`
	Icon      string     `json:"icon"`      //图标
	ParentID  int64      `json:"parentID"`  //父ID
	Path      string     `json:"path"`      //路径
	Component string     `json:"component"` //组件名
	Children  []*WebMenu `json:"children"`
}

func GetMenusByLeafs(permisses map[int64]int) []WebMenu {
	var (
		menus    []Menu
		retMenus []WebMenu
		//tempRetMaps = make(map[int64]WebMenu)
	)
	_, err := OSQL.Raw("select * from " + util.MENU_TABLE_NAME).QueryRows(&menus)
	if err != nil {
		beego.Error(err) //logs.Error(err)
	}

	beego.Info("sql menus=", menus)

	//tempMenus := groupMenus(menus)

	// for _, temp := range tempMenus {
	// 	var tempRet WebMenu
	// 	tempRet = *temp
	// 	retMenus = append(retMenus, tempRet)
	// }

	// for permissionleafId, _ := range permisses {
	// 	// if LeafFindChild(permissionleafId, tempMenus) {

	// 	// }

	// 	for _, findTmp := range tempMenus {
	// 		if LeafFindChild(permissionleafId, findTmp.Children) {
	// 			tempRetMaps[findTmp.MenuID] = *findTmp
	// 		}
	// 	}
	// }

	// for _, value := range tempRetMaps {

	// 	retMenus = append(retMenus, value)
	// }

	beego.Info("该用户所拥有的菜单=", retMenus)

	return retMenus
}

func LeafFindChild(permissionleafId int64, vs []*WebMenu) bool {
	var retBool = false
	for _, v2 := range vs {
		if len(v2.Children) > 0 {
			retBool = LeafFindChild(permissionleafId, v2.Children)
		} else {
			beego.Info("没有子节点了, local menuId=", v2.MenuID, " , need find id=", permissionleafId)
			if v2.MenuID == permissionleafId {
				retBool = true
			}
		}
	}
	return retBool
}

func GetMenus() []WebMenu {
	var (
		menus    []Menu
		retMenus []WebMenu
	)
	_, err := OSQL.Raw("select * from " + util.MENU_TABLE_NAME).QueryRows(&menus)
	if err != nil {
		beego.Error(err) //logs.Error(err)
	}

	beego.Info("sql menus=", menus)

	tempMenus := groupMenus(menus)

	for _, temp := range tempMenus {
		var tempRet WebMenu
		tempRet = *temp
		retMenus = append(retMenus, tempRet)
	}

	return retMenus
}

func groupMenus(params []Menu) (rets []*WebMenu) {
	var tempWeb []WebMenu
	for _, param := range params {
		var tempWebone WebMenu
		tempWebone.MenuID = param.MenuID
		tempWebone.Title = param.Title
		tempWebone.Icon = param.Icon
		tempWebone.ParentID = param.ParentID
		tempWebone.Path = param.Path
		tempWebone.Component = param.Component
		tempWeb = append(tempWeb, tempWebone)
	}
	pdepts := make([]*WebMenu, 0)

	for i, _ := range tempWeb {
		var a *WebMenu
		a = &tempWeb[i]
		pdepts = append(pdepts, a)
	}

	for _, pdept := range pdepts {
		if pdept != nil {
			if pdept.ParentID == 0 {
				makeTree(pdepts, pdept)
			}
		}
	}

	for _, pdept := range pdepts {
		if pdept != nil {
			if pdept.ParentID == 0 {
				rets = append(rets, pdept)
			}
		}
	}

	return rets
}

func has(v1 WebMenu, vs []*WebMenu) bool {
	var has bool
	has = false
	for _, v2 := range vs {
		v3 := *v2
		if v1.MenuID == v3.ParentID {
			has = true
			break
		}
	}
	return has

}

func makeTree(vs []*WebMenu, node *WebMenu) {
	childs := findChild(node, vs)
	for _, child := range childs {
		node.Children = append(node.Children, child)
		if has(*child, vs) {
			makeTree(vs, child)
		}
	}
}

func findChild(v *WebMenu, vs []*WebMenu) (ret []*WebMenu) {
	for _, v2 := range vs {
		if v.MenuID == v2.ParentID {
			ret = append(ret, v2)
		}
	}
	return
}

// func GetMenuByUserID(userID int64) (menu Menu, err error) {
// 	menu.UserID = userID

// 	return menu, nil
// }

// func EditMenuById(id int64) (errorCode int64) {
// 	return errorCode
// }

// func AddMenument(menu Menu) (errorCode int64) {
// 	return errorCode
// }

// func DeleteMenu(id int64) (errorCode int64) {
// 	return errorCode
// }
