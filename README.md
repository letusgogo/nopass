<p align="center">
    <picture>
        <img alt="NoPass" title="NoPass" src="assets/img/logo.png">
    </picture>
    [English](README.md) · [简体中文](README.zh-Hans.md)
</p>


# nopass
We should generate a password instead of remembering it.

# NoPass

NoPass is a password generator that avoids memorizing or storing passwords.
Instead, it generates strong and hard-to-crack passwords by remembering
some information from daily life (such as birthdays, phone numbers, email addresses, names of relatives, etc.)
and password generation rules.
Users only need to remember this information to easily reproduce the same strong passwords without worrying about password leaks.

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
