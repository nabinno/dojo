---
title: システム管理で使えるデータ分析ハンズオン：システム構成情報の収集と可視化
tags: amazon-quicksight, aws-systems-manager, 
url: -
---

# 1. セッション概要
対象者
- AWS環境の管理者、セキュリティ担当者、クラウド推進者

ゴール
- AWS Systems ManagerのInventory機能を用いてシステム構成情報を収集できる
- Amazon QuickSightを用いてインベントリ情報を可視化できる

課題リスト
- サーバの最新状態を正確にに把握したい
- 脆弱性のあるバージョンのソフトウェアがインストールされていないかすばやく確認したい

システム構成情報
当セッションではOS、導入ソフトウェア、パッチ、HW情報などのサーバ単位のシステム構成情報を取得する

# 2. IAMを作成する
1. ロール作成を選択
2. EC2の選択
3．アタッチするアクセス権限「AmazonSSMManagedInstanceCore」を選択
4. ロール名「Innovate2020-SSMAgent」で作成

# 2. 環境を準備する
## Amazon Linux
1. EC2に移動
2. インスタンスの起動を選択
3. Amazon Linuxを選択
4. インスタンスタイプはデフォルト
5. インスタンスにてIAMロール「Innovate2020-SSMAgent」を設定
6. ストレージの追加はデフォルト
7. タグの追加では「Name: Innovate2020-Linux01」
8. セキュリティグループの設定では既存のものを選択
9. インスタンスを作成

## Windows Server
1. EC2に移動
2. インスタンスの起動を選択
3. Windows Server 2019 Baseを選択
4. インスタンスタイプはデフォルト
5. インスタンスにてIAMロール「Innovate2020-SSMAgent」を設定
6. ストレージの追加はデフォルト
7. タグの追加では「Name: Innovate2020-Linux01」
8. セキュリティグループの設定では既存のものを選択
9. インスタンスを作成

# 3. AWS Systems Managerでインベントリを設定する
1. セットアップインベントリを選択
2. 名前「Innovate2020-Inventory」
3. 作成


# 4. リソースデータの同期
構成
- SSM Resource Data Syncにてインベントリの情報をS3に適宜出力
- リソースデータの同期 - Glue CrawlerがS3を12時間ごとにクロール

## S3の設定
1. S3にてバケット「innovate2020-ssm-inventory」を作成
2. バケットポリシーを[こちらのページ](https://docs.aws.amazon.com/ja_jp/systems-manager/latest/userguide/sysman-inventory-datasync.html)を参考に設定

```json
{
   "Version":"2012-10-17",
   "Statement":[
      {
         "Sid":"SSMBucketPermissionsCheck",
         "Effect":"Allow",
         "Principal":{
            "Service":"ssm.amazonaws.com"
         },
         "Action":"s3:GetBucketAcl",
         "Resource":"arn:aws:s3:::bucket-name"
      },
      {
         "Sid":" SSMBucketDelivery",
         "Effect":"Allow",
         "Principal":{
            "Service":"ssm.amazonaws.com"
         },
         "Action":"s3:PutObject",
         "Resource":[
            "arn:aws:s3:::bucket-name/bucket-prefix/*/accountid=account-id-1/*"
         ],
         "Condition":{
            "StringEquals":{
               "s3:x-amz-acl":"bucket-owner-full-control"
            }
         }
      }
   ]
}
```

## リソースデータの同期
1. Systems Managerでインベントリを選択
2. リソースデータの同期の作成を選択
3. 同期名「Innovate2020-InventoryResourceDataSync」、バケット名「Innovate2020-InventoryResourceDataSync」で作成

# 5. QuickSightで収集した情報の可視化
構成
- データカタログ -> Athena -> QuickSight

## QuickSightのアカウント作成
1. QuickSightアカウントの作成を選択
2. アカウント名「innovate2020」
3. Choose S3 bucketsにて作成したバケット「innovate2020-ssm-inventory」を選択
4. 完了

## QuickSightでチャートをつくる
1. データセットを選択
2. Athenaを選択
3. データソース名「ssm-inventory」に設定
4. データベース「innovate2020-ssm-inventory-us-east-1-database」を選択
5. テーブル「aws-application」を選択
6. データクエリを直接実行を選択
7. Visualize

# References
- https://pages.awscloud.com/rs/112-TZM-766/images/%5BH-2%5DAWSInnovate_Online_Conference_2020_Spring_handson%28SSMInventory-QuickSight%29_ans.pdf
