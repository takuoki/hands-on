# xss

## Get started

1. Run Nuxt.js

```bash
# install dependencies
$ npm install

# serve with hot reload at localhost:3000
$ npm run dev
```

2. Normal use

    Go to http://localhost:3000/mypage.

    - This application stores data in cookies.

3. Malicious use

    Go to http://localhost:3000/evil and click link.

## What's wrong

`./pages/mypage/confirm.vue`

* XSS vulnerabilities

```js
          <p>
            名前：<span v-html="name"></span>
          </p>
          <p>
            メールアドレス：<span v-html="email"></span>
          </p>
```

* Fixed code

```js
          <p>
            名前：{{ name }}
          </p>
          <p>
            メールアドレス：{{ email }}
          </p>
```

## References

- [クロスサイト・スクリプティングの説明](https://www.ipa.go.jp/security/vuln/websecurity-HTML-1_5.html)
- [Vueのテンプレート構文](https://jp.vuejs.org/v2/guide/syntax.html)
