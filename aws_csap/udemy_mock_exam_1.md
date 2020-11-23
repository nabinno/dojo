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

## 11. AWS Storage GatewayとCHAP認証
AWSとオンプレミスサイトを併用したハイブリッドアーキテクチャ
- AWS Storage Gatewayのボリュームゲートウェイ
- iSCSIを介したハイブリッド構成のデータ共有システム
- セキュリティ上の問題
  - データ共有システムのネットワークに対するDDoS攻撃
  - 不正アクセス
  - 不正傍受
=> AWS Storage Gateway, CHAP

## 12. AWS LambdaとAmazon Rekognition API
動物画像検索アプリ (画像識別)
- ユーザーが動物写真などをアップロード、類似した動物を画像検索
- 一連の写真をアップロード、特定の画像が利用された時間を検索
期待
- アプリケーションの開発と運用を低コストに実施
=> AWS Lambda, Amazon Rekognition API, Lambda -> S3

## 13. Amazon Transit Gatewayをつかったセキュリティ構成
AWS上の社内システム
- セキュリティ強化
  - 利用するVPCに侵入検知・防止システムを実装
- システム要件
  - VPC内で実行されている数百にも及ぶインスタンスを拡張できる機能が必要
  - 現在VPCは12個起動
期待
- まとめてモニタリングする効率的な方法
=> AWS Transit Gateway, Amazon VPC, IDS/IPS

## 14. AWS Snowball Edge Storage Optimizedをつかったデータ移行
要件
- オンプレのインフラとアプリケーション全般をAWSクラウドに移行
  - 合計150TBのデータ
    - タイムリーかつ費用対効果の高い方法でS3バケットに移動する必要がある
  - 既存のインターネット接続の空き容量を使用
    - データをAWSにアップロードするのには1週間以上かかると予測
=> AWS Snowball Edge Storage Optimized


## 15. AWSでスケーラビリティと弾力性を高める
3層ウェブアプリケーションとなっているニュースサイト
- オンプレミスでデプロイ
要件
- スケーラビリティと弾力性を高めるためにAWSに移行
  - トラフィック変動に自動的に対応する必要あり
    - Web層とアプリケーション層を組み合わせた読み取り専用のニュースレポートサイト
    - 予測不可能な大規模なトラフィック要求を受け取るデータベース層
=> ELB, AutoScaling, ElastiCache Redis, CloudWatch, RDS Read/Replica

## 16. AWS Security Token Service
モバイルアプリケーション
- ユーザーがS3バケット内のデータを利用する際に一時認証を利用
  - STSを利用して一時的な認証情報を取得しユーザーに渡す
  - ただし一部の一時認証情報のアクセス権限が間違っているため必要なリソースへのアクセスが提供されていないことが判明
要件
- 一時認証によって付与されたアクセス権を取り消す方法
=> AWS IAM, AWS Security Token Service

