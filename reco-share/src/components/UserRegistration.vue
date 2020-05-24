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
              <b-input placeholder="ユーザ名" v-model="userInfo.user"></b-input>
            </b-input-group>
            <b-input-group class="mt-1">
              <b-input
                placeholder="メールアドレス"
                v-model="userInfo.email"
              ></b-input>
            </b-input-group>
            <b-input-group class="mt-1">
              <b-input
                type="password"
                placeholder="パスワード(6文字以上)"
                v-model="userInfo.pass"
              ></b-input>
            </b-input-group>
            <b-input-group class="mt-1">
              <b-input
                type="password"
                placeholder="パスワード(再入力)"
                v-model="userInfo.confirmedPass"
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
    <b-modal centered ref="after-regist" ok-only>
      <h3 class="text-center mb-3">
        User registration is complete.
      </h3>

      <p class="text-center">
        reco<span class="e5b848">×</span>share
        <font-awesome-icon icon="headphones-alt" class="ml-1" />
        へのご登録ありがとうございます。
      </p>
      <template v-slot:modal-footer>
        <b-button size="sm" variant="secondary" @click="close">閉じる</b-button>
      </template>
    </b-modal>
  </div>
</template>
<script>
export default {
  data() {
    return {
      userInfo: {
        user: null,
        email: null,
        pass: null,
        confirmedPass: null,
      },
      validateErrorList: [],
    };
  },
  computed: {
    isNotInputForm() {
      return (
        // もう少しスマートに書きたい
        this.userInfo.user == null ||
        this.userInfo.user == "" ||
        this.userInfo.email == null ||
        this.userInfo.email == "" ||
        this.userInfo.pass == null ||
        this.userInfo.pass == "" ||
        this.userInfo.confirmedPass == null ||
        this.userInfo.confirmedPass == ""
      );
    },
  },
  methods: {
    close() {
      this.$refs["after-regist"].hide();
    },
    // もう少しスマートに初期化したい。
    init() {
      this.userInfo.user = null;
      this.userInfo.email = null;
      this.userInfo.pass = null;
      this.userInfo.confirmedPass = null;
      this.validateErrorList = [];
    },
    preValidate(userInfo) {
      this.validateErrorList = [];
      // パスワードと確認パスワードを比較し、同じでなければエラー
      if (userInfo.pass !== userInfo.confirmedPass) {
        this.validateErrorList.push(
          "入力されたパスワードと確認用パスワードが異なります。"
        );
      }

      // パスワードが6文字以上でなければエラー
      if (userInfo.pass.length < 6) {
        this.validateErrorList.push(
          "パスワードは6文字以上で入力してください。"
        );
      }
      // メールアドレスの形式でなければエラー
      // TODO:バリデーション対応させる。
      let regex = new RegExp(
        "/^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$/"
      );
      if (regex.test(userInfo.email)) {
        this.validateErrorList.push("メールアドレスの形式で入力してください。");
      }
      // エラーがない場合は、登録処理に進む。
      if (this.validateErrorList.length == 0) {
        this.registUser(userInfo);
      }
    },
    registUser(userInfo) {
      this.$axios
        .post("http://localhost:8082/registuser", userInfo)
        .then((r) => {
          if (r.status == 200) {
            if (r.data.resultCode == "80") {
              this.validateErrorList.push("ユーザ名が既に登録されています。");
            } else {
              this.$refs["user-regist"].hide();
              this.$refs["after-regist"].show();
            }
          }
        });
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
