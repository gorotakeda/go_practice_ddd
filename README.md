# go_practice_ddd

アプリケーション概要

ユーザーのCRUDを行うアプリケーション

このアプリケーションはDDD（ドメイン駆動設計）のアーキテクチャパターンを採用しています。以下のような階層構造になっています：

1. **プレゼンテーション層（Presentation Layer）**
   - `handler/user_handler.go`
   - HTTPリクエストの受け付けとレスポンスの返却を担当
   - Gin Webフレームワークを使用
   - エンドポイント：
     - GET /users
     - POST /users

2. **アプリケーション層（Application Layer）**
   - `service/user_service.go`
   - ユースケースの実装
   - ドメインオブジェクトの操作とビジネスロジックの調整
   - リポジトリを使用してデータの永続化を行う

3. **ドメイン層（Domain Layer）**
   - `domain/db.go`
   - ビジネスロジックの中心
   - エンティティの定義（User構造体）
   - ドメインモデルとビジネスルール

4. **インフラストラクチャ層（Infrastructure Layer）**
   - `infrastructure/db.go`
   - データベース接続の管理
   - `repository/user_repository.go`
   - データの永続化を担当
   - GORMを使用してMySQLとの通信を行う

5. **設定層（Configuration Layer）**
   - `config/config.go`
   - 環境変数の管理
   - データベース接続情報の設定

依存関係の方向：
```
Handler → Service → Repository → Infrastructure
    ↘    ↘    ↘
      Domain   Config
```

技術スタック：
1. **言語とフレームワーク**
   - Go言語
   - Gin（Webフレームワーク）
   - GORM（ORMライブラリ）

2. **データベース**
   - MySQL 8.0

3. **コンテナ化**
   - Docker
   - Docker Compose
   - マルチコンテナ構成：
     - アプリケーションコンテナ
     - データベースコンテナ

クリーンアーキテクチャの原則に従い：
1. 依存関係は内側に向かう（外層は内層に依存）
2. 内側の層は外側の層を知らない
3. ドメインロジックはインフラストラクチャの詳細から独立

これにより：
- テスタビリティの向上
- 保守性の向上
- 関心の分離
- スケーラビリティの確保

が実現されています。
