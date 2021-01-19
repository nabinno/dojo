---
title: Udemy Mock Exam 1
tags: aws-csap, amazon-web-services
url: https://www.udemy.com/course/aws-53225/learn/quiz/4723906#overview
---

# 1st time
## [x] 1. Amazon SWF
フロー
- 子プロセスを作成
- 画像識別による一時的な評価と人の目による二次的評価を実施
- 親プロセスに評価結果を返信するプロセスが必要
=> Amazon SWF

## [x] 2. AWS Server Migration Service, Amazon EC2, AWS License Manager
- DB2、SAP、Windowsオペレーティングシステムサーバーなどのライセンスを使用
- これらのライセンスごとAWSに移行する予定
=> AWS Server Migration Service, Amazon EC2, AWS License Manager

## [x] 3. AWS IAM (cross account)
WEBアプリケーションを使用したい
- アカウント内で実行されているEC2インスタンスへのAPIコマンドを発行するアクセス権が必要
- 第三者による目的外利用ができない形式で権限を付与することが必須
=> AWS IAM (cross account)

## [x] 4. Proxy Server, Amazon VPC
社内用業務WEBアプリケーション
- パブリックサブネットに設置
- EC2インスタンス
要件
- オープンなインターネットから誰でもアクセスできるようにはしたくはない
  - パッチ更新のために
    - インバウンドリクエストを受ける
    - 特定のURLへのアウトバウンドリクエストのみに限定
=> Proxy Server, Amazon VPC

## [x] 5. 最もコスト効率がよく最適なアーキテクチャ - Multi-AZ, NAT Gateway, ELB, Bastion host, Security Gropu
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

## [x] 6. IPv6の設定 - Internet Gateway -> Amazon VPC (subnet/IPv6 CIDR, Network ACL) -> EC2 (security group/IPv6 CIDR)
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
=> Internet Gateway -> Amazon VPC (subnet/IPv6 CIDR, Network ACL) -> EC2 (security group/IPv6 CIDR)

## [x] 7. AWSへの移行
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
=> Amazon EC2 (VM Import/Export)

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

## [x] 10. ロケーションベースのアラート機能 - Amazon DynamoDB, Amazon EC2, Amazon SQS, AutoScaling, AWS Lambda, Amazon Lex 
グローバルな国際決済サービス
- iOSおよびAndroidモバイル
- 飲食店クーポンモバイルアプリ
期待
- ロケーションベースのアラート機能を追加
  - GPSを利用して近隣店舗に近づいた際に
    - その店舗を紹介するチャットボットによるレコメンデーション
    - クーポン提示
=> Amazon DynamoDB, Amazon EC2, Amazon SQS, AutoScaling, AWS Lambda, Amazon Lex 

## [x] 11. AWS Storage GatewayとCHAP認証
AWSとオンプレミスサイトを併用したハイブリッドアーキテクチャ
- AWS Storage Gatewayのボリュームゲートウェイ
- iSCSIを介したハイブリッド構成のデータ共有システム
- セキュリティ上の問題
  - データ共有システムのネットワークに対するDDoS攻撃
  - 不正アクセス
  - 不正傍受
=> AWS Storage Gateway, CHAP

## [x] 12. AWS LambdaとAmazon Rekognition API
動物画像検索アプリ (画像識別)
- ユーザーが動物写真などをアップロード、類似した動物を画像検索
- 一連の写真をアップロード、特定の画像が利用された時間を検索
期待
- アプリケーションの開発と運用を低コストに実施
=> AWS Lambda, Amazon Rekognition API, Lambda -> S3

## [x] 13. Amazon Transit Gatewayをつかったセキュリティ構成
AWS上の社内システム
- セキュリティ強化
  - 利用するVPCに侵入検知・防止システムを実装
- システム要件
  - VPC内で実行されている数百にも及ぶインスタンスを拡張できる機能が必要
  - 現在VPCは12個起動
期待
- まとめてモニタリングする効率的な方法
=> AWS Transit Gateway, Amazon VPC, IDS/IPS

## [x] 14. AWS Snowball Edge Storage Optimizedをつかったデータ移行
要件
- オンプレのインフラとアプリケーション全般をAWSクラウドに移行
  - 合計150TBのデータ
    - タイムリーかつ費用対効果の高い方法でS3バケットに移動する必要がある
  - 既存のインターネット接続の空き容量を使用
    - データをAWSにアップロードするのには1週間以上かかると予測
=> AWS Snowball Edge Storage Optimized

## [x] 15. AWSでスケーラビリティと弾力性を高める
3層ウェブアプリケーションとなっているニュースサイト
- オンプレミスでデプロイ
要件
- スケーラビリティと弾力性を高めるためにAWSに移行
  - トラフィック変動に自動的に対応する必要あり
    - Web層とアプリケーション層を組み合わせた読み取り専用のニュースレポートサイト
    - 予測不可能な大規模なトラフィック要求を受け取るデータベース層
