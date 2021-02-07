---
title: Udemy Mock Exam 2
tags: aws-csap, amazon-web-services
url: https://www.udemy.com/course/aws-53225/learn/quiz/4723918#content
---

# 1st time
## 1. Amazon EC2 (security group: outbound)
パブリックサブネットで起動されたEC2インスタンスを運用
- 定期的にインターネットからソフトウェアアップグレードとパッチにアクセスできるようにする必要がある
- これらのインフラはURLを介して運用管理業者からアクセスされてモニタリングやメンテナンスが実行
課題
- セキュリティ監査を実施したところ
  - 第三者モニタリングの仕組みにおいてネットワーク上の問題点があるとの報告を受ける
  - VPC内のインスタンスからインターネットに対して
    - 他のアウトバウンド接続を明示的に拒否する設定を行う必要がある
=> Amazon EC2 (security group: outbound)

## 2. Amazon CloudWatch Metrics (HTTPCode_Backend_5XXX)
社内業務システム
- Classic Load Balancerを使用して
  - 複数のリザーブドEC2インスタンスに均等に着信トラフィックを分散
課題
- この業務システムのアプリケーションサーバーが原因と思われる断続的な使用不可が発生
- この原因を究明するために
  - 登録されたインスタンスから送信されたサーバーエラーを確認することが必要
=> HTTPCode_Backend_5XXX

## 3. Amazon CloudFront (restrict: origin server)
仮想通貨取引プラットフォーム事業の要件
- これは金融サービスとなるため高いコンプライアンス要件が仮想通貨取引所には求められており
  - マネーロンダリング防止およびテロ対策資金調達対策を実施する必要がある
- すべてのレポートファイルは
  - 特定の国・地域または個人に対して利用できないようにすることが要件
課題
- コンテンツを世界中のユーザーに低レイテンシーで配信することが必要
- ディストリビューションに関連するファイルのサブセットへのアクセスを制限することが必要
=> CloudFront (restrict: origin server)

## 4. Amazon S3 (versioning)
アート鑑賞SNSサービス
- ユーザー情報を保存するために
  - PINTOR-CONFIGという名前のプライベートS3バケットを使用
    - データを保護するために本日バージョン管理を有効にして
      - ユーザー構成情報に加えられた変更を追跡できるように設定
バージョン管理有効化時点のファイル
- ユーザー構成情報ファイル
- 共通設定ファイル
- ユーザー構成情報ファイル02
今後のイベント状況
- バージョン管理後にタスクファイルとログファイルが追加
- バージョン管理後にユーザー構成情報ファイルとユーザー構成情報ファイル02が更新
=> Amazon S3 (versioning)

## 5. Amazon Rekognition
防犯カメラの映像から万引きを特定するサービス
- 防犯カメラからのストリーミングビデオで顔認証を実施
  - 過去の万引き犯データとマッチング
要件
- デオを介してリアルタイムで迅速に顔のアドレスを指定
  - ダウンストリームの処理に適した方法で出力を保存できる必要がある
- Rekognitionを使用してサービスを開発
=> PutMedia API: Kinesis Video Streams -> Rekognition Video
=> CreateStreamProcessor: Rekognition Video
=> Rekognition Video -> Data Streams -> Data Streams Analytics

## 6. AWS Database Migration Service
MySQLデータベースをオンプレミス環境
方針
- AWSクラウドとのハイブリッド構成、MySQLをAWSに移行することを検討
要件
- オンプレミスとの同期を維持するにはAWSのDBインスタンスが必要
- DB移行後は徹底的にテストされて問題なく
  - オンプレミス環境との齟齬がなくなった段階でオンプレミスDBは利用されなくなる
- Amazon Database Migration Service（DMS）を利用した実施方法を整理
=> On-premise DB <-> RDS: DMS (VPN)
=> DMS (CPU, replication)
=> DMS (migrate existing data and replicate ongoing changes)

 
