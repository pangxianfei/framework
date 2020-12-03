

## 关于 Tmaic-go 框架

软件简介

tmaic-go 是一套简洁、GoLang API Web开发框架(GoLang API Web Framework)。它可以让你从面条一样杂乱的代码中解脱出来；它可以帮你构建一个完美的网络应用，而且每行代码都可以简洁、富于表达力。 
## Docs
#https://tmaic.com
#生成表

	db := l.DB() // *grom.DB
	
	article := models.Article{}

	db.Migrator().CreateTable(&article)

#添加数据（记录）

	db := l.DB()
	
	article := models.Article{Title: "这是标题",Body: "这是内容",Slug: 12345}

	db.Create(&article) // 通过数据的指针来创建
	
	操作方法见：https://gorm.io/zh_CN/docs/create.html
	
	
	
	
## Thanks
* gin
* gorm
* validator.v9
* viper
* big
* jwt
* i18n
* urfave/cli
* fatih/color
* golang/protobuf
* nsqio/go-nsq
* robfig/cron
* ztrue/tracerr
* go-redis/redis
* getsentry/raven-go
* iancoleman/strcase
* gorilla/websocket

## License
This project is licensed under the [MIT license](https://github.com/pangxianfei/framework/blob/main/LICENSE).

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

* https://gitee.com/pangxianfei/tmaic-go/issues
* pangxianfei@icloud.com
