# Go博客系统

## 项目结构：
```
📦go_task_4
 ┣ 📂api
 ┣ 📂assets
 ┣ 📂build
 ┣ 📂configs                    #配置文件
 ┃ ┗ 📜setting.yaml
 ┣ 📂deployments
 ┣ 📂docs                       #swagger自动生成docs
 ┃ ┣ 📜docs.go
 ┃ ┣ 📜swagger.json
 ┃ ┗ 📜swagger.yaml
 ┣ 📂examples
 ┣ 📂githooks
 ┣ 📂init                       #初始化
 ┃ ┣ 📜initdatabase.go
 ┃ ┗ 📜initviperconfig.go
 ┣ 📂internal                   #项目内部包
 ┃ ┣ 📂handler                  #gin框架 http处理器逻辑
 ┃ ┃ ┣ 📜comment_handler.go
 ┃ ┃ ┣ 📜post_handler.go
 ┃ ┃ ┗ 📜user_handler.go
 ┃ ┣ 📂model                    #数据模型
 ┃ ┃ ┣ 📜comment.go
 ┃ ┃ ┣ 📜pageresult.go
 ┃ ┃ ┣ 📜post.go
 ┃ ┃ ┗ 📜user.go
 ┃ ┣ 📂repository               #gorm框架数据库模型逻辑
 ┃ ┃ ┣ 📜comment_repository.go
 ┃ ┃ ┣ 📜post_repository.go
 ┃ ┃ ┗ 📜user_repository.go
 ┃ ┣ 📂response                 #响应模型
 ┃ ┃ ┗ 📜response.go
 ┃ ┣ 📂router                   #路由
 ┃ ┃ ┗ 📜router.go
 ┃ ┗ 📂service                  #业务数据处理逻辑
 ┃ ┃ ┣ 📜comment_service.go
 ┃ ┃ ┣ 📜post_service.go
 ┃ ┃ ┗ 📜user_service.go
 ┣ 📂logfiles                   #日志文件记录目录
 ┃ ┣ 📜2025-10-06-16.log
 ┃ ┗ 📜2025-10-10-13.log
 ┣ 📂pkg                        #项目公共包
 ┃ ┣ 📂utils                    #工具包
 ┃ ┃ ┣ 📜JWTutil.go
 ┃ ┃ ┣ 📜encryptutil.go
 ┃ ┃ ┣ 📜qpageutil.go
 ┃ ┃ ┣ 📜validationutil.go
 ┃ ┃ ┗ 📜zaplogutil.go
 ┃ ┗ 📜.DS_Store
 ┣ 📂scripts
 ┣ 📂test                       #测试包
 ┃ ┗ 📜test.go
 ┣ 📂third_party
 ┣ 📂tools
 ┣ 📂web
 ┣ 📂website
 ┣ 📜.DS_Store
 ┣ 📜app.go                     #应用程序入口
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┗ 📜readme.txt
```

## 功能特性

- 用户管理：注册、登录、个人资料管理
- 文章管理：创建、编辑、删除、查看文章
- 评论系统：发表评论、回复评论

## 技术栈

- Go语言
- Gin Web框架
- MySQL数据库
- gorm库
- JWT认证
- Viper配置管理
- zap日志库
- swagger接口文档库

## 安装依赖库： [根目录下 eg：../go_task_4/ 进入到这个根目录下]
        go get gorm.io/gorm                 ---安装gorm库
        go get gorm.io/driver/mysql         ---安装gorm库的mysql数据库驱动
        go get github.com/spf13/viper       ---安装 服务器配置管理库
        go get go.uber.org/zap              ---安装 日志库 zap库 适用于高并发高性能的服务 主流选择
        go get -u github.com/gin-gonic/gin  ---安装 gin web框架
                
                MACOS 配置JWT生成token时需要的盐（bashrc/zshrc）
                终端输入：
                echo 'export JWT_SECRET="123456789"' >> ~/.bashrc
                echo 'export JWT_SECRET="123456789"' >> ~/.zshrc
                source ~/.bashrc
                source ~/.zshrc
                检验：
                echo $JWT_SECRET
                
        
        go install github.com/swaggo/swag/cmd/swag@latest       ---swagger接口文档库
                安装好之后配置环境变量 macos为例：
                查看 GOPATH
                go env GOPATH
                把 GOPATH/bin 加入 PATH（zsh）
                echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
                source ~/.zshrc
                验证
                swag --version
                or
                which swag
        go get github.com/swaggo/gin-swagger@latest
        go get github.com/swaggo/files@latest

