<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card class="py-4 justify-center">
        <v-card-title>
          ユーザー情報確認
        </v-card-title>
        <v-card-text>
          <p>
            名前：<span v-html="name"></span>
          </p>
          <p>
            メールアドレス：<span v-html="email"></span>
          </p>
          <v-btn
            class="mr-4 primary"
            :href="`/mypage/update?name=${name}&email=${email}`"
          >
            戻る
          </v-btn>
          <v-btn
            class="primary"
            @click="update"
          >
            更新
          </v-btn>
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
  export default {
    async asyncData ({ query }) {
      return {
        name: query.name,
        email: query.email,
      }
    },
    methods: {
      update() {
        this.$cookies.set("name", this.name)
        this.$cookies.set("email", this.email)
        this.$router.push('/mypage')
      }
    }
  }
</script>
