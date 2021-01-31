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