=> ELB, AutoScaling, ElastiCache Memcached, CloudWatch, RDS Read/Replica

## [x] 16. AWS Security Token Service
モバイルアプリケーション
- ユーザーがS3バケット内のデータを利用する際に一時認証を利用
  - STSを利用して一時的な認証情報を取得しユーザーに渡す
  - ただし一部の一時認証情報のアクセス権限が間違っているため必要なリソースへのアクセスが提供されていないことが判明
要件
- 一時認証によって付与されたアクセス権を取り消す方法
=> AWS IAM, AWS Security Token Service

## [x] 17. DDos攻撃対策 - Amazon Route 53 (shuffle sharding, anycast routing) AWS Shield Advanced, AWS WAF
社内アプリケーションに対するDDoS攻撃によって大規模なシステム障害が発生
- DDoS攻撃などの外部攻撃を軽減
- 具体的に防止するべき攻撃リスト
  - DDoS攻撃
  - SYNフラッド
  - UDPリフレクション攻撃
  - SQLインジェクション
  - クロスサイトスクリプティング
  - 不正IP取得によるアカウントアクセス
  - ネットワーク情報の取得
=> Amazon Route 53 (shuffle sharding, anycast routing)
=> AWS Shield Advanced
=> AWS WAF

## [x] 18. AWS CloudFormation - AWS CloudFormation DeletionPolicy (S3 retain), Amazon RDS (snapshot)
S3とRDSを利用したデータ共有アプリケーション
- CloudFormationテンプレートを利用
- 画像をS3バケットに保存
- RDSに顧客データを記録
要件
- サービスを停止
  - いつでも再開できるように準備が必要
  - インフラを終了
  - 同時にデータを保持
=> AWS CloudFormation DeletionPolicy (S3 retain), Amazon RDS (snapshot)

## [x] 19. 誤操作対策 - AWS IAM, Amazon EC2 (tag), Amazon VPC
AWS上にエンタープライズシステム
- システムが突如停止するという障害が発生
  - 一人のエンジニアが本番環境のEC2インスタンスを誤って終了
  - 実稼働するアプリケーションにアクセスできる開発者が多数存在
=> AWS IAM, Amazon EC2 (tag), Amazon VPC

## [x] 20. Amazon CloudFrontの配置 - AWS Certificate Manager, Amazon CloudFront (Cache-control max-age directive)
美術鑑賞向けSNSサービス「PINTORアプリケーション」
- CloudFrontディストリビューションを使用
  - コンテンツ読み込み時間を短縮
- 配信されるデータは個人情報も多い
要件
- アプリケーションからユーザーへのCloudFront配信においてHTTPS通信を実施
- CloudFrontのビューアリクエストの割合を増やすことにより
  - パフォーマンスを改善
  - コストを抑える
=> AWS Certificate Manager, Amazon CloudFront (Cache-control max-age directive)

## [x] 21. Amazon EC2のパフォチュー - Amazon EC2 (Placement Group, T3, Enhanced Networking)
顧客管理向けのJavaアプリケーション
- WEBサーバーにEC2
  - 約40％のCPU使用率に相当する一定のワークロード
  - オンデマンドEC2インスタンスのフリート
  - インスタンスは複数利用
    - 通信を最適化することが求められている
- RDSに顧客の構成情報データ
=> Amazon EC2 (Placement Group, T3, Enhanced Networking)


## [x] 22. SSL証明書の適用 - AWS Certificate Manager, Elastic Load Balancer
モバイルアプリケーション
- EC2インスタンス
  - AutoScalingグループ
  - ELB
- セキュリティポリシー
  - インスタンスからVPC内の他のサービスへの全てのアウトバウンド接続について
    - インスタンスアクセス時に一意のSSL認証が利用される必要がある
=> AWS Certificate Manager, Elastic Load Balancer

## [x] 23. VPCエンドポイントをつかったS3との連携 - Amazon EC2, Amazon S3 (bucket policy), VPC Endpoint, NAT Gateway
顧客管理システム
- AWSパブリッククラウド
- ２層アプリケーション
- EC2
  - データ処理サーバー
  - S3
    - データ保存と管理
    - S3との間において毎秒5 Gbpsを超えるデータを送信
    - プライベートサブネットのアプリケーションレイヤーからS3にデータを転送
  - インスタンスの処理にはサードパーティーのソフトウェアが利用
    - 定期的にソフトウェアに対するパッチ更新が必要
=> Amazon EC2, Amazon S3 (bucket policy), VPC Endpoint, NAT Gateway

## [x] 24. データ共有システム - Amazon EC2 (Auto Scaling, Spot Instance), Amazon SQS, Amazon S3 Glacier
社内データ共有システム
- データセンターにホストして運用
- 社内データ
  - データセンターのストレージに保存
  - 中長期保存用のため迅速なデータ抽出は必要ない
  - データ処理のためにOSSメッセージングシステムを利用したジョブ管理を行っている
  - データはテープライブラリによってアーカイブされる構成をオンプレミスで実施
