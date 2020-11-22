---
title: Udemy Mock Exam 1
url: https://www.udemy.com/course/aws-53225/learn/quiz/4723906#overview
tags: aws-csap, amazon-web-services
---

# 1st time
## 1. ソリューションアーキテクチャ
フロー
- 子プロセスを作成
- 画像識別による一時的な評価と人の目による二次的評価を実施
- 親プロセスに評価結果を返信するプロセスが必要
=> AWS Step Functions

## 2. 移行方法とライセンス管理方法
- DB2、SAP、Windowsオペレーティングシステムサーバーなどのライセンスを使用
- これらのライセンスごとAWSに移行する予定
=> RDS, Windows AMI, AWS License Manager

## 3. 権限
- サードパーティのWEBアプリケーションを使用したい
- アカウント内で実行されているEC2インスタンスへのAPIコマンドを発行するアクセス権が必要
- 第三者による目的外利用ができない形式で権限を付与することが必須
=> AWS IAM, Amazon API Gateway

## 4. EC2のアクセス制限
- 社内用業務WEBアプリケーション
- パブリックサブネットに設置されたWebサーバー、EC2インスタンスで構成
- オープンなインターネットから誰でもアクセスできるようにはしたくはない
- パッチ更新のために
  - インバウンドリクエストを受ける
  - 特定のURLへのアウトバウンドリクエストのみに限定
=> Security Group

## 5. 最もコスト効率がよく最適なアーキテクチャ
ニュースメディア配信アプリケーションの構成
- 東京リージョン
- EC2インスタンス
  - ワイヤレスセンサーの管理サーバー
    - Javaフロントエンド
  - Linux バックエンドサーバー ... AZ障害により停止
	- Javaバックエンド
	- MySQLデータベース
    - 1つのAZしか利用していない（NATゲートウェイも同じ）
	- インターネットに接続してパッチをダウンロードすることが必要
	  - セキュリティ上の理由からジャンプホストへのSSHポートのみを開くことが要件
=> Multi-AZ, NAT Gateway, ELB, Bastion host, Security Gropu

## 6. IPv6の設定
WEBサービスの構成
- IPv4 CIDR（10.0.0.0/16）のVPC内に設置
  - 2つのパブリックサブネット
    - プライベートサブネットからのインターネットへの返信処理のためにNATゲートウェイを配置
  - 2つのプライベートサブネット
- データベース用のEC2インスタンス
  - NATゲートウェイに関連付けたルートテーブルを持つプライベートサブネットに配置
社内方針
- IPv6によるIP管理が実施
- 上記アプリケーションでも今後はIPv6を利用した構成に変更することが求められている
=> Egress-only Internet Gateway, Amazon VPC

## 7. AWSへの移行
現状
- WEBアプリケーション
  - オンプレミス環境
  - VMware環境内の高度にカスタマイズされたWindows VM上でアプリケーション実行
  - オンプレミスサーバーの老朽化
期待
- 新サーバーに切り替えるタイミングでAWSに移行
制約
- この移行は週末の土日2日間で実施
- 効率的で迅速な対応が不可欠
=> AWS Server Migration Service

## 8. データベース移行方法
現状
- ニュースメディア配信アプリケーション
  - EC2インスタンス
  - ELB
  - Auto-Scalingグループ
  - MySQL
期待
- MySQLからPosgreSQLへとデータベースを移行
=> AWS Database Migration Service

## 9. AWS Lambda
新規事業用アプリケーション
- RDS MySQL
  - 多数の顧客の基本情報
  - これまでの売買記録
  - 分析などに利用予定
- Lambda
  - RDSのコネクション接続
=> AWS Lambda, RDS Proxy

## 10. ロケーションベースのアラート機能
グローバルな国際決済サービス
- iOSおよびAndroidモバイル
- 飲食店クーポンモバイルアプリ
期待
- ロケーションベースのアラート機能を追加
  - GPSを利用して近隣店舗に近づいた際に
    - その店舗を紹介するチャットボットによるレコメンデーション
    - クーポン提示
=> Amazon DynamoDB, Amazon EC2, Amazon SQS, AutoScaling, AWS Lambda, Amazon Lex 
