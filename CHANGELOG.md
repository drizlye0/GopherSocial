# Changelog

## [1.2.0](https://github.com/drizlye0/GopherSocial/compare/v1.1.0...v1.2.0) (2025-06-20)


### Features

* Dockerfile ([e9d5f7e](https://github.com/drizlye0/GopherSocial/commit/e9d5f7ed4ecaf8b44ab51d866443e40169b129b1))


### Bug Fixes

* change manual version ([98cb915](https://github.com/drizlye0/GopherSocial/commit/98cb915b408922ff625c68e071b0dcc9c3299f2a))

## [1.1.0](https://github.com/drizlye0/GopherSocial/compare/v1.0.0...v1.1.0) (2025-06-20)


### Features

* action for auto update version ([8a8e190](https://github.com/drizlye0/GopherSocial/commit/8a8e190e9319969c5e878e3543c6b4ddce160eeb))

## 1.0.0 (2025-06-20)


### Features

* add a command for clean dirty db on migrations ([0000f9d](https://github.com/drizlye0/GopherSocial/commit/0000f9d437d5fc60a9fc598cf69dcbe000ddbc39))
* add automation workflow ([9cbf0d1](https://github.com/drizlye0/GopherSocial/commit/9cbf0d16b8477b039996c89d8ba04afc35bb0c55))
* add indexes for better perfomance on reads ([d4b0a3a](https://github.com/drizlye0/GopherSocial/commit/d4b0a3a42a8f3e4ebf525b347b4e4fafd07fd6c6))
* add redis chaching for auth middleware ([3827030](https://github.com/drizlye0/GopherSocial/commit/382703082eabed6bb0c53c0548ad476a0df4df21))
* add server metrics ([9e708ad](https://github.com/drizlye0/GopherSocial/commit/9e708ada809853d3d920539818714fd694f1b5af))
* add store.Comments.Create for create comments ([765daa1](https://github.com/drizlye0/GopherSocial/commit/765daa1cdbdde4acff4d959778f15713b30e397d))
* adding a jsonResponse standardized and replace writeJSON on handlers ([e7af798](https://github.com/drizlye0/GopherSocial/commit/e7af79840518d6dfcd483b15cb85237065bd8f53))
* adding structured logger to api ([9a7e886](https://github.com/drizlye0/GopherSocial/commit/9a7e886b21247660f81b38db9710323457efcdbd))
* adds contextTimeout for querys ([2335472](https://github.com/drizlye0/GopherSocial/commit/233547227cbe5f235b77d79e8f4cf43bdb0ca830))
* adds db seed for fake data ([cb17fe5](https://github.com/drizlye0/GopherSocial/commit/cb17fe56279ed01a7a1c0e7fc3bcf333d3e8da33))
* adds getPostHandler and store.GetByID ([5948710](https://github.com/drizlye0/GopherSocial/commit/5948710618a2b1c59a7e59dbcdb2b69228229313))
* adds tags,search,until,since fields in pagination filter ([2c9c37f](https://github.com/drizlye0/GopherSocial/commit/2c9c37fc616ec5beb49808793dad90736190c525))
* adds validators for payload structs ([097f56c](https://github.com/drizlye0/GopherSocial/commit/097f56c5e94c15d0823825274563c4ec839bf9a7))
* create initial auth for user creation ([d3feceb](https://github.com/drizlye0/GopherSocial/commit/d3feceb1a99cbeff871a9954f2cebc8e2fa3f482))
* create new columns for migrations and implement createPostHandler ([569bb34](https://github.com/drizlye0/GopherSocial/commit/569bb346f3e9bd1b315d87d094930b773316d892))
* graceful server shutdown ([33ee720](https://github.com/drizlye0/GopherSocial/commit/33ee7203db7ab988eed4c26ad1bbb80c63d522e2))
* implement pagination for user feed ([12c9bd2](https://github.com/drizlye0/GopherSocial/commit/12c9bd2438211f0b0bcc92c0bfac89768ca9cde4))
* implementing activation for users ([53c72df](https://github.com/drizlye0/GopherSocial/commit/53c72df7a122d901273f97bc1aa6aa34831bc48f))
* implementing rate limiter ([9dbf4b0](https://github.com/drizlye0/GopherSocial/commit/9dbf4b0f8322272b61805c9a6f2c93800a200b3f))
* implementing user feed logic ([f9671fe](https://github.com/drizlye0/GopherSocial/commit/f9671fe1b2646994d915b1d6e36a56975b033bce))
* implements basic auth ([1fc5932](https://github.com/drizlye0/GopherSocial/commit/1fc593209a471f031a669a085981080ed362a7be))
* implements caching for getUserHandler ([ce7ce34](https://github.com/drizlye0/GopherSocial/commit/ce7ce34c8d36a0505eba7e38a558b9dce870b817))
* implements createCommentHandler ([8910d71](https://github.com/drizlye0/GopherSocial/commit/8910d71791aa9e7d0d9a46c4daea52206d3fb1f5))
* implements deletePostHandler and deleteByID methods ([be583d2](https://github.com/drizlye0/GopherSocial/commit/be583d26abec925c57829890bb5d95c3862f9af8))
* implements followers table and followers handler ([10e26b3](https://github.com/drizlye0/GopherSocial/commit/10e26b37b83d62bbe947cb1fd4860b6d21d265ea))
* implements initial docs with swagger ([19cbf8d](https://github.com/drizlye0/GopherSocial/commit/19cbf8d799b016f6e1c02a28f0a17bde08fb1379))
* implements jwt for tokens auth ([fac6fed](https://github.com/drizlye0/GopherSocial/commit/fac6fed75fed3b7c3e8f56d6dec74b32b5df6bf7))
* implements since and until url Params for userFeed ([aa3c3ed](https://github.com/drizlye0/GopherSocial/commit/aa3c3ed837d187310b2b4fc6d1c0cfe8d3637e3c))
* implements user comments in posts ([a6949a0](https://github.com/drizlye0/GopherSocial/commit/a6949a07601e1a238c6ff4e04f917612a33e06ce))
* implements userContextMiddleware, getUserHandler and store.GetByID ([5bab8c2](https://github.com/drizlye0/GopherSocial/commit/5bab8c224ece040597bb71db5452b1ad6b58e8cd))
* implments role authorization ([be25020](https://github.com/drizlye0/GopherSocial/commit/be250201d1d14975b0a18dab4bb373804a3c889d))
* implments updatePostHandler, updatePost and postContextMiddleware for patch route ([f84f689](https://github.com/drizlye0/GopherSocial/commit/f84f689e7bedb021ae6030a1acdcc3312bfbff1a))
* improve error handling for http.handlers ([3c6948d](https://github.com/drizlye0/GopherSocial/commit/3c6948d228e4afac45e8a443d2397f584d79f962))
* optmistic concurrency control ([3e2dd4e](https://github.com/drizlye0/GopherSocial/commit/3e2dd4ed189af988132a4cf0bcef156c0223ff6b))
* release please script ([7b4696a](https://github.com/drizlye0/GopherSocial/commit/7b4696ae10cf18c6269ae8d5f9e70588e2ef6223))
* sending emails for auth finished ([197f9e5](https://github.com/drizlye0/GopherSocial/commit/197f9e57c8bc26ffed2cb9153ca9f81b185a9a10))
* validations for tokens and clean all todos ([39e2f7c](https://github.com/drizlye0/GopherSocial/commit/39e2f7c092d7d4e440c1ddd017016bd40111c4b3))


### Bug Fixes

* bug on GetRoleByName ([097529e](https://github.com/drizlye0/GopherSocial/commit/097529e8d5654ee9acfe82a21c663dfedffd7ae2))
* bug on token generation handler ([17b3436](https://github.com/drizlye0/GopherSocial/commit/17b34364415bd9ad94ea3a5534813a2d3d74ed76))
* change ubuntu version ([9016197](https://github.com/drizlye0/GopherSocial/commit/9016197fe4b5e5fcb9612c567342217747345b0e))
* check if redis is enabled ([1592d9c](https://github.com/drizlye0/GopherSocial/commit/1592d9c3488f1ea22a8eff683e2b8a64b1e2e23d))
* clean imports ([b6bcb7c](https://github.com/drizlye0/GopherSocial/commit/b6bcb7cb5ac8d08da7b850c06bcda9f0a2c8cc28))
* cors ([dbd7382](https://github.com/drizlye0/GopherSocial/commit/dbd7382cf3e9ed37a61d3e805918529b9b6b4f4c))
* delete DB_MIGRATOR_ADDR ([a907630](https://github.com/drizlye0/GopherSocial/commit/a907630ed7ded8e40beb929d1861e1885d54c1ad))
* fix query for user feed and fix bugs with fq.Tags ([700a366](https://github.com/drizlye0/GopherSocial/commit/700a366a40121285c328c19a44b265cf0b49bfba))
* insert query on posts storage ([33e159d](https://github.com/drizlye0/GopherSocial/commit/33e159d8b1df88d6b00b806dd7d2cc495e221eac))
* remove imports ([449e86a](https://github.com/drizlye0/GopherSocial/commit/449e86a2ccdf15baa603ce366bfcefa5b122f82f))
* status not used ([4d40f0d](https://github.com/drizlye0/GopherSocial/commit/4d40f0d44ec1a19c24208bacca464eb65eda8481))
* unused middleware code ([9ba1fb0](https://github.com/drizlye0/GopherSocial/commit/9ba1fb037e3b4bd6a085ffa40ede91294fd590e7))
* user invitationOC ([5b19f99](https://github.com/drizlye0/GopherSocial/commit/5b19f9967a1a57f502da466efe45c9839da82cee))