要件
- これらのシステムをAWSに移行
=> Amazon EC2 (Auto Scaling, Spot Instance), Amazon SQS, Amazon S3 Glacier

## [x] 25. IDフェデレーション - AWS IAM (SAML Id Provider), AWS SSO, AWS Security Token Service
エンタープライズシステムのデータセンター
- AWSクラウドに拡張するハイブリッドクラウドインフラストラクチャ
  - オンプレミス側とクラウド側で2つの個別のログインアカウントを持つ
    - 複数の資格情報を保存することを避ける必要がある
要件
- AWSリソースを管理する構成
  - 社内アカウントを使用して既にサインインしているオンプレミスユーザーが
    - 個別のIAMユーザーを作成しないこと
=> AWS IAM (SAML Id Provider), AWS SSO, AWS Security Token Service

## [x] 26. ECSのオートスケーリング
エンタープライズアプリケーション
- Amazon ECSを使用したDockerベースの
- マルチAZ構成でリードレプリカを持つRDS MySQL
  - 高可用でスケーラブル
要件
- アプリケーション層にスケーラビリティを確保
  - ECSクラスターに対するオートスケーリング設定を実施
=> AWS CloudWatch (Capacity Provider Reservation), AWS Auto Scaling, Amazon ECS (AWS ECS Cluster Auto Scaling)

## [x] 27. AWS Resource Access Manager
社内の統合管理のために全社共通のIT運用部門
- AWS Organizationsを使用
  - マルチアカウントおよびマルチリージョンを管理
  - AWSアカウントAとAWSアカウントBとAWSアカウントC
要件
- クロスアカウント処理が必要となる定期タスクの自動化設定
  - アカウントAのユーザーがアカウントBのEC2インスタンスへのアクセスを定期的に実施
=> AWS Resource Access Manager (enable-sharing-with-aws-organizations), AWS Organizations (trusted access)

## [x] 28. 安全なECS環境 - AWS Secrets Manager, AWS KMS, AWS IAM, Amazon ECS
CI/CD環境
- 開発環境などは全てDocker
- Fargate起動タイプを使用するAmazon ECSクラスター
要件
- 社内製品を販売するためのECサイトを構築
- ECSでの実装
  - 環境変数を使用
- 顧客データベースの資格情報をECサイトに提供する必要があり
  - セキュリティを徹底
    - 資格情報がデータ保持とイメージ転送が安全であることが保障
    - かつクラスター自体で表示できないように
=> AWS Secrets Manager, AWS KMS, AWS IAM, Amazon ECS

## [x] 29. VPN
ウェブベースの会計アプリケーション
- フロントサーバー群はAWSのパブリックサブネット上で利用
- 社内のネットワークからのみAWSサイト間VPN接続によって利用
要件
- 会社ではSOHOを推進
  - 外部Wifiがある環境であればどこからでもリモートで接続して作業ができる機能を実装
	- 下記事案に関してセキュリティ性能をできる限り高める
      - 外部からのアクセスが頻繁に発生すること
      - 機密性の高いデータを扱っていること
=> Amazon VPC (Private Subnet <- NAT Gateway <- Public Subnet), VPN (OpenVPN)

## [x] 30. AWS Directory Service - AWS Directory Service <-> MS Active Directory, AWS Managed Microsoft AD (SSO)
オンプレミス環境
- 以前からMicrosoft Active Directoryを使用
- すべての従業員アカウントとデバイスを管理
要件
- AWSクラウドを利用したハイブリッドアーキテクチャを採用することを決定
  - AWS Directory Serviceの設定を行うことが必要
    - 既存のWindowsアカウントパスワードを使用して様々なAWSリソースに接続して使用
    - 新規にAWSにおいてIAM管理を実施することは非効率
=> AWS Directory Service <-> MS Active Directory, AWS Managed Microsoft AD (SSO)

## [x] 31. AWS VPN CloudHub - AWS Direct Connect, AWS Site-to-Site VPN (IPSec, private link)
AWSをクラウドソリューションとして導入
- AWSとオフィスネットワークとを接続
- リモートネットワークをAmazon VPC環境に接続するための接続設定を実施
社内要件
- 予測可能なネットワークパフォーマンスを提供
- 安全なIPsec VPN接続を実現
- コスト効率の良い方法で可用性を達成
=> AWS Direct Connect, AWS Site-to-Site VPN (IPSec, private link)

## [x] 32. AssumeRoleによるユーザー認証 - AWS Security Token Service (AssumeRole, GetFederationToken)
ハイブリッドクラウドアーキテクチャを採用
- 自社ネットワークとAWSのクラウドインフラストラクチャを接続
- 既存のいくつかのデータベースを高速処理が可能なAWS上のサービスに移管
要件
- オンプレミス環境のアプリケーションからAWSリソースへとアクセスするための認証方式
- 社内ではSAML 2.0をサポートしていない社内のID認証システムによってユーザー管理を実施
=> AWS Security Token Service (AssumeRole, GetFederationToken)

