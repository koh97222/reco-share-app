<template>
  <div>
    <b-modal centered ref="user-regist" ok-only>
      <h2 class="text-center">
        reco<span class="e5b848">×</span>share
        <font-awesome-icon icon="headphones-alt" class="ml-1" />
      </h2>
      <p class="text-center">登録して友達の投稿を見てみよう。</p>
      <hr />
      <b-card class="mt-4 mb-4">
        <b-row>
          <b-col xl="1"></b-col>
          <b-col xl="10">
            <b-input-group prepend="@">
              <b-input
                id="inline-form-input-name"
                placeholder="ユーザ名"
              ></b-input>
            </b-input-group>
            <b-input-group class="mt-1">
              <b-input
                id="inline-form-input-mailaddress"
                placeholder="メールアドレス"
              ></b-input>
            </b-input-group>
            <b-input-group class="mt-1">
              <b-input
                id="text-password"
                type="password"
                placeholder="パスワード(6文字以上)"
              ></b-input>
            </b-input-group>
            <b-input-group class="mt-1">
              <b-input
                id="text-password-retyping"
                type="password"
                placeholder="パスワード(再入力)"
              ></b-input>
            </b-input-group>
          </b-col>
          <b-col xl="1"></b-col>
        </b-row>
      </b-card>
      <template v-slot:modal-footer>
        <b-button size="sm" variant="secondary" @click="ok()">
          登録する<font-awesome-icon icon="user-plus" class="ml-1" />
        </b-button>
      </template>
    </b-modal>
  </div>
</template>
<script>
export default {
  data() {
    return {};
  },
  methods: {},
};
</script>
<style scoped>
.btn {
  color: black;
  background-color: white;
  box-shadow: 2px 2px 4px #cccccc;
  width: 170px;
}
.btn:hover {
  background-color: #e5b848;
  color: black;
}
.card-body {
  background-color: #f2f2f2;
}
.e5b848 {
  color: #e5b848;
}
</style>
