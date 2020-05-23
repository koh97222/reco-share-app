<template>
  <div>
    <b-modal centered ref="user-regist" ok-only>
      <h2 class="text-center">
        reco<span class="e5b848">×</span>share
        <font-awesome-icon icon="headphones-alt" class="ml-1" />
      </h2>
      <p class="text-center">登録して友達の投稿を見てみよう。</p>
      <div
        class="text-center"
        style="color:red"
        v-for="(error, index) in validateErrorList"
        :key="index"
      >
        {{ error }}
      </div>

      <hr />
      <b-card class="mt-4 mb-4">
        <b-row>
          <b-col xl="1"></b-col>
          <b-col xl="10">
            <b-input-group prepend="@">
              <b-input
                id="inline-form-input-name"
                placeholder="ユーザ名"
                v-model="userInfo.userName"
              ></b-input>
            </b-input-group>
            <b-input-group class="mt-1">
              <b-input
                id="inline-form-input-mailaddress"
                placeholder="メールアドレス"
                v-model="userInfo.mail"
              ></b-input>
            </b-input-group>
            <b-input-group class="mt-1">
              <b-input
                id="text-password"
                type="password"
                placeholder="パスワード(6文字以上)"
                v-model="userInfo.password"
              ></b-input>
            </b-input-group>
            <b-input-group class="mt-1">
              <b-input
                id="text-password-retyping"
                type="password"
                placeholder="パスワード(再入力)"
                v-model="userInfo.confirmedPassword"
              ></b-input>
            </b-input-group>
          </b-col>
          <b-col xl="1"></b-col>
        </b-row>
      </b-card>
      <template v-slot:modal-footer>
        <b-button
          size="sm"
          variant="secondary"
          @click="preValidate(userInfo)"
          :disabled="isNotInputForm"
        >
          登録する<font-awesome-icon icon="user-plus" class="ml-1" />
        </b-button>
      </template>
    </b-modal>
  </div>
</template>
<script>
export default {
  data() {
    return {
      userInfo: {
        userName: null,
        mail: null,
        password: null,
        confirmedPassword: null,
      },
      validateErrorList: [],
    };
  },
  computed: {
    isNotInputForm() {
      return (
        // もう少しスマートに書きたい
        this.userInfo.userName == null ||
        this.userInfo.userName == "" ||
        this.userInfo.mail == null ||
        this.userInfo.mail == "" ||
        this.userInfo.password == null ||
        this.userInfo.password == "" ||
        this.userInfo.confirmedPassword == null ||
        this.userInfo.confirmedPassword == ""
      );
    },
  },
  methods: {
    // もう少しスマートに初期化したい。
    init() {
      this.userInfo.userName = null;
      this.userInfo.mail = null;
      this.userInfo.password = null;
      this.userInfo.confirmedPassword = null;
      this.validateErrorList = [];
    },
    preValidate(userInfo) {
      console.log(userInfo);
      this.validateErrorList = [];
      // パスワードと確認パスワードを比較し、同じでなければエラー
      if (userInfo.password !== userInfo.confirmedPassword) {
        this.validateErrorList.push(
          "入力されたパスワードと確認用パスワードが異なります。"
        );
      }

      // パスワードが6文字以上でなければエラー
      if (userInfo.password.length < 6) {
        this.validateErrorList.push(
          "パスワードは6文字以上で入力してください。"
        );
      }
      // メールアドレスの形式でなければエラー
      // TODO:機能しない。修正が必要。
      let regex = RegExp(
        "/^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$/"
      );
      if (!regex.test(userInfo.password)) {
        this.validateErrorList.push("メールアドレスの形式で入力してください。");
        console.log("huhu");
      }
    },
  },
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