## [x] 33. AWS Direct Connect - AWS Direct Connect (virtual private gateway), Amazon VPC (route propagation, route table), Amazon EC2
ハイブリッドクラウドアーキテクチャ
- 自社ネットワークとAWSのクラウドインフラストラクチャを接続
移行要件
- ハイブリッドクラウドを実現するためにオンプレミス環境からAWSへの Direct Connect接続を確立
- Direct Connectリンクを設定してルートをオンプレミス環境に接続
=> AWS Direct Connect (virtual private gateway), Amazon VPC (route propagation, route table), Amazon EC2

## [x] 34. Multi-AZとAuto Scaling - Amazon EC2 (spot fleet), Auto Scaling (multi-AZ)
EコマースサイトをAWSにホスト
- 3つのアベイラビリティーゾーン
  - ALBとオンデマンドEC2インスタンス
利用者が増加
- このECサイトのピーク時に処理落ちが発生
- システムの改善を依頼
要件
- 負荷のピーク時にはマルチAZに負荷を分散してオートスケール処理ができる必要
- スポットインスタンスを上手く利用してコスト最適に実現
=> Amazon EC2 (spot fleet), Auto Scaling (multi-AZ)

## [x] 35. Business Continuity Planning
事業継続性計画（BCP）ガイドライン
- 障害復旧時間（RTO）は1時間
- 目標復旧時点（RPO）は15分前
- 下記ケースにおいて推定される障害復旧時間とデータ損失
  - 災害が発生したことで停電などが発生
  - 午後2時にサーバーが停止
=> RPO 1:45, Lost 1:45-2:00

## [x] 36. Amazon DynamoDB (local secondary indexes)
業務システム
- EC2
- ELB
- DynamoDB
  - 設定時にUser_IDの主キーを持つTRANSACTIONSというテーブルを作成
  - 問題なく、IDの主キーに基づいてデータを照会できるように構成
DynamoDBテーブルの要件
- データ集計機能
  - ユーザーのアクセス頻度に応じて対象顧客をセグメンテーション
- 絞り込み検索
  - User_IDというパーティションキーに関連付け
=> Amazon DynamoDB (local secondary indexes)

## [x] 37. 分析基盤とコスト管理
アプリケーションログファイルから定期的な分析レポートを作成する監査用ログシステム
- すべてのログデータはAmazon S3バケットに収集
  - その後、毎日のAmazon EMRジョブによって分析が実行
    - 日次レポートと集計テーブルをCSV形式で生成
	- 別のS3バケットに保存
	- Amazon Redshiftデータウェアハウスに転送
課題
- 分析に利用するデータの使用頻度は不確実
- データ管理のライフサイクルポリシーをうまく設定できません
要件
- パフォーマンスやデータの整合性を損なうことなくコスト削減
=> Redshift (reserved instance)
=> EMR core/master node (reserved instance), EMR task node (spot instance)
=> S3 (Intelligent Tiering)

## [x] 38. AWS Organizations - Service Control Policy, AWS IAM
複数部門と支社でAWSサービスを利用
- 部門ごとにAWSアカウントを作成
- 各アカウント
  - その特定アカウントのみのルートアクセス権を持つシステム管理者によって管理
要件
- 全社統一でAWSアカウントを統合
  - 内部統制を強化
  - コスト削減
  - 複数のAWSアカウント全体でポリシーを集中管理
    - 特定のAWSサービスを許可または拒否
      - 個々のアカウントまたはアカウントのグループに対して
=> AWS Organizations (organization unit, service control policy)
=> AWS Organizations (organization unit), AWS IAM

## [x] 39. Amazon DynamoDBによるリージョン間のレプリケーション - Amazon DynamoDB (streams: singapore, global-table: all)
C to C専門のモバイルフリマサイト
- 複数のAWSリージョンに対してバックエンドAPIが起動
  - ユーザーに最も近いリージョンで販売および取引が処理されるようにルーティング
  - 東京リージョンから東南アジアにも展開
要件
- トランザクションがシンガポールリージョンにも自動的に複製されるレプリケーション構成を実現
=> Amazon DynamoDB (streams: singapore, global-table: all)

## [x] 40. AWS CloudTrail - [AWS Config, AWS CloudTrail] <-> AWS Organizations, AWS CloudWatch Events -> Amazon SNS (topic)
AWS Organizationsを使用
- 複数の組織単位（OU）にグループ化されたさまざまなチームや部門
- 年度ITセキュリティ監査を実施
  - 一部のメンバーアカウント内において、許可されていないサードパーティのアカウントが作成
    - 該当アカウントの責任者からは、API連携を実施する際に必要な対応であり、問題のないものと確認
