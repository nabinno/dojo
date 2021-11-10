---
title: "Windows上にてSSMセッションマネージャープロキシでDataGripからDBに接続する方法"
emoji: "📛"
type: "aws"
topic: ["ssm-session-manager", "aws", "datagrip"]
published: true
---

表題の件がネットに明示されていなかったので改めて書きます。

- [AWS SSM Session Manager Proxy : DBE-9500](https://youtrack.jetbrains.com/issue/DBE-9500)
    - `So I checked the PATH with echo, and there are many missing paths, how should I solve them?`

上記Jetbrainsのトラックを見るに、Authentication Type「 `OpenSSH config and authentication agent` 」はAWS CLIと Session Manager Pluginのパスを読み込めていないようなので、下記対処により（ターミナルでポートフォーワード処理せずに）DataGripだけで完結できます。

1. まずAWS CLIと Session Manager Pluginを環境変数Pathに設定します
2. 次にProxyCommand実行時に
    1. PowerShellの絶対パスで実行
    2. 環境変数PATHを読み込んだ後に `aws ssm start-session` を実行

```
## ssh config設定例
Host test-ssmbastion
  User ec2-user
  port 22
  ProxyCommand C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe "$env:Path = [System.Environment]::GetEnvironmentVariable('Path','Machine') + ';' + [System.Environment]::GetEnvironmentVariable('Path','User'); aws ssm start-session --target i-xxxxxxxx --document-name AWS-StartSSHSession --parameters 'portNumber=%p' --region ap-northeast-1"
  IdentityFile ~/.ssh/id_rsa
```

今回はWindows用に書きましたが、macOSでも環境変数PATHの読み込みを事前に行えば実行できると思います。
