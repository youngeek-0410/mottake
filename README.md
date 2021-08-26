# mottake

## 起動

### Backend

1. ```console
   git clone https://github.com:youngeek-0410/mottake.git
   cd mottake
   ```

1. `server/`に`config.yml`を追加

   ```yml:config.yml
   db_config:
    user: admin
    password: password
    name: mottake
    host: db
   firebase_secret: "/go/src/app/firebase-secret.json"
   mode: DEVELOP
   port: :80
   ```

1. `server/`に Firebase のサービス アカウント用の秘密鍵ファイル`firebase-secret.json`を追加
   (`config.yml`の mode が DEVELOP ならこのファイルがなくても起動する)

1. `postgres/`に`.env`を追加

   ```.env:.env
   POSTGRES_USER=admin
   POSTGRES_PASSWORD=password
   POSTGRES_DB=mottake
   ```

1. ```console
   docker-compose build
   docker-compose up
   ```

1. [localhost](http://localhost) にアクセス