要件
- 今後は許可のない外部アカウントの登録は拒否される必要がある
- 違反をモニタリングして早期に発見
- 事前に予防
=> AWS Config <-> AWS Organizations, AWS CloudWatch Events -> Amazon SNS (topic)
=> AWS CloudTrail <-> AWS Organizations, AWS CloudWatch Events -> Amazon SNS

## [x] 41. Amazon S3 (requester pays bucket)
開発部門
- 複数のAmazon S3バケットを使用
  - さまざまなデジタルアートワーク用の高解像度メディアファイルを保存
運用部門
- 別AWSアカウントにて
  - 開発部門S3バケット内のデータを利用した既存アプリケーションとの連携機能が運用
課題
- 運用部門のAWSアカウントから開発部門S3バケットの複数オブジェクトを頻繁に取得
  - 転送コストが開発部門に請求
=> Amazon S3 (requester pays bucket)

## [x] 42. Amazon VPC (overlap error of subnets)
レガシーシステムをアップグレードするためにAWSへの移行
- オンプレミスネットワークをAWSクラウドに移行
  - 以下のように構成
    - VPC（10.0.0.0/16）
      - パブリックサブネット（10.0.0.0/24）
      - プライベートサブネット（10.0.1.0/24）
      - 新しいパブリックサブネット（10.0.0.0/16）を追加
=> Amazon VPC (overlap error of subnet) (新しいパブリックサブネットと既存パブリックサブネット)

## [x] 43. ネットワークをオンプレからクラウドに移設 - BYOIP, Amazon Registrar RDAP
オンプレミスのデータセンター
- サプライチェーンアプリケーションをホスト
  - ホワイトリストに登録した信頼できるIPアドレスを使用
要件
- このアプリケーションを含めてオンプレミス環境のインフラをAWSへと移行
  - IPアドレスホワイトリストの変更を要求することなく、VPCに移行
=> BYOIP, Amazon Registrar RDAP

## [x] 44. Federation user with AWS IAM - AWS IAM (IAM SAML 2.0 IdP)
オンプレミス環境
- サードパーティーのSAML IdPを利用したログイン
  - これを利用したAWSリソースへのアクセス制御が必要
要件
- オンプレミスネットワークをAWSクラウドに接続するハイブリッドクラウドアーキテクチャ
- SAML IdPを利用したAWSリソースへのアクセス制御
=> AWS IAM (IAM SAML 2.0 IdP)

## [x] 45. AWS IAM (AssumeRoleWithWebIdentity)
写真共有アプリケーション
- 構成
  - アプリケーションサーバー: 一連のECインスタンスにELBとAutoScalingが設定
  - S3: 写真を保存
  - Amazon DynamoDB: 文字情報を保存
要件
- アプリケーションサーバーからDynamoDBテーブルへと連携する処理を実装
- このアプリケーションへのモバイル認証を実施して、DynamoDBへのアクセスすることが必要
  - Amazon Cognitoを利用しない方式
=> Web IdP, AWS IAM (AssumeRoleWithWebIdentity), Amazon DynamoDB

## [x] 46. Amazon CloudWatch Logs - Amazon CloudWatch Agent <-> Amazon EC2, CloudWatch Logs Insights
AWSにおいて決済管理システムや顧客管理ポータルを運用
- WindowsおよびLinux EC2インスタンスの毎月のパフォーマンスチェック
- 実稼働環境で実行されている200を超えるオンデマンドEC2インスタンスを利用
  - 収集分析
    - 各インスタンスのメトリクス
      - メモリ使用量
      - ディスク容量
    - さまざまなシステム詳細情報のログ
=> Amazon CloudWatch Agent <-> Amazon EC2, CloudWatch Logs Insights

## [x] 47. VPC Lambda
Lambda関数を使ったシステムコンポーネントの設計・実装
- WEB上で該当するLambda関数を実行すると
  - VPCにホストされているデータベースに処理結果を保存する機能を構築して
    - Lambda関数からVPC内のプライベートサブネットにあるデータベースにアクセスを試みましたが
      - Lambda関数は動作を停止してしまいました
=> Lambda (security group: outbound)
=> Amazon VPC, NAT Gateway

## [x] 48. Amazon CloudFront (Lambda Edge, Origin Failover)
仮想通貨取引システム
- 今年リリースしたモバイルから仮想通貨取引に参加できるアプリケーション
  - サーバレスアーキテクチャにより実装
課題
- このモバイルアプリケーションはグローバルに何10万人ものユーザーを抱えており
  - CloudFrontによってコンテンツが配信されることで最適な配信構成を実現していましたが
   - 最近になってHTTP 504エラーが時々発生しているようです
   - 特にログイン時に時間がかかっているようです
=> Amazon CloudFront (Lambda Edge)
=> Amazon CloudFront (Origin Failover)

