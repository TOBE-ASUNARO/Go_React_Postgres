# フロント用開発環境 by Docker
## 1) .env ファイルを作る。（環境設定ファイル）
1. git cloneして、できたTOPのDirにて作業（docker-compose.ymlのあるDir)
1. dot.env_CopyThisFileTo.envを、コピーして、.envファイルを作成。
1. .envの内容を編集。

## 2) docker-compose up -d
たぶん、起動すると思います。


## 3) 起動するサーバの内容
1. nginx リバースプロキシ
1. node
1. postgres; usr:root pw:password
1. swagger-edit: http://loccalhost:10081
1. swagger-ui: http://loccalhost:10082
1. swagger-mock**: http://loccalhost:10083

**
モックサーバーは、stoplight/prism:4  
/openapi/api_reference.yamlで、内容を定義
