# gintone

kintone REST API の Go クライアントライブラリ（自動生成）

## 概要

gintone は、[kintone の公式 OpenAPI 仕様](https://github.com/kintone/rest-api-spec)から自動生成される Go 言語のクライアントライブラリです。タグごとに独立したクライアントパッケージを提供し、必要な機能だけを選んで使用できます。

## 特徴

- **自動生成**: kintone の公式 OpenAPI 仕様から oapi-codegen を使用して自動生成
- **タグベース**: API タグごとに独立したクライアント（apps, fields, files, forms, plugins, record, records, spaces など）
- **常に最新**: GitHub Actions により毎日自動更新（UTC 18:00 / JST 03:00）
- **型安全**: Go の型システムを活用した安全な API 呼び出し

## インストール

```bash
go get github.com/f4ah6o/gintone
```

特定のタグのクライアントのみを使用:

```bash
# 例: records クライアントのみ
go get github.com/f4ah6o/gintone/tags/records
```

## 使い方

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/f4ah6o/gintone/tags/records"
)

func main() {
    // クライアントの作成
    client, err := records.NewClient("https://your-domain.cybozu.com",
        records.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
            req.Header.Set("X-Cybozu-API-Token", "your-api-token")
            return nil
        }))
    if err != nil {
        log.Fatal(err)
    }

    // API の呼び出し
    // （具体的な使用方法は生成されたクライアントの型定義を参照）
}
```

## 利用可能なクライアント

以下のタグごとにクライアントが生成されています:

- `tags/apps` - アプリケーション管理
- `tags/fields` - フィールド管理
- `tags/files` - ファイル管理
- `tags/forms` - フォーム管理
- `tags/plugins` - プラグイン管理
- `tags/record` - レコード操作（単一）
- `tags/records` - レコード操作（複数）
- `tags/spaces` - スペース管理

## プロジェクト構造

```
.
├── cmd/
│   └── sanitizeopenapi/    # OpenAPI 仕様のサニタイゼーションツール
├── tags/                    # タグごとのクライアント
│   ├── apps/
│   ├── fields/
│   ├── files/
│   ├── forms/
│   ├── plugins/
│   ├── record/
│   ├── records/
│   └── spaces/
└── .github/
    └── workflows/
        └── generate-kintone-tags.yaml  # 自動生成ワークフロー
```

## 開発

### クライアントの再生成

GitHub Actions により自動的に実行されますが、手動で実行することも可能です:

1. [kintone/rest-api-spec](https://github.com/kintone/rest-api-spec) をクローン
2. OpenAPI 仕様をサニタイズ:
   ```bash
   go run ./cmd/sanitizeopenapi -in <spec-file> -out <output-file>
   ```
3. タグごとにクライアントを生成（詳細は `.github/workflows/generate-kintone-tags.yaml` を参照）

### sanitizeopenapi ツール

kintone の OpenAPI 仕様には、oapi-codegen と互換性のない `format: boolean` が含まれている場合があります。このツールは、そのような互換性のない定義を自動的に削除します。

```bash
go run ./cmd/sanitizeopenapi -in input.yaml -out output.yaml
```

## ライセンス

MIT License - 詳細は [LICENSE](LICENSE) ファイルを参照してください。

## クレジット

- [kintone REST API Specification](https://github.com/kintone/rest-api-spec)
- [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)