## [x] 49. Amazon CloudFront (AES256, Origin Access Identity)
動画再生アプリケーション
- 動画データをS3に保存
- EC2インスタンスによる動画処理
- グローバルにユーザーに利用してもらう配信プラットフォームであるため
  - CloudFrontを前面に設定
課題
- 動画配信のセキュリティ制御を実装
要件
- 動画配信において
  - 配信者をアプリケーションの会員ユーザーに限定することが必要であり
  - 暗号化によって保存データを保護する必要があります
=> Amazon CloudFront (AES256, Origin Access Identity)

## [x] 50. Placement Group (stop, add & start Amazon EC2)
多層Webアプリケーション
- WEB層において
  - プレースメントグループを構成している9つのEC2インスタンスが実行
  - EC2インスタンスの処理負荷が増加
    - このプレイスメントグループに対して2つの新しいインスタンスを追加
=> Placement Group (stop, add & start Amazon EC2)

## [x] 51. Amazon Kinesis Data Streams (FIFO)
IoTデータによる農業データ管理システム
- 毎日の実行タスク
  - その日の農地のかかる土壌および水分データを取得
  - 最適な育成環境であるかを管理
  - 機械学習によってレコメンデーションを実行
要件
- この機能を実施するためには
  - 下記2つのトランザクション処理を実施することが必要
    - リアルタイム土壌分析処理
    - リアルタイム栄養素分析処理
- 2つのトランザクション機能が効率的にデータを処理できるようにするには
  - 同じトランザクションデータが確実に配信されてシリアル順でデータ順序が保証されていることが必要
=> Amazon Kinesis Data Streams

## 52. Amazon VPC (DHCP, domain-name-servers=AmazonProvidedDNS)
ネットワーク構成を実装
- DHCPは構成情報をTCP/IPネットワーク上のホストに提供
- DHCPオプションの最初のセットを作成してAmazonのDNSサーバーを利用してVPCに関連付けましたが
  - エラーが発生
=> Amazon VPC (DHCP, domain-name-servers=AmazonProvidedDNS)

## 53. AWS IAM (service principal: apigateway.amazon.com)
タスク管理アプリケーションサービス
- DynamoDBテーブル
- アプリケーションのデータ処理において新規のサーバレス機能を実装
  - APIゲートウェイからLambda関数を呼び出すことで
    - DynamoDBテーブルのデータを取得してデータ集計
  - 実装にはLambda関数によるDynamoDBテーブルへのアクセスのためには
    - Lambda関数に対してIAMロールを設定することが必要となり
	- 現在設定
=> AWS Lambda, AWS IAM (service principal: apigateway.amazon.com)
=> AWS Lambda, Amazon DynamoDB, IAM (action: [dynamodb:GetItem, dynamodb:PutItem])

## 54.Amazon EC2 (reserved instance)
金融システム向けのAWSクラウド環境
- 会社にはシステム開発・運用の各段階を分けるために3つの統合された請求先アカウントがあります
- 利用しているアカウントは
  - 開発
    - AZ：ap-northeast-1d
    - 3つのm4.largeのリザーブドインスタンス
    - 開発アカウントで実行されているインスタンスはありません
  - テスト
  - 本番環境用
    - AZ：ap-northeast-1d
    - 5つのm4.largeインスタンス 利用中
=> Amazon VPC (AZ: ap-northeast-1d), Amazon EC2 (reserved instance)

## 55. Amazon CloudFront (Viewer Protocol Policy)
新しいSNSアプリケーション
- 日常の写真などを共有したりメッセージを発信
- EC2インスタンス
  - 2つのアベイラビリティーゾーンにデプロイ
  - Auto Scalingグループ
  - ELB
  - CloudFront - 静的なコンテンツを配信
課題
- HTTPS/SSLを利用していないためGoogle検索ランキングが低くなっている
=> Amazon CloudFront (Viewer Protocol Policy, HTTPS Only)
=> Amazon CloudFront (Viewer Protocol Policy, Redirect HTTP to HTTPS)
=> Amazon CloudFront (Viewer Protocol Policy, SSL/TLS)

## 56. aws auto-scaling terminate-instance-in-auto-scaling-group --instance-id
現在開発している顔認証システム
- オンデマンドEC2インスタンス
  - Auto Scalingグループを使用
  - エラーを引き起こしている特定のインスタンスが1つあり
    - これを迅速に終了する必要があります
=> aws auto-scaling terminate-instance-in-auto-scaling-group --instance-id

## 57. Elastic IP Address
不動産ポータルサイト
- ECインスタンスベースのWebサーバー
  - パブリックサブネットに設定
  - 全てのIPアドレスからトラフィックを受信できるように設定
