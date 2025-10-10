1 根目录下 eg：../go_task_4/ 进入到这个根目录下
    安装依赖库：
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

        swagger接口文档库
        go install github.com/swaggo/swag/cmd/swag@latest
        安装好之后配置环境变量 macos为例：
        # 查看 GOPATH
        go env GOPATH
        # 把 GOPATH/bin 加入 PATH（zsh）
        echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
        source ~/.zshrc
        # 验证
        swag --version
        # or
        which swag

        go get github.com/swaggo/gin-swagger@latest
        go get github.com/swaggo/files@latest

