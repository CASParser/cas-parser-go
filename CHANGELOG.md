# Changelog

## 0.8.0 (2026-05-01)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/CASParser/cas-parser-go/compare/v0.7.0...v0.8.0)

### Features

* **go:** add default http client with timeout ([beb6a1f](https://github.com/CASParser/cas-parser-go/commit/beb6a1f6e2ec642dac7e6d5050a781cb081871a8))
* support setting headers via env ([c5ffd5f](https://github.com/CASParser/cas-parser-go/commit/c5ffd5f6015f40bdfe27a34f630c59f7c2c187f4))


### Chores

* avoid embedding reflect.Type for dead code elimination ([0b78868](https://github.com/CASParser/cas-parser-go/commit/0b78868a151d1a5a401906190cf202a4854310b0))
* **internal:** more robust bootstrap script ([58cdef6](https://github.com/CASParser/cas-parser-go/commit/58cdef67485d0890acb2b0263053a79260a3fe23))

## 0.7.0 (2026-04-19)

Full Changelog: [v0.6.1...v0.7.0](https://github.com/CASParser/cas-parser-go/compare/v0.6.1...v0.7.0)

### Features

* **api:** api update ([f1e6d47](https://github.com/CASParser/cas-parser-go/commit/f1e6d47f6e70df14a15cd32c475809d2673cd84a))
* **api:** api update ([f51a84f](https://github.com/CASParser/cas-parser-go/commit/f51a84fb9a2e2b45cc9c332953df042d0ae381ff))

## 0.6.1 (2026-03-31)

Full Changelog: [v0.6.0...v0.6.1](https://github.com/CASParser/cas-parser-go/compare/v0.6.0...v0.6.1)

### Chores

* **ci:** support opting out of skipping builds on metadata-only commits ([35efa48](https://github.com/CASParser/cas-parser-go/commit/35efa48691feeb6e6ed97d2349e7c7cb1f11caf2))
* update docs for api:"required" ([14c880c](https://github.com/CASParser/cas-parser-go/commit/14c880c0fd219543f84f2b6d2d303e2fac946274))

## 0.6.0 (2026-03-27)

Full Changelog: [v0.5.3...v0.6.0](https://github.com/CASParser/cas-parser-go/compare/v0.5.3...v0.6.0)

### Features

* **internal:** support comma format in multipart form encoding ([6c39815](https://github.com/CASParser/cas-parser-go/commit/6c3981526f178833589ac8991f3e39ac3bffb2cf))


### Bug Fixes

* prevent duplicate ? in query params ([f7cfb89](https://github.com/CASParser/cas-parser-go/commit/f7cfb89b74abdde929b9bf09e7f461f0a59701cb))


### Chores

* **ci:** skip lint on metadata-only changes ([cb79b17](https://github.com/CASParser/cas-parser-go/commit/cb79b17587f210a297a451bde4d02934c0e7de67))
* **client:** fix multipart serialisation of Default() fields ([17d404d](https://github.com/CASParser/cas-parser-go/commit/17d404d8858084b9b71419138e72de0dc7e6f65b))
* **internal:** support default value struct tag ([6da3050](https://github.com/CASParser/cas-parser-go/commit/6da30504423b6f369d2a1a3e9ddb24263856db9c))
* **internal:** update gitignore ([22ed58c](https://github.com/CASParser/cas-parser-go/commit/22ed58cca58650432ef3cfaefbe3cb80d85d60ee))
* remove unnecessary error check for url parsing ([7cb383a](https://github.com/CASParser/cas-parser-go/commit/7cb383af66275655755216a51b0d0af6fcd3b09e))

## 0.5.3 (2026-03-17)

Full Changelog: [v0.5.2...v0.5.3](https://github.com/CASParser/cas-parser-go/compare/v0.5.2...v0.5.3)

### Chores

* **internal:** minor cleanup ([5dcbbb1](https://github.com/CASParser/cas-parser-go/commit/5dcbbb199f0fc37d6676c9d259205773c794406f))
* **internal:** tweak CI branches ([6377b4a](https://github.com/CASParser/cas-parser-go/commit/6377b4aa8da9f70467ba34dba7a1ddf4ffcc3743))
* **internal:** use explicit returns ([edbbf41](https://github.com/CASParser/cas-parser-go/commit/edbbf41442d2577fb04903e619b9b67554d10209))
* **internal:** use explicit returns in more places ([af11842](https://github.com/CASParser/cas-parser-go/commit/af11842d738d5907bde4ba49fcd17679a037dc3e))

## 0.5.2 (2026-03-07)

Full Changelog: [v0.5.1...v0.5.2](https://github.com/CASParser/cas-parser-go/compare/v0.5.1...v0.5.2)

### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([2087492](https://github.com/CASParser/cas-parser-go/commit/20874926fcaea5a7dfac7cd4bde7cbfa3d58f547))
* **internal:** codegen related update ([1a6c299](https://github.com/CASParser/cas-parser-go/commit/1a6c2997eeb31bda0fe28435dc58d99e3b1914d9))

## 0.5.1 (2026-03-03)

Full Changelog: [v0.5.0...v0.5.1](https://github.com/CASParser/cas-parser-go/compare/v0.5.0...v0.5.1)

### Chores

* **internal:** codegen related update ([30c321b](https://github.com/CASParser/cas-parser-go/commit/30c321bc71574888f59d690cd0bfc984bcbf1855))
* **internal:** move custom custom `json` tags to `api` ([54c32e5](https://github.com/CASParser/cas-parser-go/commit/54c32e5ef17c4c66f25b0c3efa3090827b1f6cb5))

## 0.5.0 (2026-02-23)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/CASParser/cas-parser-go/compare/v0.4.0...v0.5.0)

### Features

* **api:** manual updates ([e6e91f2](https://github.com/CASParser/cas-parser-go/commit/e6e91f2b8484a34b154fd55f9297183a9ff9870c))

## 0.4.0 (2026-02-23)

Full Changelog: [v0.3.1...v0.4.0](https://github.com/CASParser/cas-parser-go/compare/v0.3.1...v0.4.0)

### Features

* **api:** api update ([b60c76b](https://github.com/CASParser/cas-parser-go/commit/b60c76b95425a9fe76c8c14de99154b64ae5f460))
* **api:** api update ([bd3e431](https://github.com/CASParser/cas-parser-go/commit/bd3e431ba72031908dbee6cf4bd6ea2f7902a887))
* **api:** api update ([2a405e1](https://github.com/CASParser/cas-parser-go/commit/2a405e1083d8cdcb597dfcd9a7031a0bfaac037a))
* **api:** manual updates ([fb1332a](https://github.com/CASParser/cas-parser-go/commit/fb1332a253d0c529f15a279ab4efad27d1d313d8))

## 0.3.1 (2026-02-20)

Full Changelog: [v0.3.0...v0.3.1](https://github.com/CASParser/cas-parser-go/compare/v0.3.0...v0.3.1)

### Bug Fixes

* allow canceling a request while it is waiting to retry ([63aa0f0](https://github.com/CASParser/cas-parser-go/commit/63aa0f04dc1cd584e50deb4ba2f529ae54d0e624))
* **client:** use correct format specifier for header serialization ([9b3f4fd](https://github.com/CASParser/cas-parser-go/commit/9b3f4fd9c3c01fcebacb10306c468b1019888e8b))


### Chores

* **internal:** remove mock server code ([2c3a380](https://github.com/CASParser/cas-parser-go/commit/2c3a380445b0150b03cee471ac806dc1710ad8fc))
* update mock server docs ([059ab0b](https://github.com/CASParser/cas-parser-go/commit/059ab0ba58a86bbeacf3482d64a0e44a1157355e))

## 0.3.0 (2026-02-14)

Full Changelog: [v0.2.2...v0.3.0](https://github.com/CASParser/cas-parser-go/compare/v0.2.2...v0.3.0)

### Features

* **api:** manual updates ([5811aac](https://github.com/CASParser/cas-parser-go/commit/5811aac49429fcacb608f6c78baa59bc8cd02be7))
* **api:** manual updates ([1d4b13f](https://github.com/CASParser/cas-parser-go/commit/1d4b13fcd33b0a79a8ebef02a2d85d2399783851))

## 0.2.2 (2026-02-14)

Full Changelog: [v0.2.1...v0.2.2](https://github.com/CASParser/cas-parser-go/compare/v0.2.1...v0.2.2)

### Chores

* update SDK settings ([5ab428d](https://github.com/CASParser/cas-parser-go/commit/5ab428d440f6d14a609c03aeed33db018c6f638d))

## 0.2.1 (2026-02-14)

Full Changelog: [v0.2.0...v0.2.1](https://github.com/CASParser/cas-parser-go/compare/v0.2.0...v0.2.1)

### Chores

* configure new SDK language ([d3eaa94](https://github.com/CASParser/cas-parser-go/commit/d3eaa94f97fdc89fdd307a7e53d6bb48cf8b60f8))

## 0.2.0 (2026-02-14)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/CASParser/cas-parser-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** manual updates ([0e09dc2](https://github.com/CASParser/cas-parser-go/commit/0e09dc270811ed3b6e2d30164ae03d45d258ab26))


### Chores

* update SDK settings ([3e236d2](https://github.com/CASParser/cas-parser-go/commit/3e236d27c63d4b603c610cebd8c5a68adbcdd372))

## 0.1.0 (2026-02-14)

Full Changelog: [v0.0.3...v0.1.0](https://github.com/CASParser/cas-parser-go/compare/v0.0.3...v0.1.0)

### Features

* **api:** api update ([321c851](https://github.com/CASParser/cas-parser-go/commit/321c851d5ce4729962559904ec4e13db1d8febdd))
* **api:** api update ([58bd341](https://github.com/CASParser/cas-parser-go/commit/58bd3413a106d7e2e99011664649d9ea80f47286))
* **api:** api update ([58f353e](https://github.com/CASParser/cas-parser-go/commit/58f353ea9385874ac4f84ae9107bcee481fd79f9))
* **api:** api update ([4de0c4b](https://github.com/CASParser/cas-parser-go/commit/4de0c4b7936cbe00370980ab20aad02586090fcd))
* **api:** api update ([c4b3ed4](https://github.com/CASParser/cas-parser-go/commit/c4b3ed40ff9e5d6ad1ac4ca57e0219d184038d20))
* **api:** api update ([4cafc4e](https://github.com/CASParser/cas-parser-go/commit/4cafc4e634e376310e635eeeacd9c6bc1656d7ca))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([48f9ffd](https://github.com/CASParser/cas-parser-go/commit/48f9ffda3a0053adcf1a68c5378bbeac9de9027c))
* use slices.Concat instead of sometimes modifying r.Options ([c30d79a](https://github.com/CASParser/cas-parser-go/commit/c30d79aa4eb1f8b4d793fe5e641db7d0f385a9dc))


### Chores

* bump minimum go version to 1.22 ([54f1ae1](https://github.com/CASParser/cas-parser-go/commit/54f1ae1bedb207621a92b70a64220456781a8e37))
* do not install brew dependencies in ./scripts/bootstrap by default ([a665c6a](https://github.com/CASParser/cas-parser-go/commit/a665c6aa0d14754a7f4ea644b2b53ac49dec0f0c))
* **internal:** codegen related update ([b0d066b](https://github.com/CASParser/cas-parser-go/commit/b0d066be8e48533c4e0bf3f004c3c6ad96af0008))
* **internal:** grammar fix (it's -&gt; its) ([0af28e9](https://github.com/CASParser/cas-parser-go/commit/0af28e95581cfd4c2d9f729ee59f8213253da54a))
* update more docs for 1.22 ([fc3b56e](https://github.com/CASParser/cas-parser-go/commit/fc3b56e861603f65fd1a61d6c843b220e94154cc))
* update SDK settings ([819a23b](https://github.com/CASParser/cas-parser-go/commit/819a23b89c3fcbe421721b9c7061f6d2d975a589))

## 0.0.3 (2025-09-06)

Full Changelog: [v0.0.2...v0.0.3](https://github.com/CASParser/cas-parser-go/compare/v0.0.2...v0.0.3)

### Bug Fixes

* close body before retrying ([891bbc6](https://github.com/CASParser/cas-parser-go/commit/891bbc6d9b7b2840269c438f9242274fbab9fe97))
* **internal:** unmarshal correctly when there are multiple discriminators ([08c4f8a](https://github.com/CASParser/cas-parser-go/commit/08c4f8ad7fa898837bef8f1a8898f3e8b56329d9))


### Chores

* **internal:** codegen related update ([86015d2](https://github.com/CASParser/cas-parser-go/commit/86015d220838b6d6d72472ba134308e8c851100b))

## 0.0.2 (2025-08-18)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/CASParser/cas-parser-go/compare/v0.0.1...v0.0.2)

### Chores

* configure new SDK language ([330270c](https://github.com/CASParser/cas-parser-go/commit/330270ce22229cbe4b11925afbc071b28c9399da))
* update SDK settings ([5930893](https://github.com/CASParser/cas-parser-go/commit/5930893d2344d77bbac819d9dd66ff3f5074f072))