要件
- さらに追加で1つのバックエンド処理が実行することが必要
  - このバックエンド接続では
    - EC2インスタンスが選択したIP範囲からのみSSHトラフィックを受信することが必要
  - これらの処理は全ての1つのEC2インスタンス上で実現することが必要
  - したがって、1つのEC2インスタンスに対して
    - 2つのパブリックIPアドレスを利用することが要件
      - バックエンド処理用のIPアドレス
      - トラフィック制御用のIPアドレス
=> Elastic IP Address, EC2

## 58. Elastic Load Balancing (path-based routing)
eコマースアプリケーション
- 複数のEC2インスタンス
- ELB
- 複数のデバイスプラットフォームをサポート
  - モバイルやPC端末など多様なデバイスからアクセスして利用される予定
要件
- SSLによるセキュアな通信方式を設定することが必要
=> Elastic Load Balancing (path-based routing)

## 59. Amazon EC2, Auto Scaling
AIベースの交通監視アプリケーション
- このシステムは都市全体で使用されるため下記のような状態は大きな問題
  - 間違った情報を利用
  - 途中で障害が発生して不必要なダウンタイムが発生
要件
- システムダウンを極力回避するために
  - 可用性と耐障害性を高める必要
- アプリケーションの処理にはEC2インスタンスサーバー
- データベース処理にはSQLクエリ処理
=> Elastic Load Balancing, Amazon EC2, Auto Scaling, Amazon Aurora (multi-AZ), Amazon Route 53 (alias record)

## 60. Amazon VPC (Secondary IPv4 CIDR)
保険をAPIで提供するアプリケーション
- IPv4 CIDRブロック10.0.0.0/24のVPCに設置
  - IPアドレスが枯渇
要件
- VPC CIDR範囲を拡張するが必要
=> Amazon VPC (Secondary IPv4 CIDR)

## 61. AWS Key Management Service, AWS CloudHSM, Amazon EBS
健康管理アプリケーション
- 複数のEC2インスタンス
  - ELB
  - AutoScalingグループ
  - EBS - 機密性の高い健康記録データ保存
要件
- セキュリティコンプライアンスの一環として下記が義務付けられている
  - クラウドインフラストラクチャに保存されているすべてのデータ適切に保護あるいは暗号化
  - 暗号化方式を検討
=> AWS CloudHSM, EBS
=> AWS Key Management Service, EBS

## 62. LDAP with AWS Security Token Service
顧客管理用のWEBアプリケーション
- ハイブリッドアーキテクチャ
  - データベースがAWSのRDSに移行
要件
- オンプレミス環境にあるWEBアプリケーション
  - LDAP（Lightweight Directory Access Protocol）サーバーによる認証
  - AWS上にあるRDSの顧客データにアクセスすることが必要
=> Identity broker, LDAP authentication, AWS Security Token Service (AssumeRole), Amazon RDS

## 63. AWS CloudFormation, AWS Service Catalog
統合管理・一括請求の仕組みを構築
- 複数アカウントを管理するためにAWS Organizationsを利用
- 組織内のすべてのリソースを適切に管理するには
  - すべてのアカウントでリソースが作成されたときに
    - タグが常に追加されるようにする必要
=> AWS Service Catalog
=> AWS CloudFormation

## 64. Elastic Load Balancing (without sticky session), Amazon EC2
仲介業者を介さずに物件の内見予約が出来るアプリケーション
- ２つのAZに展開された8のECインスタンスによって構成
- 負荷テストを実施したところ負荷分散が達成されていない
  - ユーザーセッションは両方のAZのEC2インスタンスに均等に分散
  - 負荷テストツールのトラフィックは１つのAZにあるEC2インスタンスのみを利用
=> Elastic Load Balancing (without sticky session), Amazon EC2
=> Distributed stress testing

## 65. AWS CloudFormation (CreationPolicy: { ResourceSignal: Timeout })
CloudFormationを利用した環境構築
- スタック全体の起動に失敗した場合に
  - 自動的にロールバックするCloudFormationテンプレートを準備する必要がある
- アプリケーションスタックを適切に実行するには
  - 最初に前提条件パッケージをインストールする必要がある
- 展開が完了するには約1時間ほどかかる場合がある
=> AWS CloudFormation (CreationPolicy: { ResourceSignal: Timeout })

## 66. Amazon CloudFront (Origin Access Identity), Amazon S3
S3を利用したコンテンツ共有の仕組みを構築
- ユーザーからのアクセスに限定する必要
  - インスタンスのIPアドレスからのみオブジェクトにアクセスできるようにする設定することが要件
=> Amazon CloudFront (Origin Access Identity), Amazon S3

## 67. RD Gateway Server
インターネットを介してVPCに接続し
- EC2インスタンスを管理
  - パブリックサブネット
  - プライベートサブネット
要件
- Microsoft Remote Desktop Protocol（RDP）アクセスによるリモートアクセスを実現したい
  - EC2インスタンスへのインターネットアクセスを制限して安全なものにするため
    - Bastionホストを利用したインターネットアクセスを構成する必要
