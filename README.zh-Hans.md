# NoPass

NoPass 是一个密码生成器，旨在避免记忆或保存密码，而是通过记住一些生活中的信息（如生日、手机号、邮箱、亲人的名字等）
和密码生成规则来生成强大且难以破解的密码。用户只需记住这些信息，即可轻松重现相同的强密码，无需担心密码泄露风险。

## 背景
我不想记忆我的各个网站的密码，特别是一般不同网站的使用不同的密码。
除了我的脑子，哪里都不安全无论是第三方网站，还是一些在线或者离线的笔记。

## 核心功能

1. 输入生活信息：用户可以提供与其生活息息相关的信息，如生日、手机号、邮箱、亲人的名字等，作为密码生成器的基础输入。

2. 配置文件加载：项目支持从 JSON 配置文件中加载参数，如密码长度、字符类型等。

3. 哈希和转换：将输入的生活信息进行哈希和转换操作，生成具有高度随机性的密码。

4. 四种字符类型：生成的密码将包含数字、大写字母、小写字母和特殊符号，以提高其复杂性。

5. 可重现性：当用户使用相同的生活信息时，将始终生成相同的强密码。

## 基本实现思想

NoPass 使用哈希函数将输入的生活信息映射到不同类型的字符。项目将原始密码的每个字符依次映射到小写字母、大写字母、数字和特殊符号，然后循环进行。
通过这种方式，NoPass 可以确保生成的密码具有高度随机性和复杂性，同时避免了密码泄露风险。

## 安装
如果你安装了 go 和 make, 可以使用下面的命令安装。或者你可以从 release 页面下载对应的二进制文件。


mac os:
```bash
go install fyne.io/fyne/v2/cmd/fyne@latest
sudo make install
```

linux:
```bash
go get fyne.io/fyne/v2@latest && go install fyne.io/fyne/v2/cmd/fyne@latest
sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev
sudo make install-linux
```

windows:
```bash
go install fyne.io/fyne/v2/cmd/fyne@latest
make win
```

## 使用
直接执行 nopass gen 或者 nopass 快速开始使用，更多参数可以使用 nopass -h 查看。
```bash
nopass
```
```bash
nopass gen
```

## 配置文件
可以使用 -c 参数指定配置文件。
nopass -c config.yaml

可以自己修改配置文件生成适合自己的规则, 默认使用 default 规则。也可以通过指令指定使用的规则。
nopass gen -r simple
```yaml
rules:
    default:
        - name: luckNum
        hint: please input a fixed number
        - name: webSite
        hint: please input web site like google
        - name: genMonth
        hint: please input the mouth of the password generated on this website, like 202101

    simple:
        - name: luckNum
        hint: please input your luck number, like 618

    difficult:
        - name: luckNum
        hint: please input your luck number, like 618
        - name: webSite
        hint: please input web site like google
        - name: birthday
        hint: please input birthday like 19900101
        - name: email
        hint: please input email like helloworldyong9@gmail
        - name: momName
        hint: please input mom name like Julie
```