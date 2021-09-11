# フロント用開発環境 by Docker

## 1) .env ファイルを作る。（環境設定ファイル）

1. git clone して、できた TOP の Dir にて作業（docker-compose.yml のある Dir)
1. dot.env_CopyThisFileTo.env を、コピーして、.env ファイルを作成。
1. .env の内容を編集。

## 2) docker-compose up -d

たぶん、起動すると思います。

## 3) 起動するサーバの内容

1. nginx リバースプロキシ\* : localhost:8888
1. node : localhost:3000
1. postgres; usr:root pw:password, port 5432
1. swagger-edit: http://loccalhost:10081
1. swagger-ui: http://loccalhost:10082
1. swagger-mock**: http://loccalhost:10083

- node.js 側での設定も必要  
  具体的には、vue cli なら、カレントに、vue.config.js をつくり、

```shell
module.exports = {
  devServer: {
    disableHostCheck: true
  }
}
```

を記載

**
モックサーバーは、stoplight/prism:4  
/openapi/api_reference.yaml で、内容を定義

---
# Vite にて、環境構築。
ネット上の記事を参考に構築  
設定ファイルを記載  
チーム開発を想定して、.vscodeを共有、
.editconfigも設定

## React-TS scaffold and prettier + eslint
```shell
yarn init
npm init @vitejs/app
```
https://qiita.com/hedrall/items/f2239f98f1c0fdd14b07

## eslint
```shell
npx eslint --init
```
https://zenn.dev/thiragi/scraps/8e988668dbc860

## prettier and eslint & vscode
```shell
yarn add -D prettier eslint-config-prettier
yarn add -D eslint-plugin-react-hooks
````
( https://qiita.com/sprout2000/items/ee4fc97f83f45ba1d227 )

## vscode
prettierの動作を無効化するファイルを設定
https://dev.classmethod.jp/articles/the-setting-to-disable-formatting-for-certain-languages-does-not-work-after-updating-prettier-in-vscode/

プラグインを推奨表示させる
https://qiita.com/k_bobchin/items/717c216ddc29e5fbcd43