=> Amazon EC2 (private subnet), RD Gateway Server/Amazon EC2 (security group)

## 68. Amazon CloudWatch Dashboards (cross-region)
Eコマースサイト
- マルチリージョン
- EC2
要件
- EC2インスタンス全てのパフォーマンスを監視する必要
=> Amazon CloudWatch Dashboards (cross-region)
 
## 69. Amazon Cognito, AWS STS (AssumeRole, AssumeRoleWithWebIdentity)
データ共有システム
- ユーザーが直接アップロードしたデータを単一のS3バケットに保存
- ユーザーはS3バケットから直接自分がアップロードしたデータを表示およびダウンロードできる
要件
- 社員数1万ものユーザーがいるためこれらのデータ処理は可能な限り安全に実施される必要
=> Amazon RDS, AWS IAM, Amazon Cognito, AWS Security Token Service (AssumeRoleWithWebIdentity), Amazon S3
=> Amazon RDS, AWS IAM, AWS Security Token Service (AssumeRole), Amazon S3

## 70. Amazon VPC (Subnet, Network ACL), Amazon EC2 (Placement Group, Security Group), AWS Config
金融機関の決済システム
- 2つのグループに分類されたEC2インスタンス群
  - 同じアベイラビリティゾーン（AZ）
  - 2つの異なるサブネット
  - 1つのグループのEC2インスタンス群
    - データベースを実行
  - もう1つのグループのEC2インスタンス群
    - データベースからデータを取得するWebサーバー
要件
- このWEBアプリケーションが正しく機能するためには
  - EC2インスタンス群内での通信パフォーマンスを最適なものとした上で
  - それぞれのEC2インスタンスの一連のグループが相互に接続できることを確認する必要
- インスタンスに関連付けられたセキュリティ構成の変更履歴を追跡する必要
=> Amazon VPC (Subnet, Network ACL), Amazon EC2 (Placement Group, Security Group), AWS Config

## 71. Amazon RDS (snapshot), Amazon Aurora (multi cluster, read replica), Amazon CloudFront
英語ニュースサイト
- 各記事には多数の画像が含まれ、そのコンテンツは少なくとも200語以上あります
- 新しい記事は最初の1か月間で最も閲覧されており
  - 著者は公開後最初の1か月で頻繁に記事を更新する傾向があります
- データベースにはRDS MySQLを利用しており
  - ニュース記事のデータ処理にはクエリ処理が多数利用されるため
    - リレーショナルモデルが必要不可欠です
課題
- 最近になってこのニュースメディアの利用者が急増しており
  - コンテンツの読込時間が長いというクレームが発生するようになりました
=> Amazon RDS (snapshot), Amazon Aurora (multi cluster, read replica), Amazon CloudFront

## 72. Amazon Redshift (auto snapshot), Amazon Redshift (disater recovery)
仮想通貨取引プラットフォーム
- 取引実行データの分析においてRedshiftクラスターを実行
要件
- Redshiftの災害対応の構成を検討
  - リージョン内のAZが停止した際に即時に対応できる構成
  - リージョン自体が停止したい際には1日で回復できる構成
=> Amazon Redshift (auto snapshot), Amazon Redshift (disater recovery)
=> Amazon Redshift (multi AZ)

## 73. Network ACL (inbound: port 1024-65535)
WEBアプリケーション
- Amazon EC2インスタンス
  - 複数のアベイラビリティーゾーン
  - ALB
  - Auto Scalingグループでホスト
- HTTPおよびHTTPSトラフィックを許可するために
  - ポート80および443でのインバウンドトラフィックを許可
    - ALBおよびEC2インスタンスの両方のネットワークACL
    - セキュリティグループ
課題
- インターネットからWEBアプリケーションへと接続することができない
=> Network ACL (inbound: port 1024-65535)

## 74. SSM Agent, AWS Systems Manager Automation
オンライン決済プラットフォーム
- EC2インスタンス
  - Auto Scalingグループ
  - ELB
  - AWS Systems Managerを使用してEC2インスタンスのグループを監視および処理
要件
- これらのインスタンスにメンテナンスやOSパッチなどのバッチ操作がある場合は
  - Systems Managerを使用してこれらのアクティビティを自動的に実施する設定が必要
=> EC2 (SSM Agent)
=> AWS Systems Manager Automation

## 75. Amazon EC2 (AMI, key pairs)
AWSの既存リソースの一部を別リージョンに移行
1. すべてのAmazon Machine Image（AMI）を東京リージョンからシンガポールリージョンにコピー
  - しかしながら、AMIのコピーするだけでは該当EC2インスタンスにアクセスすることができない
  - シンガポールリージョンに向けてコピーされたAMIを起動する際には
    - 最適なPEMキーを指定して起動することが必要
      - 会社の管理方針としてPEMキーを単一のキーで一元的に利用することが求められている
=> Amazon EC2 (AMI, key pairs)
