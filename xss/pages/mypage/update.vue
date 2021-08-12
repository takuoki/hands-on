<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card class="py-4 justify-center">
        <v-card-title>
          ユーザー情報修正
        </v-card-title>
        <v-card-text>
          <v-text-field
            v-model="name"
            label="名前"
          ></v-text-field>
          <v-text-field
            v-model="email"
            label="メールアドレス"
          ></v-text-field>
          <v-btn
            class="primary"
            :href="`/mypage/confirm?name=${name}&email=${email}`"
          >
            確認
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
    }
  }
</script>
