# go-learning-webapp-auth

## About This

https://auth0.com/blog/authentication-in-golang/

こちらの記事のサーバ側処理を学習のために写経したもの。

JWTによる認証は以下の記事を参考にした。

https://blog.motikan2010.com/entry/2017/05/12/jwt-go%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B

またget-tokenには、Basic認証を実装している。

## API Call

```
$ curl -v http://localhost:8080/status
$ curl -v http://localhost:8080/products # not authorized error.
$ curl -u username:password -v http://localhost:8080/get-token
$ TOKEN=上記で取得したトークン
$ curl -v http://localhost:8080/products -H "Authorization: $TOKEN"
```
