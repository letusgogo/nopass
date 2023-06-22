<p align="center">
    <picture>
        <img alt="NoPass" title="NoPass" src="assets/img/logo.png">
    </picture>
    <p align="center">
            <a href="README.md">English</a>
            ·
            <a href="README.zh-Hans.md">简体中文</a>
   </p>
</p>

# nopass
We should generate a password instead of remembering it.

# NoPass

NoPass is a password generator that avoids memorizing or storing passwords.
Instead, it generates strong and hard-to-crack passwords by remembering
some information from daily life (such as birthdays, phone numbers, email addresses, names of relatives, etc.)
and password generation rules.
Users only need to remember this information to easily reproduce the same strong passwords without worrying about password leaks.

## Background

I don't want to remember the passwords for my various websites, especially since different websites generally use different passwords.

Nowhere but my brain is safe from third-party websites or some notes online or offline.

## Core Features

1. Input daily life information: Users can provide information closely related to their lives,
such as birthdays, phone numbers, email addresses, names of relatives, etc., as the basic input for the password generator.

2. Configuration file loading: The project supports loading parameters from a YAML configuration file, such as password length and character types.

3. Hashing and transformation: Perform hashing and transformation operations on the input daily life information to generate highly random passwords.

4. Four character types: The generated password will contain numbers, uppercase letters, lowercase letters, and special symbols to increase its complexity.

5. Reproducibility: When users use the same daily life information, the same strong passwords will always be generated.

## Basic Implementation Idea

NoPass uses a hashing function to map the input daily life information to different types of characters.
In this way, NoPass ensures that the generated password has high randomness and complexity while avoiding the risk of password leaks.

## Installation
If you have installed Go and Make, you can use the following commands to install the program. Alternatively, you can download the corresponding binary files from the release page.
macOS:
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
## Usage
Run nopass gen or nopass directly to get started quickly. For more options, use nopass -h.
```bash
nopass
```
```bash
nopass gen
```
## Configuration File
You can use the -c option to specify the configuration file.
```bash
nopass -c config.yaml
```
You can modify the configuration file to create rules that suit your needs. The default rule is default.
You can also specify the rule to be used with the command.
```bash
nopass gen -r simple
```

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