# GoAutoGPT

[English Version README](./README_EN.md)

[中文版 README](README.md)

## Usage

Using ChatGPT to implement AI programming, input a name to automatically complete architecture design, function decomposition, and code programming.

Usage:

1. `git clone git@github.com:NiuStar/GoAutoGPT.git`

2. `cd GoAutoGPT`

3. `go build`

4. Modify the `UserId` and `src` in the config file under the `config` directory. The `userId` is the unique identifier of the user you applied for, and `src` is the directory where the code is generated.

Currently, the default configuration in `config` includes a test account with `userId` 1. If you need another account, please email me at yjkj2@qq.com to communicate.

5. `./GoCodeGPT create -p WashingMachineReservationManagementSystem -d This is a brief description of the system for testing purposes`

It takes the time for a cup of coffee. If the creation is successful, you will see the generated code in `$src`.

If an error occurs, please use:

`./GoCodeGPT generate -p 'projectId you want to generate'`

The `projectId` can be found by running:

`./GoCodeGPT list`

Find the UUID of the project you want to generate, and the `nameEn` is the code directory name in `$src`.

We have switched to the cloudcone server. If you cannot ping 198.211.18.242, please use 🛫.

## Future Plans:

1. More user-friendly interaction methods

2. More new features

3. AI programming plans for web, Android, and iOS
