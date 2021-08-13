## 平台简介
* 基于Gin的后台管理系统
* 前端采用ruoyi-ui 、Vue、Element UI。
* 后端采用GO语言 框架 Gin。
* 本项目由猴酷团队开发。

## 内置功能

1.  用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2.  部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3.  岗位管理：配置系统用户所属担任职务。
4.  菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5.  角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6.  字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7.  参数管理：对系统动态配置常用参数。

## github地址
[github:https://github.com/druidcaesa/monkey-admin](https://github.com/druidcaesa/monkey-admin)
## 码云地址
[码云https://gitee.com/termites/monkey-admin](https://gitee.com/termites/monkey-admin)

## 演示地址
[http://www.monkeycool.cn](http://www.monkeycool.cn)
账号：admin  密码：admin123
## 配置
项目数据库文件 /data/db.sql 创建数据库导入后修改配置/config/config-*.ini


##运行
go run main.go 直接访问http://localhost:8080

账号：admin  密码：admin123

项目为前后端分离，前端代码在monkey-ui目录下 

##docker镜像构建
[docker安装使用请参考官方](https://www.docker.com/)
1. 根据情况修改Dockerfile文件
2. 在项目根目录下使用命令docker build -t <你要出的进行名>:<版本号> .

## 演示图

![WX20210804-171404.png](https://i.loli.net/2021/08/04/kOY2tHZfAMTKFXQ.png)
![WX20210804-172034.png](https://i.loli.net/2021/08/04/twOd9v1qRP7cZlX.png)
![WX20210804-172645.png](https://i.loli.net/2021/08/04/RUsB8qwlZuaEdDL.png)
![WX20210804-172633.png](https://i.loli.net/2021/08/04/P7swv5g9o3L4JIp.png)
![WX20210804-172615.png](https://i.loli.net/2021/08/04/ng9A1FpmlbcaWJT.png)
![WX20210804-172714.png](https://i.loli.net/2021/08/04/Y4TDF3rq98Akhao.png)
![WX20210804-172735.png](https://i.loli.net/2021/08/04/YbaczLMherSJxR6.png)
![WX20210804-172751.png](https://i.loli.net/2021/08/04/lnHoRfxwESa27DM.png)
![WX20210804-172726.png](https://i.loli.net/2021/08/04/CkMHJvW4RDiQfnj.png)
![WX20210804-172549.png](https://i.loli.net/2021/08/04/NA87nPgubIhatLr.png)
![WX20210804-172802.png](https://i.loli.net/2021/08/04/JLZ2YbHuq3BDxUh.png)
![WX20210804-172702.png](https://i.loli.net/2021/08/04/nclOa39fs4AxhUP.png)


## 感谢(排名不分先后)
> Gin框架 [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
>
> gotool[https://github.com/druidcaesa/gotool](https://github.com/druidcaesa/gotool)
> 
> RuoYi-Vue [https://gitee.com/y_project/RuoYi-Vue](https://gitee.com/y_project/RuoYi-Vue)
>
>jwt [https://github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
>
>excelize [https://github.com/qax-os/excelize](https://github.com/qax-os/excelize)
>
>xorm [https://github.com/go-xorm/xorm](https://github.com/go-xorm/xorm)

## 感觉对自己学习和工作有用的，麻烦给一个小小的star

## 免责声明：
> 1、monkey-admin仅限自己学习使用，一切商业行为与monkey-admin无关。

> 2、用户不得利用monkey-admin从事非法行为，用户应当合法合规的使用，发现用户在使用产品时有任何的非法行为，monkey-admin有权配合有关机关进行调查或向政府部门举报，monkey-admin不承担用户因非法行为造成的任何法律责任，一切法律责任由用户自行承担，如因用户使用造成第三方损害的，用户应当依法予以赔偿。

> 3、所有与使用monkey-admin相关的资源直接风险均由用户承担。 
