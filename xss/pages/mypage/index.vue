<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card class="py-4 justify-center">
        <v-card-title>
          ユーザー情報
        </v-card-title>
        <v-card-text>
          <p>
            名前：{{ name }}
          </p>
          <p>
            メールアドレス：{{ email }}
          </p>
          <v-btn
            class="primary"
            :href="`/mypage/update?name=${name}&email=${email}`"
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
    async asyncData ({ app }) {
      return {
        name: app.$cookies.get("name"),
        email: app.$cookies.get("email"),
      }
    },
  }
</script>
